package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash hashing password using bcrypt
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword validate password using bcrypt
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Encrypt use for encrypt
func Encrypt(pass string) (string, error) {
	hashedPassword, err := Hash(pass)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
