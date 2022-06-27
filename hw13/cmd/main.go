package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/api"
	"go-search/hw13/pkg/crawler/spider"
	_ "go-search/hw13/pkg/docs"
	"go-search/hw13/pkg/index/cache"
	"go-search/hw13/pkg/searcher"
	"go-search/hw13/pkg/storage/memstore"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// @title           Go Search
// @version         1.0
// @description     This is simple search.
// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	app := searcher.New(
		spider.New(),
		memstore.New(),
		cache.New(),
		[]string{"https://go.dev", "https://golang.org"},
		3,
	)

	log.Println("Start site scanning...")

	docs, err := app.ScanForDocuments()
	if err != nil {
		log.Fatal("Critical error: ", err)
		return
	}

	app.AddDocumentsToStorage(docs)
	app.AddDocumentsToIndex(docs)

	fmt.Println("Site scanning finished")

	r := mux.NewRouter()
	api.New(r, app.Storage())

	go func() {
		log.Fatal(http.ListenAndServe("localhost:8080", r))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("Server shutting down...")
}
