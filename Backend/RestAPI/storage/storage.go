package storage

import (
	"database/sql"
	"fmt"
	"log"
	models "main/models"
	"os"

	"github.com/joho/godotenv"
)

var (
	LibraryDB *sql.DB
)

func loadDotenv() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("[RestAPI]: Error in loading .env file: ", err)
	}

	log.Println("[RestAPI]: Successfully loaded info from .env")
}

func ConnectDataBase() *sql.DB {

	loadDotenv()

	db, err := sql.Open("mysql", os.Getenv("DB_DNS"))
	if err != nil {
		log.Fatal("[RestAPI]: Error in opening the library database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("[RestAPI]: Error in library database pinging", err)
	}

	log.Print("[RestAPI]: Succsessfully connected to the library")

	return db
}

func LoadBooks(c chan []models.Book, errChan chan error) {

	bookTable, err := LibraryDB.Query("select * from books")
	if err != nil {
		c <- nil
		errChan <- err
		return
	}

	defer bookTable.Close()

	var currentBookId int = 1
	var books []models.Book

	for bookTable.Next() {

		var currentBook models.Book
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

func FindBooksWithTitle(title string, c chan []models.Book, errChan chan error) {

	rows, err := LibraryDB.Query("select * from books where title like ?", "%"+title+"%")
	if err != nil {
		c <- nil
		errChan <- err
		return
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {

		var b models.Book
		err := rows.Scan(&b.Id, &b.Title, &b.Author, &b.ISBN)
		if err != nil {
			continue
		}

		books = append(books, b)
	}

	c <- books
	errChan <- rows.Err()
}

func InsertBook(b *models.Book, errChan chan error) {

	result, err := LibraryDB.Exec("insert into books (title, author, isbn) values (?, ?, ?)", b.Title, b.Author, b.ISBN)
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

func EditBookFromDB(id int, field, newValue string, errChan chan error) {
	_, err := LibraryDB.Exec(fmt.Sprintf("update books set %s = ? where id = ?", field), newValue, id)
	errChan <- err
}
