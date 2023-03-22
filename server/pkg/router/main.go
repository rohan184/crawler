package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rohan184/server/pkg/database"
	"golang.org/x/net/html"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"0.0.0.0"})
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "hello world")
	})

	r.POST("/post", func(c *gin.Context) {
		url, wc := service()

		err := database.Insert(url, wc)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(200, gin.H{"msg": "success"})
	})

	r.GET("/result", func(c *gin.Context) {
		res, err := database.Query()
		c.JSON(200, gin.H{"msg": fmt.Sprintf("res:%v, err:%v", res, err)})
	})
	return r
}

func service() (string, int) {
	url := "https://www.suitejar.com/blog"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var count int
	var visitNode func(*html.Node)
	visitNode = func(n *html.Node) {
		if n.Type == html.TextNode {
			words := strings.Fields(n.Data)
			count += len(words)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visitNode(c)
		}
	}

	visitNode(doc)

	return url, count
}
