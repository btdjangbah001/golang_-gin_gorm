package controllers

import (
	"net/http"

	"github.com/btdjangbah001/gin-basics/models"
	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	books, err := models.FindAllBooks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": books})
}
