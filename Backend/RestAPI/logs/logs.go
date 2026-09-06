package logs

import (
	"log"
	"os"
)

var (
	ErrorLog *os.File
)

func Setup() error {

	ErrorLog, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	log.SetOutput(ErrorLog)
	return nil
}
