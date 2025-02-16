include .envrc

# ==================================================================================== # 
# HELPERS
# ==================================================================================== #

## help: print the help message
.PHONY: help
help:
	@echo 'help'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure?[y/N]' && read ans && [ $${ans:N} = y ]

# ==================================================================================== # 
# DEVELOPMENT
# ==================================================================================== #


## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api


## db/sql: connect the database using psql
.PHONY: db/sql
db/psql:
	psql ${GREENLIGHT_DB_DSN}

## db/migration/new name=$1: create a new database migration
.PHONY: db/migration/new
db/migration/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migration/up: apply all up database migrations...
.PHONY: db/migration/up
db/migration/up: confirm
	@echo 'Running migrations...'
	migrate -path ./migrations -database ${GREENLIGHT_DB_DSN} up


# ==================================================================================== # 
# QUALITY CONTROL
# ==================================================================================== #

.PHONY: audit
audit: vendor
	@echo "Formatting code..."
	go fmt ./...
	@echo "Vetting code..."
	go vet ./...
	staticcheck ./...
	@echo "Running tests..."
	go test -race -vet=off ./...

.PHONY: vendor
vendor:
	@echo "Tidying and verifying module dependencies..."
	go mod tidy
	go mod verify
	@echo "Vendoring dependencies..."
	go mod vendor


# ==================================================================================== # 
# BUILD
# ==================================================================================== #

current_time = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

## build/api: build the cmd/api application

.PHONY: build/api
build/api:
	@echo "Building cmd/api..."
	go build -ldflags=${linker_flags} -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/api ./cmd/api


# ==================================================================================== # 
# PRODUCTION
# ==================================================================================== #

production_host_ip = '3.93.241.94'

## production/connect: connect to the production server
.PHONY: production/connect
production/connect:
	ssh -i ~/.ssh/greenlight.pem -t greenlight@${production_host_ip}

## production/deploy/api: deploy the api to production
.PHONY: production/deploy/api
production/deploy/api:
	rsync -e "ssh -i ~/.ssh/greenlight.pem"  -P ./bin/linux_amd64/api greenlight@${production_host_ip}:~
	rsync -e "ssh -i ~/.ssh/greenlight.pem"  -rP --delete ./migrations greenlight@${production_host_ip}:~
	rsync -e "ssh -i ~/.ssh/greenlight.pem" -P ./remote/production/api.service greenlight@${production_host_ip}:~
	rsync -e "ssh -i ~/.ssh/greenlight.pem" -P ./remote/production/Caddyfile greenlight@${production_host_ip}:~
	ssh -i ~/.ssh/greenlight.pem -t greenlight@${production_host_ip} \
	'migrate -path ~/migrations -database $$GREENLIGHT_DB_DSN up \
	&& sudo mv ~/api.service /etc/systemd/system/ \
	&& sudo systemctl enable api \
	&& sudo systemctl restart api \
	&& sudo mv ~/Caddyfile /etc/caddy \
	&& sudo systemctl reload caddy \
	'