package utils

import (
	"errors"
	"fmt"
	"strings"
)

func ValidateUpdate(attributes map[string]interface{}, allowedKeys map[string]bool) error {
	var invalidKeys []string

	for key := range attributes {
		if !allowedKeys[key] {
			invalidKeys = append(invalidKeys, key)
		}
	}

	if len(invalidKeys) > 0 {
		return errors.New(fmt.Sprintf("chaves não permitidas: %s", strings.Join(invalidKeys, ", ")))
	}

	return nil
}

func KeysExist(m map[string]interface{}, keys ...string) bool {
	for _, key := range keys {
		if _, exists := m[key]; !exists {
			return false
		}
	}
	return true
}
