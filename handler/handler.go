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

func ShowOpportunityHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ShowOpportunityHandler",
	})
}

func UpdateOpportunityHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "UpdateOpportunityHandler",
	})
}

func DeleteOpportunityHandler(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"message": "DeleteOpportunityHandler",
	})
}

func ListOpportunitiesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ListOpportunitiesHandler",
	})
}