package server

import "github.com/gorilla/mux"

type API interface {
	Name() string
	PathPrefix() string
	RegisterHandlers(router *mux.Router)
}
