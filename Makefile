.PHONY: help build clean test proto generate migrate seed docker k8s deploy

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-25s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Proto generation
proto-install: ## Install protobuf tools
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest

proto-generate: ## Generate Go code from protobuf definitions
	cd proto && buf generate

proto-lint: ## Lint protobuf files
	cd proto && buf lint

proto-breaking: ## Check for breaking changes in protobuf
	cd proto && buf breaking --against .git

proto-all: proto-generate proto-lint ## Generate and lint protobuf files

# Build commands
build-all: ## Build all Go services
	./scripts/build-all-services.sh

build-go: ## Build all Go services (alternative)
	./scripts/build-go-services.sh

build-auth: ## Build auth service only
	cd go/services/auth-service && go build -o ../../../bin/auth-service ./cmd/auth-service

build-user: ## Build user service only
	cd go/services/user-service && go build -o ../../../bin/user-service ./cmd/user-service

build-wallet: ## Build wallet service only
	cd go/services/wallet-service && go build -o ../../../bin/wallet-service ./cmd/wallet-service

build-game-engine: ## Build game engine service
	cd go/services/game-engine-service && go build -o ../../../bin/game-engine-service ./cmd/game-engine-service

build-card-games: ## Build card games service
	cd go/services/card-games-service && go build -o ../../../bin/card-games-service ./cmd/card-games-service

build-dice-games: ## Build dice games service
	cd go/services/dice-games-service && go build -o ../../../bin/dice-games-service ./cmd/dice-games-service

build-slot-games: ## Build slot games service
	cd go/services/slot-games-service && go build -o ../../../bin/slot-games-service ./cmd/slot-games-service

build-betting: ## Build betting service
	cd go/services/betting-service && go build -o ../../../bin/betting-service ./cmd/betting-service

build-multiplayer: ## Build multiplayer service
	cd go/services/multiplayer-service && go build -o ../../../bin/multiplayer-service ./cmd/multiplayer-service

build-tournament: ## Build tournament service
	cd go/services/tournament-service && go build -o ../../../bin/tournament-service ./cmd/tournament-service

build-jackpot: ## Build jackpot service
	cd go/services/jackpot-service && go build -o ../../../bin/jackpot-service ./cmd/jackpot-service

build-leaderboard: ## Build leaderboard service
	cd go/services/leaderboard-service && go build -o ../../../bin/leaderboard-service ./cmd/leaderboard-service

build-game-registry: ## Build game registry service
	cd go/services/game-registry-service && go build -o ../../../bin/game-registry-service ./cmd/game-registry-service

build-rng: ## Build RNG service
	cd go/services/rng-service && go build -o ../../../bin/rng-service ./cmd/rng-service

build-notification: ## Build notification service
	cd go/services/notification-service && go build -o ../../../bin/notification-service ./cmd/notification-service

build-chat: ## Build chat service
	cd go/services/chat-service && go build -o ../../../bin/chat-service ./cmd/chat-service

build-live-dealer: ## Build live dealer service
	cd go/services/live-dealer-service && go build -o ../../../bin/live-dealer-service ./cmd/live-dealer-service

build-loyalty: ## Build loyalty service
	cd go/services/loyalty-service && go build -o ../../../bin/loyalty-service ./cmd/loyalty-service

build-referral: ## Build referral service
	cd go/services/referral-service && go build -o ../../../bin/referral-service ./cmd/referral-service

build-banner: ## Build banner service
	cd go/services/banner-service && go build -o ../../../bin/banner-service ./cmd/banner-service

build-winners-showcase: ## Build winners showcase service
	cd go/services/winners-showcase-service && go build -o ../../../bin/winners-showcase-service ./cmd/winners-showcase-service

build-sports-betting: ## Build sports betting service
	cd go/services/sports-betting-service && go build -o ../../../bin/sports-betting-service ./cmd/sports-betting-service

build-merchant: ## Build merchant service
	cd go/services/merchant-service && go build -o ../../../bin/merchant-service ./cmd/merchant-service

build-agent: ## Build agent service
	cd go/services/agent-service && go build -o ../../../bin/agent-service ./cmd/agent-service

build-player-gateway: ## Build player API gateway
	cd go/gateways/player-api-gateway && go build -o ../../../bin/player-api-gateway ./cmd/player-api-gateway

build-admin-gateway: ## Build admin API gateway
	cd go/gateways/admin-api-gateway && go build -o ../../../bin/admin-api-gateway ./cmd/admin-api-gateway

build-merchant-gateway: ## Build merchant API gateway
	cd go/gateways/merchant-api-gateway && go build -o ../../../bin/merchant-api-gateway ./cmd/merchant-api-gateway

build-agent-gateway: ## Build agent API gateway
	cd go/gateways/agent-api-gateway && go build -o ../../../bin/agent-api-gateway ./cmd/agent-api-gateway

build-java: ## Build all Java Spring Boot services
	mvn -f java/services/payment-service/pom.xml clean package -DskipTests
	mvn -f java/services/bonus-service/pom.xml clean package -DskipTests
	mvn -f java/services/affiliate-service/pom.xml clean package -DskipTests
	mvn -f java/services/commission-service/pom.xml clean package -DskipTests

build-python: ## Build all Python FastAPI services
	@echo "Python services don't require building, but installing dependencies..."
	pip install -r python/services/aml-service/requirements.txt
	pip install -r python/services/fraud-service/requirements.txt
	pip install -r python/services/risk-service/requirements.txt
	pip install -r python/services/kyc-service/requirements.txt

