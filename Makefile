SHELL := /bin/sh

postgres:
	sudo docker run --name postgres12 -p 5432:5432 \
		-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=secret \
		-d postgres:12-alpine

createdb: 
	sudo docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	sudo docker exec -it postgres12 dropdb simple_bank

migrate-up:
	migrate -path db/migration \
	 -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" \
	 -verbose up

migrate-up-1:
	migrate -path db/migration \
	 -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" \
	 -verbose up 1

migrate-down:
	migrate -path db/migration \
	 -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" \
	 -verbose down

migrate-down-1:
	migrate -path db/migration \
	 -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" \
	 -verbose down 1

sqlc:
	sqlc generate

test: 
	go test -timeout 30s -v -cover ./tests/...

clean:
	docker stop postgres12

docs: 
	swag init --dir ./src/

server: docs
	go run src/main.go

mock:
	mockgen -package mockdb -destination tests/mock/store.go simple-bank/src/dao Store

.PHONY: postgres createdb dropdb migrate-up migrate-down sqlc test clean docs mock migrate-up-1 migrate-down-1
