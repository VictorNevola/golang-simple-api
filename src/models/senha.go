package models

//Senha representa o formatado da requisicao de alteracao de senha
type Senha struct {
	Nova   string `json:"nova"`
	Antiga string `json:"antiga"`
}
