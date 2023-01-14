package models

type RuleType string

const (
	MinSize         RuleType = "minSize"
	MinUppercase    RuleType = "minUppercase"
	MinLowercase    RuleType = "minLowercase"
	MinDigit        RuleType = "minDigit"
	MinSpecialChars RuleType = "minSpecialChars"
	NoRepeted       RuleType = "noRepeted"
)