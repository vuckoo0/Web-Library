package api

import (
	models "main/models"
	storage "main/storage"

	"github.com/gin-gonic/gin"
)

func HandlePOST(c *gin.Context) {

	var newBook models.Book
	errorChanel := make(chan error)

	err := c.ShouldBindJSON(&newBook)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	go storage.InsertBook(&newBook, errorChanel)

	err = <-errorChanel
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, newBook)
}

func HandlePOSTRegister(c *gin.Context) {

	var newUser models.User

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	errorChanel := make(chan error)
	go storage.AddUserToDB(&newUser, errorChanel)

	err = <-errorChanel
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"id":        newUser.Id,
		"name":      newUser.Name,
		"privilege": newUser.Privilege,
	})
}