build-mobile-android: ## Build Android app
	cd mobile/android && ./gradlew assembleRelease

build-mobile-ios: ## Build iOS app
	cd mobile/ios && xcodebuild -scheme RoyalPlatform -configuration Release archive

build-frontend: ## Build admin panel frontend
	cd frontend/admin-panel && npm install && npm run build

build-full: proto-generate build-all build-java build-frontend ## Build everything

# Test commands
test: ## Run all tests
	go test ./go/... -v

test-unit: ## Run unit tests only
	go test ./go/... -short -v

test-integration: ## Run integration tests
	go test ./go/... -tags=integration -v

test-coverage: ## Run tests with coverage
	go test ./go/... -coverprofile=coverage.out -covermode=atomic
	go tool cover -html=coverage.out -o coverage.html

test-auth: ## Test auth service
	cd go/services/auth-service && go test -v ./...

test-user: ## Test user service
	cd go/services/user-service && go test -v ./...

test-wallet: ## Test wallet service
	cd go/services/wallet-service && go test -v ./...

# Database commands
migrate-up: ## Run all database migrations up
	./scripts/migrate.sh up

migrate-down: ## Rollback last migration
	./scripts/migrate.sh down

migrate-create: ## Create new migration file
	./scripts/migrate.sh create $(name)

migrate-status: ## Show migration status
	./scripts/migrate.sh status

migrate-reset: ## Reset database (drop and recreate)
	./scripts/migrate.sh reset

seed: ## Seed database with test data
	./scripts/seed-data.sh

backup: ## Backup database
	./scripts/backup-restore.sh backup

restore: ## Restore database from backup
	./scripts/backup-restore.sh restore $(file)

# Docker commands
docker-build: ## Build all Docker images
	docker-compose -f docker/docker-compose.yml build

docker-build-dev: ## Build dev Docker images
	docker-compose -f docker/docker-compose.dev.yml build

docker-up: ## Start all services with Docker
	docker-compose -f docker/docker-compose.yml up -d

docker-down: ## Stop all Docker services
	docker-compose -f docker/docker-compose.yml down

docker-up-dev: ## Start dev services with Docker
	docker-compose -f docker/docker-compose.dev.yml up -d

docker-down-dev: ## Stop dev Docker services
	docker-compose -f docker/docker-compose.dev.yml down

docker-logs: ## View logs from all services
	docker-compose -f docker/docker-compose.yml logs -f

docker-ps: ## List running containers
	docker-compose -f docker/docker-compose.yml ps

docker-clean: ## Remove all containers and volumes
	docker-compose -f docker/docker-compose.yml down -v --remove-orphans
	docker system prune -f

# Kubernetes commands
k8s-apply: ## Apply Kubernetes manifests
	kubectl apply -f k8s/base/

k8s-delete: ## Delete Kubernetes manifests
	kubectl delete -f k8s/base/

k8s-scale: ## Scale a deployment
	kubectl scale deployment $(service) --replicas=$(replicas)

k8s-logs: ## View logs from a pod
	kubectl logs -f deployment/$(service)

k8s-exec: ## Execute command in a pod
	kubectl exec -it deployment/$(service) -- /bin/sh

helm-install: ## Install with Helm
	helm install royal-platform helm/royal-platform --namespace royal-platform --create-namespace

helm-upgrade: ## Upgrade Helm release
	helm upgrade royal-platform helm/royal-platform --namespace royal-platform

helm-uninstall: ## Uninstall Helm release
	helm uninstall royal-platform --namespace royal-platform

# Deployment
deploy-dev: ## Deploy to development environment
	./scripts/k3s-deploy.sh dev

deploy-staging: ## Deploy to staging environment
	./scripts/k3s-deploy.sh staging

deploy-prod: ## Deploy to production environment
	./scripts/k3s-deploy.sh prod

# Utility commands
clean: ## Clean build artifacts
	rm -rf bin/*
	rm -rf go/*/services/*/*/dist
	rm -rf frontend/*/dist
	rm -rf coverage.out coverage.html

setup: ## Run initial setup
	./scripts/setup.sh

setup-dev: ## Setup development environment
	./scripts/setup-dev.sh

fmt: ## Format Go code
	go fmt ./go/...

lint: ## Lint Go code
	golangci-lint run ./go/...

vet: ## Vet Go code
	go vet ./go/...

tidy: ## Tidy Go modules
	go mod tidy

generate-mock: ## Generate mocks for testing
	mockgen -source=./go/services/auth-service/internal/service/auth.go -destination=./go/services/auth-service/internal/mocks/mock_auth.go

openapi: ## Generate OpenAPI specs
	cd proto && buf generate --template buf.gen.yaml

docs: ## Generate documentation
	@echo "Generating documentation..."
	# Add documentation generation commands here

health: ## Check health of all services
	@echo "Checking service health..."
	curl -s http://localhost:8080/health || echo "Auth Gateway: DOWN"
	curl -s http://localhost:8081/health || echo "Admin Gateway: DOWN"
	curl -s http://localhost:8082/health || echo "Merchant Gateway: DOWN"

version: ## Show version info
	@echo "Royal Platform Build System"
	@echo "Go version: $$(go version)"
	@echo "Docker version: $$(docker --version)"
	@echo "Kubernetes version: $$(kubectl version --client --short 2>/dev/null || echo 'Not installed')"

# Quick development workflow
dev: proto-generate build-all docker-up-dev ## Full dev workflow: generate, build, start

reset: docker-clean clean migrate-reset seed docker-up-dev ## Reset everything and start fresh
