package httpmethods

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (b *Book) InsertBook() error {

	result, err := LibraryDB.Exec("insert into books (title, author, isbn) values (?, ?, ?)", b.Title, b.Author, b.ISBN)
	if err != nil {
		log.Print("[RestAPI]: Error in inerting book into library database: ", err)
		return err
	}

	b.Id, err = result.LastInsertId()
	if err != nil {
		log.Print("[RestAPI]: Error in Id-ing inserted book: ", err)
		return err
	}

	return nil
}

func HandlePOST(c *gin.Context) {

	var newBook Book

	err := c.ShouldBindJSON(&newBook)

	if err != nil {

		log.Print("[RestAPI]: Error in loading book sent from frontend: ", err)
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	err = newBook.InsertBook()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, newBook)
}
