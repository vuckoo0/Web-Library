package main

import (
	httpmethods "main/httpmethods"
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

	router.GET("/books", httpmethods.HandleGET)
	router.POST("/books", httpmethods.HandlePOST)
	router.GET("/books/search", httpmethods.HandleGETWithURLParams)
	router.PUT("/books", httpmethods.HandlePUT)

	router.Run(":8080")
}
