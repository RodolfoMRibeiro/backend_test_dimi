package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimAllSpacesInString(t *testing.T) {

	t.Run("verify if the func is usefull to trim all spaces", func(t *testing.T) {

		recivedBrokenString := TrimAllSpacesInString("C@l%i*e    n(tA   ut#he*  n t!  ica $tio$%nTo@  ken    ")

		expectedCleanString := "C@l%i*en(tAut#he*nt!ica$tio$%nTo@ken"

		assert.Equal(t, expectedCleanString, recivedBrokenString)

	})
}

func TestRevomeSpecialChars(t *testing.T) {

	t.Run("verify if regex trim all non-letter character", func(t *testing.T) {

		recivedBrokenString := RevomeSpecialChars("C@l%i*en(tAut#he*nt!ica$tio$%nTo@kenáç")

		expectedOnlyletters := "ClientAuthenticationToken"

		assert.Equal(t, expectedOnlyletters, recivedBrokenString)
	})
}

func TestLetOnlyNumbers(t *testing.T) {

	t.Run("verify if regex trim all non-number character", func(t *testing.T) {

		recivedBrokenString := LetOnlyNumbers("1g2h3#4m5n6v7c8U9$10")

		expectedOnlyNumbers := "12345678910"

		assert.Equal(t, expectedOnlyNumbers, recivedBrokenString)
	})

}

func TestVerifyingCPForCNPJ(t *testing.T) {

	t.Run("verify if CPF is valid", func(t *testing.T) {

		_, recivedDocValidation := VerifyingCPForCNPJ("05293365029")

		expectedValidation := true

		assert.Equal(t, expectedValidation, recivedDocValidation)
	})

	t.Run("verify if CPF isn't valid", func(t *testing.T) {

		_, recivedDocValidation := VerifyingCPForCNPJ("05293365429")

		expectedValidation := false

		assert.Equal(t, expectedValidation, recivedDocValidation)
	})

	t.Run("verify if CNPJ is valid", func(t *testing.T) {

		_, recivedDocValidation := VerifyingCPForCNPJ("82007188000164")

		expectedValidation := true

		assert.Equal(t, expectedValidation, recivedDocValidation)
	})

	t.Run("verify if CNPJ isn't valid", func(t *testing.T) {

		_, recivedDocValidation := VerifyingCPForCNPJ("82017188000164")

		expectedValidation := false

		assert.Equal(t, expectedValidation, recivedDocValidation)
	})
}
