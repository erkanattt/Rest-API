package main

import (
	"github.com/erkanattt/go-rest-crud-project/internal/db"
	"github.com/erkanattt/go-rest-crud-project/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	r.Static("/static", "./frontend/static")

	r.GET("/", func(c *gin.Context) {
		c.File("../frontend/index.html")
	})

	r.GET("/login", func(c *gin.Context) {
		c.File("../frontend/login.html")
	})

	r.GET("/register", func(c *gin.Context) {
		c.File("../frontend/register.html")
	})

	r.GET("/admin", func(c *gin.Context) {
		c.File("../frontend/admin.html")
	})

	routes.SetupRoutes(r, db.DB)

	r.Run(":8080")
}
