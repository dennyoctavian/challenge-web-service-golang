package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const PORT = ":8000"

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"name_book"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var db *gorm.DB
var err error

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=project2 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&Book{})
	r := gin.Default()
	r.GET("/books", getAllBook)
	r.GET("/books/:id", getBookById)
	r.POST("/books", addBook)
	r.PUT("/books/:id", updateBookById)
	r.DELETE("/books/:id", deleteBookById)
	r.Run()
}
func getAllBook(c *gin.Context) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, books)
	}
}

func getBookById(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Book
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, book)
	}

}

func addBook(c *gin.Context) {
	var book Book
	c.BindJSON(&book)
	db.Create(&book)
	c.JSON(200, book)
}

func updateBookById(c *gin.Context) {
	var book Book
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&book).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&book)
	db.Save(&book)
	c.JSON(200, book)
}

func deleteBookById(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Book
	d := db.Where("id = ?", id).Delete(&book)
	fmt.Println(d)
	c.JSON(200, gin.H{"message": "Book beleted successfully"})

}
