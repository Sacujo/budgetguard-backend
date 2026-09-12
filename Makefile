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