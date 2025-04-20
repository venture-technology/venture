package utils

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

func MakeHash(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ValidateHash(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func ValidatePassword(password string) (bool, string) {
	var errors []string

	if utf8.RuneCountInString(password) < 8 {
		errors = append(errors, "password must be at least 8 characters long")
	}

	if matched, _ := regexp.MatchString(`[A-Z]`, password); !matched {
		errors = append(errors, "password must contain at least one capital letter")
	}

	if matched, _ := regexp.MatchString(`[a-z]`, password); !matched {
		errors = append(errors, "password must contain at least one lowercase letter")
	}

	if matched, _ := regexp.MatchString(`[0-9]`, password); !matched {
		errors = append(errors, "password must contain at least one number")
	}

	if matched, _ := regexp.MatchString(`[!@#\$%\^&\*(),.?":{}|<>]`, password); !matched {
		errors = append(errors, "password must contain at least one special character")
	}

	return len(errors) == 0, strings.Join(errors, "\n")
}
