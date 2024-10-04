package user

import (
	"context"

	"github.com/d4499/jager/internal/database/db"
)

type UserService struct {
	db *db.Queries
}

func NewUserService(db *db.Queries) *UserService {
	return &UserService{
		db: db,
	}
}

func (u *UserService) GetUserByEmail(email string) (db.User, error) {
	user, err := u.db.GetUserByEmail(context.Background(), email)
	if err != nil {
	}

	return user, nil
}
