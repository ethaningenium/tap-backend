package services

import (
	"encoding/json"

	"net/http"
	jwt "tap/internal/libs/jwt"
	m "tap/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) AuthGoogle(token string) (accessToken string, refreshToken string, err error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	var user map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return  "", "", err
	}


	refreshToken = jwt.CreateRefresh(user["email"].(string))

	userWithToken := m.RegisterResponse{
		ID: primitive.NewObjectID(),
		Name: user["name"].(string),
		Email: user["email"].(string),
		Password: "googleAuthUser",
		RefreshToken: refreshToken,
		IsEmailVerified: true,
	}
	id, err := s.repo.Users.CreateNewUser(userWithToken)
	if err != nil {
		if err.Error() != "User already exists"{
			return "", "", err
		}
	}

	accessToken = jwt.CreateAccess(id, user["email"].(string), user["name"].(string))
	return refreshToken, accessToken, nil
}