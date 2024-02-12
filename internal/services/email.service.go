package services

import (
	"fmt"
	"tap/config"
	"tap/internal/libs/email"
	m "tap/internal/models"
)

func (s *Service) SendEmail(emailBody m.EmailRequest) error {
	location := config.Location()
	user, err := s.repo.Users.GetUserByEmail(emailBody.Email)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("<h1>Your verification link is:</h1><p>%s/verify/%s?redirect=%s</p>",location, user.VerifyCode, emailBody.Redirect)
	return email.SendHtml("Verification", msg, emailBody.Email)
}

func (s *Service) VerifyEmail(verifyCode string) error {
	return s.repo.Users.SetVerifiedTrue(verifyCode)
}