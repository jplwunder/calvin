# syntax=docker/dockerfile:1

# pull official golang base image
FROM golang:1.24

# set the Current Working Directory inside the container
WORKDIR /app

# copy the go mod and sum files into the WORKDIR
COPY go.mod go.sum ./

# download all dependencies
RUN go mod download

# copy all files with the go extension from the current directory to the WORKDIR inside the container
COPY *.go ./

# compile application in a static application binary named contacts-manager
RUN go build -o /contacts-manager

# tell docker what command to run when image is used to start a container
CMD ["/contacts-manager"]

EXPOSE 8080