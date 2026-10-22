include ./.env


REDIS_URL = redis://$(REDIS_USER):$(REDIS_PASS)@$(REDIS_HOST):$(REDIS_PORT)
DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

MIGRATE_PATH=db/migrations
SEEDER_PATH=db/seeders/seeder.sql

migrate-up:
	migrate -database $(DB_URL) -path $(MIGRATE_PATH) up

migrate-down:
	migrate -database $(DB_URL) -path $(MIGRATE_PATH) down $(N)

swag:
	swag init -g cmd/main.go -o docs

up:
	docker compose -f deployment/docker-compose.yml up -d --build

down:
	docker compose -f deployment/docker-compose.yml down --rmi local