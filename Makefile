include .env
export

server:
	go run ./cmd/api

migrateup:
	migrate -path ./migrations -database $(GREENLIGHT_DB_DSN) -verbose up
	
migratedown:
	migrate -path ./migrations -database $(GREENLIGHT_DB_DSN) -verbose down

.PHONY: migrateup migratedown server