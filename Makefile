DB_URL=postgresql://root:167916@localhost:5432/roomate?sslmode=disable

new_migration:
	migrate create -ext sql -dir repository/migration -seq -digits 2 $(name)

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root roomate

dropdb:
	docker exec -it postgres16 dropdb roomate

migrateup:
	migrate -path repository/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path repository/migration -database "$(DB_URL)" -verbose down

pull_sqlc:
	docker pull sqlc/sqlc

sqlc:
	docker run --rm -v .:/src -w /src sqlc/sqlc generate

.PHONY: new_migration createdb dropdb migrateup migratedown pull_sqlc sqlc