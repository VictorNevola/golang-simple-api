package controllers

import (
	"api/src/autentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Login é responsável por fazer o login do usuário
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.Usuario
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repositorie := repositories.NewRepositorieUsers(db)
	userInDB, err := repositorie.FindByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = security.CheckHash(userInDB.Senha, user.Senha); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := autentication.CreateToken(userInDB.ID)

	w.Write([]byte(token))

}
