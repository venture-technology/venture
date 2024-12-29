package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func MakeHash(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
