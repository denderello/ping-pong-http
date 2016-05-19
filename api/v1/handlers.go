package v1

import "net/http"

type ResponseMessage struct {
	Message string `json:"message"`
}

func (v1 *V1API) PingPongHandler(w http.ResponseWriter, req *http.Request) {
	v1.Logger.Info("Received ping request. Sending pong message.")
	writeJSON(ResponseMessage{"pong"}, w)
}
