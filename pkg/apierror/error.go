package apierror

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Codigo   int    `json:"codigo"`
	Mensagem string `json:"mensagem"`
}

func NewAPIError(w http.ResponseWriter, codigo int, mensagem string) {
	w.Header().Set("content-type", "application/json")

	apiErr := APIError{
		Codigo:   codigo,
		Mensagem: mensagem,
	}

	data, err := json.Marshal(&apiErr)
	if err != nil {
		w.Write([]byte(`{\"codigo\":500,\"mensagem\":\"falha na requisição.\"}`))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
	w.WriteHeader(codigo)
}
