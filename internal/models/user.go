package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RegisterBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	RefreshToken string `json:"access_token"`
}

type RegisterResponse struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	RefreshToken string `json:"access_token"`
}

type LoginBody struct {
	Email string `json:"email"`
	Password string `json:"password"`
}




