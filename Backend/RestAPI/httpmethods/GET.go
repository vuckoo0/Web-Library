package httpmethods

import (
	recorder "main/recorder"

	"github.com/gin-gonic/gin"
)

func loadBooks(c chan []Book, errChan chan error) {

	bookTable, err := recorder.LibraryDB.Query("select * from books")
	if err != nil {
		c <- nil
		errChan <- err
		return
	}

	defer bookTable.Close()

	var currentBookId int = 1
	var books []Book

	for bookTable.Next() {

		var currentBook Book
		err := bookTable.Scan(&currentBook.Id, &currentBook.Title, &currentBook.Author, &currentBook.ISBN)
		if err != nil {
			continue
		}

		currentBookId += 1
		books = append(books, currentBook)
	}

	err = bookTable.Err()
	if err != nil {
		c <- nil
		errChan <- err
		return
	}

	c <- books
	errChan <- nil
}

func HandleGET(c *gin.Context) {

	bookChanel := make(chan []Book, 1)
	errorChanel := make(chan error, 1)

	go loadBooks(bookChanel, errorChanel)

	books := <-bookChanel
	c.JSON(200, books)
}
