package main

import (
	"github.com/btdjangbah001/gin-basics/controllers"
	"github.com/btdjangbah001/gin-basics/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	r := router.Group("api/v1")

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBookById)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	models.ConnectDatabase()

	router.Run()
}
