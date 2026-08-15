package recorder

import (
	"database/sql"
	"log"
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
