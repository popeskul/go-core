package main

import (
	"flag"
	"fmt"
	"go-search/pkg/crawler"
	"go-search/pkg/crawler/spider"
	"go-search/pkg/index"
	"go-search/pkg/index/cache"
	"go-search/pkg/storage"
	"go-search/pkg/storage/memstore"
	"log"
	"math/rand"
	"time"
)

type searcher struct {
	storage storage.Interface
	index   index.Interface
	scanner crawler.Interface
	sites   []string
	depth   int
}

func main() {
	app := new()
	app.run()

	fmt.Println("Indexes: \n", app.index)

	for {
		fmt.Println("Enter memstore: ")

		var userInput string
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			return
		}

		ids := app.index.Search(userInput)
		fmt.Println("Found ids: ", ids)

		fmt.Println("Results:")
		for _, doc := range app.storage.Search(ids) {
			fmt.Println("- ", doc)
		}
	}
}

func new() *searcher {
	searcher := searcher{}
	searcher.scanner = spider.New()
	searcher.storage = memstore.New()
	searcher.index = cache.New()
	searcher.sites = []string{"https://go.dev", "https://golang.org"}
	searcher.depth = 3

	return &searcher
}

func (s *searcher) run() {
	randInit := rand.NewSource(time.Now().UnixNano())

	searchPtr := flag.String("s", "", "Search")
	flag.Parse()
	if *searchPtr == "" {
		flag.PrintDefaults()
		fmt.Println("Exit")
		return
	}

	fmt.Printf("Request in progress: %s...\n", *searchPtr)

	for _, url := range s.sites {
		docs, errs := s.scanner.Scan(url, s.depth)
		if errs != nil {
			log.Println(errs)
		}

		for _, doc := range docs {
			doc.ID = rand.New(randInit).Intn(1000)
			s.storage.Add([]crawler.Document{doc})
			s.index.Add([]crawler.Document{doc})
		}
	}

	log.Println("Website scanning completed")
}
