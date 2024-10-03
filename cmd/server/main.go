package main

import (
	"log"
	"net/http"
	"os"
	"time"

	postgres "github.com/d4499/jager/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
}

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DATABSE_URL")

	pool := postgres.NewPostgres(dbUrl)

	srv := newServer(serverConfig{
		addr: ":8080",
		pool: pool,
	})

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Unable to start server: %s", err)
	}

	return nil
}

type serverConfig struct {
	addr string
	pool *pgxpool.Pool
}

func newServer(conf serverConfig) *http.Server {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(middleware.Logger)

	registerRoutes(r)

	return &http.Server{
		Addr:         conf.addr,
		Handler:      r,
		IdleTimeout:  time.Minute,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
	}
}
