docker-run:
	docker run --name postgres12  -p 9090:5432 -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres:12.16-alpine 

.PHONY: docker-run
include .env