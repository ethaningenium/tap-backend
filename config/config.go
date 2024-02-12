package config

import (
	"errors"
	"fmt"

	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)
type App struct{
	port string
	location string
	envorment string
}

type DB struct{
	connectionurl string
	name string
}

type JWT struct{
	key string
}

type SMTP struct{
	email string
	password string
}
type Client struct{
	home string
}

type Google struct{
	config *oauth2.Config
	state string
}

type Config struct {
	app *App
	db *DB
	jwt *JWT
	smtp *SMTP
	client *Client
	google *Google
}

var config Config

func InitConfig() error {
	err := godotenv.Load(".env.local")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			return errors.New("Error loading .env file")
		}
	}
	app := &App{
		port: os.Getenv("PORT"),
		location: os.Getenv("LOCATION"),
		envorment: os.Getenv("ENVIRONMENT"),
	}
	db := &DB{
		connectionurl: os.Getenv("DB_URL"),
		name: os.Getenv("DB_NAME"),
	}
	jwt := &JWT{
		key: os.Getenv("JWT_KEY"),
	}
	smtp := &SMTP{
		email: os.Getenv("SMTP_EMAIL"),
		password: os.Getenv("SMTP_PASSWORD"),
	}
	client := &Client{
		home: os.Getenv("CLIENT_HOME"),
	}
	googleConfig := &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s/auth/google/callback", app.location),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	config = Config{
		app: app,
		db: db,
		jwt: jwt,
		smtp: smtp,
		client: client,
		google: &Google{
			config: googleConfig,
			state: "random",
		},
	}
	return nil
}

