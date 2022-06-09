package api

import (
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/crawler"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApi_search(t *testing.T) {
	type want struct {
		status int
		body   string
	}
	type args struct {
		query  string
		fields []crawler.Document
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "search with results",
			args: args{
				query: "test",
				fields: []crawler.Document{
					{Title: "test"},
					{Title: "test2"},
					{Title: "another"},
				},
			},
			want: want{
				status: http.StatusOK,
				body:   `[{"ID":0,"URL":"","Title":"test","Body":""},{"ID":0,"URL":"","Title":"test2","Body":""}]`,
			},
		},
		{
			name: "search without results",
			args: args{
				query: "bad",
				fields: []crawler.Document{
					{Title: "test"},
					{Title: "test2"},
				},
			},
			want: want{
				status: http.StatusNotFound,
				body:   "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/v1/search", nil)
			req.Header.Set("content-type", "application/json")
			req = mux.SetURLVars(req, map[string]string{"query": tt.args.query})

			rr := httptest.NewRecorder()
			api := &Api{
				Docs: tt.args.fields,
			}
			api.search(rr, req)

			if rr.Code != tt.want.status {
				t.Errorf("Expected %d, got %d", tt.want.status, rr.Code)
			}

			if !strings.Contains(rr.Body.String(), tt.want.body) {
				t.Errorf("Expected body to contain %s", tt.want.body)
			}
		})
	}
}
