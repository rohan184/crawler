package service

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func Service(url string) (string, int) {
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
