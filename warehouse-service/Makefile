.PHONY: run test local-db db/migrate

run:
	air -c .air.toml

test:
	godotenv -f .env go test -cover ./...

unit-test:
	@mkdir coverage || true
	-go test -race -v -coverprofile=coverage/coverage.txt.tmp -count=1  ./...
	@cat coverage/coverage.txt.tmp | grep -v "go_boilerplate_" > coverage/coverage.txt
	@go tool cover -func=coverage/coverage.txt
	@go tool cover -html=coverage/coverage.txt -o coverage/index.html

open-coverage:
	@open coverage/index.html

local-db:
	docker-compose down
	docker-compose up -d

db/migrate:
	go run ./cmd/migrate