package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PageFromBody struct {
	ID string `json:"id" bson:"page_id"`
	Title string `json:"title" bson:"title"`
	Address string `json:"address" bson:"address"`
	Bricks []Brick `json:"bricks" bson:"bricks"`
}


type PageRequest struct {
	ID string `json:"id" bson:"page_id"`
	Title string `json:"title" bson:"title"`
	Address string `json:"address" bson:"address"`
	Bricks []Brick `json:"bricks" bson:"bricks"`
	User primitive.ObjectID `json:"user" bson:"user_id"`
}