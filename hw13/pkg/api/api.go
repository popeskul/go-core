package api

import (
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/crawler"
	"go-search/hw13/pkg/storage"
	"go-search/hw13/pkg/storage/memstore"
	"net/http"
	"sync"
)

type Api struct {
	Router *mux.Router
	Docs   []crawler.Document
	store  storage.Interface
	mux    sync.Mutex
}

func New(router *mux.Router, docs []crawler.Document) *Api {
	s := &Api{
		Docs:   docs,
		Router: router,
		store:  memstore.New(),
		mux:    sync.Mutex{},
	}
	s.store.Add(docs)
	s.routes(s.Router)
	return s
}

func (api *Api) routes(r *mux.Router) {
	r.HandleFunc("/api/v1/search/{query}", api.search).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/docs", api.docs).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/docs/{id}", api.findDoc).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/docs/{id}", api.createDoc).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/docs/{id}", api.putDoc).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/docs/{id}", api.patchDoc).Methods(http.MethodPatch)
}
