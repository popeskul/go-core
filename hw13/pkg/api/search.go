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

// @Summary Search documents
// @Tags search
// @Description Search documents
// @Accept  json
// @Produce  json
// @Param query path string true "Search query"
// @Success 200 {integer} crawler.Document
// @Failure 400 {string} string "Bad request"
// @Router /search/{query} [get]
func (api *Api) search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["query"]

	store := cache.New()
	store.Add(api.store.GetAll())

	ids := store.Search(query)
	if len(ids) == 0 {
		w.WriteHeader(http.StatusOK)
	}

	doc, err := search(query, api.store.GetAll())
	fmt.Println("searching for", query)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	_ = json.NewEncoder(w).Encode(doc)
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
