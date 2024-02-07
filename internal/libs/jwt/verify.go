package jwt

import (
	"errors"
	"tap/cfg"

	"github.com/dgrijalva/jwt-go"
)


type TokenClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func VerifyToken(tokenString string) (*TokenClaims, error) {
	var secretKey = []byte(cfg.JwtKey())
	// Разбираем токен и проверяем его подпись
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
	})
	if err != nil {
			return nil, errors.New("Invalid token")
	}

	// Проверяем, действителен ли токен
	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
			return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}