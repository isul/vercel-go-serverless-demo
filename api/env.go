package handler

import (
	"os"
	"net/http"
	"net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("REQUEST_TIMEOUT: " + os.Getenv("REQUEST_TIMEOUT")))
}
