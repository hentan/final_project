COMPOSE_FILE=docker-compose.yml
ENV_FILE=.env

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
	docker-compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) down

logs:
	@echo "Showing logs..."
	docker-compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) logs -f

ps:
	@echo "Showing container status..."
	docker-compose -f $(COMPOSE_FILE) --env-file $(ENV_FILE) ps
