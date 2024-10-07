set dotenv-load

migrate:
  cd ./db/migrations && goose postgres $DATABASE_URL up 

down:
  cd ./db/migrations && goose postgres $DATABASE_URL down

status:
  cd ./db/migrations && goose postgres $DATABASE_URL status

generate:
  cd ./internal/database && sqlc generate
