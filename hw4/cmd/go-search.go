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
	"math/rand"
	"os"
	"time"
)

type searcher struct {
	storage storage.Interface
	index   index.Interface
	scanner crawler.Interface
	sites   []string
	depth   int
}

const fileName = "storage.json"

func main() {
	app := New()

	// flags
	searchPtr := flag.String("s", "", "Search")
	flag.Parse()
	if *searchPtr == "" {
		flag.PrintDefaults()
		fmt.Println("Exit")
		return
	}

	fmt.Printf("Request in progress: %s...\n", *searchPtr)

	// scanning
	docs, err := app.scan(app.sites, app.depth)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	for _, doc := range docs {
		app.storage.Add([]crawler.Document{doc})
		app.index.Add([]crawler.Document{doc})
	}

	fmt.Println("Indexes: \n", app.index)

	app.userInput()
}

func New() *searcher {
	searcher := searcher{}
	searcher.scanner = spider.New()
	searcher.storage = memstore.New()
	searcher.index = cache.New()
	searcher.sites = []string{"https://go.dev", "https://golang.org"}
	searcher.depth = 3

	return &searcher
}

func (s *searcher) userInput() {
	for {
		fmt.Println("Enter store index: ")

		var userInput string
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			return
		}

		ids := s.index.Search(userInput)
		fmt.Println("Found ids: ", ids)

		fmt.Println("Results:")
		for _, doc := range s.storage.Search(ids) {
			fmt.Println("- ", doc)
		}
	}
}

func (s *searcher) scan(urls []string, depth int) ([]crawler.Document, error) {
	if isEmptyFile(fileName) {
		docs := s.scanUrls(urls, depth)

		f, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		_, err = s.storage.Write(f, docs)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		return docs, nil
	}

	var docs []crawler.Document

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	docs, err = s.storage.Read(f)
	if err != nil {
		return nil, err
	}

	return docs, nil
}

func (s *searcher) scanUrls(urls []string, depth int) []crawler.Document {
	randInit := rand.NewSource(time.Now().UnixNano())
	var allDocs []crawler.Document

	for _, url := range urls {
		docs, errs := s.scanner.Scan(url, depth)
		if errs != nil {
			fmt.Println("Error scan: ", errs)
		}

		for i := range docs {
			docs[i].ID = rand.New(randInit).Intn(1000)
		}

		allDocs = append(docs, docs...)
	}

	return allDocs
}

func isEmptyFile(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return true
	}

	f, err := os.Stat(fileName)
	if err != nil {
		return true
	}

	if f.Size() == 0 {
		return true
	}

	return false
}
