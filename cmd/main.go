package main

import (
	"github.com/erkanattt/go-rest-crud-project/internal/db"
	"github.com/erkanattt/go-rest-crud-project/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r, db.DB)
	r.Run(":8080")
}
