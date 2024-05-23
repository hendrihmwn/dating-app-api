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

prepare:
	go mod tidy
	go mod vendor
	go install github.com/vektra/mockery/v2@v2.38.0 1> /dev/null
	go install gotest.tools/gotestsum@latest 1> /dev/null
	go install github.com/boumenot/gocover-cobertura@latest 1> /dev/null
	go install github.com/ggere/gototal-cobertura@latest 1> /dev/null

install:
	set -e mkdir target/bin
	go build -o target/bin/dating-app serve/main.go
	cp target/bin/* /usr/local/bin/

run:
	go run serve/main.go

mock:
	go fmt ./...
	rm -rf mocks
	mockery --all --keeptree --dir app

test:
	go fmt ./...
	gotestsum --format testname --junitfile junit.xml -- -coverprofile=coverage.lcov.info -covermode count ./app/...
	gocover-cobertura < coverage.lcov.info > coverage.xml
	gototal-cobertura < coverage.xml

up:
	docker compose up
