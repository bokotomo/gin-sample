init:
	make create_db
	cp ./infrastructure/env/.env.sample .env
	make migrate
	make seed
	make test_migrate

lint:
	gofmt -w ./driver ./repository ./domain ./http ./model ./router ./usecase ./infrastructure/sql main.go

check:
	gofmt -s -d driver ./driver ./repository ./domain ./http ./model ./router ./usecase ./infrastructure/sql main.go

test:
	go test -v gin-sample/usecase/test
	go test -v gin-sample/repository/test

migrate:
	go run ./infrastructure/sql/migrate.go

test_migrate:
	go run ./infrastructure/sql/test_migrate.go

seed:
	go run ./infrastructure/sql/seed.go

clear_db:
	go run ./infrastructure/sql/clear_database.go

create_db:
	mysql -uroot -pdev -hdb < ./infrastructure/sql/create_databases.sql
