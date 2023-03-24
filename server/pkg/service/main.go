package service

import (
	"log"
	"net/http"
	"strings"

	"github.com/rohan184/server/pkg/resources"
	"golang.org/x/net/html"
)

func GetInsight(url string) *resources.Insight {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	imgSrcs := findMediaSrcs(doc)
	if len(imgSrcs) == 0 {
		log.Println("no image found")
	}

	webLinks := findWebLinks(doc)
	if len(webLinks) == 0 {
		log.Println("no web links found")
	}

	count := countWords(doc)

	return &resources.Insight{
		URL:        url,
		WordCount:  count,
		WebLinks:   webLinks,
		MediaLinks: imgSrcs,
	}

}

func findMediaSrcs(n *html.Node) []string {
	var srcs []string

	if n.Type == html.ElementNode && n.Data == "img" {
		for _, attr := range n.Attr {
			if attr.Key == "src" {
				srcs = append(srcs, attr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		srcs = append(srcs, findMediaSrcs(c)...)
	}
	return srcs
}

func findWebLinks(n *html.Node) []string {
	var srcs []string

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				srcs = append(srcs, attr.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		srcs = append(srcs, findWebLinks(c)...)
	}
	return srcs
}

func countWords(n *html.Node) int {
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
	visitNode(n)
	return count
}
