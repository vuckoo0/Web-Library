package main

import (
	api "main/api"
	logs "main/logs"
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
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:    []string{"Content-Type", "Authorization"},
	}))

	logs.SetupErrorLogs()

	storage.LibraryDB = storage.ConnectDataBase()
	defer storage.LibraryDB.Close()

	router.GET("/books", middleware.AuthenticationMiddleware(), api.HandleLoadingBooks)
	router.GET("/books/search", middleware.AuthenticationMiddleware(), api.HandleBookSearch)

	router.POST("/login", api.HandleLogIn)
	router.POST("/register", api.HandleSignUp)

	router.POST("/books", middleware.AuthenticationMiddleware(), middleware.PrivilegeAuthorization(0), api.HandleAddingBook)
	router.PATCH("/books", middleware.AuthenticationMiddleware(), middleware.PrivilegeAuthorization(0), api.HandleBookFieldEdit)

	router.Run(":8080")
}
