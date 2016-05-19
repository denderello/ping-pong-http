package v1

import "github.com/gorilla/mux"

type V1API struct{}

func (v1 V1API) Name() string {
	return "V1"
}

func (v1 V1API) PathPrefix() string {
	return "/v1"
}

func (v1 V1API) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/ping", PingPongHandler)
}
