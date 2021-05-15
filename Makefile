.PHONY: build
build:
	go build -v ./cmd/main.go


.PHONY: migrate
migrate:
	migrate -path ./schema -database 'postgres://postgres:password@localhost:5436/postgres?sslmode=disable' up


.PHONY: run
run:
	go run ./cmd/main.go


.PHONY: swag
swag:
	swag init -g cmd/main.go

.DEFAULT_GOAL := run

