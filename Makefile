DB_URL=root:secret@tcp(localhost:3306)/db_aggregator_dev

postgres:
	docker run --name mariadb-local -p 5432:5432 -e MARIADB_ROOT_PASSWORD=secret -d mariadb:10.2.11

compose-up:
	docker-compose up --build -d

compose-down:
	docker-compose down

run:
	go run ./cmd/

.PHONY: postgres createdb migrateup migrateup1 migratedown migratedown1 new_migration re-db run compose-up compose-down
