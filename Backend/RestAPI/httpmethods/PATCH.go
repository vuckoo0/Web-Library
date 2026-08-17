package httpmethods

import (
	"fmt"
	"log"
	"main/recorder"
	"strconv"

	"github.com/gin-gonic/gin"
)

func editBookFromDB(id int, field, newValue string, errChan chan error) {
	_, err := recorder.LibraryDB.Exec(fmt.Sprintf("update books set %s = ? where id = ?", field), newValue, id)
	errChan <- err
}

func bookID(title string) ([]int, error) {

	result, err := recorder.LibraryDB.Query("select id where title = ?", title)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	IDs := []int{}
	for result.Next() {

		var currentID int
		err = result.Scan(&currentID)

		if err != nil {
			return nil, err
		}

		IDs = append(IDs, currentID)
	}

	err = result.Err()
	if err != nil {
		return nil, err
	}

	return IDs, nil
}

func HandlePATCH(c *gin.Context) {

	Id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Book Id"})
		return
	}

	var edit struct {
		Field    string `json:"field"`
		NewValue string `json:"new_value"`
	}

	err = c.ShouldBindJSON(&edit)

	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	errorChanel := make(chan error, 1)
	go editBookFromDB(Id, edit.Field, edit.NewValue, errorChanel)

	err = <-errorChanel
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	c.JSON(200, edit)
}
