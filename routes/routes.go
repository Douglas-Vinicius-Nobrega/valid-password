package routes

import (
	"log"
	"net/http"
	"valid-password/controllers"

	"github.com/gorilla/mux"
)

func HandleRequest() { // requisições
	r := mux.NewRouter() // nosso mapeamento de rotas
	r.HandleFunc("/verify", controllers.CheckPasswords).Methods(http.MethodPost) // nossa rota,  usando método POST 
	log.Fatal(http.ListenAndServe(":8080", r)) 
}