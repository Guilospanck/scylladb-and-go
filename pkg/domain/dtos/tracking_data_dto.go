package dtos

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type TrackingDataDTO struct {
	FirstName       string    `json:"firstName" validate:"required"`
	LastName        string    `json:"lastName" validate:"required"`
	Timestamp       time.Time `json:"timestamp" validate:"required"`
	Location        string    `json:"location"`
	Speed           float64   `json:"speed"`
	Heat            float64   `json:"heat"`
	TelepathyPowers int       `json:"telepathyPowers"`
}

type TrackingDataPrimaryKeyDTO struct {
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Timestamp time.Time `json:"timestamp" validate:"required"`
}

type TrackingDataPartitionKeyDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

func ParseJson(data []byte, dto interface{}, dtoName string) error {
	err := json.Unmarshal(data, &dto)
	if err != nil {
		return err
	}

	err = isValid(dto, dtoName)
	if err != nil {
		return err
	}

	return nil
}

func isValid(dto interface{}, dtoName string) error {
	v := validator.New()
	err := v.Struct(dto)
	if err != nil {
		return fmt.Errorf("Error during %s validation: %s", dtoName, err.Error())
	}

	return nil
}
