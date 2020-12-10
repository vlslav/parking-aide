package db_storage

import (
	"context"

	"github.com/pkg/errors"
)

func (s *Storage) SaveUser(ctx context.Context, user *User) error {
	if _, err := s.db.ExecContext(ctx, `INSERT INTO user (user_id, registered_at) VALUES($1, $2) ON CONFLICT(user_id) DO NOTHING;`,
		user.UserID,       // $1
		user.RegisteredAt, // $2

	); err != nil {
		return errors.Wrap(err, "can't insert user")
	}

	return nil
}
