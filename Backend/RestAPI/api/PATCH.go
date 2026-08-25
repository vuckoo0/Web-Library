package api

import (
	"log"
	"main/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	go storage.EditBookFromDB(Id, edit.Field, edit.NewValue, errorChanel)

	err = <-errorChanel
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	c.JSON(200, edit)
}
