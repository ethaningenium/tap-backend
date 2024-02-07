package services

import (
	"errors"
	bc "tap/internal/libs/bcrypt"
	jwt "tap/internal/libs/jwt"
	m "tap/internal/models"
)

func (s *Service) Register(user m.UserRegister) ( refreshToken string , accessToken string,err error) {

	// Create JWT
	refreshToken, err = jwt.CreateRefreshToken(user.Email)
	if err != nil {
		
		return "", "", errors.New("Error creating refresh token")
	}

	hashedPassword, err := bc.HashPassword(user.Password)
	if err != nil {
	
		return "", "", errors.New("Error hashing password")
	}

	userWithToken := m.UserRegisterRequest{
		Name: user.Name,
		Email: user.Email,
		Password: hashedPassword,
		RefreshToken: refreshToken,
	}
	
	err = s.repo.Users.CreateNewUser(userWithToken)
	if err != nil {
		
		return "", "", errors.New("Error creating user")
	}

	accessToken, err = jwt.CreateAccessToken(user.Email)
	if err != nil {
		
		return "", "", errors.New("Error creating access token")
	}

	return refreshToken, accessToken, nil
}

func (s *Service) Login(email string, password string) (name string ,refreshToken string , accessToken string,err error) {
	user, err := s.repo.Users.GetUserByEmail(email)
	if err != nil {
		return "", "", "", errors.New("User not found")
	}
	if err := bc.CheckPasswordHash(password, user.Password); err != nil {
		return "", "", "", errors.New("Invalid password")
	}
	refreshToken, err = jwt.CreateRefreshToken(user.Email)
	if err != nil {
		return "", "", "", errors.New("Error creating refresh token")
	}
	accessToken, err = jwt.CreateAccessToken(user.Email)
	if err != nil {
		return "", "", "", errors.New("Error creating access token")
	}
	err = s.repo.Users.SetNewRefreshToken(user.Email, refreshToken)
	if err != nil {
		return "", "", "", errors.New("Error setting refresh token")
	}
	return user.Name, refreshToken,accessToken, nil
}

func (s *Service) Getme(refreshToken string) (email string, name string, accessToken string , err error) {
	user, err := s.repo.Users.GetUserByRefreshToken(refreshToken)
	if err != nil {
		return 	"", "", "", errors.New("User not found")
	}
	accessToken, err = jwt.CreateAccessToken(user.Email)
	if err != nil {
		return "", "", "", errors.New("Error creating access token")
	}
	return user.Email, user.Name, accessToken, nil
}
