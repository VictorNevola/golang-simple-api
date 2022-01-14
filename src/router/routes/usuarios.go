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
		isPrivate:   true,
	},
	{
		URI:         "/usuarios/{usuarioId}",
		Method:      http.MethodGet,
		HandlerFunc: controllers.BuscaUsuario,
		isPrivate:   true,
	},
	{
		URI:         "/usuarios/{usuarioId}",
		Method:      http.MethodPut,
		HandlerFunc: controllers.AtualizarUsuario,
		isPrivate:   true,
	},
	{
		URI:         "/usuarios/{usuarioId}",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeletarUsuario,
		isPrivate:   true,
	},
	{
		URI:         "/usuarios/{idToFollow}/seguir",
		Method:      http.MethodPost,
		HandlerFunc: controllers.SeguirUsuario,
		isPrivate:   true,
	},
	{
		URI:         "/usuarios/{idToUnFollow}/parar-de-seguir",
		Method:      http.MethodPost,
		HandlerFunc: controllers.PararDeSeguirUsuario,
		isPrivate:   true,
	},
	{
		URI:         "/usuarios/{usuarioId}/seguidores",
		Method:      http.MethodGet,
		HandlerFunc: controllers.BuscarSeguidores,
		isPrivate:   true,
	},
	{
		URI:         "/usuarios/{usuarioId}/seguindo",
		Method:      http.MethodGet,
		HandlerFunc: controllers.BuscarSeguindo,
		isPrivate:   true,
	},
	{
		URI:         "/usuarios/{usuarioId}/atualizar-senha",
		Method:      http.MethodPost,
		HandlerFunc: controllers.AtualizarSenha,
		isPrivate:   true,
	},
}
