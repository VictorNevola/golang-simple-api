package controllers

import (
	"api/src/autentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//CriarPublicacao cria uma publicacao no bd
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	userId, err := autentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
	}

	var publicacao models.Publicacao
	if err = json.Unmarshal(body, &publicacao); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	publicacao.AutorId = userId

	if err = publicacao.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	repositorie := repositories.PublicacaoRepository(db)
	publicacao.ID, err = repositorie.Create(publicacao)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
	}

	responses.JSON(w, http.StatusCreated, publicacao)

}

//BuscaPublicacoes busca que apareciam no feed do usuario
func BuscaPublicacoes(w http.ResponseWriter, r *http.Request) {

}

//BuscaPublicacao busca uma publicacao
func BuscaPublicacao(w http.ResponseWriter, r *http.Request) {

}

//AtualizarPublicacao atualiza uma publicacao
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

//DeletarPublicacao exclui os dados de uma publicao
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
