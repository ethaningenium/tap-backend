package services

import (
	m "tap/internal/models"
)

func (s *Service) GetPageByAddress(address string) (m.PageRequest, error) {
	return s.repo.Pages.GetByAddress(address)
}

func (s *Service) CreatePage(page m.PageRequest) error {
	return s.repo.Pages.CreateNewPage(page)
}