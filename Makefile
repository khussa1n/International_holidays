.PHONY:
.SILENT:

build:
	GOOS=linux go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

migrate:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up

drop:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' down

docker-compose-run:
	docker-compose up --build bot
