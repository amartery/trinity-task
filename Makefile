.PHONY: build
build:
	go build -v ./cmd/main.go


.PHONY: test
test:
	go test -v -race -timeout 30s ./...


.PHONY: run
run:
	go run ./cmd/main.go

.DEFAULT_GOAL := run

