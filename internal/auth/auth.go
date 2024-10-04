package auth

import (
	"github.com/d4499/jager/internal/database/db"
	"github.com/d4499/jager/internal/email"
	"github.com/d4499/jager/internal/user"
)

type AuthService struct {
	db    *db.Queries
	email email.EmailClient
	uSvc  user.UserService
}

func NewAuthService(db *db.Queries, email email.EmailClient, uSvc user.UserService) AuthService {
	return AuthService{
		db:    db,
		email: email,
		uSvc:  uSvc,
	}
}
