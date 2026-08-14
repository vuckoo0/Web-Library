package main

import (
	httpmethods "main/RestAPI/httpmethods"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// func FindBookWithTitle(bookTitle string) ([]Book, error) {

// 	booksWithTitle, err := libraryDB.Query("select * from books where title = ?", bookTitle)

// 	if err != nil {
// 		log.Print("[RestAPI]: Error in finding book into library database: ", err)
// 		return nil, err
// 	}

// 	defer booksWithTitle.Close()

// 	currentBookId := 1
// 	books := []Book{}

// 	for booksWithTitle.Next() {

// 		var currentBook Book
// 		err = booksWithTitle.Scan(&currentBook.Id, &currentBook.Title, &currentBook.Author, &currentBook.ISBN)

// 		if err != nil {
// 			log.Print("[RestAPI]: Error in reading book with Id: ", currentBookId)
// 			continue
// 		}

// 		currentBookId += 1
// 		books = append(books, currentBook)
// 	}

// 	err = booksWithTitle.Err()
// 	if err != nil {
// 		log.Print("[RestAPI]: Unexpected error:", err)
// 		return nil, err
// 	}

// 	return books, nil
// }

// func HandleFIND(c *gin.Context) {

// 	var bookTitle string

// 	err := c.ShouldBindJSON(bookTitle)

// 	if err != nil {

// 		log.Print("[RestAPI]: Error in loading book title from frontend: ", err)
// 		c.JSON(400, gin.H{"error": err.Error()})

// 		return
// 	}

// 	books, err := FindBookWithTitle(bookTitle)

// 	if err != nil {

// 		log.Print("[RestAPI]: Error in finding books with a title: ", err)
// 		c.JSON(400, gin.H{"error": err.Error()})

// 		return
// 	}

// 	c.JSON(200, books)
// }

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.Default())

	httpmethods.LibraryDB = httpmethods.ConnectDataBase()
	defer httpmethods.LibraryDB.Close()

	router.GET("/books", httpmethods.HandleGET)
	router.POST("/books", httpmethods.HandlePOST)

	router.Run(":8080")
}
