package entities

import "time"

type TrackingDataEntity struct {
	FirstName       string    `db:"first_name"`
	LastName        string    `db:"last_name"`
	Timestamp       time.Time `db:"timestamp"`
	Location        string    `db:"location"`
	Speed           float64   `db:"speed"`
	Heat            float64   `db:"heat"`
	TelepathyPowers int       `db:"telepathy_powers"`
}
