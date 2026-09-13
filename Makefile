include .env
export

.PHONY: up down run test

up:
	docker compose up -d --wait

down:
	docker compose down

run:
	go run ./cmd/api

test:
	go test -v ./... -rase

migrate: up
	goose -dir migrations postgres "$(DB_URL)" up

migrate-down:
	goose -dir migrations postgres "$(DB_URL)" down