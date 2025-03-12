#!/bin/bash

# Load environment variables
set -a
source .env
set +a

# Build connection string
DSN="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE}"

# Run migrations
migrate -path ./migrations -database "$DSN" "$@"