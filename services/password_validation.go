package services

import (
	"strings"
	"unicode"
	"unicode/utf8"
	"valid-password/models"
)

const special =  `!@#$%^&*()-+\/{}[]`
 
type Validation struct {}

func (v Validation) PasswordValid(vPassword *models.ValidationMessege) models.PVerify {
	result := models.PVerify{}
	for _ , rule := range vPassword.Rules {
		switch (rule.Rule ) {
		case models.MinSize:
			if !v.minSize(vPassword.Password, rule.Value) {
				result.NoMatch = append(result.NoMatch, models.MinSize)
			}

		case models.MinUppercase:
			if !v.minUppercase(vPassword.Password, rule.Value) {
				result.NoMatch = append(result.NoMatch, models.MinUppercase)
			}

		case models.MinLowercase:
			if !v.minLowercase(vPassword.Password, rule.Value) {
				result.NoMatch = append(result.NoMatch, models.MinLowercase)
			}

		case models.MinDigit:
			if !v.minDigit(vPassword.Password, rule.Value) {
				result.NoMatch = append(result.NoMatch, models.MinDigit)
			}

		case models.MinSpecialChars:
			if !v.minSpecialChars(vPassword.Password, rule.Value) {
				result.NoMatch = append(result.NoMatch, models.MinSpecialChars)
			}
			
		case models.NoRepeted:
			if !v.noRepeted(vPassword.Password, rule.Value) {
				result.NoMatch = append(result.NoMatch, models.NoRepeted)
			}
		}
	}
	result.Verify = len(result.NoMatch) == 0
	return result
}

func (v Validation) minDigit(password string, min int) bool {
	counter := 0
	for _, ch := range password {
		if unicode.IsDigit(ch) {
			counter++
		}
	}
	return counter >= min
}

func (v Validation) minLowercase(password string, min int) bool {
	counter := 0
	for _, ch := range password {
		if unicode.IsLower(ch) {
			counter++
		}
	}
	return counter >= min
}

func (v Validation) minSize(password string, min int) bool {
	return utf8.RuneCountInString(password) >= min
}

func (v Validation) minSpecialChars(password string, min int) bool {
	counter := 0
	for _, ch := range password {
		if strings.Contains(special, string(ch)) {
			counter++
		}
	}
	return counter >= min
}

func (v Validation) minUppercase(password string, min int) bool {
	counter := 0
	for _, ch := range password {
		if unicode.IsUpper(ch) {
			counter++
		}
	}
	return counter >= min
}

func (v Validation) noRepeted(password string, min int) bool {
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			return false
		}
	}
	return true
}








