package services

import (
	"tap/internal/libs/primitive"
	m "tap/internal/models"
)

func (s *Service) CheckAddress(address string) (bool, error) {
	return s.repo.Pages.CheckAddressExists(address)
}

func (s *Service) GetPages(userId string) ([]m.PageRequest, error) {
	userIdobject , err := primitive.GetObject(userId)
	if err != nil {
		return nil, err
	}
	
	return s.repo.Pages.GetAll(userIdobject)
}

func (s *Service) GetPageByAddress(address string) (m.PageRequest, error) {
	return s.repo.Pages.GetByAddress(address)
}

func (s *Service) CreatePage(page m.PageFromBody, userId string) error {
	HexId, err := primitive.GetObject(userId)
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

func (s *Service) UpdatePage(page m.PageFromBody, userId string) error {
	HexId, err := primitive.GetObject(userId)
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
	
	return s.repo.Pages.UpdatePage(requestPage)
}

func (s *Service) UpdatePageMeta(page m.PageMetaData, userId string) error {
	return s.repo.Pages.UpdatePageMeta(page)
}