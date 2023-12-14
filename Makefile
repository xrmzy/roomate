postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=167916 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root roomate

dropdb:
	docker exec -it postgres16 dropdb roomate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb test