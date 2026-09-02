package main

import (
	api "main/api"
	middleware "main/middleware"
	storage "main/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://127.0.0.1:5500"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	storage.LibraryDB = storage.ConnectDataBase()
	defer storage.LibraryDB.Close()

	router.GET("/books", middleware.AuthenticationMiddleware(), api.HandleGET)
	router.GET("/books/search", middleware.AuthenticationMiddleware(), api.HandleGETWithURLParams)

	router.POST("/login", api.HandleLogIn)
	router.POST("/register", api.HandleSignUp)
	router.POST("/books", api.HandlePOST)

	router.PATCH("/books", api.HandlePATCH)

	router.Run(":8080")
}
