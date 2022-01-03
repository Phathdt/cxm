package auth

import "errors"

var (
	ErrUserExist          = errors.New("user exists")
	ErrPasswordNotMatch   = errors.New("password not match")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidAccessToken = errors.New("invalid access token")
)
