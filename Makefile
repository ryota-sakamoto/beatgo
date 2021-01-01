.PHONY: dev
dev:
	go run cmd/beatgo/main.go

.PHONY: test
test:
	go test -v ./...
