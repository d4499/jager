package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/d4499/jager/internal/auth"
	postgres "github.com/d4499/jager/internal/database"
	"github.com/d4499/jager/internal/database/db"
	"github.com/d4499/jager/internal/email"
	"github.com/d4499/jager/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/resend/resend-go/v2"
)

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
}

type serverConfig struct {
	addr  string
	pool  *pgxpool.Pool
	email *resend.Client
}

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	client := resend.NewClient(os.Getenv("RESEND_APIKEY"))

	pool := postgres.NewPostgres(dbUrl)

	srv := newServer(serverConfig{
		addr:  ":8080",
		pool:  pool,
		email: client,
	})

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Unable to start server: %s", err)
	}

	return nil
}

func newServer(conf serverConfig) *http.Server {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(middleware.Logger)

	db := db.New(conf.pool)
	client := email.NewEmailClient(conf.email)
	userSvc := user.NewUserService(db)
	authSvc := auth.NewAuthService(db, *client, *userSvc)
	auth.NewAuthRoutes(authSvc).Register(r)

	return &http.Server{
		Addr:         conf.addr,
		Handler:      r,
		IdleTimeout:  time.Minute,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
	}
}
