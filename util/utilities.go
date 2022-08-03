package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/paemuri/brdoc"
	"gorm.io/gorm"
)

func ParseMapToJson(mp map[string]string) string {
	str, _ := json.Marshal(mp)
	return string(str)
}

func TrimAllSpacesInString(str string) string {
	return strings.Replace(str, " ", "", -1)
}

func RevomeSpecialChars(str string) string {
	regx := regexp.MustCompile(`[^ A-Za-z0-9]`)
	return regx.ReplaceAllString(str, "")
}

func LetOnlyNumbers(str string) string {
	regx := regexp.MustCompile(`[^ 0-9]`)
	return regx.ReplaceAllString(str, "")
}

// ---------------------------------< validate func >--------------------------------- \\

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func VerifyingCPForCNPJ(str string) (string, bool) {
	return str, brdoc.IsCPF(str) || brdoc.IsCNPJ(str)
}

func ShowErrors(err error) {
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

func ContainsError(err error) bool {
	return err != nil
}

func ValidateTransacion(tx *gorm.DB, err error) error {
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
