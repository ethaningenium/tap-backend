package cfg

import (
	"errors"

	"os"

	"github.com/joho/godotenv"
)

type Congif struct {
	Port string
	DB string
	JwtKey string
}

var config Congif

func InitConfig() error {
	err := godotenv.Load(".env.local")

	if err != nil {
		return errors.New("Error loading .env file")
	}
	port := os.Getenv("PORT")
	dburl := os.Getenv("MONGO_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	config = Congif{
		Port: port,
		DB: dburl,
		JwtKey: jwtSecret,
	}
	if port == "" || dburl == "" || jwtSecret == "" {
		return errors.New("Error loading .env variables")
	}

	return nil
}

func Port () string {
	return config.Port
}

func DB() string {
	return config.DB
}

func JwtKey() string {
	return config.JwtKey
}