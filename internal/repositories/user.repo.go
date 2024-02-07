package repositories

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	m "tap/internal/models"
)

type UserRepo struct {
	coll *mongo.Collection
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
		coll: database,
	}
}


func (repo *UserRepo) CreateNewUser( user m.UserRegisterRequest )  error {
	_, err := repo.coll.InsertOne(context.Background(), user)
    if err != nil {
        return errors.New("Error on create user")
    }
	return nil
}

func (repo *UserRepo) GetUserByEmail(email string) (m.UserRegisterRequest, error) {
	var user m.UserRegisterRequest
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
			if err == mongo.ErrNoDocuments {
					// Пользователь с указанным email не найден
					return m.UserRegisterRequest{}, errors.New("user not found")
			}
			// Произошла ошибка при выполнении запроса к базе данных
			return m.UserRegisterRequest{}, err
	}
	
	return user, nil
}


func (repo *UserRepo) SetNewRefreshToken(email string, refreshToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.coll.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"refreshtoken": refreshToken}})
	if err != nil {
			if err == mongo.ErrNoDocuments {
					return errors.New("user not found")
			}
			return err
	}
	return nil
}

func (repo *UserRepo) GetUserByRefreshToken(refreshToken string) (m.UserRegisterRequest, error) {
	var user m.UserRegisterRequest
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repo.coll.FindOne(ctx, bson.M{"refreshtoken": refreshToken}).Decode(&user)
	if err != nil {
			if err == mongo.ErrNoDocuments {
					// Пользователь с указанным email не найден
					return m.UserRegisterRequest{}, errors.New("user not found")
			}
			// Произошла ошибка при выполнении запроса к базе данных
			return m.UserRegisterRequest{}, err
	}
	
	return user, nil
}