package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpportunityHandler(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"message": "DeleteOpportunityHandler",
	})
}