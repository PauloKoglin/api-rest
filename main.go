package main

import (
	auth "api-rest/authentication"
	"api-rest/endpoint/servicos"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/token", auth.CreateTokenEndpoint).Methods("GET")
	router.HandleFunc("/servicos", auth.ValidateMiddleware(servicos.GetServicos)).Methods("GET")
	router.HandleFunc("/servicos/{id}", auth.ValidateMiddleware(servicos.GetServico)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
