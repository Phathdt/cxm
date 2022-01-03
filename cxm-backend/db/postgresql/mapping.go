package postgresql

import (
	"cxm-auth/auth"
)

func (u User) MapToEntity() auth.User {
	return auth.User{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
