package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/d4499/jager/internal/database/db"
	"github.com/d4499/jager/internal/email"
	"github.com/d4499/jager/internal/user"
	"github.com/jackc/pgx/v5/pgtype"
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

func isSessionExpired(timestamp pgtype.Timestamp) bool {
	return time.Now().After(timestamp.Time)
}

func (a *AuthService) SessionMiddleware(authSvc *AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sessionCookie, err := r.Cookie("jager_session")
			if err != nil {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			session, err := a.GetSession(sessionCookie.Value)
			exp := isSessionExpired(session.ExpiresAt)
			if exp {
				fmt.Println("session expired")
				next.ServeHTTP(w, r)
				return
			}

			if err != nil {
				fmt.Println("unable to retrieve session: %w", err)
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), "session", session)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
