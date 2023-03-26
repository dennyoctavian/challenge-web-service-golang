package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const PORT = ":8000"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var books = []Book{}

func main() {
	r := gin.Default()
	r.GET("/books", getAllBook)
	r.GET("/books/:id", getBookById)
	r.POST("/books", addBook)
	r.PUT("/books/:id", updateBookById)
	r.DELETE("/books/:id", deleteBookById)
	r.Run()
}
func getAllBook(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBookById(c *gin.Context) {
	bookID := c.Param("id")

	bookIDInteger, _ := strconv.ParseInt(bookID, 10, 64)

	var bookData Book

	for _, book := range books {
		if int(bookIDInteger) == book.ID {
			bookData = book
			break
		}
	}

	if bookData.ID == 0 {
		var message = fmt.Sprintf("Not found book with id %s", bookID)
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": message,
		})
		return
	} else {
		c.JSON(http.StatusOK, bookData)
		return
	}

}

func addBook(c *gin.Context) {
	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	books = append(books, newBook)
	c.JSON(http.StatusOK, newBook)
}

func updateBookById(c *gin.Context) {
	bookID := c.Param("id")
	bookIDInteger, _ := strconv.ParseInt(bookID, 10, 64)

	var updatedBook Book

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	updatedBook.ID = int(bookIDInteger)

	for i, book := range books {
		if int(bookIDInteger) == book.ID {
			books[i] = updatedBook
			break
		}
	}
	c.String(http.StatusOK, "Updated")
}

func deleteBookById(c *gin.Context) {
	bookID := c.Param("id")
	bookIDInteger, _ := strconv.ParseInt(bookID, 10, 64)
	condition := false
	var bookIndex int

	for i, book := range books {
		if int(bookIDInteger) == book.ID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	copy(books[bookIndex:], books[bookIndex+1:])
	books[len(books)-1] = Book{}
	books = books[:len(books)-1]

	c.String(http.StatusOK, "Deleted")

}
