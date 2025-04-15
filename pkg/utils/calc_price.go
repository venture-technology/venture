package utils

func CalculateContract(distance, amount float64) float64 {
	if distance < 2 {
		return 200
	}
	diff := distance - 2
	return 200 + (amount * diff)
}
