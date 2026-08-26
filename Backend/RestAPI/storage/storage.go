package storage

import (
	"database/sql"
	"fmt"
	"log"
	"main/config"
	models "main/models"

	"golang.org/x/crypto/bcrypt"
)

var (
	LibraryDB *sql.DB
)

func ConnectDataBase() *sql.DB {

	db, err := sql.Open("mysql", config.Config().DB_DSN)
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

func IsUsernameTaken(username string) (bool, error) {

	var id int
	err := LibraryDB.QueryRow("select id from users where `Name` = ?", username).Scan(&id)

	switch err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return true, nil
	default:
		return false, err
	}
}

func AddUserToDB(u *models.User, errorChanel chan error) {

	usedUsername, err := IsUsernameTaken(u.Name)
	if err != nil {
		errorChanel <- err
		return
	}

	if usedUsername {
		errorChanel <- fmt.Errorf("Username already taken!")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		errorChanel <- err
		return
	}

	u.Password = string(hashedPassword[:])

	result, err := LibraryDB.Exec("insert into users (`name`, `password`, privilege) values (?, ?, ?)", u.Name, u.Password, u.Privilege)
	if err != nil {
		errorChanel <- err
		return
	}

	u.Id, err = result.LastInsertId()
	if err != nil {
		errorChanel <- err
		return
	}

	errorChanel <- nil
}
