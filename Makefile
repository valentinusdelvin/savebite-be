run:
	@go run cmd/api/main.go

compose-up:
	@docker-compose up -d --build