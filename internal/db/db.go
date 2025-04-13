package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/erkanattt/go-rest-crud-project/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

func InitDB() {
	dbHost := "localhost"
	dbName := "postgres"
	dbUser := "postgres"
	dbPass := "7982"
	dbPort := "5432"
	sslmode := "disable"

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, sslmode)

	// Raw SQL connection for migrate
	sqlDB, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Golang Migrate
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:../internal/db/migrations", "postgres", driver)
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

	if err := DB.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
		log.Fatal("Failed to auto migrate models: ", err)
	}
}
