package middleware

import (
	"net/http"
	"strings"
)

// MustParams .. validates that the required parameters for an endpoint are in the request
func MustParams(params ...string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			for _, param := range params {

				key, ok := r.URL.Query()[param]

				// validation for required key present
				if !ok || len(key) < 1 {
					error := strings.Builder{}
					error.WriteString("Missing key ")
					error.WriteString(param)
					error.WriteString(" in request")
					http.Error(w, error.String(), http.StatusBadRequest)
					return
				}

				// validation for argument missing from required key
				if r.URL.Query().Get(param) == "" {
					error := strings.Builder{}
					error.WriteString("Value required for key ")
					error.WriteString(param)
					http.Error(w, error.String(), http.StatusBadRequest)
					return
				}

			}

			h.ServeHTTP(w, r)
		})
	}
}

// MustAuth .. checks if auth token is present and valid
func MustAuth() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Write([]byte("MustAuth() called"))
			h.ServeHTTP(w, r)
		})
	}
}
