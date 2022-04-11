package main

import (
	"flag"
	"fmt"
	"go-search/pkg/crawler"
	"go-search/pkg/crawler/spider"
	"log"
	"os"
	"strings"
)

const maxDepth = 3

var urls = []string{"https://go.dev"}

func main() {
	searchPtr := flag.String("s", "", "Search")
	flag.Parse()

	if *searchPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("Request in progress: %s...\n", *searchPtr)

	pages := scan(urls)
	for _, p := range pages {
		if findByAndWith(p.Title, *searchPtr) {
			printText(p)
		}
	}
}

func scan(urls []string) []crawler.Document {
	var result []crawler.Document
	s := spider.New()

	for _, url := range urls {
		pages, err := s.Scan(url, maxDepth)
		if err != nil {
			log.Print(err)
		}

		result = append(result, pages...)
	}

	return result
}

func findByAndWith(text, search string) bool {
	return strings.Contains(strings.ToLower(text), strings.ToLower(search))
}

func printText(p crawler.Document) {
	fmt.Println("----")
	fmt.Println("Title ", p.Title)
	fmt.Println("URL ", p.URL)
	fmt.Println("----")
}
