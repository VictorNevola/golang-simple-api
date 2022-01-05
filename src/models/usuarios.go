package models

import (
	"errors"
	"net/mail"
	"strings"
	"time"
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

	if _, err := mail.ParseAddress(usuario.Email); err != nil {
		return errors.New("email é obrigatório, não pode ser vazio ou é invalido")
	}

	return nil
}

//Prepare é responsável por preparar o usuário para ser salvo no banco de dados
func (Usuario *Usuario) Prepare(step string) error {
	if err := Usuario.validate(step); err != nil {
		return err
	}

	Usuario.format()
	return nil
}
func (usuario *Usuario) format() {

	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Senha = strings.TrimSpace(usuario.Senha)

}
