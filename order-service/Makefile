.PHONY: run test local-db lint db/migrate

run:
	air -c .air.toml

test:
	godotenv -f .env go test -cover ./...

local-db:
	docker-compose down
	docker-compose up -d

db/migrate:
	go run ./cmd/migrate