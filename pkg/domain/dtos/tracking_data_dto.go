package dtos

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type Timestamp time.Time

type TrackingDataDTO struct {
	FirstName       string     `json:"firstName" validate:"required"`
	LastName        string     `json:"lastName" validate:"required"`
	Timestamp       *Timestamp `json:"timestamp" validate:"required"`
	Location        string     `json:"location"`
	Speed           float64    `json:"speed"`
	Heat            float64    `json:"heat"`
	TelepathyPowers int        `json:"telepathyPowers"`
}

type TrackingDataPrimaryKeyDTO struct {
	FirstName string     `json:"firstName" validate:"required"`
	LastName  string     `json:"lastName" validate:"required"`
	Timestamp *Timestamp `json:"timestamp" validate:"required"`
}

type TrackingDataPartitionKeyDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

func (timestamp *Timestamp) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}

	const timeLayout = "2006-01-02 15:04:05 -0700 MST"
	t, err := time.Parse(timeLayout, value)
	if err != nil {
		return err
	}

	*timestamp = Timestamp(t)
	return nil
}

func ParseJson(data []byte, dto interface{}, dtoName string) error {
	err := json.Unmarshal(data, dto)
	if err != nil {
		return fmt.Errorf("[ParseJson] Error: %s\n", err.Error())
	}

	err = isValid(dto, dtoName)
	if err != nil {
		return fmt.Errorf("[isValid] Error: %s\n", err.Error())
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
