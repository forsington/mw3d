build:
	go build -o bin/mw3d
run:
	go run ./...

.PHONY: lint
lint:
	golangci-lint run --timeout 3m --fix ./...

.PHONY: test
test:
	 go test ./... -coverprofile=coverage.out

.PHONY: ready
ready: lint test build
