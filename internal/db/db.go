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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, sslmode)

	sqlDB, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, sslmode,
	))
	if err != nil {
		log.Fatal(err)
	}

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

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = gormDB

	if err := DB.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
		log.Fatal("Миграция кезінде қате болды: ", err)
	}

}
