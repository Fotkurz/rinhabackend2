package dominio

type Conta struct {
	Id        int `json:"id"`
	Saldo     int `json:"saldo"`
	Limite    int `json:"limite"`
	ClienteId int `json:"-"`
}
