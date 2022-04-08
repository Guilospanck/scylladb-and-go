package dtos

import "time"

type TrackingDataDTO struct {
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Timestamp       time.Time `json:"timestamp"`
	Location        string    `json:"location"`
	Speed           float64   `json:"speed"`
	Heat            float64   `json:"heat"`
	TelepathyPowers int       `json:"telepathyPowers"`
}

type TrackingDataPrimaryKeyDTO struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Timestamp time.Time `json:"timestamp"`
}

type TrackingDataPartitionKeyDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
