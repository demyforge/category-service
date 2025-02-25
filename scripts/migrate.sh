#!/bin/bash

# Loading variables from stub.toml
CONFIG_FILE="configs/stub.toml"

if ! command -v yq &> /dev/null; then
    echo "Error: 'yq' is required. Install it with 'brew install yq' (macOS) або 'sudo apt install yq' (Linux)."
    exit 1
fi

DB_HOST=$(yq '.postgres.Host' "$CONFIG_FILE")
DB_NAME=$(yq '.postgres.Database' "$CONFIG_FILE")
DB_USER=$(yq '.postgres.Username' "$CONFIG_FILE")
DB_PASSWORD=$(yq '.postgres.Password' "$CONFIG_FILE")

# Checking whether all variables are set
if [ -z "$DB_USER" ] || [ -z "$DB_PASSWORD" ] || [ -z "$DB_NAME" ] || [ -z "$DB_HOST" ]; then
    echo "Missing required database configurations in $CONFIG_FILE"
    exit 1
fi

# Setting up
MIGRATIONS_DIR="migrations"
DATABASE_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}/${DB_NAME}?sslmode=disable"

# Checking the availability of the migrate utility
if ! command -v migrate &> /dev/null; then
    echo "Error: 'migrate' was not found. Install it with 'go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest'"
    exit 1
fi

# Command processing
case "$1" in
    create)
        if [ -z "$2" ]; then
            echo "Usage: $0 create <name>"
            exit 1
        fi
        migrate create -ext sql -dir "$MIGRATIONS_DIR" -seq "$2"
        ;;
    up)
        if [ -z "$2" ]; then
            migrate -database "$DATABASE_URL" -path "$MIGRATIONS_DIR" up
        else
            migrate -database "$DATABASE_URL" -path "$MIGRATIONS_DIR" up "$2"
        fi
        ;;
    down)
        if [ -z "$2" ]; then
            migrate -database "$DATABASE_URL" -path "$MIGRATIONS_DIR" down
        else
            migrate -database "$DATABASE_URL" -path "$MIGRATIONS_DIR" down "$2"
        fi
        ;;
    version)
        migrate -database "$DATABASE_URL" -path "$MIGRATIONS_DIR" version
        ;;
    *)
        echo "Usage: $0 {create|up|down|version}"
        exit 1
        ;;
esac
