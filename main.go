package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
}

func run() error {
	srv := newServer(serverConfig{
		addr: ":8080",
	})

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Unable to start server: %s", err)
	}

	return nil
}

type serverConfig struct {
	addr string
}

func newServer(conf serverConfig) *http.Server {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(middleware.Logger)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return &http.Server{
		Addr:         conf.addr,
		Handler:      r,
		IdleTimeout:  time.Minute,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
	}
}
