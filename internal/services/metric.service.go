package services

import (
	m "tap/internal/models"
)

func (s *Service) CreateMetric(metric m.Metric) error {
	return s.repo.Metrics.CreateMetric(metric)
}
