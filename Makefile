include .env
export

.PHONY: up down run test migrate migrate-down

up:
	docker compose up -d --wait

down:
	docker compose down

run:
	go run ./cmd/api

test:
	go test -v ./... -race

migrate: up
	goose -dir migrations postgres "$(DB_URL)" up

migrate-down: up
	goose -dir migrations postgres "$(DB_URL)" down