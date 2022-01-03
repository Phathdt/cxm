package auth

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, username string, password string) error
	GetUser(ctx context.Context, id int64) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}
