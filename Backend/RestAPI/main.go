package main

import (
	api "main/api"
	storage "main/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(cors.Default())

	storage.LibraryDB = storage.ConnectDataBase()
	defer storage.LibraryDB.Close()

	router.GET("/books", api.HandleGET)
	router.GET("/books/search", api.HandleGETWithURLParams)

	router.POST("/login", api.HandleLogIn)
	router.POST("/register", api.HandleSignUp)
	router.POST("/books", api.HandlePOST)

	router.PATCH("/books", api.HandlePATCH)

	router.Run(":8080")
}
