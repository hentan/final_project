.PHONY: build

d.up: 
	docker-compose up -d

d.down: 
	docker-compose down

build:
	go build -v ./cmd/api/

run: d.up
		go run ./cmd/api/

.DEFAULT_GOAL:= build