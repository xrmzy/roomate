#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database postgresql://root:167916@postgres:5432/roomate?sslmode=disable -verbose up

echo "start the app"
exec "$@"
