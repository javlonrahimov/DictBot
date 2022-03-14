# Include variables from the .envrc file
include .envrc


# ==================================================================================== #
# HELPERS
# ==================================================================================== #


## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #


## run/api: run the cmd/api application
.PHONY: run/bot
run/bot:
	go run ./cmd/bot -db-dsn=${DICT_DB_DSN}?sslmode=disable -bot-token=${DICT_BOT_TOKEN}

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	psql ${DICT_DB_DSN}?sslmode=disable

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations'
	migrate -path ./migrations -database ${DICT_DB_DSN}?sslmode=disable up


# ==================================================================================== #
# BUILD
# ==================================================================================== #


current_time = $(shell date --iso-8601=seconds)
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags=${linker_flags} -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/api ./cmd/api


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #


## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...


## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor
