package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(conn string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), conn)
	if err != nil {
		log.Fatalf("Unable to connect to postgres: %v", err)
	}

	return pool
}
