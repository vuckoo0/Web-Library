package logs

import (
	"log"
	"os"
)

func SetupErrorLogs() {

	f, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("failed to open log file: ", err)
	}
	log.SetOutput(f)
}
