package webapp

import (
	"github.com/gorilla/mux"
	"go-search/hw12/pkg/crawler"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		docs   []crawler.Document
		router *mux.Router
	}{
		{
			docs:   []crawler.Document{},
			router: mux.NewRouter(),
		},
		{
			docs:   []crawler.Document{crawler.Document{}},
			router: mux.NewRouter(),
		},
	}

	for _, tt := range tests {
		s := New(tt.router, tt.docs)
		if reflect.DeepEqual(s.docs, tt.docs) == false {
			t.Errorf("New(%v, %v) got %v, expected %v", tt.docs, tt.router, s.docs, tt.docs)
		}
		if s.router != tt.router {
			t.Errorf("New(%v, %v) got %v, expected %v", tt.docs, tt.router, s.router, tt.router)
		}
	}
}

func TestIndexHandler(t *testing.T) {
	// TODO: implement request handler with recorder
	var tests = []struct {
		docs   []crawler.Document
		router *mux.Router
	}{
		{
			docs: []crawler.Document{
				{
					ID:    0,
					Title: "Title 0",
					Body:  "Body 0",
					URL:   "https://example.com/0",
				},
				{
					ID:    1,
					Title: "Title 1",
					Body:  "Body 1",
					URL:   "https://example.com/1",
				},
			},
			router: mux.NewRouter(),
		},
		{
			docs:   []crawler.Document{crawler.Document{}},
			router: mux.NewRouter(),
		},
	}

	for _, tt := range tests {
		s := New(tt.router, tt.docs)
		t.Log(s.docs, tt.docs)
		if reflect.DeepEqual(s.docs, tt.docs) == false {
			t.Errorf("New(%v, %v) got %v, expected %v", tt.docs, tt.router, s.docs, tt.docs)
		}
		if s.router != tt.router {
			t.Errorf("New(%v, %v) got %v, expected %v", tt.docs, tt.router, s.router, tt.router)
		}
	}
}

func TestSearchIndexHandler(t *testing.T) {
	t.Errorf("Not implemented")
}

func TestDocsHandler(t *testing.T) {
	t.Errorf("Not implemented")
}
