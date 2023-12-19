package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/create-opportunity", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"message": "POST Opportunity",
			})
		})

		v1.GET("/opportunity", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET Opportunity",
			})
		})

		v1.PUT("/update-opportunity", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "PUT Opportunity",
			})
		})

		v1.DELETE("/delete-opportunity", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "DELETE Opportunity",
			})
		})

		v1.GET("/opportunities", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET Opportunities",
			})
		})
	}
}