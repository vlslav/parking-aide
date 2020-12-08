package db_storage

import (
	"context"
	"time"
)

func (s *Storage)GetUser(ctx context.Context, userID int64) (*User, error) {

	return &User{
		UserID:       42,
		RegisteredAt: time.Now(),
	}, nil
}
