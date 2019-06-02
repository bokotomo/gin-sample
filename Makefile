init:
	cp ./infrastructure/env/.env.sample .env

lint:
	gofmt -w driver repository domain http model router usecase clear_database.go migrate.go seed.go main.go

check:
	gofmt -s -d driver repository domain http model router usecase clear_database.go migrate.go seed.go main.go
