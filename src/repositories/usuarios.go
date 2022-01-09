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

//FindByNameOrNick retorna todos usuario que atendam ao nome ou nick
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

//FindByID retorna um usuario pelo id
func (userRepo *usuarios) FindByID(userID uint64) (models.Usuario, error) {
	rows, err := userRepo.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?", userID)
	if err != nil {
		return models.Usuario{}, err
	}
	defer rows.Close()

	var user models.Usuario

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Nome, &user.Nick, &user.Email, &user.CriadoEm); err != nil {
			return models.Usuario{}, err
		}
	}

	return user, nil
}

//Update atualiza um usuario
func (userRepo *usuarios) Update(ID uint64, user models.Usuario) error {
	statement, err := userRepo.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Nome, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

//Delete deleta um usuario
func (userRepo *usuarios) Delete(ID uint64) error {
	statement, err := userRepo.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

//FindByEmail retorna um usuario pelo email
func (userRepo *usuarios) FindByEmail(email string) (models.Usuario, error) {
	rows, err := userRepo.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if err != nil {
		return models.Usuario{}, err
	}
	defer rows.Close()

	var user models.Usuario

	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Senha); err != nil {
			return models.Usuario{}, err
		}
	}

	return user, nil
}

//Follow permite seguir um usuario
func (userRepo *usuarios) Follow(userID uint64, followID uint64) error {
	statement, err := userRepo.db.Prepare("INSERT IGNORE INTO seguidores (usuario_id, seguidor_id) VALUES (?, ?)")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(userID, followID); err != nil {
		return err
	}

	return nil
}

//Unfollow permite que um usuario pare de seguir o outro
func (userRepo *usuarios) Unfollow(userID uint64, followID uint64) error {
	statement, err := userRepo.db.Prepare("DELETE FROM seguidores WHERE usuario_id = ? AND seguidor_id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(userID, followID); err != nil {
		return err
	}

	return nil
}

//FindFollowers retorna todos os usuarios que seguem um usuario
func (userRepo *usuarios) FindFollowers(userID uint64) ([]models.Usuario, error) {
	rows, err := userRepo.db.Query(`
		SELECT u.id, u.nome, u.nick, u.email, u.criadoEm 
		from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.Usuario

	for rows.Next() {
		var user models.Usuario

		if err := rows.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CriadoEm,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
