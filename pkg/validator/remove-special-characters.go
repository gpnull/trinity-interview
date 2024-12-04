package validator

import (
	"regexp"
	"strings"
)

func RemoveSpecialCharacters(input string) string {
	// Remove whitespace and line breaks
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "\n", "")

	// Remove special characters using regular expression
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	input = reg.ReplaceAllString(input, "")

	return input
}
