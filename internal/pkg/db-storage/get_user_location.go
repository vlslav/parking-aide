package db_storage

import "context"

func (s *Storage)GetUserLocation(ctx context.Context, userID int64) (*UserLocation, error) {

	return &UserLocation{}, nil
}

