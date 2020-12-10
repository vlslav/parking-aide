package db_storage

import "github.com/jmoiron/sqlx"

func New(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}
