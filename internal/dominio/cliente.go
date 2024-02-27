package dominio

type Cliente struct {
	Id    int   `json:"id"`
	Conta Conta `json:"conta"`
}
