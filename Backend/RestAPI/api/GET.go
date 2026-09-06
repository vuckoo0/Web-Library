package api

import (
	models "main/models"
	storage "main/storage"

	"github.com/gin-gonic/gin"
)

func HandleLoadingBooks(c *gin.Context) {

	bookChanel := make(chan []models.Book, 1)
	errorChanel := make(chan error, 1)

	go storage.LoadBooks(bookChanel, errorChanel)

	books := <-bookChanel
	c.JSON(200, books)
}

func HandleBookSearch(c *gin.Context) {

	title := c.Query("title")

	booksChanel := make(chan []models.Book, 1)
	errorChanel := make(chan error, 1)

	go storage.FindBooksWithTitle(title, booksChanel, errorChanel)

	books := <-booksChanel
	err := <-errorChanel

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, books)
}
