package routes

import (
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

	for _, route := range routes {
		r.HandleFunc(route.URI, route.HandlerFunc).Methods(route.Method)
	}

	return r
}
