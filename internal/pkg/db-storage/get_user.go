package db_storage

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
)

func (s *Storage) GetUser(ctx context.Context, userID int64) (*User, error) {
	user := &User{}
	if err := s.db.GetContext(ctx, user,
		`SELECT user_id, registered_at FROM users WHERE user_id=$1`,
		userID, // $1
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.Wrap(err, "can't select user")
	}

	return user, nil
}
