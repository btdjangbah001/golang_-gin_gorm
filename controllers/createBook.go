package controllers

import (
	"net/http"

	"github.com/btdjangbah001/gin-basics/models"
	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	newBook := models.Book{Title: input.Title, Author: input.Author}

	_, err = newBook.SaveBook()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"data": newBook})
}
