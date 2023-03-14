package tool

import (
	"crypto/sha256"
	"encoding/hex"
)

// Math data sha356 value
func MathSha256(data []byte) string {
	str := sha256.Sum256(data)
	return hex.EncodeToString(str[:])
}
