package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger is a middleware handler that does request logging and tracing
type Logger struct {
	handler http.Handler
	logger  *log.Logger
}

// ModifiedResponseWriter embeds ResponseWriter and adds a field to capture the status code so we can read it after the request
type ModifiedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewModifiedResponseWriter .. constructor
func NewModifiedResponseWriter(w http.ResponseWriter) *ModifiedResponseWriter {
	return &ModifiedResponseWriter{w, http.StatusOK}
}

// WriteHeader .. added capturing the status code
func (mrw *ModifiedResponseWriter) WriteHeader(statusCode int) {
	mrw.statusCode = statusCode
	mrw.ResponseWriter.WriteHeader(statusCode)
}

// ServeHTTP wrapper implementing the wrapper Logger struct
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.logger.Printf("%s | %s | %s\n", r.RemoteAddr, r.RequestURI, r.Method)

	mrw := NewModifiedResponseWriter(w)
	l.handler.ServeHTTP(mrw, r)
	l.logger.Printf("%v | %v | %s | %s | %s\n", time.Since(start), mrw.statusCode, r.RequestURI, r.RemoteAddr, r.Method)
}

// NewLogger constructs a new middleware Logger handler
func NewLogger(h http.Handler, logger *log.Logger) *Logger {
	return &Logger{h, logger}
}
