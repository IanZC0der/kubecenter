swag:
	@echo "Generating Swagger documentation..."
	@swag init

run:
	@echo "starting gin server..."
	@go run main.go