package representacao

import "time"

type ExtratoDeCliente struct {
	Saldo             SaldoDoExtrato `json:"saldo"`
	UltimasTransacoes []TransacaoDoExtrato
}

type TransacaoDoExtrato struct {
	Valor       int       `json:"int"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
}

type SaldoDoExtrato struct {
	Total  int       `json:"total"`
	Limite int       `json:"limite"`
	Data   time.Time `json:"data_extrato"`
}
