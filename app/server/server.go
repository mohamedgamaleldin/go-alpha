package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Server is an embedded type to enable adding some useful utility functions while still accessing http.Server methods
type Server struct {
	http.Server
}

// InitServer .. Initializes the server
func InitServer(logger *log.Logger, addr string) {
	// logger stuff
	logger.Println("server is starting")

	// router stuff with wrapper handler functions for the middleware
	router := MountRoutes(logger)

	// create the server
	server := &Server{
		Server: http.Server{
			Addr:     addr,
			Handler:  router,
			ErrorLog: logger,
		},
	}

	// start the server
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("server is shutting down")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)

		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("could not gracefully shutdown: %v\n", err)
		}

		close(done)
	}()

	logger.Println("server handling requests on ", server.Addr)
	server.StartServer(logger)

	<-done
	logger.Println("server stopped")
}

// StartServer uses the embedded server type and starts it - ignoring server closed errors as they are handled gracefully
func (server *Server) StartServer(logger *log.Logger) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("could not listen on %s: %v", server.Addr, err)
	}
}
