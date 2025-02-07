package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, u *User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := s.db.QueryRowContext(ctx,
		query,
		u.Username, u.Email, u.Password,
	).Scan(&u.ID, &u.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
