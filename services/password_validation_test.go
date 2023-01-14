package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDigit(t *testing.T) {
	validation := Validation{}

	assert.True(t, validation.minDigit("TesteSenhaForte!123&", 3))
}

func TestMinLowerCase(t *testing.T) {
	validation := Validation{}

	assert.True(t, validation.minLowercase("TesteSenhaForte!123&", 2))
}

func TestMinSize(t *testing.T) {
	validation := Validation{}

	assert.True(t, validation.minSize("TesteSenhaForte!123&", 8))
}

func TestMinSpecialChars(t *testing.T) {
	validation := Validation{}
	
	assert.True(t, validation.minSpecialChars("TesteSenhaForte!123&", 2))
}

func TestMinUppercase(t *testing.T) {
	validation := Validation{}

	assert.True(t, validation.minUppercase("TesteSenhaForte!123&", 2))
}

func TestNoRepeted(t *testing.T) {
	validation := Validation{}

	assert.True(t, validation.noRepeted("TesteSenhaForte!123&", 0))
}