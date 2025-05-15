package utils

func CalculateContract(distance float64, amount int64) int64 {
	if distance < 2 {
		return 20000
	}
	diff := distance - 2
	return 20000 + int64(float64(amount)*diff)
}
