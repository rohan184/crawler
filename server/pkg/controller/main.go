package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rohan184/server/pkg/database"
	"github.com/rohan184/server/pkg/resources"
	"github.com/rohan184/server/pkg/service"
)

func Insight(c *gin.Context) {
	var r resources.RequestBody
	if err := c.ShouldBind(&r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	resp := service.GetInsight(r.URL)

	err := database.Insert(resp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "result": resp})
}

func GetInsights(c *gin.Context) {
	res, err := database.Query()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": res})
}

func RemoveInsight(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	database.DeleteInsight(id)
}

func FavoriteInsight(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	err := database.MarkInsightFav(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
}
