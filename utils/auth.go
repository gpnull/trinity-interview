package utils

import (
	"time"

	errorType "trinity/errors"

	"github.com/golang-jwt/jwt"
)

var jwtSecretKey string

func InitJwt(initJwtSecretKey string) {
	jwtSecretKey = initJwtSecretKey
}

// GenerateJWTToken generates a JWT token with the given device Id
func GenerateJWTToken(email string) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
		Issuer:    "your-app",                            // Issuer of the token
		Subject:   email,                                 // Subject (device Id)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ParseJWTToken parses a JWT token and returns the device Id contained in it
func ParseJWTToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errorType.ErrInvalidToken
}
