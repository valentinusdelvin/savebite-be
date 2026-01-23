run:
	@go run cmd/api/main.go

compose-up:
	@$(DOCKER_COMPOSE_CMD) up --detach --build