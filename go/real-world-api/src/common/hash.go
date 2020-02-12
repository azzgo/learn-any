package common

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hash godoc
func Hash(s string) string {
	sha := sha256.New()
	sha.Write([]byte(s))
	sha.Write([]byte("salt"))
	hashed := sha.Sum(nil)
	return hex.EncodeToString(hashed[:])
}
