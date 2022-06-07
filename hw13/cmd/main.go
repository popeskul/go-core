package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/crawler"
	"go-search/hw13/pkg/crawler/spider"
	_ "go-search/hw13/pkg/docs"
	"go-search/hw13/pkg/index"
	"go-search/hw13/pkg/index/cache"
	"go-search/hw13/pkg/storage"
	"go-search/hw13/pkg/storage/memstore"
	"go-search/hw13/pkg/webapp"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

// @title           Go Search
// @version         1.0
// @description     This is simple search.
// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	app := New()
	log.Println("Start site scanning...")

	docs, err := app.scan()
	if err != nil {
		log.Fatal("Critical error: ", err)
		return
	}

	app.storage.Add(docs)
	app.index.Add(docs)

	fmt.Println("Site scanning finished")

	r := mux.NewRouter()
	webapp.New(r, docs)

	go func() {
		log.Fatal(http.ListenAndServe("localhost:8080", r))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("Server shutting down...")
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

func (s *searcher) scan() ([]crawler.Document, error) {
	docs, err := s.read(fileName)
	if err == nil {
		return docs, nil
	}

	docs = s.scanUrls()
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

func (s *searcher) scanUrls() []crawler.Document {
	randInit := rand.NewSource(time.Now().UnixNano())
	var allDocs []crawler.Document

	for _, url := range s.sites {
		docs, errs := s.scanner.Scan(url, s.depth)
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
