package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

//Route represents all routes for the API
type Route struct {
	URI         string
	Method      string
	HandlerFunc func(http.ResponseWriter, *http.Request)
	isPrivate   bool
}

//Config inserts all routes into the router
func Config(r *mux.Router) *mux.Router {
	routes := rotasUsuarios
	routes = append(routes, rotasLogin)
	routes = append(routes, rotasPublicacoes...)

	for _, route := range routes {
		if route.isPrivate {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.HandlerFunc))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.HandlerFunc)).Methods(route.Method)
		}
	}

	return r
}
