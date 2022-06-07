package webapp

import (
	"github.com/gorilla/mux"
	"go-search/hw12/pkg/crawler"
	_ "go-search/hw12/pkg/testing_init"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	webapp Server
	docs   []crawler.Document
)

func TestMain(m *testing.M) {
	docs = []crawler.Document{
		{ID: 0, URL: "https://go.dev", Title: "The Go Programming Language", Body: "The Go Programming Language"},
		{ID: 1, URL: "https://go.dev", Title: "Some title", Body: "Some body"},
	}

	r := mux.NewRouter()
	webapp = New(r, docs)
	webapp.routes()
	m.Run()
}

func TestIndexHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/index", nil)
	req.Header.Set("content-type", "text/html")

	rr := httptest.NewRecorder()

	webapp.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, rr.Code)
	}

	if !strings.Contains(rr.Body.String(), docs[0].Title) {
		t.Errorf("Expected body to contain 'The Go Programming Language'")
	}
}

func TestSearchIndexHandler(t *testing.T) {
	tests := []struct {
		name     string
		req      string
		docs     []crawler.Document
		expected string
		status   int
	}{
		{
			name:     "title",
			req:      "go",
			docs:     docs,
			expected: "The Go Programming Language",
			status:   http.StatusOK,
		},
		{
			name:     "not found",
			req:      "oops",
			docs:     docs,
			expected: "Not found",
			status:   http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/index/"+tt.req, nil)
			req.Header.Set("content-type", "text/html")

			rr := httptest.NewRecorder()

			webapp.router.ServeHTTP(rr, req)

			if rr.Code != tt.status {
				t.Errorf("Expected %d, got %d", http.StatusOK, rr.Code)
			}

			if !strings.Contains(rr.Body.String(), tt.expected) {
				t.Errorf("Expected body to contain 'The Go Programming Language'")
			}
		})
	}
}

func TestDocsHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)
	req.Header.Set("content-type", "text/html")

	rr := httptest.NewRecorder()

	webapp.router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, rr.Code)
	}

	if !strings.Contains(rr.Body.String(), docs[0].Title) {
		t.Errorf("Expected body to contain 'The Go Programming Language'")
	}
}
