package webapp

import (
	"github.com/gorilla/mux"
	"go-search/hw12/pkg/crawler"
	"net/http"
)

type Server struct {
	router *mux.Router
	docs   []crawler.Document
}

func New(router *mux.Router, docs []crawler.Document) Server {
	s := Server{
		docs:   docs,
		router: router,
	}
	s.routes(s.router)
	return s
}

func (s *Server) routes(r *mux.Router) {
	r.HandleFunc("/index", s.IndexHandler).Methods(http.MethodGet)
	r.HandleFunc("/index/{search}", s.SearchIndexHandler).Methods(http.MethodGet)
	r.HandleFunc("/docs", s.DocsHandler).Methods(http.MethodGet)
}
