package jwt

import (
	"log"
	"tap/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)



func Create(id string, email string, name string) (string) {
	var secretKey = []byte(config.JWTKey())
	// Создаем новый токен
	token := jwt.New(jwt.SigningMethodHS256)
	// Устанавливаем клеймы (payload) токена
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["email"] = email
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 24 * 60).Unix() // Токен действителен в течение 24 часов

	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
			log.Fatal("Error creating access token: ", err)
	}

	return tokenString
}


