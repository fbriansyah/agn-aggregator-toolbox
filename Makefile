DB_URL=root:secret@tcp(localhost:3306)/db_aggregator_dev

mariadb:
	docker run --name mariadb-local -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mariadb:10.2.11 

sqlc-win:
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate
	
compose-up:
	docker-compose up --build -d

compose-down:
	docker-compose down

run:
	go run ./cmd/

.PHONY: mariadb run compose-up compose-down
