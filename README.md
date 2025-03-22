# Green Light



##  migration 
Create
`migrate create -seq -ext=.sql -dir=./migrations create_movies_table`

Execute
`migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up`


- Third-party staticcheck tool to carry out some additional static analysis checks.
`go install honnef.co/go/tools/cmd/staticcheck@latest`



## Setup



TODO
[ ] Creating an end-to-end test for the GET /v1/healthcheck endpoint to verify that the headers and response body are what you expect.
[ ] Creating a unit-test for the rateLimit() middleware to confirm that it sends a 429 Too Many Requests response after a certain number of requests.
[ ] Creating an end-to-end integration test, using a test database instance, which confirms that the authenticate() and requirePermission() middleware work together correctly to allow or disallow access to specific endpoints.



## Prerequisites

Docker installed

Docker Compose installed

.env file configured with the necessary environment variables

## 1. Build the Docker Image

Run the following command to build the API Docker image:

docker build -t greenlight-api .

## 2. Run the Container (Standalone)

To run the container without Docker Compose:

docker run -p 4000:4000 --env-file .env greenlight-api

This will start the API and bind it to port 4000.

## 3. Running with Docker Compose

Docker Compose sets up both the API and PostgreSQL database.

### Step 1: Ensure .env is configured

Create a .env file in the project root with the following content:

```
POSTGRES_USER=<username>
POSTGRES_PASSWORD=<password>
POSTGRES_DB=greenlight
DB_DSN=postgres://<username>:<password>@db:5432/greenlight?sslmode=disable
```

### Step 2: Start the Services

Run the following command to start all services:

`docker compose up -d`

This will:

Start a PostgreSQL database

Build and run the Greenlight API

Apply database migrations

### Step 3: Verify Running Services

Check running containers:

`docker ps`

Check logs for the API:

`docker logs -f <container_id>`

### Step 4: Stop the Services

To stop and remove all containers:

`docker compose down`

## 4. Database Migrations

The migrations container runs automatically when starting Docker Compose. If needed, you can manually run migrations:

`docker compose run migrate`


## 5. Troubleshooting

Check logs: `docker compose logs -f api`

Check database health: `docker compose ps`

Manually connect to PostgreSQL:

`docker compose exec db psql -U greenlight -d greenlight`

## 6. Accessing the API

Once running, the API is accessible at:

`http://localhost:4000`

## 7. Cleaning Up

To remove containers, networks, and volumes:

`docker compose down -v`


