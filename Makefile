
postgreinit:
	docker run --name postgres162 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:16-alpine
postgrepsql:
	docker exec -it postgres162 psql
createdb:
	docker exec -it postgres162 createdb --username=root --owner=root chat-go
dropdb:
	docker exec -it postgres162 dropdb chat-go
migrateup:
	migrate -path db/migrations -database "postgresql://root:123456@localhost:5433/chat-go?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:123456@localhost:5433/chat-go?sslmode=disable" -verbose down
server:
	go run cmd/main.go
.PHONY: postgreinit, postgrepsql, createdb, dropdb, migrateup, migratedown, server