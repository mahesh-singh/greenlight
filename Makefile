-include .env
export

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
# prereqs
# ==================================================================================== #

# prereqs: check for docker, docker-compose and make
.PHONY: prereqs
prereqs:
	@echo "Checking prerequisites..."
	@command -v docker >/dev/null 2>&1 || { echo "Docker is not installed. Please install Docker first."; exit 1; }
	@command -v docker-compose >/dev/null 2>&1 || { echo "Docker Compose is not installed. Please install Docker Compose first."; exit 1; }
	@command -v make >/dev/null 2>&1 || { echo "Make is not installed. Please install Make first."; exit 1; }
	@echo "All prerequisites are installed."


# env/init: alidate environment configuration, create .env  if not exist
.PHONY: env/init
env/init:
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo "Created .env file from template. Please review and update the configurations."; \
	else \
		echo ".env file already exists."; \
	fi


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
# DOCKER
# ==================================================================================== #

## docker/build: Build docker image
.PHONY: docker/build
docker/build:
	@echo "Building docker images..."
	docker compose build

## docker/up: Start Docker container
.PHONY: docker/up
docker/up:
	@echo "Starting Docker containers..."
	docker compose up -d

## docker/down: Stop docker containers
.PHONY: docker/down
docker/down:
	@echo "Stopping Docker containers..."
	docker compose down

## docker/clean: Remove docker containers, network and volume
.PHONY: docker/clean
docker/clean:
	@echo "Removing Docker containers, networks, and volumes..."
	docker compose down -v


## docker/logs: Show logs of running containers
.PHONY: docker/logs
docker/logs:
	@echo "Show logs..."
	docker compose logs -f


## docker/db/start: Start DB container only
.PHONY: docker/db/start
docker/db/start:
	@echo "Starting DB container..."
	@if [ -z "$$(docker-compose ps -q db 2>/dev/null)" ] || [ -z "$$(docker ps -q --no-trunc | grep "$$(docker-compose ps -q db 2>/dev/null)" 2>/dev/null)" ]; then \
		docker compose up -d db; \
		echo "Waiting for DB to be ready..."; \
		sleep 5; \
	else \
		echo "DB container is already running."; \
	fi

## docker/migrate: Run database migrations in the container
.PHONY: docker/migrate
docker/migrate:
	@echo "Checking if migrations are already applied..."
	@if ! docker compose run --rm migrate -path ./migrations -database "${DB_DSN}" version | grep -q "No migrations"; then \
		echo "Running database migrations..."; \
		docker compose run --rm migrate -path ./migrations -database "${DB_DSN}" up; \
	else \
		echo "Migrations are already up-to-date."; \
	fi


## docker/api/start: Start REST API container with dependencies
.PHONY: docker/api/start
docker/api/start: docker/db/start docker/migrate docker/build
	@echo "Starting REST API container..."
	@if [ -z "$$(docker-compose ps -q api 2>/dev/null)" ] || [ -z "$$(docker ps -q --no-trunc | grep "$$(docker-compose ps -q api 2>/dev/null)" 2>/dev/null)" ]; then \
		docker compose up -d api; \
	else \
		echo "API container is already running."; \
		echo "Restarting API container..."; \
		docker compose restart api; \
	fi
	@echo "API is now running at http://localhost:4000"

## docker/restart: Restart all Docker services
.PHONY: docker/restart
docker/restart: docker/down docker/up

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


# ==================================================================================== # 
# KUBERNETES
# ==================================================================================== #

## k8s/cluster/start: Start minikube cluster with 3 nodes
.PHONY: k8s/cluster/start
k8s/cluster/start:
	@echo "Starting Minikube cluster with 3 nodes..."
	minikube start --nodes 3 -p greenlight-cluster
	@echo "Enabling ingress addon..."
	minikube addons enable ingress -p greenlight-cluster
	@echo "Minikube cluster with ingress is ready!"

## k8s/cluster/stop: Stop minikube cluster
.PHONY: k8s/cluster/stop
k8s/cluster/stop:
	@echo "Stopping Minikube cluster..."
	minikube stop -p greenlight-cluster

## k8s/cluster/delete: Delete minikube cluster
.PHONY: k8s/cluster/delete
k8s/cluster/delete:
	@echo "Deleting Minikube cluster..."
	minikube delete -p greenlight-cluster

## k8s/nodes/label: Add appropriate labels to nodes
.PHONY: k8s/nodes/label
k8s/nodes/label:
	@echo "Labeling nodes..."
	$(eval NODE_A=$(shell kubectl get nodes -o jsonpath='{.items[0].metadata.name}'))
	$(eval NODE_B=$(shell kubectl get nodes -o jsonpath='{.items[1].metadata.name}'))
	$(eval NODE_C=$(shell kubectl get nodes -o jsonpath='{.items[2].metadata.name}'))
	kubectl label node $(NODE_A) type=application
	kubectl label node $(NODE_B) type=database
	kubectl label node $(NODE_C) type=dependent_services
	@echo "Nodes labeled successfully."
	kubectl get nodes --show-labels

## k8s/deploy: Deploy the application to the Kubernetes cluster
.PHONY: k8s/deploy
k8s/deploy:
	@echo "Deploying application to Kubernetes cluster..."
	chmod +x ./k8s/scripts/deploy.sh
	./k8s/scripts/deploy.sh

## k8s/status: Check the status of the Kubernetes deployment
.PHONY: k8s/status
k8s/status:
	@echo "Checking deployment status..."
	kubectl get pods -o wide
	kubectl get services
	kubectl get deployments
	

## k8s/dashboard: Open Kubernetes dashboard
.PHONY: k8s/dashboard
k8s/dashboard:
	@echo "Opening Kubernetes dashboard..."
	minikube dashboard -p greenlight-cluster

## k8s/setup: Setup the entire Kubernetes environment
.PHONY: k8s/setup
k8s/setup: k8s/cluster/start k8s/nodes/label k8s/deploy
	@echo "Kubernetes environment setup complete!"