package webapp

import (
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/crawler"
	"go-search/hw13/pkg/storage/memstore"
	_ "go-search/hw13/pkg/testing_init"
	"os"
	"testing"
)

var (
	webapp *Server
	docs   []crawler.Document
)

func TestMain(m *testing.M) {
	store := memstore.New()
	docs = []crawler.Document{
		{ID: 0, URL: "https://go.dev", Title: "The Go Programming Language", Body: "The Go Programming Language"},
		{ID: 1, URL: "https://go.dev", Title: "Some title", Body: "Some body"},
	}
	store.Add(docs)

	router := mux.NewRouter()
	webapp = New(router, store)
	os.Exit(m.Run())
}
