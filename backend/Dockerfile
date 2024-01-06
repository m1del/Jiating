FROM golang:1.21-alpine3.19

WORKDIR /app

# install bash
RUN apk add --no-cache bash

# copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application's code
COPY . .

# Install Air for live reloading
RUN go install github.com/cosmtrek/air@latest

# Expose port (e.g., 8080) to the outside world
EXPOSE 3000

# Command to run the application
CMD ["air"]
