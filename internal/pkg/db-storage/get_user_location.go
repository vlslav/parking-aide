package db_storage

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

func (s *Storage) GetUserLocation(ctx context.Context, userID int64) (*UserLocation, error) {
	userLocation := &UserLocation{}
	if err := s.db.GetContext(ctx, userLocation,
		`SELECT user_id, latitude, longitude FROM user_location WHERE user_id=$1;`,
		userID, // $1
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "can't select user location")
	}

	return userLocation, nil
}
