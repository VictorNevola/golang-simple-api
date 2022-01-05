package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

//NewRepositorieUsers cria um novo repositorio de usuarios
func NewRepositorieUsers(db *sql.DB) *usuarios {
	return &usuarios{db}
}

//Create cria um usuario
func (userRepo *usuarios) Create(usuario models.Usuario) (uint64, error) {

	statement, err := userRepo.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)")
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

//Find retorna todos usuario que atendam ao nome ou nick
func (userRepo *usuarios) FindByNameOrNick(nameOrNick string) ([]models.Usuario, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := userRepo.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ?", nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []models.Usuario

	for rows.Next() {
		var usuario models.Usuario

		err = rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm)
		if err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil

}
