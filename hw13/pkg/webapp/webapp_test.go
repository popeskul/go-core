package webapp

import (
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/crawler"
	_ "go-search/hw13/pkg/testing_init"
	"os"
	"reflect"
	"testing"
)

var (
	webapp *Server
	docs   []crawler.Document
)

func TestMain(m *testing.M) {
	docs = []crawler.Document{
		{ID: 0, URL: "https://go.dev", Title: "The Go Programming Language", Body: "The Go Programming Language"},
		{ID: 1, URL: "https://go.dev", Title: "Some title", Body: "Some body"},
	}

	router := mux.NewRouter()
	webapp = New(router, docs)
	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	if reflect.TypeOf(webapp) != reflect.TypeOf(&Server{}) {
		t.Error("New() should return Server type")
	}

	if webapp.api.Router == nil {
		t.Error("New() should return not nil router")
	}

	if webapp.api.Docs == nil {
		t.Error("New() should return not nil docs")
	}
}
