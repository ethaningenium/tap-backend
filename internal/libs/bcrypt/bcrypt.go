package bcrypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		log.Fatal("Error hashing password: ", err)
	}
	return string(hashedPassword)
}

func CheckPasswordHash(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}