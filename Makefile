createdb :
	docker exec -it postgres12 createdb --username=root --owner=root simplebank
dropdb:
	docker exec -it postgres12 dropdb  simplebank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go bank-transaction/db/sqlc Store

.PHONY:postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock