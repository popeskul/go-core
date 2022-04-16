package main

import (
	"flag"
	"fmt"
	"go-search/pkg/crawler"
	"go-search/pkg/crawler/spider"
	"go-search/pkg/index/storage"
	"log"
	"os"
)

const maxDepth = 3

var urls = []string{"https://go.dev"}

func main() {
	searchPtr := flag.String("s", "", "FindIndexIds")
	flag.Parse()

	if *searchPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("Request in progress: %s...\n", *searchPtr)

	store := storage.New()
	pages := scan(urls)
	store.Add(pages)

	fmt.Println("Store: ", store)
	fmt.Println("Enter index: ")

	var userInput string
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		return
	}

	ids := store.FindIndexIds(userInput)
	results := search(ids, pages)
	fmt.Println("Results:", results)
}

func scan(urls []string) []crawler.Document {
	var result []crawler.Document
	pageId := 500 // for sorting example we can use pageId-- or rand.Intn

	s := spider.New()

	for _, url := range urls {
		pages, err := s.Scan(url, maxDepth)
		if err != nil {
			log.Print(err)
			continue
		}

		for _, page := range pages {
			page.ID = pageId
			pageId--
			result = append(result, page)
		}
	}

	return result
}

func search(ids []int, pages []crawler.Document) []crawler.Document {
	var result []crawler.Document

	for _, id := range ids {
		for _, page := range pages {
			if page.ID == id {
				result = append(result, page)
			}
		}
	}

	return result
}
