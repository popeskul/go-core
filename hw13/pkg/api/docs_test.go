package api

import (
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
	req.Header.Set("content-type", "application/json")

	rr := httptest.NewRecorder()
	api.docs(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, rr.Code)
	}

	if !strings.Contains(rr.Body.String(), docs[0].Title) {
		t.Errorf("Expected body to contain " + docs[0].Title)
	}
}

func TestApi_find(t *testing.T) {
	tests := []struct {
		name   string
		query  string
		status int
	}{
		{
			name:   "valid",
			query:  "0",
			status: http.StatusOK,
		},
		{
			name:   "invalid",
			query:  "",
			status: http.StatusBadRequest,
		},
		{
			name:   "invalid",
			query:  "123123",
			status: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/docs/find/%s", tt.query), nil)
			req.Header.Set("content-type", "application/json")
			req = mux.SetURLVars(req, map[string]string{"id": tt.query})

			rr := httptest.NewRecorder()
			api.findDoc(rr, req)

			if rr.Code != tt.status {
				t.Errorf("Expected %d, got %d", tt.status, rr.Code)
			}

			if !strings.Contains(rr.Body.String(), tt.query) {
				t.Errorf("Expected body to contain " + tt.query)
			}
		})
	}
}

func TestApi_createDoc(t *testing.T) {
	docGood := `{"URL":"","Title":"test","Body":""}`
	docBad := `{"URL":"","Title":"","Body":""}`

	type args struct {
		docJson string
	}

	type want struct {
		status int
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "valid",
			args: args{
				docJson: docGood,
			},
			want: want{
				status: http.StatusCreated,
			},
		},
		{
			name: "invalid",
			args: args{
				docJson: docBad,
			},
			want: want{
				status: http.StatusBadRequest,
			},
		},
		{
			name: "vary bad",
			args: args{
				docJson: `asdfdsfkal;k`,
			},
			want: want{
				status: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/docs", strings.NewReader(tt.args.docJson))
			req.Header.Set("content-type", "application/json")

			rr := httptest.NewRecorder()
			api.createDoc(rr, req)

			if rr.Code != tt.want.status {
				t.Errorf("Expected %d, got %d", tt.want.status, rr.Code)
			}
		})
	}
}

func TestApi_put(t *testing.T) {
	originalDoc := crawler.Document{
		URL:   "https://www.google.com",
		Title: "Google",
	}

	type args struct {
		docJson string
		id      string
	}

	type want struct {
		status int
		doc    crawler.Document
	}

	docGood := `{"URL":"","Title":"test","Body":""}`
	docBad := `{"URL":"","Title":"","Body":""}`

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "found",
			args: args{
				docJson: docGood,
				id:      "0",
			},
			want: want{
				status: http.StatusOK,
				doc:    originalDoc,
			},
		},
		{
			name: "not found",
			args: args{
				docJson: docBad,
				id:      "10000",
			},
			want: want{
				status: http.StatusNotFound,
			},
		},
		{
			name: "very bad json",
			args: args{
				docJson: `asd`,
				id:      "10000",
			},
			want: want{
				status: http.StatusBadRequest,
			},
		},
		{
			name: "bad id",
			args: args{
				docJson: docGood,
				id:      "asdasd",
			},
			want: want{
				status: http.StatusBadRequest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := strings.NewReader(tt.args.docJson)
			path := fmt.Sprintf("/api/v1/docs/%s", tt.args.id)
			req := httptest.NewRequest(http.MethodPut, path, body)
			req = mux.SetURLVars(req, map[string]string{"id": tt.args.id})
			req.Header.Set("content-type", "application/json")

			rr := httptest.NewRecorder()
			api.putDoc(rr, req)

			checkStatusCode := rr.Code != tt.want.status

			if checkStatusCode {
				t.Errorf("Expected %d, got %d", tt.want.status, rr.Code)
			}

			if checkStatusCode && tt.want.doc.URL != originalDoc.URL {
				t.Errorf("Expected URL to be %s, got %s", originalDoc.URL, tt.want.doc.URL)
			}
		})
	}
}

func TestApi_patch(t *testing.T) {
	resetDB()

	type args struct {
		id      string
		docJson string
	}

	type want struct {
		status  int
		docJson string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "changed title",
			args: args{
				id:      "0",
				docJson: `{"Title":"test"}`,
			},
			want: want{
				docJson: `{"Title":"test"}`,
				status:  http.StatusOK,
			},
		},
		{
			name: "doesn't change anything",
			args: args{
				id:      "0",
				docJson: `{"Title":""}`,
			},
			want: want{
				docJson: `{"Title":""}`,
				status:  http.StatusOK,
			},
		},
		{
			name: "bad id",
			args: args{
				id:      "asdasd",
				docJson: `{"Title":""}`,
			},
			want: want{
				docJson: `{"Title":""}`,
				status:  http.StatusBadRequest,
			},
		},
		{
			name: "not found by id",
			args: args{
				id:      "123",
				docJson: `{"Title":""}`,
			},
			want: want{
				status: http.StatusNotFound,
			},
		},
		{
			name: "bod request body",
			args: args{
				id:      "123",
				docJson: `{""}`,
			},
			want: want{
				status: http.StatusBadRequest,
			},
		},
	}

	for _, tt := range tests {
		docJson := `{"URL":"https://www.google.com","Title":"Google","Body":""}`

		t.Run(tt.name, func(t *testing.T) {
			body := strings.NewReader(tt.args.docJson)
			path := fmt.Sprintf("/api/v1/docs/%s", tt.args.id)
			req := httptest.NewRequest(http.MethodPatch, path, body)
			req = mux.SetURLVars(req, map[string]string{"id": tt.args.id})
			req.Header.Set("content-type", "application/json")

			rr := httptest.NewRecorder()
			api.patchDoc(rr, req)

			checkStatusCode := rr.Code != tt.want.status

			if checkStatusCode {
				t.Errorf("Expected %d, got %d", tt.want.status, rr.Code)
			}

			if checkStatusCode && tt.want.docJson != docJson {
				t.Errorf("Expected %s, got %s", docJson, tt.want.docJson)
			}
		})
	}
}

func TestApi_deleteDoc(t *testing.T) {
	type args struct {
		id string
	}

	type want struct {
		status int
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "delete by id",
			args: args{
				id: "0",
			},
			want: want{
				status: http.StatusOK,
			},
		},
		{
			name: "not found by id",
			args: args{
				id: "123",
			},
			want: want{
				status: http.StatusNotFound,
			},
		},
		{
			name: "bad id",
			args: args{
				id: "aaa",
			},
			want: want{
				status: http.StatusBadRequest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := fmt.Sprintf("/api/v1/docs/%s", tt.args.id)
			req := httptest.NewRequest(http.MethodDelete, path, nil)
			req = mux.SetURLVars(req, map[string]string{"id": tt.args.id})

			rr := httptest.NewRecorder()
			api.deleteDoc(rr, req)

			if rr.Code != tt.want.status {
				t.Errorf("Expected %d, got %d", tt.want.status, rr.Code)
			}
		})
	}
}

func resetDB() {
	api.Docs = docs
	api.store = memstore.New()
	api.store.Add(docs)
}
