package email

import (
	"net/smtp"
	"tap/config"
)

func SendHtml(subject string , html string, to string) error {
	email := config.SmtpEmail()
	password := config.SmtpPassword()
	auth := smtp.PlainAuth(
		"",
		email,
		password,
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + html

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		email,
		[]string{to},
		[]byte(msg),
	)

	return err
}