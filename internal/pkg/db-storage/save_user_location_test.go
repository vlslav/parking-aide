package db_storage

import (
	"context"
)

func (s *TestStorageSuite) TestSaveUserLocation() {
	ctx := context.Background()
	const userID int64 = 42
	const latitude, longitude float64 = 0.123, 0.234

	s.Run("Локация успешно сохранена", func() {
		err := s.storage.SaveUserLocation(ctx, &UserLocation{
			UserID:    userID,
			Latitude:  latitude,
			Longitude: longitude,
		})
		s.Require().NoError(err)

		userLocation := s.getUserLocation()
		s.EqualValues(userID, userLocation["user_id"])
		s.EqualValues(latitude, userLocation["latitude"])
		s.EqualValues(longitude, userLocation["longitude"])
	})
}

func (s *TestStorageSuite) getUserLocation() map[string]interface{} {
	userLocation := make(map[string]interface{})
	s.db.Get(userLocation, `SELECT * FROM user_location;`)

	return userLocation
}
