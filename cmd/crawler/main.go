package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Parse the url flag
	urlPtr := flag.String("url", "no default", "url to crawl")
	flag.Parse()

	if *urlPtr == "no default" {
		fmt.Println("Please provide a url to crawl")
		return
	}

	c := colly.NewCollector()

	// Open urls.txt for writing
	file, err := os.Create("urls.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer file.Close()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// If the url starts with http, then add it to urls.txt
		if strings.HasPrefix(e.Attr("href"), "http") {
			file.WriteString(e.Attr("href") + "\n")
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(*urlPtr)
}
