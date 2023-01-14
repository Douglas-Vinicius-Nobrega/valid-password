package models

type ValidationRule struct {
	Rule  RuleType `json:"rule"`
	Value int      `json:"value"`
}
