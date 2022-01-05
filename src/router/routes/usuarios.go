package routes

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Route{
	{
		URI:         "/usuarios",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CriarUsuario,
		isPrivate:   false,
	},
	{
		URI:         "/usuarios",
		Method:      http.MethodGet,
		HandlerFunc: controllers.BuscaUsuarios,
		isPrivate:   false,
	},
	{
		URI:         "/usuarios/{usuarioId}",
		Method:      http.MethodGet,
		HandlerFunc: controllers.BuscaUsuario,
		isPrivate:   false,
	},
	{
		URI:         "/usuarios/{usuarioId}",
		Method:      http.MethodPut,
		HandlerFunc: controllers.AtualizarUsuario,
		isPrivate:   false,
	},
	{
		URI:         "/usuarios/{usuarioId}",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeletarUsuario,
		isPrivate:   false,
	},
}
