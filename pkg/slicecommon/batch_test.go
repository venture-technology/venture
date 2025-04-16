package slicecommon_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/venture-technology/venture/pkg/slicecommon"
)

func TestBatchSlice_Ints(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	batchSize := 3

	expected := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7},
	}

	result, err := slicecommon.BatchSlice(input, batchSize)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestBatchSlice_ExactDiv(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	batchSize := 2

	expected := [][]string{
		{"a", "b"},
		{"c", "d"},
	}

	result, err := slicecommon.BatchSlice(input, batchSize)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestBatchSlice_BatchSizeGreaterThanInput(t *testing.T) {
	input := []int{1, 2, 3}
	batchSize := 10

	expected := [][]int{
		{1, 2, 3},
	}

	result, err := slicecommon.BatchSlice(input, batchSize)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestBatchSlice_EmptyInput(t *testing.T) {
	input := []int{}
	batchSize := 3

	result, err := slicecommon.BatchSlice(input, batchSize)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestBatchSlice_InvalidBatchSize(t *testing.T) {
	input := []int{1, 2, 3}
	batchSize := 0

	result, err := slicecommon.BatchSlice(input, batchSize)
	assert.Error(t, err)
	assert.Nil(t, result)
}
