package api

import (
	"log"
	models "main/models"
	storage "main/storage"

	"github.com/gin-gonic/gin"
)

func HandleLoadingBooks(ctx *gin.Context) {

	bookChanel := make(chan []models.Book, 1)
	errorChanel := make(chan error, 1)

	go storage.LoadBooks(bookChanel, errorChanel)

	err := <-errorChanel
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	books := <-bookChanel
	ctx.JSON(200, books)
}

func HandleBookSearch(c *gin.Context) {

	title := c.Query("title")

	booksChanel := make(chan []models.Book, 1)
	errorChanel := make(chan error, 1)

	go storage.FindBooksWithTitle(title, booksChanel, errorChanel)
	err := <-errorChanel

	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	books := <-booksChanel
	c.JSON(200, books)
}
