package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToPascalCase(s string) string {
	caser := cases.Title(language.BrazilianPortuguese)
	result := caser.String(strings.ToLower(s))

	return result
}
