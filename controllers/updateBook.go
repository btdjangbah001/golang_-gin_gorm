package controllers

import (
	"net/http"

	"github.com/btdjangbah001/gin-basics/models"
	"github.com/gin-gonic/gin"
)

func UpdateBook(c *gin.Context) {
	book, err := models.FindBook(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cannot find book with id " + c.Param("id")})
		return
	}

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// models.DB.Model(book).Updates(input)  // returns old book 🤔
	_, err = book.UpdateBook(&input)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": book})
}
