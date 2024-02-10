package services

import (
	"tap/internal/repositories"
)


type Service struct {
	repo *repositories.Repository
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{
		repo: repository,
	}
}

