package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpportunityHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "CreateOpportunityHandler",
	})
}
