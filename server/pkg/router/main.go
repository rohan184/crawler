package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rohan184/server/pkg/controller"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/insight", controller.Insight)
	r.GET("/insights", controller.GetInsights)
	r.DELETE("/insight/:id", controller.RemoveInsight)
	r.PUT("/insight/:id", controller.FavouriteInsight)

	return r
}
