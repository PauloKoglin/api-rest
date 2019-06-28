package servicos

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Servico representa um servico LH
type Servico struct {
	ID   string `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
}

// GetServicos responde ao endpoint uma lista de servicos
func GetServicos(w http.ResponseWriter, r *http.Request) {
	var servicos []Servico
	servicos = append(servicos, Servico{ID: "1", Nome: "ExecutaServicos"})
	servicos = append(servicos, Servico{ID: "2", Nome: "Escriturar"})
	json.NewEncoder(w).Encode(servicos)
}

// GetServico retorna o servico com ID enviado na requisição
func GetServico(w http.ResponseWriter, r *http.Request) {
	var servicos []Servico
	servicos = append(servicos, Servico{ID: "1", Nome: "ExecutaServicos"})
	servicos = append(servicos, Servico{ID: "2", Nome: "Escriturar"})

	var body = mux.Vars(r)
	for _, item := range servicos {
		if item.ID == body["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
