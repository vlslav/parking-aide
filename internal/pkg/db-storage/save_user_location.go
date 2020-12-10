package db_storage

import (
	"context"

	"github.com/pkg/errors"
)

func (s *Storage) SaveUserLocation(ctx context.Context, userLocation *UserLocation) error {
	if _, err := s.db.ExecContext(ctx, `INSERT INTO user_location (user_id, latitude, longitude) VALUES($1, $2, $3);`,
		userLocation.UserID,    // $1
		userLocation.Latitude,  // $2
		userLocation.Longitude, // $3
	); err != nil {
		return errors.Wrap(err, "can't insert user location")
	}

	return nil
}
