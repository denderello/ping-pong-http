package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func writeJSON(data interface{}, writer http.ResponseWriter) {
	body, err := json.Marshal(data)
	if err != nil {
		jsonError(err, writer)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(body)
}

type apiError struct {
	Message string `json:"error"`
	Code    string `json:"code"`
}

func jsonError(err error, writer http.ResponseWriter) {
	errorCode := 500
	errorMessage := fmt.Sprintf("Unknown server error. %s", err)

	log.Printf("An error occurred: %+v", err)

	switch err := err.(type) {
	case HTTPError:
		errorCode = err.StatusCode()
		errorMessage = err.Error()
	}

	body, err := json.Marshal(apiError{
		Message: errorMessage,
		Code:    strconv.Itoa(errorCode),
	})
	if err != nil {
		panic(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(errorCode)
	writer.Write(body)
}
