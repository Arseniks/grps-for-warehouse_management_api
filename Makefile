include .env

service-build:
	docker-compose build

service-up: service-build
	docker-compose up

up_migrations:
	migrate -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path db/migrations up

down_migrations:
	migrate -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path db/migrations down
