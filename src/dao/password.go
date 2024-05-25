package dao

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword will return the bcrypt hash value of the password
func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %s", err)
	}
	return string(hash), nil
}

// CheckPassword checks if the password is correct or not
func CheckPassword(hash, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
}
