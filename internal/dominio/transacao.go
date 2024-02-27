package dominio

type Transacao struct {
	Id        int    `json:"id"`
	IdConta   int    `json:"id_conta"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}
