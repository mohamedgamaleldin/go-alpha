package handlers

import (
	"net/http"
)

// GetTokenHandler is a handler function for '/get-token' route
func GetTokenHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			http.Error(w, "No POST", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)

	})
}
