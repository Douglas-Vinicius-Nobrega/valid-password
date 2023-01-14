package routes

import (
	"log"
	"net/http"
	"valid-password/controllers"

	"github.com/gorilla/mux"
)

// HandleRequest ela vai lidar com as requisições
func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/verify", controllers.CheckPasswords).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r)) 
}