package validate

import (
	"unicode"
	"unicode/utf8"
)

func PasswordMinLength(password string, minLength int) bool {
	return utf8.RuneCountInString(password) >= minLength
}

func PasswordNumeric(password string) bool {
	for _, v := range password {
		if !unicode.IsDigit(v) {
			return false
		}
	}
	return true
}
