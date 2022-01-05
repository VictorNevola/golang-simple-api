package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//Generate generates a router
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
