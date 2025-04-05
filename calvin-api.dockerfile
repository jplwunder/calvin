# syntax=docker/dockerfile:1

# pull official golang base image
FROM golang:1.24

# Install air for live reloading
RUN go install github.com/air-verse/air@latest

# set the Current Working Directory inside the container
WORKDIR /app

# copy the go mod and sum files into the WORKDIR
COPY go.mod go.sum ./

# download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Command to run air for live reloading
CMD ["air", "-c", ".air.toml"]

EXPOSE 8080