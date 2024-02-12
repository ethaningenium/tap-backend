package services

import (
	"errors"
	bc "tap/internal/libs/bcrypt"
	jwt "tap/internal/libs/jwt"
	"tap/internal/libs/random"
	m "tap/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) Register(user m.RegisterBody) ( refreshToken string , accessToken string,err error) {
	refreshToken = jwt.CreateRefresh(user.Email)
	hashedPassword := bc.HashPassword(user.Password)

	userWithToken := m.RegisterResponse{
		ID: primitive.NewObjectID(),
		Name: user.Name,
		Email: user.Email,
		Password: hashedPassword,
		RefreshToken: refreshToken,
		VerifyCode: random.GenerateRandomString(36),
	}
	
	id, err := s.repo.Users.CreateNewUser(userWithToken)
	if err != nil {
		return "", "", err
	}

	accessToken = jwt.CreateAccess(id, user.Name, user.Email)
	return refreshToken, accessToken, nil
}

func (s *Service) Login(user m.LoginBody) (name string ,refreshToken string , accessToken string,err error) {
	myUser, err := s.repo.Users.GetUserByEmail(user.Email)
	if err != nil {
		return "", "", "", errors.New("User not found")
	}
	if err := bc.CheckPasswordHash(user.Password, myUser.Password); err != nil {
		return "", "", "", errors.New("Invalid password")
	}
	refreshToken = jwt.CreateRefresh(user.Email)
	accessToken = jwt.CreateAccess(myUser.ID.Hex(), myUser.Name, myUser.Email)
	err = s.repo.Users.SetNewRefreshToken(user.Email, refreshToken)
	if err != nil {
		return "", "", "", errors.New("Error setting refresh token")
	}
	return myUser.Name, refreshToken,accessToken, nil
}

func (s *Service) Getme(email string) (name string, accessToken string , err error) {
	user, err := s.repo.Users.GetUserByEmail(email)
	if err != nil {
		return  "", "", errors.New("User not found")
	}
	accessToken = jwt.CreateAccess(user.ID.Hex(), user.Name, user.Email)
	return  user.Name, accessToken, nil
}
