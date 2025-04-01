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

func KeysExist(m map[string]interface{}, keys []string) bool {
	for _, key := range keys {
		if _, exists := m[key]; !exists {
			return false
		}
	}
	return true
}

func ValidateRequiredGroup(m map[string]interface{}, fields []string) error {
	var presentFields []string

	for _, field := range fields {
		if _, exists := m[field]; exists {
			presentFields = append(presentFields, field)
		}
	}

	if len(presentFields) == 0 {
		return nil
	}

	if len(presentFields) < len(fields) {
		missingFields := findMissingFields(fields, presentFields)
		return errors.New("os seguintes campos são obrigatórios: " + strings.Join(missingFields, ", "))
	}

	return nil
}

func findMissingFields(allFields, presentFields []string) []string {
	var missing []string
	presentSet := make(map[string]bool)

	for _, field := range presentFields {
		presentSet[field] = true
	}

	for _, field := range allFields {
		if !presentSet[field] {
			missing = append(missing, field)
		}
	}

	return missing
}
