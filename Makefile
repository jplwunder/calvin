include .env

build:
	go build -o ${BINARY} ./cmd/api/

start:
	./${BINARY}

restart: build start