package models

import "time"

type Metric struct {
	PageID string    `json:"page_id" bson:"page_id"`
	Date   time.Time `json:"date" bson:"date"`
	Type   string    `json:"type" bson:"type"`
	Value  string    `json:"value" bson:"value"`
}
