package routes

import (
	"api/src/controllers"
	"net/http"
)

var rotasLogin = Route{
	URI:         "/login",
	Method:      http.MethodPost,
	HandlerFunc: controllers.Login,
	isPrivate:   false,
}
