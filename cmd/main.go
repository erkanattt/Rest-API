package main

import (
	"github.com/erkanattt/go-rest-crud-project/internal/models"
	"github.com/erkanattt/go-rest-crud-project/internal/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Postgres контейнерінің атауын қолдану
	db, err := gorm.Open(postgres.Open("postgres://postgres:7982@postgres:5432/postgres?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Error on migrating to the DB", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":8080")
}
