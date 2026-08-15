package httpmethods

import (
	"log"
	recorder "main/recorder"

	"github.com/gin-gonic/gin"
)

func (b *Book) InsertBook(errChan chan error) {

	result, err := recorder.LibraryDB.Exec("insert into books (title, author, isbn) values (?, ?, ?)", b.Title, b.Author, b.ISBN)
	if err != nil {
		log.Print("[RestAPI]: Error in inerting book into library database: ", err)
		errChan <- err
	}

	b.Id, err = result.LastInsertId()
	if err != nil {
		log.Print("[RestAPI]: Error in Id-ing inserted book: ", err)
		errChan <- err
	}

	errChan <- nil
}

func HandlePOST(c *gin.Context) {

	var newBook Book
	errorChanel := make(chan error)

	err := c.ShouldBindJSON(&newBook)

	if err != nil {

		log.Print("[RestAPI]: Error in loading book sent from frontend: ", err)
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	newBook.InsertBook(errorChanel)

	err = <-errorChanel
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, newBook)
}
