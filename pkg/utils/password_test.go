package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeHash(t *testing.T) {
	password := "myPassword"
	_, err := MakeHash(password)
	if err != nil {
		t.Errorf(fmt.Sprintf("make hash func error %s", err.Error()))
	}
	assert.Nil(t, err)
}

func Test_ValidateHash(t *testing.T) {
	password := "myPassword"
	pwd, err := MakeHash(password)
	if err != nil {
		t.Errorf(fmt.Sprintf("make hash func error %s", err.Error()))
	}
	err = ValidateHash(pwd, password)
	if err != nil {
		t.Errorf(fmt.Sprintf("validate hash func error %s", err.Error()))
	}
	assert.Nil(t, err)
}
