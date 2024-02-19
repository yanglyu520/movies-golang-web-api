# include variables from .envrc file
include .envrc

# ======================================================== #
# HELPERS
# ======================================================== #
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ======================================================== #
# DEVELOPMENT
# ======================================================== #
.PHONY: swag/fmt
swag/fmt:
	@swag fmt -d cmd/api,pkg

.PHONY: swag/init
swag/init:
	@swag init -g cmd/api/main.go -o cmd/api/dist --ot yaml

.PHONY: run
run:
	cd cmd/api && go build . && ./api

.PHONY: build
build:
	cd cmd/api && go build .


.PHONY: all
all: swag/fmt swag/init run
# ======================================================== #
# Quality Control
# ======================================================== #
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	staticcheck ./...

.PHONY: lint
lint:
	@golangci-lint run --disable-all --enable gci --fix
	@golangci-lint run
