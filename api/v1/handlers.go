package v1

import "net/http"

type ResponseMessage struct {
	Message string `json:"message"`
}

func PingPongHandler(w http.ResponseWriter, req *http.Request) {
	writeJSON(ResponseMessage{"pong"}, w)
}
