package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"go-search/hw12/pkg/crawler"
	"go-search/hw12/pkg/crawler/spider"
	"go-search/hw12/pkg/index"
	"go-search/hw12/pkg/index/cache"
	"go-search/hw12/pkg/storage"
	"go-search/hw12/pkg/storage/memstore"
	"go-search/hw12/pkg/webapp"
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
	log.Println("Start site scanning...")

	docs, err := app.scan(app.sites, app.depth)
	if err != nil {
		log.Fatal("Critical error: ", err)
		return
	}

	for _, doc := range docs {
		app.storage.Add([]crawler.Document{doc})
		app.index.Add([]crawler.Document{doc})
	}
	fmt.Println("Site scanning finished")

	r := mux.NewRouter()
	webapp.New(r, docs)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
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

func (s *searcher) scan(urls []string, depth int) ([]crawler.Document, error) {
	docs, err := s.read(fileName)
	if err == nil {
		return docs, nil
	}

	docs = s.scanUrls(urls, depth)
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = s.storage.Write(f, docs)
	if err != nil {
		return nil, err
	}

	return docs, nil
}

func (s *searcher) read(fileName string) ([]crawler.Document, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	docs, err := s.storage.Read(f)
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
