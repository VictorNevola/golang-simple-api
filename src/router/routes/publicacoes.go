package routes

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Route{
	{
		URI:         "/publicacoes",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CriarPublicacao,
		isPrivate:   true,
	},
	{
		URI:         "/publicacoes",
		Method:      http.MethodGet,
		HandlerFunc: controllers.BuscaPublicacoes,
		isPrivate:   true,
	},
	{
		URI:         "/publicacoes/{publicaoId}",
		Method:      http.MethodGet,
		HandlerFunc: controllers.BuscaPublicacao,
		isPrivate:   true,
	},
	{
		URI:         "/publicacoes/{publicaoId}",
		Method:      http.MethodPut,
		HandlerFunc: controllers.AtualizarPublicacao,
		isPrivate:   true,
	},
	{
		URI:         "/publicacoes/{publicaoId}",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeletarPublicacao,
		isPrivate:   true,
	},
}
