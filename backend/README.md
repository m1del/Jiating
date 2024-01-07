# Jiating Backend

This backend is part of the Jiating website and is designed to work with the client built using React. It's built with Go, using the go-chi framework for the HTTP services (the REST API) with a PostgresQL database. It is dockerized for easy development and deployment.

## Getting Started

These instructions will help you set up the backend services on your local machine for development and testing purposes.

### Prequisites
- Docker
- Docker Compose
#### Optional
- Go version 1.21.5
- psql 15.5 (recommended)

## Using Docker for Development

This project uses Docker to simplify dependency management and to ensure a consistent development environment. Here's how to get started:


Start the application
```
make up
```
Start the application (in detached mode)
```
make up-d
```

Stop the application
```
make down
```

Rebuilding and restarting the application:
```
make rebuild
```

Running Tests
```
make test
```

Clean up (remove containers, networks, and volumes)
```
make clean
```

# Deployment

lol todo ;-;
