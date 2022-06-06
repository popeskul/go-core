package webapp

import (
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/api"
	"go-search/hw13/pkg/crawler"
)

type Server struct {
	api *api.Api
}

func New(router *mux.Router, docs []crawler.Document) *Server {
	s := &Server{
		api: api.New(router, docs),
	}

	return s
}
