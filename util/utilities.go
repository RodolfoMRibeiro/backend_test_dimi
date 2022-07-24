package util

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/paemuri/brdoc"
)

func TrimAllSpacesInString(str string) string {

	return fmt.Sprint(strings.Replace(str, " ", "", -1))
}

func RevomeSpecialChars(str string) string {
	// take care about some UTF-8 characters ( ex.: á ç õ ù... )

	regx := regexp.MustCompile(`[^ A-Za-z0-9]`)

	return fmt.Sprint(regx.ReplaceAllString(str, ""))
}

func LetOnlyNumbers(str string) string {

	regx := regexp.MustCompile(`[^ 0-9]`)

	return fmt.Sprint(regx.ReplaceAllString(str, ""))
}

func VerifyingCPForCNPJ(str string) (string, bool) {

	switch {

	case brdoc.IsCPF(str):
		return str, true

	case brdoc.IsCNPJ(str):
		return str, true

	default:
		return `invalid CPF or CNPJ`, false

	}
}

func PresentatePanicErros(err error) {
	if err != nil {
		panic(err)
	}
}

func PresentateErros(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
