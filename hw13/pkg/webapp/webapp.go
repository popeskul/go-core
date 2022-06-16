package webapp

import (
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/api"
	"go-search/hw13/pkg/storage"
)

type Server struct {
	api *api.Api
}

func New(router *mux.Router, docs storage.Interface) *Server {
	s := &Server{
		api: api.New(router, docs),
	}

	return s
}
