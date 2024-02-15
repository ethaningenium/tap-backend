package services

import (
	"errors"
	bc "tap/internal/libs/bcrypt"
	jwt "tap/internal/libs/jwt"
	"tap/internal/libs/random"
	m "tap/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) Register(user m.RegisterBody) (token string,err error) {
	
	hashedPassword := bc.HashPassword(user.Password)

	userWithToken := m.RegisterResponse{
		ID: primitive.NewObjectID(),
		Name: user.Name,
		Email: user.Email,
		Password: hashedPassword,
		VerifyCode: random.GenerateRandomString(36),
	}
	
	id, err := s.repo.Users.CreateNewUser(userWithToken)
	if err != nil {
		return  "", err
	}

	token = jwt.Create(id, user.Name, user.Email)
	return token, nil
}

func (s *Service) Login(user m.LoginBody) (name string ,token string,err error) {
	myUser, err := s.repo.Users.GetUserByEmail(user.Email)
	if err != nil {
		return "", "", errors.New("User not found")
	}
	if err := bc.CheckPasswordHash(user.Password, myUser.Password); err != nil {
		return "", "", errors.New("Invalid password")
	}
	
	token = jwt.Create(myUser.ID.Hex(), myUser.Name, myUser.Email)
	return myUser.Name,token, nil
}

func (s *Service) Getme(userId string) (user m.RegisterResponse , err error) {
	user, err = s.repo.Users.GetOneByID(userId)
	if err != nil {
		return  m.RegisterResponse{},  errors.New("User not found")
	}
	
	return  user, nil
}
