package db_storage

import "context"

func (s *TestStorageSuite) TestDeleteUserLocation() {
	ctx := context.Background()
	const userID int64 = 42
	const latitude, longitude float64 = 0.123, 0.234

	s.Run("Локация успешно удалена", func() {
		s.prepareUserLocation(userID, latitude, longitude)

		err := s.storage.DeleteUserLocation(ctx, &UserLocation{
			UserID:    userID,
			Latitude:  latitude,
			Longitude: longitude,
		})
		s.Require().NoError(err)

		count := s.getUserLocationCount()
		s.EqualValues(0, count)
	})
}

func (s *TestStorageSuite) getUserLocationCount() int64 {
	count := int64(0)
	s.db.Get(&count, `SELECT COUNT(*) FROM user_location;`)

	return count
}
