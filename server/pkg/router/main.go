package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rohan184/server/pkg/controller"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "hello world")
	})

	r.POST("/post", controller.PostResult)
	r.GET("/result", controller.Result)

	return r
}
