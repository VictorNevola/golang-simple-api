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
	"strconv"

	"github.com/gorilla/mux"
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
	userId, err := autentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.PublicacaoRepository(db)
	publications, err := repository.Busca(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publications)

}

//BuscaPublicacao busca uma publicacao
func BuscaPublicacao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publiId, erro := strconv.ParseInt(params["publicaoId"], 10, 64)
	if erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.PublicacaoRepository(db)
	publicacao, ererr := repository.BuscaPorID(uint64(publiId))
	if ererr != nil {
		responses.Error(w, http.StatusInternalServerError, ererr)
		return
	}

	responses.JSON(w, http.StatusOK, publicacao)

}

//AtualizarPublicacao atualiza uma publicacao
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

//DeletarPublicacao exclui os dados de uma publicao
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
