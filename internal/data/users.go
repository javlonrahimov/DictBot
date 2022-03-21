package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrDuplicateUser = errors.New("duplicate user")
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	Version   string `json:"version"`
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Insert(user *User) error {
	query := `
	insert into users (id, first_name, last_name, username)
	values ($1, $2, $3, $4)
	returning created_at`

	args := []interface{}{user.ID, user.FirstName, user.LastName, user.Username}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.CreatedAt)
	if err != nil {
		switch {
		case err.Error() == "pq: duplicate key value violates unique constraint \"users_pkey\"":
			return ErrDuplicateUser
		default:
			return err
		}
	}

	return nil
}
