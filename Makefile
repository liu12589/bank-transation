createdb :
	docker exec -it postgres12 createdb --username=root --owner=root simplebank
dropdb:
	docker exec -it postgres12 dropdb  simplebank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

.PHONY:postgres createdb dropdb migrateup migratedown sqlc test