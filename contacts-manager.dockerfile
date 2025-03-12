# syntax=docker/dockerfile:1

# pull official golang base image
FROM golang:1.24

# Install reflex for hot reloading
RUN go install github.com/cespare/reflex@latest

# set the Current Working Directory inside the container
WORKDIR /app

# copy the go mod and sum files into the WORKDIR
COPY go.mod go.sum ./

# download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Command to run reflex for hot reloading
CMD ["reflex", "-r", "\\.go$$", "-s", "--", "sh", "-c", "go run ."]

EXPOSE 8080