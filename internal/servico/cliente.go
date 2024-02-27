package servico

import (
	"errors"

	"github.com/Fotkurz/rinhabackend2/internal/representacao"
)

type IServicoCliente interface {
	Extrato(id int) (representacao.ExtratoDeCliente, error)
}

type servicoCliente struct {
}

func NewServicoCliente() IServicoCliente {
	return &servicoCliente{}
}

func (s *servicoCliente) Extrato(id int) (representacao.ExtratoDeCliente, error) {
	return representacao.ExtratoDeCliente{}, errors.New("not implemented")
}
