package httpmethods

import (
	"main/recorder"

	"github.com/gin-gonic/gin"
)

func findBooksWithTitle(title string, c chan []Book, errChan chan error) {

	rows, err := recorder.LibraryDB.Query("select * from books where title like ?", "%"+title+"%")
	if err != nil {
		c <- nil
		errChan <- err
		return
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {

		var b Book
		err := rows.Scan(&b.Id, &b.Title, &b.Author, &b.ISBN)
		if err != nil {
			continue
		}

		books = append(books, b)
	}

	c <- books
	errChan <- rows.Err()
}

func HandleGETWithURLParams(c *gin.Context) {

	title := c.Query("title")

	booksChanel := make(chan []Book, 1)
	errorChanel := make(chan error, 1)

	go findBooksWithTitle(title, booksChanel, errorChanel)

	books := <-booksChanel
	err := <-errorChanel

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, books)
}
