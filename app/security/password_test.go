package security

import (
	"fmt"
	"testing"
)

// TestGet get cache by key
func TestHashAndSalt(t *testing.T) {
	password := "password"
	hashed := HashAndSalt([]byte(password))
	fmt.Println(hashed)
}
