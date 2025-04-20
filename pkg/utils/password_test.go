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

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantOk   bool
		wantErr  string
	}{
		{
			name:     "Valid password",
			password: "Abcdef1!",
			wantOk:   true,
			wantErr:  "",
		},
		{
			name:     "Too short",
			password: "Ab1!",
			wantOk:   false,
			wantErr:  "password must be at least 8 characters long",
		},
		{
			name:     "Missing uppercase",
			password: "abcdef1!",
			wantOk:   false,
			wantErr:  "password must contain at least one capital letter",
		},
		{
			name:     "Missing lowercase",
			password: "ABCDEF1!",
			wantOk:   false,
			wantErr:  "password must contain at least one lowercase letter",
		},
		{
			name:     "Missing number",
			password: "Abcdefg!",
			wantOk:   false,
			wantErr:  "password must contain at least one number",
		},
		{
			name:     "Missing special character",
			password: "Abcdefg1",
			wantOk:   false,
			wantErr:  "password must contain at least one special character",
		},
		{
			name:     "Completely invalid",
			password: "abc",
			wantOk:   false,
			wantErr: "password must be at least 8 characters long\n" +
				"password must contain at least one capital letter\n" +
				"password must contain at least one number\n" +
				"password must contain at least one special character",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotErr := ValidatePassword(tt.password)

			if gotOk != tt.wantOk {
				t.Errorf("gotOk = %v, want %v", gotOk, tt.wantOk)
			}

			if tt.wantErr != "" && gotErr != tt.wantErr {
				t.Errorf("gotErr = %q, want %q", gotErr, tt.wantErr)
			}

			if tt.wantErr == "" && gotErr != "" {
				t.Errorf("expected no error, but got: %q", gotErr)
			}
		})
	}
}
