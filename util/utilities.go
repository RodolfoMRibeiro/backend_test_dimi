package util

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
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

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
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

func BadRequest(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
}
