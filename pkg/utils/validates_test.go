package utils

import "testing"

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
