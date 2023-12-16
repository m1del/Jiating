FROM golang:1.19.0

# Set working directory using linux file conventions
WORKDIR /usr/src/app

# Package for hot reloading go file
RUN go install github.com/cosmtrek/air@latest

# Copy stuff from home directory to docker container
COPY . .

# Make sure go dependencies installed correctly
RUN go mod tidy