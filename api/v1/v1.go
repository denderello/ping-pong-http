package v1

import (
	"github.com/denderello/ping-pong-http/log"
	"github.com/denderello/ping-pong-http/server"
)

type V1API struct {
	Logger log.Logger
}

func (v1 V1API) Name() string {
	return "V1"
}

func (v1 V1API) PathPrefix() string {
	return "/v1"
}

func (v1 V1API) RegisterHandlers(router server.APIRouter) {
	router.HandleFunc("/ping", v1.PingPongHandler)
}
