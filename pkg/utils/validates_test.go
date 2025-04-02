package utils

import (
	"errors"
	"strings"
	"testing"
)

func TestValidateUpdate(t *testing.T) {
	allowedKeys := map[string]bool{
		"name":  true,
		"email": true,
	}

	tests := []struct {
		name       string
		attributes map[string]interface{}
		expectErr  bool
	}{
		{
			name: "Valid attributes",
			attributes: map[string]interface{}{
				"name":  "John",
				"email": "john@example.com",
			},
			expectErr: false,
		},
		{
			name: "Invalid attribute",
			attributes: map[string]interface{}{
				"username": "john_doe",
			},
			expectErr: true,
		},
		{
			name: "Mix of valid and invalid attributes",
			attributes: map[string]interface{}{
				"name":     "John",
				"location": "USA",
			},
			expectErr: true,
		},
		{
			name:       "Empty attributes",
			attributes: map[string]interface{}{},
			expectErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUpdate(tt.attributes, allowedKeys)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
		})
	}
}

func TestKeysExist(t *testing.T) {
	tests := []struct {
		name   string
		m      map[string]interface{}
		keys   []string
		expect bool
	}{
		{"All keys exist", map[string]interface{}{"a": 1, "b": 2}, []string{"a", "b"}, true},
		{"Missing one key", map[string]interface{}{"a": 1}, []string{"a", "b"}, false},
		{"Empty keys list", map[string]interface{}{"a": 1}, []string{}, true},
		{"Empty map", map[string]interface{}{}, []string{"a"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeysExist(tt.m, tt.keys); got != tt.expect {
				t.Errorf("KeysExist() = %v, want %v", got, tt.expect)
			}
		})
	}
}

func TestValidateRequiredGroup(t *testing.T) {
	tests := []struct {
		name      string
		m         map[string]interface{}
		fields    []string
		expectErr error
	}{
		{"All fields present", map[string]interface{}{"a": 1, "b": 2}, []string{"a", "b"}, nil},
		{"No fields present", map[string]interface{}{}, []string{"a", "b"}, nil},
		{"Some fields missing", map[string]interface{}{"a": 1}, []string{"a", "b"}, errors.New("os seguintes campos são obrigatórios: b")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRequiredGroup(tt.m, tt.fields)
			if (err != nil) != (tt.expectErr != nil) || (err != nil && tt.expectErr != nil && !strings.Contains(err.Error(), tt.expectErr.Error())) {
				t.Errorf("ValidateRequiredGroup() error = %v, want %v", err, tt.expectErr)
			}
		})
	}
}
