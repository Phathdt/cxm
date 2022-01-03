package authpg

import (
	"context"
	"database/sql"

	"cxm-auth/auth"
	"cxm-auth/db/postgresql"
)

type UserRepository struct {
	q *postgresql.Queries
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{q: postgresql.New(db)}
}

func (r UserRepository) CreateUser(ctx context.Context, username string, password string) error {
	_, err := r.q.InsertUser(ctx, postgresql.InsertUserParams{
		Username: username,
		Password: password,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) GetUser(ctx context.Context, id int64) (*auth.User, error) {
	u, err := r.q.GetUser(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, auth.ErrUserNotFound
		}

		return nil, err
	}

	user := u.MapToEntity()

	return &user, nil
}

func (r UserRepository) GetUserByUsername(ctx context.Context, username string) (*auth.User, error) {
	u, err := r.q.GetUserByUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, auth.ErrUserNotFound
		}

		return nil, err
	}

	user := u.MapToEntity()

	return &user, nil
}
