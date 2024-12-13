BINARY_NAME=inventory_management
GO_ENV=development
VERSION=1.0.0

#Target
.PHONY: all run air build clean env

all: build

migrate: env
	go run cmd/api/main.go migrate

run: env
	go run cmd/api/main.go start

run-air: env
	air start

build: env
	go build -o $(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

env:
	@echo "Environment: $(GO_ENV)"
	@echo "Version: $(VERSION)"