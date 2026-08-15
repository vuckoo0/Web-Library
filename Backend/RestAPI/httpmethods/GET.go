package httpmethods

import (
	"log"
	recorder "main/recorder"

	"github.com/gin-gonic/gin"
)

func loadBooks(c chan []Book) {

	bookTable, err := recorder.LibraryDB.Query("select * from books")
	if err != nil {
		log.Print("[RestAPI]: Error in quering library database: ", err)
		c <- nil
	}

	defer bookTable.Close()

	var currentBookId int = 1
	var books []Book

	for bookTable.Next() {

		var currentBook Book
		err := bookTable.Scan(&currentBook.Id, &currentBook.Title, &currentBook.Author, &currentBook.ISBN)
		if err != nil {
			log.Print("[RestAPI]: Error in reading book with Id: ", currentBookId)
			continue
		}

		currentBookId += 1
		books = append(books, currentBook)
	}

	err = bookTable.Err()
	if err != nil {
		log.Print("[RestAPI]: Unexpected error:", err)
		c <- nil
	}

	c <- books
}

func HandleGET(c *gin.Context) {

	bookChanel := make(chan []Book)
	go loadBooks(bookChanel)

	books := <-bookChanel
	c.JSON(200, books)
}
