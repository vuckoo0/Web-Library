package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Book struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

var (
	libraryDB *sql.DB
)

func (b Book) String() string {
	return fmt.Sprintf("%d  | %s | %s | %s\n", b.Id, b.Title, b.Author, b.ISBN)
}

func LoadDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[RestAPI]: Error in loading .env file: ", err)
	}

	log.Println("[RestAPI]: Succesfully loaded info from .env")
}

func ConnectDataBase() *sql.DB {

	LoadDotenv()

	db, err := sql.Open("mysql", os.Getenv("DB_DNS"))
	if err != nil {
		log.Fatal("[RestAPI]: Error in opening the library database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("[RestAPI]: Error in library database pinging", err)
	}

	log.Print("[RestAPI]: Succsesfully connected to the library")

	return db
}

func (b *Book) InertBook() error {

	result, err := libraryDB.Exec("insert into books (title, author, isbn) values (?, ?, ?)", b.Title, b.Author, b.ISBN)
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

func LoadBooks() []Book {

	bookTable, err := libraryDB.Query("select * from books")
	if err != nil {
		log.Print("[RestAPI]: Error in quering library database: ", err)
		return nil
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
		log.Print("[RestAPI]: Unecspected error:", err)
		return nil
	}

	return books
}

func HandleGET(c *gin.Context) {
	books := LoadBooks()
	c.JSON(200, books)
}

func HandlePOST(c *gin.Context) {

	var newBook Book

	err := c.ShouldBindJSON(&newBook)
	if err != nil {
		log.Print("[RestAPI]: Error in loading book sent from frontend: ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = newBook.InertBook()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, newBook)
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.Default())

	libraryDB = ConnectDataBase()
	defer libraryDB.Close()

	router.GET("/books", HandleGET)
	router.POST("/books", HandlePOST)

	router.Run(":8080")
}
