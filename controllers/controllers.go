package controllers

import (
	"encoding/json"
	"net/http"
	"valid-password/models"
	"valid-password/services"
)

// CheckPasswords quando chegar uma requisição lá no nosso arquivo de “routes.go” vai ser o responsável, quem vai controlar essa página vai ser essa nossa função.
func CheckPasswords(w http.ResponseWriter, r *http.Request) {
	dados := &models.ValidationMessege{} 
	dados.Rules = []models.ValidationRule{}
	json.NewDecoder(r.Body).Decode(dados)
	validador := services.Validation{}

	json.NewEncoder(w).Encode(validador.PasswordValid(dados)) 
}