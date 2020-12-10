package db_storage

import (
	"context"

	"github.com/pkg/errors"
)

func (s *Storage) DeleteUserLocation(ctx context.Context, userLocation *UserLocation) error {
	if _, err := s.db.ExecContext(ctx, `DELETE FROM user_location WHERE user_id=$1;`,
		userLocation.UserID, // $1
	); err != nil {
		return errors.Wrap(err, "can't delete user location")
	}

	return nil
}
