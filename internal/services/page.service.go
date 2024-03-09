package services

import (
	m "tap/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) CheckAddress(address string) (bool, error) {
	return s.repo.Pages.CheckAddressExists(address)
}

func (s *Service) GetPages(userId string) ([]m.PageRequest, error) {
	
	userIdobject , err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	
	return s.repo.Pages.GetAll(userIdobject)
}

func (s *Service) GetPageByAddress(address string) (m.PageRequest, error) {
	return s.repo.Pages.GetByAddress(address)
}

func (s *Service) CreatePage(page m.PageFromBody, userId string) error {
	HexId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	requestPage := m.PageRequest{
		ID: page.ID,
		Title: page.Title,
		Address: page.Address,
		Theme: page.Theme,
		Favicon: page.Favicon,
		Bricks: page.Bricks,
		User: HexId,
	}
	
	return s.repo.Pages.CreateNewPage(requestPage)
}

func (s *Service) UpdatePage(page m.PageRequest) error {
	
	return s.repo.Pages.UpdatePage(page)
}

func (s *Service) DeletePage(address string, userId string) error {
	HexId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	return s.repo.Pages.DeletePage(address, HexId)
}

func (s *Service) UpdatePageMeta(page m.PageMetaData, userId string) error {
	return s.repo.Pages.UpdatePageMeta(page)
}