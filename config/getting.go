package config

import "golang.org/x/oauth2"

func Location() string {
	return config.app.location
}

func Envorment() string {
	return config.app.envorment
}

func Port() string {
	return config.app.port
}

func ConnectionUrl() string {
	return config.db.connectionurl
}

func DBName() string {
	return config.db.name
}

func JWTKey() string {
	return config.jwt.key
}

func SmtpEmail() string {
	return config.smtp.email
}

func SmtpPassword() string {
	return config.smtp.password
}

func ClientHome() string {
	return config.client.home
}

func GoogleState() string {
	return config.google.state
}

func GoogleConfig() *oauth2.Config {
	return config.google.config
}
