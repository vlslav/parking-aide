package db_storage

import (
	"context"
	"time"
)

func (s *TestStorageSuite) TestGetUser() {
	ctx := context.Background()
	const userID int64 = 42

	s.Run("Пользователь существует", func() {
		s.prepareUser(userID)

		user, err := s.storage.GetUser(ctx, userID)
		s.Require().NoError(err)

		s.EqualValues(userID, user.UserID)
	})

	s.Run("Пользователя не существует", func() {
		user, err := s.storage.GetUser(ctx, 43)
		s.Nil(user)

		s.EqualValues(ErrNotFound, err)
	})
}

func (s *TestStorageSuite) prepareUser(userID int64) {
	s.db.Exec(`INSERT INTO users (user_id, registered_at) VALUES ($1,  $2)`, userID, time.Now())
}
