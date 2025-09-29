# Oracle Pipeline Makefile
# Provides convenient commands for development and deployment

.PHONY: help build test clean deploy run stop logs status

# Default target
help:
	@echo "Oracle Pipeline - Available Commands:"
	@echo ""
	@echo "Development:"
	@echo "  build          Build all services"
	@echo "  test           Run all tests"
	@echo "  clean          Clean build artifacts"
	@echo ""
	@echo "Deployment:"
	@echo "  deploy         Deploy contracts to Sepolia"
	@echo "  run            Start the entire system"
	@echo "  stop           Stop all services"
	@echo "  restart        Restart all services"
	@echo ""
	@echo "Monitoring:"
	@echo "  logs           View all service logs"
	@echo "  status         Check service status"
	@echo "  test-system    Run system tests"
	@echo ""
	@echo "Individual Services:"
	@echo "  backend        Start only backend service"
	@echo "  updater        Start only updater service"
	@echo "  contracts      Build and test contracts"
	@echo ""

# Build all services
build:
	@echo "Building all services..."
	docker-compose build

# Run all tests
test:
	@echo "Running tests..."
	@echo "Testing backend..."
	cd backend && go test ./...
	@echo "Testing contracts..."
	cd contracts && forge test
	@echo "Testing system..."
	./scripts/test-system.sh

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	docker-compose down -v
	docker system prune -f
	cd backend && go clean
	cd updater && go clean
	cd contracts && forge clean

# Deploy contracts to Sepolia
deploy:
	@echo "Deploying contracts to Sepolia..."
	./deployments/deploy-contracts.sh

# Start the entire system
run:
	@echo "Starting Oracle Pipeline system..."
	./deployments/run-system.sh

# Stop all services
stop:
	@echo "Stopping all services..."
	docker-compose down

# Restart all services
restart: stop run

# View all service logs
logs:
	@echo "Viewing service logs..."
	docker-compose logs -f

# Check service status
status:
	@echo "Checking service status..."
	docker-compose ps
	@echo ""
	@echo "Service Health:"
	@echo "Backend: $(shell curl -s http://localhost:8080/health > /dev/null 2>&1 && echo "✅ Healthy" || echo "❌ Unhealthy")"
	@echo "Prometheus: $(shell curl -s http://localhost:9090/-/healthy > /dev/null 2>&1 && echo "✅ Healthy" || echo "❌ Unhealthy")"
	@echo "Grafana: $(shell curl -s http://localhost:3000/api/health > /dev/null 2>&1 && echo "✅ Healthy" || echo "❌ Unhealthy")"

# Run system tests
test-system:
	@echo "Running system tests..."
	./scripts/test-system.sh

# Start only backend service
backend:
	@echo "Starting backend service..."
	docker-compose up -d postgres redis nats
	sleep 10
	docker-compose up -d backend

# Start only updater service
updater:
	@echo "Starting updater service..."
	docker-compose up -d nats
	sleep 5
	docker-compose up -d updater

# Build and test contracts
contracts:
	@echo "Building and testing contracts..."
	cd contracts && forge build
	cd contracts && forge test
	cd contracts && forge coverage

# Development helpers
dev-backend:
	@echo "Starting backend in development mode..."
	cd backend && go run cmd/main.go

dev-updater:
	@echo "Starting updater in development mode..."
	cd updater && go run cmd/main.go

# Database operations
db-migrate:
	@echo "Running database migrations..."
	docker-compose exec postgres psql -U oracle_user -d oracle_db -f /docker-entrypoint-initdb.d/init-db.sql

db-reset:
	@echo "Resetting database..."
	docker-compose down -v
	docker-compose up -d postgres
	sleep 10
	make db-migrate

# Monitoring
monitor:
	@echo "Opening monitoring dashboards..."
	@echo "Prometheus: http://localhost:9090"
	@echo "Grafana: http://localhost:3000"
	@echo "NATS: http://localhost:8222"

# Security scan
security:
	@echo "Running security scans..."
	cd backend && go list -json -deps ./... | nancy sleuth
	cd contracts && forge test --gas-report

# Performance test
perf-test:
	@echo "Running performance tests..."
	@echo "Testing API performance..."
	ab -n 1000 -c 10 http://localhost:8080/price
	@echo "Testing database performance..."
	docker-compose exec postgres pgbench -U oracle_user -d oracle_db -c 10 -j 2 -T 30

# Backup
backup:
	@echo "Creating backup..."
	mkdir -p backups
	docker-compose exec postgres pg_dump -U oracle_user oracle_db > backups/db_backup_$(shell date +%Y%m%d_%H%M%S).sql
	docker-compose exec redis redis-cli --rdb backups/redis_backup_$(shell date +%Y%m%d_%H%M%S).rdb

# Restore
restore:
	@echo "Restoring from backup..."
	@echo "Available backups:"
	@ls -la backups/
	@echo "Please specify backup file: make restore-backup FILE=backup_file.sql"

restore-backup:
	@echo "Restoring database from $(FILE)..."
	docker-compose exec -T postgres psql -U oracle_user -d oracle_db < backups/$(FILE)

# Update dependencies
update:
	@echo "Updating dependencies..."
	cd backend && go get -u ./... && go mod tidy
	cd updater && go get -u ./... && go mod tidy
	cd contracts && forge update

# Format code
format:
	@echo "Formatting code..."
	cd backend && go fmt ./...
	cd updater && go fmt ./...
	cd contracts && forge fmt

# Lint code
lint:
	@echo "Linting code..."
	cd backend && golangci-lint run
	cd updater && golangci-lint run
	cd contracts && forge test --gas-report

# Generate documentation
docs:
	@echo "Generating documentation..."
	cd contracts && forge doc
	@echo "Documentation generated in contracts/docs/"

# Show environment info
env:
	@echo "Environment Information:"
	@echo "Docker: $(shell docker --version)"
	@echo "Docker Compose: $(shell docker-compose --version)"
	@echo "Go: $(shell go version)"
	@echo "Foundry: $(shell forge --version 2>/dev/null || echo "Not installed")"
	@echo "Node: $(shell node --version 2>/dev/null || echo "Not installed")"
	@echo ""
	@echo "Environment Variables:"
	@echo "ETH_PRIVATE_KEY: $(shell echo ${ETH_PRIVATE_KEY:0:10}...)"
	@echo "ETHERSCAN_API_KEY: $(shell echo ${ETHERSCAN_API_KEY:0:10}...)"

