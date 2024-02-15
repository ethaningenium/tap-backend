package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RegisterBody struct {
	Name  string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type GoogleRegisterBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}


type RegisterResponse struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsEmailVerified bool `json:"isemailverified"`
	VerifyCode string `json:"verifycode"`
}

type LoginBody struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}




