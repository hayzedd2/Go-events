package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email, userName string, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"userName": userName,
		"userId":   userId,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", errors.New("could not parse token")
	}

	IsTokenValid := parsedToken.Valid

	if !IsTokenValid {
		return "", errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claim")
	}
	// email := claims["email"].(string)
	// userName := claims["userName"].(string)
	// userId := (claims["userId"].(float64))
	userId, ok := claims["userId"].(string)
	if !ok {
		return "", errors.New("userId in token is not a valid string")
	}
	return userId, nil
}
