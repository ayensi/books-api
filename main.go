package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("book/byId/:id", getBookById)
	router.POST("/book/create", createBook)
	router.PATCH("book/update/:id", updateBookById)
	router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookById(c *gin.Context) {
	book, err := bookById(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	}
	c.IndentedJSON(http.StatusFound, book)
}

func updateBookById(c *gin.Context) {
	book, err := bookById(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	}

	c.BindJSON(&book)

	c.IndentedJSON(http.StatusOK, book)

}

func bookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}
