package utils

import "fmt"

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
