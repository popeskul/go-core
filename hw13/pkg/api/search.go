package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/crawler"
	"go-search/hw13/pkg/index/cache"
	"net/http"
	"strings"
)

func (api *Api) search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["query"]

	store := cache.New()
	store.Add(api.Docs)

	ids := store.Search(query)
	if len(ids) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	doc, err := search(query, api.Docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(doc)
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
		return nil, fmt.Errorf("no results found")
	}

	return res, nil
}
