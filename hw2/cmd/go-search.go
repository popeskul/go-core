package main

import (
	"flag"
	"fmt"
	"go-search/pkg/crawler"
	"go-search/pkg/crawler/spider"
	"go-search/pkg/index/storage"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

const maxDepth = 3

var urls = []string{"https://go.dev", "https://golang.org"}

func main() {
	searchPtr := flag.String("s", "", "Search")
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

	fmt.Println("Results:")

	for _, id := range ids {
		fmt.Println("- ", search(id, pages))
	}
}

func scan(urls []string) []crawler.Document {
	var result []crawler.Document
	randInit := rand.NewSource(time.Now().UnixNano())

	s := spider.New()

	for _, url := range urls {
		pages, err := s.Scan(url, maxDepth)
		if err != nil {
			log.Print(err)
			continue
		}

		for _, page := range pages {
			page.ID = rand.New(randInit).Intn(1000)
			result = append(result, page)
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}

func search(id int, pages []crawler.Document) *crawler.Document {
	index := sort.Search(len(pages), func(index int) bool { return pages[index].ID >= id })

	if index >= len(pages) || pages[index].ID != id {
		return nil
	}

	return &pages[index]
}
