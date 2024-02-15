package services

import (
	"encoding/json"
	"fmt"

	"net/http"
	jwt "tap/internal/libs/jwt"
	m "tap/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) AuthGoogle(googleToken string) (token string, err error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + googleToken)
	if err != nil {
		return "", err
	}
	fmt.Println(googleToken)
	defer resp.Body.Close()

	var user map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return  "", err
	}

	oldUser,err := s.repo.Users.GetUserByEmail(user["email"].(string))
	if err != nil {
		if err.Error() != "user not found"{
			return  "", err
		}
	}
	if oldUser.ID.Hex() != "" {
		token = jwt.Create(oldUser.ID.Hex(), user["email"].(string), user["name"].(string))
		return token, nil
	}

	
	userWithToken := m.RegisterResponse{
		ID: primitive.NewObjectID(),
		Name: user["name"].(string),
		Email: user["email"].(string),
		Password: "googleAuthUser",
		IsEmailVerified: true,
	}
	id, err := s.repo.Users.CreateNewUser(userWithToken)
	if err != nil {
		if err.Error() != "User already exists"{
			return  "", err
		}
	}

	token = jwt.Create(id, user["email"].(string), user["name"].(string))
	return token, nil
}