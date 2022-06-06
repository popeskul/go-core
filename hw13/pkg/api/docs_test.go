package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/crawler"
	"go-search/hw13/pkg/storage/memstore"
	http "net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "go-search/hw13/pkg/testing_init"
)

func TestApi_docs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/docs", nil)
	req.Header.Set("content-type", "text/html")

	rr := httptest.NewRecorder()
	api.docs(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, rr.Code)
	}

	if !strings.Contains(rr.Body.String(), docs[0].Title) {
		t.Errorf("Expected body to contain " + docs[0].Title)
	}
}

func TestApi_createDoc(t *testing.T) {
	tests := []struct {
		name   string
		doc    crawler.Document
		status int
	}{
		{
			name: "valid",
			doc: crawler.Document{
				Title: "test",
			},
			status: http.StatusOK,
		},
		{
			name: "invalid",
			doc: crawler.Document{
				Title: "",
			},
			status: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/docs", strings.NewReader(`{"title": "`+tt.doc.Title+`", "url": "`+tt.doc.URL+`"}`))
			req.Header.Set("content-type", "application/json")

			rr := httptest.NewRecorder()
			api.createDoc(rr, req)

			if rr.Code != tt.status {
				t.Errorf("Expected %d, got %d", tt.status, rr.Code)
			}

			if !strings.Contains(rr.Body.String(), tt.doc.Title) {
				t.Errorf("Expected body to contain " + tt.doc.Title)
			}
		})
	}
}

func TestApi_put(t *testing.T) {
	tests := []struct {
		name   string
		doc    crawler.Document
		status int
	}{
		{
			name: "valid",
			doc: crawler.Document{
				ID:    1,
				Title: "test1",
			},
			status: http.StatusOK,
		},
		{
			name: "invalid",
			doc: crawler.Document{
				ID: 10000,
			},
			status: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := strings.NewReader(`{"title": "` + tt.doc.Title + `", "url": "` + tt.doc.URL + `"}`)
			path := fmt.Sprintf("/api/v1/docs/%d", tt.doc.ID)
			req := httptest.NewRequest(http.MethodPut, path, body)
			req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprintf("%d", tt.doc.ID)})
			req.Header.Set("content-type", "application/json")

			rr := httptest.NewRecorder()
			api.putDoc(rr, req)

			if rr.Code != tt.status {
				t.Errorf("Expected %d, got %d", tt.status, rr.Code)
			}
		})
	}
}

func TestApi_patch(t *testing.T) {
	resetDB()

	tests := []struct {
		name   string
		doc    crawler.Document
		status int
		want   crawler.Document
	}{
		{
			name: "changed title",
			doc: crawler.Document{
				ID:    0,
				Title: "test00",
			},
			status: http.StatusOK,
			want: crawler.Document{
				ID:    0,
				Title: "test00",
			},
		},
		{
			name: "doesn't change anything",
			doc: crawler.Document{
				ID:    1,
				Title: "",
			},
			status: http.StatusOK,
			want:   api.Docs[1],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := strings.NewReader(`{"title": "` + tt.doc.Title + `"}`)
			path := fmt.Sprintf("/api/v1/docs/%d", tt.doc.ID)
			req := httptest.NewRequest(http.MethodPatch, path, body)
			req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprintf("%d", tt.doc.ID)})
			req.Header.Set("content-type", "application/json")

			rr := httptest.NewRecorder()
			api.patchDoc(rr, req)

			if rr.Code != tt.status {
				t.Errorf("Expected %d, got %d", tt.status, rr.Code)
			}

			var got crawler.Document
			err := json.Unmarshal(rr.Body.Bytes(), &got)
			if err != nil {
				t.Errorf("Unmarshal: %v", err)
			}

			if got.Title != tt.want.Title {
				t.Errorf("Want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestApi_deleteDoc(t *testing.T) {
	tests := []struct {
		name   string
		doc    crawler.Document
		status int
	}{
		{
			name: "valid",
			doc: crawler.Document{
				ID:    0,
				Title: "test0",
			},
			status: http.StatusOK,
		},
		{
			name: "invalid",
			doc: crawler.Document{
				ID: 10000,
			},
			status: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/api/v1/docs/%d", tt.doc.ID)
			req := httptest.NewRequest(http.MethodDelete, path, nil)
			req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprintf("%d", tt.doc.ID)})

			rr := httptest.NewRecorder()
			api.deleteDoc(rr, req)

			if rr.Code != tt.status {
				t.Errorf("Expected %d, got %d", tt.status, rr.Code)
			}
		})
	}
}

func resetDB() {
	api.Docs = docs
	api.store = memstore.New()
	api.store.Add(docs)
}
