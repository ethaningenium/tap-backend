package jwt

import (
	"tap/cfg"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = cfg.JwtKey()

func CreateAccessToken(email string) (string, error) {
	// Создаем новый токен
	token := jwt.New(jwt.SigningMethodHS256)
	// Устанавливаем клеймы (payload) токена
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour).Unix() // Токен действителен в течение 24 часов

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
			return "", err
	}

	return tokenString, nil
}

func CreateRefreshToken(email string) (string, error) {
	// Создаем новый токен
	token := jwt.New(jwt.SigningMethodHS256)
	// Устанавливаем клеймы (payload) токена
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix() // Токен действителен в течение 24 часов

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
			return "", err
	}

	return tokenString, nil
}
