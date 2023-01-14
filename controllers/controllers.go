package controllers

import (
	"encoding/json"
	"net/http"
	"valid-password/models"
	"valid-password/services"
)

func CheckPasswords(w http.ResponseWriter, r *http.Request) { // funcionalalidade de expor todos os dados na tela
	dados := &models.ValidationMessege{}
	dados.Rules = []models.ValidationRule{}
	json.NewDecoder(r.Body).Decode(dados)
	validador := services.Validation{}

	json.NewEncoder(w).Encode(validador.PasswordValid(dados)) // encodar dos meus modelos, todos itens
}