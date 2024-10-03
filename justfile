set dotenv-load

migrate:
  cd ./db/migrations && goose postgres $DATABASE_URL up 

generate:
  cd ./internal/database && sqlc generate
