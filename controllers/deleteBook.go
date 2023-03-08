package controllers

import (
	"net/http"

	"github.com/btdjangbah001/gin-basics/models"
	"github.com/gin-gonic/gin"
)

func DeleteBook(c *gin.Context) {
	book, err := models.FindBook(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cannot find book with id " + c.Param("id")})
		return
	}

	book, err = book.DeleteBook()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}
