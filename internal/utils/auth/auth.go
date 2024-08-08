package auth

import (
	"crypto/sha256"
	"fmt"
)

/**
* Auth Helpers
**/

// Hashing Password
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}
