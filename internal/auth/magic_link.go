package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/d4499/jager/internal"
	"github.com/d4499/jager/internal/database/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/resend/resend-go/v2"
)

func GenerateRandomToken(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(b)
	return token, nil
}

func newTimestamp(d time.Duration) pgtype.Timestamp {
	now := time.Now()

	futureTime := now.Add(d)
	return pgtype.Timestamp{
		Time:  futureTime,
		Valid: true,
	}
}

func (a *AuthService) SendMagicLink(email string) error {
	t, err := GenerateRandomToken(32)
	if err != nil {
		log.Print("Unable to generate random token")
	}

	expiresAt := time.Minute * 15

	ml, err := a.db.CreateMagicLink(context.Background(), db.CreateMagicLinkParams{
		ID:        internal.NewCUID(),
		Email:     email,
		Token:     t,
		ExpiresAt: newTimestamp(expiresAt),
	})
	if err != nil {
		return fmt.Errorf("Unable to create magic link: %w", err)
	}

	magicLink := fmt.Sprintf("http://localhost:5173/verify?token=%v", ml.Token)

	htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Magic Link Login</title>
    <style>
        body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
        .container { max-width: 600px; margin: 0 auto; padding: 20px; }
        .button { display: inline-block; padding: 10px 20px; background-color: #007bff; color: #ffffff; text-decoration: none; border-radius: 5px; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Magic Link Login</h1>
        <p>Hello,</p>
        <p>You requested a magic link to log in to your account. Click the button below to log in:</p>
        <p>
            <a href="%s" class="button">Log In</a>
        </p>
        <p>If you didn't request this link, you can safely ignore this email.</p>
        <p>This link will expire in 15 minutes for security reasons.</p>
    </div>
</body>
</html>
`, magicLink)

	a.email.SendEmail(resend.SendEmailRequest{
		From:    "Acme <onboarding@resend.dev>",
		To:      []string{email},
		Html:    htmlContent,
		Subject: "Jager Magic Link",
	})

	return nil
}

func isMagicLinkValid(timestamp pgtype.Timestamp) bool {
	return time.Now().After(timestamp.Time)
}

func (a *AuthService) VerifyMagicLink(token string) (db.User, error) {
	// Verify token from magic link
	link, err := a.db.GetMagicLinkByToken(context.Background(), token)
	if err != nil {
		return db.User{}, fmt.Errorf("Unable to find magic magic link")
	}

	// Check if link is not expired
	valid := isMagicLinkValid(link.ExpiresAt)
	if !valid {
		return db.User{}, fmt.Errorf("Magic link has expired")
	}

	// Check if user already exists
	u, err := a.getOrCreateUser(link.Email)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to get user: %v", err)
	}

	// delete magic link
	err = a.db.DeleteMagicLink(context.Background(), link.ID)
	if err != nil {
		return db.User{}, fmt.Errorf("unable to delete magic link: %v", err)
	}

	return u, nil
}

func (a *AuthService) getOrCreateUser(email string) (db.User, error) {
	u, err := a.db.GetUserByEmail(context.Background(), email)
	if err == nil {
		return u, nil
	}

	newUser, err := a.db.CreateUser(context.Background(), db.CreateUserParams{
		ID:    internal.NewCUID(),
		Email: email,
	})
	if err != nil {
		return db.User{}, fmt.Errorf("failed to create user: %v", err)
	}

	return newUser, nil
}
