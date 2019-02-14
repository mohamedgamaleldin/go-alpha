package server

import (
	"log"
	"net/http"

	"github.com/mohamedgamaleldin/go-alpha/app/handlers"
	"github.com/mohamedgamaleldin/go-alpha/app/middleware"
)

// Route route definition
type Route struct {
	path       string
	middleware []middleware.Adapter
	handler    func() http.Handler
}

// RoutesList .. slice array of Routes
type RoutesList []Route

// BuildRoutes .. builds a slice array of Routes
func BuildRoutes() *RoutesList {

	// index route - does nothing
	indexRoute := Route{
		path:       "/",
		middleware: []middleware.Adapter{},
		handler:    handlers.IndexHandler,
	}
	routesList := &RoutesList{indexRoute}

	return routesList
}

// MountRoutes .. mounts all routes to a new router and returns a reference to it
func MountRoutes(logger *log.Logger) http.Handler {
	routes := BuildRoutes()

	router := http.NewServeMux()

	for _, route := range *routes {
		router.Handle(route.path, middleware.Adapt(route.handler(), route.middleware...))
	}

	return middleware.NewLogger(router, logger)
}
