package router

import (
	"github.com/gbmoraes-dev/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/create-opportunity", handler.CreateOpportunityHandler)

		v1.GET("/opportunity", handler.ShowOpportunityHandler)

		v1.PUT("/update-opportunity", handler.UpdateOpportunityHandler)

		v1.DELETE("/delete-opportunity", handler.DeleteOpportunityHandler)

		v1.GET("/opportunities", handler.ListOpportunitiesHandler)
	}
}