package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/d4499/jager/internal"
	"github.com/d4499/jager/internal/database/db"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *AuthService) CreateSession(userId string) (db.Session, error) {
	monthFromNow := time.Now().AddDate(0, 1, 0)
	session, err := s.db.CreateSession(context.Background(), db.CreateSessionParams{
		ID:     internal.NewCUID(),
		UserID: userId,
		ExpiresAt: pgtype.Timestamp{
			Time:  monthFromNow,
			Valid: true,
		},
	})
	if err != nil {
		return db.Session{}, fmt.Errorf("unable to create session: %w", err)
	}

	return session, nil
}

func (a *AuthService) CreateSessionCookie(value string) http.Cookie {
	maxAge := time.Hour * 24 * 30
	c := http.Cookie{
		Name:     "jager_session",
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
		MaxAge:   int(maxAge.Seconds()),
	}

	return c
}

func (a *AuthService) GetSession(id string) (db.Session, error) {
	session, err := a.db.GetSession(context.Background(), id)
	if err != nil {
		return db.Session{}, fmt.Errorf("unable to get session: %w", err)
	}

	return session, nil
}
