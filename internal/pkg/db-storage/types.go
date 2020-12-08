package db_storage

import (
	"time"
)

type (
	Storage struct {

	}

	User struct {
		UserID int64
		RegisteredAt time.Time
	}

	UserLocation struct {
		UserID int64
		Latitude float64
		Longitude float64
	}
)
