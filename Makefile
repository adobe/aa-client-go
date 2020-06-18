.PHONY: all
all: fmt vet lint test

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: vet
vet:
	@go vet -all ./...

.PHONY: lint
lint:
	@golint ./...

.PHONY: test
test:
	@go test ./analytics -cover

.PHONY: coverage
coverage:
	@go test -coverprofile=coverage.out ./analytics
	@go tool cover -html=coverage.out
