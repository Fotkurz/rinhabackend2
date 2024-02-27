package controlador

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Fotkurz/rinhabackend2/internal/db"
	"github.com/Fotkurz/rinhabackend2/internal/dominio"
	"github.com/Fotkurz/rinhabackend2/pkg/apierror"
	"github.com/go-chi/chi"
)

var clientes = []dominio.Cliente{
	{
		Id: 1,
		Conta: dominio.Conta{
			Id:        1,
			Saldo:     1000,
			Limite:    10000,
			ClienteId: 1,
		},
	}, {
		Id: 2,
		Conta: dominio.Conta{
			Id:        2,
			Saldo:     1000,
			Limite:    10000,
			ClienteId: 2,
		},
	}, {
		Id: 3,
		Conta: dominio.Conta{
			Id:        2,
			Saldo:     1000,
			Limite:    10000,
			ClienteId: 3,
		},
	},
}

func Extrato(w http.ResponseWriter, r *http.Request) {
	raw := chi.URLParam(r, "id")
	clienteId, err := strconv.ParseInt(raw, 10, 36)
	if raw == "" || err != nil {
		apierror.NewAPIError(w, http.StatusBadRequest, "id de cliente inválido.")
		return
	}

	fmt.Println(clienteId)
	db.CONN.Begin()
	db.CONN.Query("SELECT * FROM ")

}
