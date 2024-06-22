package main

import (
	"github.com/alkatraz445/bookstore-go/models"

	"github.com/alkatraz445/bookstore-go/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type", "Accept"},
		ExposeHeaders: []string{"X-Total-Count"},
	}))

	models.ConnectDatabase()

	router.GET("/books", controllers.FindBooks)
	router.GET("/books/:id", controllers.FindBook)
	router.POST("/books", controllers.CreateBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run()
}
