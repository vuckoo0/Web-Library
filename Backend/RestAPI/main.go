package main

import (
	api "main/api"
	recorder "main/recorder"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(cors.Default())

	recorder.LibraryDB = recorder.ConnectDataBase()
	defer recorder.LibraryDB.Close()

	router.GET("/books", api.HandleGET)
	router.POST("/books", api.HandlePOST)
	router.GET("/books/search", api.HandleGETWithURLParams)
	router.PATCH("/books", api.HandlePATCH)

	router.Run(":8080")
}
