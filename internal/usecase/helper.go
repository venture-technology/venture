package usecase

import (
	"log"
)

func CalculateContract(distance, amount float64) float64 {

	log.Print(distance)

	if distance < 2 {
		return 200
	}

	diff := distance - 2

	return 200 + (amount * diff)

}
