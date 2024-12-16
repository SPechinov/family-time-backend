package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateJWT(secretKey string, exp time.Duration, body map[string]string) (string, error) {
	mapClaims := jwt.MapClaims{"exp": time.Now().Add(exp).Unix()}

	for key, value := range body {
		mapClaims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func IsValidJWT(secretKey, token string) (bool, *jwt.Token) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !jwtToken.Valid {
		return false, jwtToken
	}

	return true, jwtToken
}
