package authuc

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"cxm-auth/module/auth"

	"github.com/dgrijalva/jwt-go/v4"
)

type AuthUseCase struct {
	userRepo       auth.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(userRepo auth.UserRepository, hashSalt string, signingKey []byte, expireDuration time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration * time.Second,
	}
}

func (au *AuthUseCase) SignUp(ctx context.Context, username, password string) error {
	user, _ := au.userRepo.GetUserByUsername(ctx, username)
	if user != nil {
		return auth.ErrUserExist
	}

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(au.hashSalt))

	return au.userRepo.CreateUser(ctx, username, fmt.Sprintf("%x", pwd.Sum(nil)))
}

func (au *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	user, err := au.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(au.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	if password != user.Password {
		return "", auth.ErrPasswordNotMatch
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(au.expireDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(au.signingKey)
}

func (au *AuthUseCase) GetUser(ctx context.Context, username string) (*auth.User, error) {
	return au.userRepo.GetUserByUsername(ctx, username)
}
