package router

import (
	"api/src/router/routes"
	"github.com/gorilla/mux"
)

// Build will return a router with routes configurations
func Build() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
