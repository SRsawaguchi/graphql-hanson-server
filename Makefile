CONTAINER_NAME := hackernews
SRV_PORT := 3020
CLIENT_ORIGIN := http://localhost:8080
DB_HOST := localhost
DB_PASS := dbpass
DB_PORT := 3306
DB_USER := hackernews_app
DB_NAME := hackernews

.PHONY: run
run:
	env DB_HOST=$(DB_HOST) \
		DB_PORT=$(DB_PORT) \
		DB_PASSWORD=$(DB_PASS) \
		DB_USER=$(DB_USER) \
		DB_NAME=$(DB_NAME) \
		PORT=$(SRV_PORT) \
		CLIENT_ORIGIN=$(CLIENT_ORIGIN) \
		go run server.go

.PHONY: rundb
rundb:
	docker run -d \
    -p $(DB_PORT):5432 \
    --name $(CONTAINER_NAME) \
    -v $(PWD)/docker/postgres:/var/lib/postgresql/data \
    -e POSTGRES_USER=$(DB_USER) \
    -e POSTGRES_PASSWORD=$(DB_PASS) \
    -e POSTGRES_DB=$(DB_NAME) \
    postgres

.PHONY: startdb
startdb:
	docker start $(CONTAINER_NAME)

.PHONY: stopdb
stopdb:
	docker stop $(CONTAINER_NAME)

.PHONY: setup
setup:
	env PGPASSWORD=$(DB_PASS) psql \
		-h $(DB_HOST) \
		-p $(DB_PORT) \
		-U $(DB_USER)  \
		-w -d $(DB_NAME) -f ./deployments/1_ddl.sql

.PHONY: clear
clear:
	env PGPASSWORD=$(DB_PASS) psql \
		-h $(DB_HOST) \
		-p $(DB_PORT) \
		-U $(DB_USER)  \
		-w -d $(DB_NAME) -f ./deployments/teardown.sql

.PHONY: migrate
migrate:
	make clear
	make setup
