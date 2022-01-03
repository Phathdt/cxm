package auth

import "context"

type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	GetUser(ctx context.Context, username string) (*User, error)
}
