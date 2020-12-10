package db_storage

import (
	"context"
)

func (s *TestStorageSuite) TestGetUserLocation() {
	ctx := context.Background()
	const userID int64 = 42
	const latitude, longitude float64 = 0.123, 0.234

	s.Run("Локация существует", func() {
		s.prepareUserLocation(userID, latitude, longitude)

		userLocation, err := s.storage.GetUserLocation(ctx, userID)
		s.Require().NoError(err)

		s.EqualValues(userID, userLocation.UserID)
		s.EqualValues(longitude, userLocation.Longitude)
		s.EqualValues(latitude, userLocation.Latitude)
	})

	s.Run("Локации не существует", func() {
		userLocation, err := s.storage.GetUserLocation(ctx, 43)
		s.Nil(userLocation)

		s.EqualValues(ErrNotFound, err)
	})
}

func (s *TestStorageSuite) prepareUserLocation(userID int64, latitude, longitude float64) {
	s.db.Exec(`INSERT INTO user_location (user_id, latitude, longitude) VALUES ($1, $2, $3)`, userID, latitude, longitude)
}
