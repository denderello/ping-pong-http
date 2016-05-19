package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type APIRouter interface {
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route
}

type API interface {
	Name() string
	PathPrefix() string
	RegisterHandlers(router APIRouter)
}
