postgres:
	docker run --name seasalon --network be-seasalon-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

network:
	docker network create be-seasalon-network

schema:
	migrate create -ext sql -dir db/migrations -seq example_schema 

createdb:
	docker exec -it seasalon createdb --username=seasalon --owner=seasalon seasalon

dropdb:
	docker exec -it postgres-test dropdb seasalon

migrateup:
	migrate -path db/migrations -database "postgresql://seasalon:seasalon@localhost:5432/seasalon?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/seasalon?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run ./cmd/main.go

mock:
	mockgen -package mockdb -destination db/mock/user.go gitlab/go-prolog-api/example/db/sqlc Users

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock network createschema


