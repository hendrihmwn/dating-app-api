include .env
export

DATABASE_URL=postgres://${DATABASE_USER}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}?sslmode=${DATABASE_SSL_MODE}

migration-setup:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migration:
	migrate create -ext sql -dir migrations $(timestamp)${name}

migrate:
	migrate -database ${DATABASE_URL} -path migrations up

rollback:
	migrate -database ${DATABASE_URL} -path migrations down

run:
	go run serve/main.go

mock:
	go fmt ./...
	rm -rf mocks
	mockery --all --keeptree --dir app
