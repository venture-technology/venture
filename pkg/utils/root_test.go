package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAbsPath(t *testing.T) {
	t.Run("encontra basepath", func(t *testing.T) {
		absPath, err := GetAbsPath()
		fmt.Println(absPath)
		assert.NoError(t, err)
	})
}
