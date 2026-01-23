run:
	@go run cmd/api/main.go

up:
	@$(DOCKER_COMPOSE_CMD) up --detach --build