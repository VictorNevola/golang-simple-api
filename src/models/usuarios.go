package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuario
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}

func (usuario *Usuario) validate(step string) error {

	if usuario.Nome == "" {
		return errors.New("nome é obrigatório e não pode ser vazio")
	}
	if usuario.Nick == "" {
		return errors.New("nick é obrigatório e não pode ser vazio")
	}
	if usuario.Senha == "" && step == "create" {
		return errors.New("senha é obrigatório e não pode ser vazio")
	}

	if usuario.Email == "" {
		return errors.New("email é obrigatório e não pode ser vazio")
	}

	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		return errors.New("email inválido")
	}

	return nil
}

//Prepare é responsável por preparar o usuário para ser salvo no banco de dados
func (Usuario *Usuario) Prepare(step string) error {
	if err := Usuario.validate(step); err != nil {
		return err
	}

	if err := Usuario.format(step); err != nil {
		return err
	}

	return nil
}

func (usuario *Usuario) format(step string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Senha = strings.TrimSpace(usuario.Senha)

	if step == "create" {
		passwordHash, err := security.Hash(usuario.Senha)
		if err != nil {
			return err
		}

		usuario.Senha = string(passwordHash)

	}

	return nil
}
