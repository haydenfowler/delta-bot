.PHONY: help build run test clean dev deps lint fmt

# Variables
BINARY_NAME=delta-bot
BINARY_PATH=./bin/$(BINARY_NAME)
MAIN_PATH=./cmd/main.go

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

deps: ## Download Go modules
	go mod download
	go mod tidy

build: deps ## Build the application
	@mkdir -p bin
	go build -o $(BINARY_PATH) $(MAIN_PATH)

run: build ## Build and run the application
	$(BINARY_PATH)

dev: ## Run the application with live reload (requires air)
	@which air > /dev/null || (echo "Installing air..." && go install github.com/air-verse/air@latest)
	air

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint: ## Run linter (requires golangci-lint)
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run

fmt: ## Format code
	go fmt ./...
	go mod tidy

clean: ## Clean build artifacts
	rm -rf bin/
	rm -f coverage.out coverage.html

# Docker targets
docker-build: ## Build Docker image
	docker build -t $(BINARY_NAME) .

docker-run: ## Run development setup (app only)
	docker-compose up --build

docker-run-detached: ## Run development setup in background
	docker-compose up --build -d

docker-run-monitoring: ## Run with full monitoring (sidecars)
	docker-compose --profile monitoring up --build

docker-run-monitoring-detached: ## Run with monitoring in background
	docker-compose --profile monitoring up --build -d

docker-down: ## Stop Docker containers
	docker-compose down

docker-down-monitoring: ## Stop containers with monitoring profile
	docker-compose --profile monitoring down

docker-logs: ## View Docker container logs
	docker-compose logs -f

docker-logs-monitoring: ## View all logs including monitoring
	docker-compose --profile monitoring logs -f

docker-restart: ## Restart Docker containers
	docker-compose restart

docker-restart-monitoring: ## Restart containers with monitoring
	docker-compose --profile monitoring restart

docker-clean: ## Clean up Docker images and containers
	docker-compose --profile monitoring down --volumes --remove-orphans
	docker-compose down --volumes --remove-orphans
	docker system prune -f

# Health check
health: ## Check application health
	@curl -s http://localhost:8080/health | jq . || echo "Application not running or jq not installed"

# Terraform utility targets (safe - no state changes)
tf-fmt: ## Format Terraform files
	cd terraform && terraform fmt -recursive

tf-validate: ## Validate Terraform configuration
	cd terraform && terraform validate

# Environment-specific Terraform targets
tf-init-prod: ## Initialize Terraform for production
	cd terraform && terraform init -backend-config="environments/backend-production.conf"

tf-plan-prod: ## Plan Terraform changes for production
	cd terraform && terraform plan -var-file="environments/production.tfvars"

tf-apply-prod: ## Apply Terraform changes for production
	cd terraform && terraform apply -var-file="environments/production.tfvars"

tf-destroy-prod: ## Destroy production infrastructure
	cd terraform && terraform destroy -var-file="environments/production.tfvars"

tf-init-test: ## Initialize Terraform for testing
	cd terraform && terraform init -backend-config="environments/backend-testing.conf"

tf-plan-test: ## Plan Terraform changes for testing
	cd terraform && terraform plan -var-file="environments/testing.tfvars"

tf-apply-test: ## Apply Terraform changes for testing
	cd terraform && terraform apply -var-file="environments/testing.tfvars"

tf-destroy-test: ## Destroy testing infrastructure
	cd terraform && terraform destroy -var-file="environments/testing.tfvars"

tf-output-prod: ## Show Terraform outputs for production
	cd terraform && terraform output

tf-output-test: ## Show Terraform outputs for testing  
	cd terraform && terraform output

# AWS/Docker deployment targets
docker-push: ## Build and push Docker image to ECR
	@echo "Run 'make tf-output-prod' or 'make tf-output-test' to get ECR repository URL"
	@echo "Then run the AWS ECR login and docker push commands"

aws-deploy: ## Deploy to AWS - use tf-apply-prod or tf-apply-test
	@echo "❌ Use environment-specific commands:"
	@echo "   make tf-apply-test   # Deploy to testing"
	@echo "   make tf-apply-prod   # Deploy to production"