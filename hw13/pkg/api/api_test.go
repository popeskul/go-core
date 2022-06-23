package api

import (
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/crawler"
	"go-search/hw13/pkg/storage/memstore"
	"testing"
)

var (
	api  *Api
	docs []crawler.Document
)

func TestMain(m *testing.M) {
	store := memstore.New()
	docs = []crawler.Document{
		{ID: 0, URL: "https://go.dev", Title: "Some title 0", Body: "Some body 0"},
		{ID: 1, URL: "https://go.dev", Title: "Some title 1", Body: "Some body 1"},
	}
	store.Add(docs)

	r := mux.NewRouter()
	api = New(r, store)
	api.store.Add(docs)
	api.routes()
	m.Run()
}
