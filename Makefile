.PHONY: init
init:
	make create_db
	cp ./infrastructure/env/.env.sample .env
	make migrate
	make seed
	make test_migrate

.PHONY: fmt
fmt:
	gofmt -w -s .

.PHONY: lint
lint:
	golint ./...

.PHONY: check
check:
	gofmt -s -d driver ./driver ./repository ./domain ./http ./model ./router ./usecase ./infrastructure/sql main.go

.PHONY: test
test:
	go test -v gin-sample/usecase/test
	go test -v gin-sample/repository/test

.PHONY: migrate
migrate:
	go run ./infrastructure/sql/migrate.go

.PHONY: test_migrate
test_migrate:
	go run ./infrastructure/sql/test_migrate.go

.PHONY: seed
seed:
	go run ./infrastructure/sql/seed.go

.PHONY: clear_db
clear_db:
	go run ./infrastructure/sql/clear_database.go

.PHONY: create_db
create_db:
	mysql -uroot -pdev -hgo-sample-api-db < ./infrastructure/sql/create_databases.sql
