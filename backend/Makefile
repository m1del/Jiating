# Makefile for a Go server and postgresdb with Docker

# environment variables
include .env
export

# start Docker containers
up:
	@docker compose up
up-d:
	@docker compose up -d

# connect to database
db:
	@PGPASSWORD=${DB_PASSWORD} psql -h localhost -p ${DB_PORT} -U ${DB_USERNAME} -d ${DB_DATABASE}

# connect to server
server:
	@docker exec -it jiating_server bash

# stop Docker containers (if in detached mode)
down:
	@docker compose down

# rebuild and start Docker containers
rebuild:
	@docker-compose up -d --build

# test the application inside Docker
test:
	@docker compose exec server go test ./tests -v

# Clean up
clean:
	@docker compose down
	@docker system prune -f
	@docker volume prune -f

.PHONY: up up-d down db server rebuild test clean
