package models

import "time"

//Publicao representa uma publicacao feita por um usuario
type Publicacao struct {
	ID       uint64    `json:"id,omitempty"`
	Title    string    `json:"title,omitempty"`
	Conteudo string    `json:"Conteudo,omitempty"`
	AutorId  uint64    `json:"autorId,omitempty"`
	Curtidas uint64    `json:"curtidas"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}
