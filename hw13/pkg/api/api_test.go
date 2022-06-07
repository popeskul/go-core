package api

import (
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/crawler"
	"os"
	"testing"
)

var (
	api  *Api
	docs []crawler.Document
)

func TestMain(m *testing.M) {
	docs = []crawler.Document{
		{ID: 0, URL: "https://go.dev", Title: "Some title 0", Body: "Some body 0"},
		{ID: 1, URL: "https://go.dev", Title: "Some title 1", Body: "Some body 1"},
	}

	r := mux.NewRouter()
	api = New(r, docs)
	api.store.Add(docs)
	api.routes()
	os.Exit(m.Run())
}
