package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Users *UserRepo
	Pages *PageRepo
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Users: NewUserRepo(db),
		Pages: NewPageRepo(db),
	}
}
