.PHONY: dev
dev:
	go run cmd/beatgo/main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build -o bin/beatgo cmd/beatgo/main.go
