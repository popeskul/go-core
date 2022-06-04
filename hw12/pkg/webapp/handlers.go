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

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = t.Execute(w, s.docs)
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

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res, err := search(searchKeyword, s.docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, res)
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
		return nil, fmt.Errorf("no results")
	}

	return res, nil
}
