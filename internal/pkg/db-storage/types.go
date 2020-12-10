package db_storage

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"time"
)

var ErrNotFound = errors.New("not found")

type (
	Storage struct {
		db *sqlx.DB
	}

	User struct {
		UserID       int64     `db:"user_id"`
		RegisteredAt time.Time `db:"registered_at"`
	}

	UserLocation struct {
		UserID    int64   `db:"user_id"`
		Latitude  float64 `db:"latitude"`
		Longitude float64 `db:"longitude"`
	}
)
