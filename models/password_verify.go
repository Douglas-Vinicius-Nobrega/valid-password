package models

type PVerify struct {
	Verify  bool       `json:"verify"`
	NoMatch []RuleType `json:"noMatch"`
}