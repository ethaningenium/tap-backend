package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	m "tap/internal/models"
)

type MetricRepo struct {
	*mongo.Collection
}

func NewMetricRepo(db *mongo.Database) *MetricRepo {
	database := db.Collection("metrics")
	return &MetricRepo{
		database,
	}
}

func (repo *MetricRepo) CreateMetric(metric m.Metric) error {
	_, err := repo.InsertOne(context.Background(), metric)
	if err != nil {
		return err
	}
	return nil
}
