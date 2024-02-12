package jwt

import (
	"errors"
	"tap/config"

	"github.com/dgrijalva/jwt-go"
)


type AccessClaims struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func VerifyAccess(accessToken string) (*AccessClaims, error) {
	var secretKey = []byte(config.JWTKey())
	// Разбираем токен и проверяем его подпись
	token, err := jwt.ParseWithClaims(accessToken, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
	})
	if err != nil {
			return nil, errors.New("Invalid token")
	}

	// Проверяем, действителен ли токен
	claims, ok := token.Claims.(*AccessClaims)
	if !ok || !token.Valid {
			return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

func VerifyRefresh(refreshToken string) (*RefreshClaims, error) {
	var secretKey = []byte(config.JWTKey())
	// Разбираем токен и проверяем его подпись
	token, err := jwt.ParseWithClaims(refreshToken, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
	})
	if err != nil {
			return nil, errors.New("Invalid token")
	}

	// Проверяем, действителен ли токен
	claims, ok := token.Claims.(*RefreshClaims)
	if !ok || !token.Valid {
			return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}