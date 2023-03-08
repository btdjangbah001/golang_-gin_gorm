package controllers

import (
	"net/http"

	"github.com/btdjangbah001/gin-basics/models"
	"github.com/gin-gonic/gin"
)

func FindBookById(c *gin.Context) {
	book, err := models.FindBook(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Cannot find book with id " + c.Param("id")})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": &book})
}
