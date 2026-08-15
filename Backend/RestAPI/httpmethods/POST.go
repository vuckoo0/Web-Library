package httpmethods

import (
	recorder "main/recorder"

	"github.com/gin-gonic/gin"
)

func (b *Book) InsertBook(errChan chan error) {

	result, err := recorder.LibraryDB.Exec("insert into books (title, author, isbn) values (?, ?, ?)", b.Title, b.Author, b.ISBN)
	if err != nil {
		errChan <- err
		return
	}

	b.Id, err = result.LastInsertId()
	if err != nil {
		errChan <- err
		return
	}

	errChan <- nil
}

func HandlePOST(c *gin.Context) {

	var newBook Book
	errorChanel := make(chan error)

	err := c.ShouldBindJSON(&newBook)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	go newBook.InsertBook(errorChanel)

	err = <-errorChanel
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, newBook)
}
