package handlers

import (
	"net/http"
)

// IndexHandler is a handler function for '/' route
func IndexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("/ called."))

	})
}
