swag:
	@echo "Generating Swagger documentation..."
	@swag init --md ./swagmd

run:
	@echo "starting gin server..."
	@go run main.go