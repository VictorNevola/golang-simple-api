package repositories

import (
	"api/src/models"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

//NewRepositorieUsers cria um novo repositorio de usuarios
func NewRepositorieUsers(db *sql.DB) *usuarios {
	return &usuarios{db}
}

//Create cria um usuario
func (user *usuarios) Create(usuario models.Usuario) (uint64, error) {

	statement, err := user.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}
