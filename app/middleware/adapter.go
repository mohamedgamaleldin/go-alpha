package middleware

import (
	"net/http"
)

// Adapter is a type alias for handler wrapper functions
type Adapter func(http.Handler) http.Handler

// Adapt is the middleware chaining implementation
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}

	return h
}
