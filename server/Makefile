postgresinit:
	 docker run  --name postgres15 -p 5433:5432 -e POSTGRES_USER-postgres -e  POSTGRES_PASSWORD=Whobay123@ -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql -U postgres

createdb:
	docker exec -it postgres15 createdb --username=postgres --owner=postgres go-chat

dropdb:
	docker exec -it postgres15 dropdb go-chat

migrateup:
	 migrate -path db/migrations -database "postgres://postgres:Whobay123@@localhost:5433/go-chat?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migrations -database "postgres://postgres:Whobay123@@localhost:5433/go-chat?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown