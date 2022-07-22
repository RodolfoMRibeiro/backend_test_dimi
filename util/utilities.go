package util

import (
	"fmt"
	"regexp"
	"strings"
)

func TrimAllSpacesInString(str string) string {

	return fmt.Sprint(strings.Replace(str, " ", "", -1))
}

func RevomeSpecialChars(str string) string {
	// take care about some characters of UTF-8 package ( ex.: á ç õ ù... )

	regx := regexp.MustCompile(`[^ A-Za-z0-9]`)

	return fmt.Sprint(regx.ReplaceAllString(str, ""))
}

func LetOnlyNumbers(str string) string {

	regx := regexp.MustCompile(`[^ 0-9]`)

	return fmt.Sprint(regx.ReplaceAllString(str, ""))
}
