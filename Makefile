COMPOSE_FILE=docker-compose.yml
ENV_FILE=configs/api.env

include configs/api.env
export
DATABASE_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@db:5432/$(POSTGRES_DB)?sslmode=disable

.PHONY: all build up down logs

all: build up

build:
	@echo "Building Docker images..."
	docker-compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) build

up:
	@echo "Starting Docker containers..."
	docker-compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) up -d

down:
	@echo "Stopping Docker containers..."
	docker-compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) down -v

logs:
	@echo "Showing logs..."
	docker-compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) logs -f

ps:
	@echo "Showing container status..."
	docker-compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) ps

mig:
	@echo "Running migrations..."
	docker-compose exec app migrate -path db/migration -database $(DATABASE_URL) -verbose up
mig-down:
	@echo "Rolling back last migration..."
	docker-compose exec app migrate -path db/migration -database "$(DATABASE_URL)" down 1
