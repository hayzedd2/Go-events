package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", errors.New("could not hash password")
	}
	return string(passwordByte), nil
}

func ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
