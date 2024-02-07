package services

import (
	bc "tap/internal/libs/bcrypt"
	jwt "tap/internal/libs/jwt"
	m "tap/internal/models"
)

func (s *Service) Register(user m.UserRegister) ( refreshToken string , accessToken string,err error) {

	// Create JWT
	refreshToken, err = jwt.CreateRefreshToken(user.Email)
	if err != nil {
		return "", "", err
	}

	hashedPassword, err := bc.HashPassword(user.Password)
	if err != nil {
		return "", "", err
	}

	userWithToken := m.UserRegisterRequest{
		Name: user.Name,
		Email: user.Email,
		Password: hashedPassword,
		RefreshToken: refreshToken,
	}
	
	err = s.repo.Users.CreateNewUser(userWithToken)
	if err != nil {
		return "", "", err
	}

	accessToken, err = jwt.CreateAccessToken(user.Email)
	if err != nil {
		return "", "", err
	}

	return refreshToken, accessToken, nil
}

func (s *Service) Login(email string, password string) (name string ,refreshToken string , accessToken string,err error) {
	user, err := s.repo.Users.GetUserByEmail(email)
	if err != nil {
		return "", "", "", err
	}

	if err := bc.CheckPasswordHash(password, user.Password); err != nil {
		return "", "", "", err
	}

	refreshToken, err = jwt.CreateRefreshToken(user.Email)
	if err != nil {
		return "", "", "", err
	}
	accessToken, err = jwt.CreateAccessToken(user.Email)
	if err != nil {
		return "", "", "", err
	}
	err = s.repo.Users.SetNewRefreshToken(user.Email, refreshToken)
	if err != nil {
		return "", "", "", err
	}
	return user.Name, refreshToken,accessToken, nil
}
