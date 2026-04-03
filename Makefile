include .env

build:
	go build -o out

run:
	go run .

generate:
	sqlc generate

migrate-up:
	goose -dir sql/schema postgres $(DB_URL) up

migrate-down:
	goose -dir sql/schema postgres $(DB_URL) down

.PHONY: build run generate migrate-up migrate-down
