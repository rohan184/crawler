package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohan184/server/pkg/database"
	"github.com/rohan184/server/pkg/resources"
	"github.com/rohan184/server/pkg/service"
)

func PostResult(c *gin.Context) {
	var r resources.RequestBody
	if err := c.ShouldBind(&r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	url, wc := service.Service(r.URL)

	err := database.Insert(url, wc)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "result": fmt.Sprint(url, wc)})
}

func Result(c *gin.Context) {
	res, err := database.Query()
	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("res:%v, err:%v", res, err)})
}
