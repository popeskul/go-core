package webapp

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-search/hw12/pkg/crawler"
	"go-search/hw12/pkg/index/cache"
	"html/template"
	"net/http"
	"strings"
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

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := template.Must(template.ParseFiles("templates/index.html")).Execute(w, s.docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) SearchIndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchKeyword := vars["search"]

	store := cache.New()
	store.Add(s.docs)

	ids := store.Search(searchKeyword)
	if len(ids) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	res, err := search(searchKeyword, s.docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = template.Must(template.ParseFiles("templates/index.html")).Execute(w, res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) DocsHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/docs.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = t.Execute(w, s.docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func search(req string, docs []crawler.Document) (res []crawler.Document, err error) {
	for _, p := range docs {
		if strings.Contains(strings.ToLower(p.Title), req) {
			res = append(res, p)
		}
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("<!DOCTYPE html>\n<html>\n<head>\n<title>404</title>\n</head>\n<body>\n<h1>404</h1>\n<p>Not found</p>\n</body>\n</html>>")
	}

	return res, nil
}
