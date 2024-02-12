package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	m "tap/internal/models"
)

type UserRepo struct {
	*mongo.Collection
}

func NewUserRepo(db *mongo.Database) *UserRepo {
	database := db.Collection("users")

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1}, 
		Options: options.Index().SetUnique(true),
	}

	_, err := database.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal(err)
	}
	return &UserRepo{
		database,
	}
}


func (repo *UserRepo) CreateNewUser( user m.RegisterResponse )  (string ,error) {
	_ , err := repo.InsertOne(context.Background(), user)
  if err != nil {
		if mongo.IsDuplicateKeyError(err){
			return "", errors.New("User already exists")
		}
    return "", errors.New("Error on create user")
  }
	
	return user.ID.Hex(), nil
}

func (repo *UserRepo) GetUserByEmail(email string) (m.RegisterResponse, error) {
	var user m.RegisterResponse
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
			if err == mongo.ErrNoDocuments {
					// Пользователь с указанным email не найден
					return m.RegisterResponse{}, errors.New("user not found")
			}
			// Произошла ошибка при выполнении запроса к базе данных
			return m.RegisterResponse{}, err
	}
	
	return user, nil
}


func (repo *UserRepo) SetNewRefreshToken(email string, refreshToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"refreshtoken": refreshToken}})
	if err != nil {
			if err == mongo.ErrNoDocuments {
					return errors.New("user not found")
			}
			return err
	}
	return nil
}

func (repo *UserRepo) SetVerifiedTrue(verifyCode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.UpdateOne(ctx, bson.M{"verifycode": verifyCode}, bson.M{"$set": bson.M{"isemailverified": true}})

	if err != nil {
			fmt.Println(err)
			if err == mongo.ErrNoDocuments {
					return errors.New("user not found")
			}
			return err
	}
	return nil
}
