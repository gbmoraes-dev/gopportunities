package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpportunityHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "UpdateOpportunityHandler",
	})
}