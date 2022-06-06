package api

import (
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/crawler"
	"os"
	"reflect"
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
	api.routes(r)
	os.Exit(m.Run())
}

func TestApiNew(t *testing.T) {
	if reflect.TypeOf(api) != reflect.TypeOf(&Api{}) {
		t.Error("New() should return API type")
	}

	if api.Router == nil {
		t.Error("New() should return not nil router")
	}

	if api.Docs == nil {
		t.Error("New() should return not nil docs")
	}
}
