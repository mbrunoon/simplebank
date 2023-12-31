postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.3-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb: 
	docker exec -it postgres15 dropdb simple_bank

migrate-up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose down

sqlc:
	docker run --rm -v "${CURDIR}:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrate-up migrate-down sqlc test