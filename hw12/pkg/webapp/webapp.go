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
	r.HandleFunc("/index", s.getIndex).Methods(http.MethodGet)
	r.HandleFunc("/docs", s.getDocs).Methods(http.MethodGet)
}

func (s *Server) getIndex(w http.ResponseWriter, r *http.Request) {
	store := cache.New()
	store.Add(s.docs)

	vars := mux.Vars(r)
	ids := store.Search(vars["query"])
	if len(ids) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	for _, id := range ids {
		doc := search(id, s.docs)
		fmt.Println(doc)
	}

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

func (s *Server) getDocs(w http.ResponseWriter, r *http.Request) {
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

func search(req int, docs []crawler.Document) (res []string) {
	for _, p := range docs {
		fmt.Println(p, req)
		if strings.Contains(strings.ToLower(p.Title), fmt.Sprint(req)) {
			res = append(res, fmt.Sprintf("Document: '%s' (%s)\n", p.Title, p.URL))
		}
	}

	if len(res) == 0 {
		res = append(res, "Nothing found")
	}
	return res
}
