.PHONY: migrate-up, migrate-down, run

migrate-up:
	migrate -path db/migrations -database $(DATABASE_URL) up

migrate-down:
	migrate -path db/migrations -database $(DATABASE_URL) down 1

run:
	go run cmd/app/main.go

test:
	go test ./...
