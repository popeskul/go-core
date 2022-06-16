package api

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "go-search/hw13/pkg/docs"
	"go-search/hw13/pkg/storage"
	"net/http"
)

type Api struct {
	r     *mux.Router
	store storage.Interface
}

func New(router *mux.Router, storage storage.Interface) *Api {
	s := &Api{
		r:     router,
		store: storage,
	}
	s.routes()
	return s
}

func (api *Api) routes() {
	api.r.Use(requestIDMiddleware)
	api.r.Use(logMiddleware)

	api.r.HandleFunc("/api/v1/search/{query}", api.search).Methods(http.MethodGet)
	api.r.HandleFunc("/api/v1/docs", api.docs).Methods(http.MethodGet)
	api.r.HandleFunc("/api/v1/docs/{id}", api.findDoc).Methods(http.MethodGet)
	api.r.HandleFunc("/api/v1/docs", api.createDoc).Methods(http.MethodPost)
	api.r.HandleFunc("/api/v1/docs/{id}", api.putDoc).Methods(http.MethodPut)
	api.r.HandleFunc("/api/v1/docs/{id}", api.patchDoc).Methods(http.MethodPatch)

	api.r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}
