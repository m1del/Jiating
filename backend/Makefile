# Makefile for a Go server and postgresdb with Docker

# start Docker containers
run:
	docker compose up
run-d:
	docker compose up -d

# stop Docker containers (if in detached mode)
down:
	docker compose down

# rebuild and start Docker containers
rebuild:
	docker-compose up -d --build

# test the application inside Docker
test:
	docker compose exec server go test ./tests -v

# Clean up
clean:
	docker-compose down
	docker system prune -f
	docker volume prune -f

.PHONY: up up-d down rebuild test clean
