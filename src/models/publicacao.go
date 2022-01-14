package models

import (
	"errors"
	"strings"
	"time"
)

//Publicao representa uma publicacao feita por um usuario
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadoEm  time.Time `json:"criadoEm,omitempty"`
}

//Prepare prepara a publicacao para ser salva no bd
func (publicacao *Publicacao) Prepare() error {
	if err := publicacao.validar(); err != nil {
		return err
	}

	publicacao.formatar()
	return nil

}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("titulo nao pode ser vazio")
	}

	if publicacao.Conteudo == "" {
		return errors.New("conteudo nao pode ser vazio")
	}

	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
