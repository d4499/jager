package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AuthRoutes struct {
	a AuthService
}

func NewAuthRoutes(a AuthService) *AuthRoutes {
	return &AuthRoutes{
		a: a,
	}
}

func (a *AuthRoutes) Register(r *chi.Mux) {
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/magic", a.handleSendMagicLink)
		r.Post("/magic/verify", a.handleVerifyMagicLink)
	})
}

type MagicLinkLoginParams struct {
	Email string `json:"email"`
}

func (a *AuthRoutes) handleSendMagicLink(w http.ResponseWriter, r *http.Request) {
	var magicLinkLoginParams MagicLinkLoginParams
	if err := json.NewDecoder(r.Body).Decode(&magicLinkLoginParams); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}
	err := a.a.SendMagicLink(magicLinkLoginParams.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("magic link sent"))
}

type VerifyMagicLinkParams struct {
	Token string `json:"token"`
}

func (a *AuthRoutes) handleVerifyMagicLink(w http.ResponseWriter, r *http.Request) {
	var verifyMagicLinkParams VerifyMagicLinkParams
	if err := json.NewDecoder(r.Body).Decode(&verifyMagicLinkParams); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	u, err := a.a.VerifyMagicLink(verifyMagicLinkParams.Token)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s, err := a.a.CreateSession(u.ID)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// set session cookie
	c := a.a.CreateSessionCookie(s.ID)
	http.SetCookie(w, &c)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successfully verified magic link"))
}
