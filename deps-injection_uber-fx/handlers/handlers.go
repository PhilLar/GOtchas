package handlers

import (
	"github.com/PhilLar/GOtchas/deps-injection_uber-fx/logger"
	"io"
	"net/http"
)

func NewHandler(l logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.Println("Handler called")
		io.WriteString(w, "Hello, LOGGER")
	})
}