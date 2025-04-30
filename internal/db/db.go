package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var DB *gorm.DB

func InitDB() {
	_ = godotenv.Load()
	databaseName := "postgres"
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	sslmode := "disable"
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)

	sqlDB, err := sql.Open(databaseName, dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	migrationsPath := filepath.Join(cwd, "..", "internal", "db", "migrations")
	migrationsPath = filepath.ToSlash(migrationsPath)
	migrationsURL := fmt.Sprintf("file://%s", strings.TrimPrefix(migrationsPath, "/"))
	m, err := migrate.NewWithDatabaseInstance(migrationsURL, databaseName, driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = gormDB
}
