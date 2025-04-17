package utils

import (
	"fmt"
	"regexp"
)

func BuildAddress(street, number, complement, zip string) string {
	if complement == "" {
		return fmt.Sprintf(
			"%s, %s, %s",
			street,
			number,
			zip,
		)
	}
	return fmt.Sprintf(
		"%s, %s, %s, %s",
		street,
		number,
		complement,
		zip,
	)
}

func ValidadeZip(zip string) error {
	rule := regexp.MustCompile(`^\d{8}$`)
	status := rule.MatchString(zip)

	if !status {
		return fmt.Errorf("invalid zip")
	}

	return nil
}
