package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

//Logger escrever informacoes das requisicoes no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

//Authenticate verifica se o usuario fazendo a requisicao esta autenticado
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("validando token")
		next(w, r)
	}
}
