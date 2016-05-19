package v1

import (
	"fmt"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	time.Sleep(10 * time.Second)
	fmt.Fprintf(w, "Hello world!\n")
}
