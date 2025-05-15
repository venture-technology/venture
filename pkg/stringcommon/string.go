package stringcommon

import (
	"encoding/json"
)

func Empty(s string) bool {
	return s == ""
}

func RawMessage(input any) (string, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	return string(raw), nil
}
