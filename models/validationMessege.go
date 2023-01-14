package models

type ValidationMessege struct {
	Password string           `json:"password"`
	Rules    []ValidationRule `json:"rules"`
}