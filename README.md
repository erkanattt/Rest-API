tags:
- golang
- rest-api 
- gin 
- gorm

Интернет дүкенге арналған REST API

Бұл құжат интернет дүкенге арналған REST API жобасына түсініктеме береді. Жоба Go тілінде жазылған, Gin фреймворкі мен GORM кітапханасы қолданылған. Дерекқор ретінде PostgreSQL пайдаланылады.

1. Gin фреймворкі

Gin — Go тілінде жазылған, HTTP сервер құруға арналған жеңіл әрі жылдам фреймворк. Оның көмегімен API жасау оңай.

Артықшылықтары:
- Жылдам жұмыс істейді (net/http негізінде)
- JSON-пен жұмыс істеу оңай 
- Middleware қосуға болады 

Орнату

go get -u github.com/gin-gonic/gin
Қарапайым сервер мысалы
package main
import "github.com/gin-gonic/gin"
func main() {
r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "API жұмыс істеп тұр"})
    })

    r.Run(":8080")
}

2. GORM — Go үшін ORM

GORM — Go тілінде дерекқорлармен жұмыс істеуге арналған ORM кітапханасы.

Орнату

go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

Дерекқорға қосылу

import (
"gorm.io/driver/postgres"
"gorm.io/gorm"
"log"
)

func main() {
dsn := "host=localhost user=postgres password=7982 dbname=postgres port=5432 sslmode=disable"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
log.Fatal("Қосылу қатесі:", err)
}
}

3. Өнім моделі мен кесте құру

Модель

type Product struct {
ID          uint   `gorm:"primaryKey"`
Name        string `json:"name"`
Description string `json:"description"`
Price       uint   `json:"price"`
}

Миграция

db.AutoMigrate(&Product{})

4. Негізгі CRUD маршруттар

Барлық өнімді алу

r.GET("/products", func(c *gin.Context) {
var products []Product
db.Find(&products)
c.JSON(200, products)
})

Бір өнімді көру

r.GET("/products/:id", func(c *gin.Context) {
var product Product
if err := db.First(&product, c.Param("id")).Error; err != nil {
c.JSON(404, gin.H{"error": "Тауар табылмады"})
return
}
c.JSON(200, product)
})

Жаңа өнім қосу

r.POST("/products", func(c *gin.Context) {
var product Product
if err := c.ShouldBindJSON(&product); err != nil {
c.JSON(400, gin.H{"error": err.Error()})
return
}
db.Create(&product)
c.JSON(201, product)
})

Өнімді жаңарту

r.PUT("/products/:id", func(c *gin.Context) {
var product Product
if err := db.First(&product, c.Param("id")).Error; err != nil {
c.JSON(404, gin.H{"error": "Тауар табылмады"})
return
}
var input Product
if err := c.ShouldBindJSON(&input); err != nil {
c.JSON(400, gin.H{"error": err.Error()})
return
}
db.Model(&product).Updates(input)
c.JSON(200, product)
})

Өнімді жою

r.DELETE("/products/:id", func(c *gin.Context) {
var product Product
if err := db.First(&product, c.Param("id")).Error; err != nil {
c.JSON(404, gin.H{"error": "Тауар табылмады"})
return
}
db.Delete(&product)
c.JSON(200, gin.H{"message": "Өнім жойылды"})
})

5. Тексеру мысалы (Postman арқылы)

GET http://localhost:8080/products

POST http://localhost:8080/products

{
"name": "Ноутбук",
"description": "Жаңа үлгі",
"price": 250000
}

6. Қорытынды

Бұл жобада интернет дүкенге арналған REST API құрдым. Gin мен GORM көмегімен тауарларды қосу, көру, жаңарту, жою сияқты әрекеттер орындалады. Код құрылымы қарапайым және кеңейтуге ыңғайлы.

