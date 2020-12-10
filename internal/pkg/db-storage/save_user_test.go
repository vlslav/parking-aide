package db_storage

import (
	"context"
	"time"
)

func (s *TestStorageSuite) TestSaveUser() {
	ctx := context.Background()
	const userID int64 = 42

	s.Run("Пользователь успешно сохранен", func() {
		err := s.storage.SaveUser(ctx, &User{
			UserID:       userID,
			RegisteredAt: time.Now(),
		})
		s.Require().NoError(err)

		user := s.getUser()
		s.EqualValues(userID, user["user_id"])
	})
}

func (s *TestStorageSuite) getUser() map[string]interface{} {
	user := make(map[string]interface{})
	s.db.Get(user, `SELECT * FROM users;`)

	return user
}
