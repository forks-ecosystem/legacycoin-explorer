#!/bin/sh
set -e

cd "$(dirname "$0")"

echo '==> Stopping and removing container/network...'
docker compose down

echo '==> Removing old explorer images...'
docker rmi -f legacycoin-explorer:latest 2>/dev/null || true
docker image prune -f

echo '==> Rebuilding image...'
docker compose build explorer

echo '==> Starting container...'
docker compose up -d explorer

echo '==> Done. legacycoin-explorer rebuilt and running.'