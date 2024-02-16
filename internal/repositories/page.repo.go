package repositories

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	m "tap/internal/models"
)

type PageRepo struct {
	*mongo.Collection
}

func NewPageRepo(db *mongo.Database) *PageRepo {
	database := db.Collection("pages")


	indexModel := mongo.IndexModel{
		Keys:    bson.M{"page_id": 1}, 
		Options: options.Index().SetUnique(true),
	}

	_, err := database.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal(err)
	}

	indexModel = mongo.IndexModel{
		Keys:    bson.M{"address": 1}, 
		Options: options.Index().SetUnique(true),
	}

	_, err = database.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatal(err)
	}
	return &PageRepo{
		database,
	}
}

func (repo *PageRepo) GetAll( userID primitive.ObjectID ) ([]m.PageRequest, error) {
	var pages []m.PageRequest
	cursor, err := repo.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &pages); err != nil {
		return nil, err
	}
	return pages, nil
}

func (repo *PageRepo) CreateNewPage( page m.PageRequest ) error {
	_, err := repo.InsertOne(context.Background(), page)
		if err != nil {
				return err
		}
	return nil
}

func (repo *PageRepo) UpdatePage( page m.PageRequest ) error {
	_, err := repo.UpdateOne(context.Background(), bson.M{"page_id": page.ID}, bson.M{"$set": page})
		if err != nil {
				return errors.New("Error on update page")
		}
	return nil
}

func (repo *PageRepo) GetPagesByUserID( userID primitive.ObjectID ) ([]m.PageRequest, error) {
	var pages []m.PageRequest
	cursor, err := repo.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &pages); err != nil {
		return nil, err
	}
	return pages, nil
}

func (repo *PageRepo) GetPageByID( pageID string ) (m.PageRequest, error) {
	var page m.PageRequest
	err := repo.FindOne(context.Background(), bson.M{"page_id": pageID}).Decode(&page)
	if err != nil {
		return m.PageRequest{}, err
	}
	return page, nil
}

func (repo *PageRepo) DeletePage( pageID string ) error {
	_, err := repo.DeleteOne(context.Background(), bson.M{"page_id": pageID})
		if err != nil {
				return errors.New("Error on delete page")
		}
	return nil
}

func (repo *PageRepo) GetByAddress( address string ) (m.PageRequest, error) {
	var page m.PageRequest
	err := repo.FindOne(context.Background(), bson.M{"address": address}).Decode(&page)
	if err != nil {
		return m.PageRequest{}, err
	}
	return page, nil
}

func(repo *PageRepo) CheckAddressExists(address string) (bool, error) {
	
	ctx := context.TODO()

	// Поиск документа с заданным адресом
	filter := bson.M{"address": address}
	var result m.PageRequest
	err := repo.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
			// Документ с таким адресом не найден
			return false, nil
	} else if err != nil {
			// Произошла ошибка при выполнении запроса к базе данных
			return false, err
	}

	// Документ с таким адресом найден
	return true, nil
}