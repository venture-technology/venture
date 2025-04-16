package slicecommon

import "fmt"

// BatchSlice divides a slice into batches of a specified size.

func BatchSlice[T any](inputSlice []T, batchSize int) ([][]T, error) {
	if batchSize <= 0 {
		return nil, fmt.Errorf("batch size must be greater than 0")
	}

	var batches [][]T
	for startIndex := 0; startIndex < len(inputSlice); startIndex += batchSize {
		endIndex := startIndex + batchSize
		if endIndex > len(inputSlice) {
			endIndex = len(inputSlice)
		}
		batches = append(batches, inputSlice[startIndex:endIndex])
	}

	return batches, nil
}
