package api

import (
	"fmt"
	"log"
	"main/auth"
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

func HandleSignUp(c *gin.Context) {

	var newUser models.User

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	errorChanel := make(chan error)
	go storage.AddUserToDB(&newUser, errorChanel)

	err = <-errorChanel
	fmt.Println(err)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{})
}

func HandleLogIn(c *gin.Context) {

	var loggingUser models.User

	err := c.ShouldBindJSON(&loggingUser)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = auth.CheckUserCredentials(&loggingUser)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.GenerateToken(loggingUser.Id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"id":        loggingUser.Id,
		"name":      loggingUser.Name,
		"privilege": loggingUser.Privilege,
		"token":     token,
	})
}
