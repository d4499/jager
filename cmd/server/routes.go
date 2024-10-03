package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func registerRoutes(r *chi.Mux) {
	r.Get("/healthz", healthcheck)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok"))
}
