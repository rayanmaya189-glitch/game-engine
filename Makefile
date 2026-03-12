# ===========================================
# Game Engine Casino - Makefile
# ===========================================

# Go build settings
GO := go
GOCMD := $(GO)
GO_BUILD := $(GOCMD) build
GO_TEST := $(GOCMD) test
GO_LINT := $(GOCMD) vet
GO_MOD := $(GOCMD) mod
GO_RUN := $(GOCMD) run

# Java build settings
JAVA := java
MVN := mvn
MVN_BUILD := $(MVN) clean package -DskipTests
MVN_TEST := $(MVN) test
MVN_RUN := $(MVN) spring-boot:run

# Python build settings
PYTHON := python3
PIP := pip
PIP_INSTALL := $(PIP) install -r requirements.txt
PYTHON_RUN := $(PYTHON) -m uvicorn

# Node/npm settings
NPM := npm
NPM_BUILD := $(NPM) run build
NPM_DEV := $(NPM) run dev

# Binary output directory
BIN_DIR := bin

# ===========================================
# All Services (by technology)
# ===========================================

# Go Services
GO_SERVICES := \
	gateway \
	auth-service \
	user-service \
	wallet-service \
	game-registry \
	card-games \
	dice-games \
	slot-games \
	rng-service \
	betting-service \
	tournament-service \
	jackpot-service \
	live-dealer-service \
	sports-betting-service \
	multiplayer-service \
	chat-service \
	notification-service \
	merchant-service \
	agent-service \
	loyalty-service

# Java Services
JAVA_SERVICES := \
	payment-service \
	bonus-service \
	commission-service \
	affiliate-service

# Python Services
PYTHON_SERVICES := \
	kyc-service \
	aml-service \
	fraud-service

# All services
ALL_SERVICES := $(GO_SERVICES) $(JAVA_SERVICES) $(PYTHON_SERVICES)

# Docker settings
DOCKER := docker
DOCKER_COMPOSE := docker-compose
IMAGE_NAME := game-engine
IMAGE_TAG := latest

# Proto settings
PROTO_DIR := proto
PROTO_OUT_DIR := services
PROTOC := protoc
PROTOC_GEN_GO :=protoc-gen-go
PROTOC_GEN_GRPC_GO :=protoc-gen-go-grpc
PROTO_INCLUDES := -I$(PROTO_DIR) -I$(PROTO_DIR)/common

# Color codes for output
RED := \033[0;31m
GREEN := \033[0;32m
YELLOW := \033[0;33m
BLUE := \033[0;34m
CYAN := \033[0;36m
NC := \033[0m # No Color

# Default target
.PHONY: all
all: build

# ===========================================
# Build Targets - All Services
# ===========================================

.PHONY: build
build: build-go build-java build-python
	@echo "$(GREEN)All services built successfully!$(NC)"

# ---- Go Services ----
.PHONY: build-go
build-go: build-gateway build-auth build-user build-wallet build-game-registry build-card-games build-dice-games build-slot-games build-rng build-betting build-tournament build-jackpot build-live-dealer build-sports-betting build-multiplayer build-chat build-notification build-merchant build-agent build-loyalty
	@echo "$(GREEN)All Go services built successfully!$(NC)"

.PHONY: gateway
gateway:
	@echo "$(CYAN)Building gateway service...$(NC)"
	cd services/gateway && $(GO_BUILD) -o ../../$(BIN_DIR)/gateway .
	@echo "$(GREEN)Gateway built successfully!$(NC)"

.PHONY: auth-service
auth-service:
	@echo "$(CYAN)Building auth-service...$(NC)"
	cd services/auth-service && $(GO_BUILD) -o ../../$(BIN_DIR)/auth-service .
	@echo "$(GREEN)Auth-service built successfully!$(NC)"

.PHONY: user-service
user-service:
	@echo "$(CYAN)Building user-service...$(NC)"
	cd services/user-service && $(GO_BUILD) -o ../../$(BIN_DIR)/user-service .
	@echo "$(GREEN)User-service built successfully!$(NC)"

.PHONY: wallet-service
wallet-service:
	@echo "$(CYAN)Building wallet-service...$(NC)"
	cd services/wallet-service && $(GO_BUILD) -o ../../$(BIN_DIR)/wallet-service .
	@echo "$(GREEN)Wallet-service built successfully!$(NC)"

.PHONY: game-registry
game-registry:
	@echo "$(CYAN)Building game-registry service...$(NC)"
	cd services/game-registry && $(GO_BUILD) -o ../../$(BIN_DIR)/game-registry .
	@echo "$(GREEN)Game-registry built successfully!$(NC)"

.PHONY: card-games
card-games:
	@echo "$(CYAN)Building card-games service...$(NC)"
	cd services/card-games && $(GO_BUILD) -o ../../$(BIN_DIR)/card-games .
	@echo "$(GREEN)Card-games built successfully!$(NC)"

.PHONY: dice-games
dice-games:
	@echo "$(CYAN)Building dice-games service...$(NC)"
	cd services/dice-games && $(GO_BUILD) -o ../../$(BIN_DIR)/dice-games .
	@echo "$(GREEN)Dice-games built successfully!$(NC)"

.PHONY: slot-games
slot-games:
	@echo "$(CYAN)Building slot-games service...$(NC)"
	cd services/slot-games && $(GO_BUILD) -o ../../$(BIN_DIR)/slot-games .
	@echo "$(GREEN)Slot-games built successfully!$(NC)"

.PHONY: rng-service
rng-service:
	@echo "$(CYAN)Building rng-service...$(NC)"
	cd services/rng-service && $(GO_BUILD) -o ../../$(BIN_DIR)/rng-service .
	@echo "$(GREEN)Rng-service built successfully!$(NC)"

.PHONY: betting-service
betting-service:
	@echo "$(CYAN)Building betting-service...$(NC)"
	cd services/betting-service && $(GO_BUILD) -o ../../$(BIN_DIR)/betting-service .
	@echo "$(GREEN)Betting-service built successfully!$(NC)"

.PHONY: tournament-service
tournament-service:
	@echo "$(CYAN)Building tournament-service...$(NC)"
	cd services/tournament-service && $(GO_BUILD) -o ../../$(BIN_DIR)/tournament-service .
	@echo "$(GREEN)Tournament-service built successfully!$(NC)"

.PHONY: jackpot-service
jackpot-service:
	@echo "$(CYAN)Building jackpot-service...$(NC)"
	cd services/jackpot-service && $(GO_BUILD) -o ../../$(BIN_DIR)/jackpot-service .
	@echo "$(GREEN)Jackpot-service built successfully!$(NC)"

.PHONY: live-dealer-service
live-dealer-service:
	@echo "$(CYAN)Building live-dealer-service...$(NC)"
	cd services/live-dealer-service && $(GO_BUILD) -o ../../$(BIN_DIR)/live-dealer-service .
	@echo "$(GREEN)Live-dealer-service built successfully!$(NC)"

.PHONY: sports-betting-service
sports-betting-service:
	@echo "$(CYAN)Building sports-betting-service...$(NC)"
	cd services/sports-betting-service && $(GO_BUILD) -o ../../$(BIN_DIR)/sports-betting-service .
	@echo "$(GREEN)Sports-betting-service built successfully!$(NC)"

.PHONY: multiplayer-service
multiplayer-service:
	@echo "$(CYAN)Building multiplayer-service...$(NC)"
	cd services/multiplayer-service && $(GO_BUILD) -o ../../$(BIN_DIR)/multiplayer-service .
	@echo "$(GREEN)Multiplayer-service built successfully!$(NC)"

.PHONY: chat-service
chat-service:
	@echo "$(CYAN)Building chat-service...$(NC)"
	cd services/chat-service && $(GO_BUILD) -o ../../$(BIN_DIR)/chat-service .
	@echo "$(GREEN)Chat-service built successfully!$(NC)"

.PHONY: notification-service
notification-service:
	@echo "$(CYAN)Building notification-service...$(NC)"
	cd services/notification-service && $(GO_BUILD) -o ../../$(BIN_DIR)/notification-service .
	@echo "$(GREEN)Notification-service built successfully!$(NC)"

.PHONY: merchant-service
merchant-service:
	@echo "$(CYAN)Building merchant-service...$(NC)"
	cd services/merchant-service && $(GO_BUILD) -o ../../$(BIN_DIR)/merchant-service .
	@echo "$(GREEN)Merchant-service built successfully!$(NC)"

.PHONY: agent-service
agent-service:
	@echo "$(CYAN)Building agent-service...$(NC)"
	cd services/agent-service && $(GO_BUILD) -o ../../$(BIN_DIR)/agent-service .
	@echo "$(GREEN)Agent-service built successfully!$(NC)"

.PHONY: loyalty-service
loyalty-service:
	@echo "$(CYAN)Building loyalty-service...$(NC)"
	cd services/loyalty-service && $(GO_BUILD) -o ../../$(BIN_DIR)/loyalty-service .
	@echo "$(GREEN)Loyalty-service built successfully!$(NC)"

# ---- Java Services ----
.PHONY: build-java
build-java: build-payment build-bonus build-commission build-affiliate
	@echo "$(GREEN)All Java services built successfully!$(NC)"

.PHONY: payment-service
payment-service:
	@echo "$(YELLOW)Building payment-service...$(NC)"
	cd services/payment-service && $(MVN_BUILD)
	@echo "$(GREEN)Payment-service built successfully!$(NC)"

.PHONY: bonus-service
bonus-service:
	@echo "$(YELLOW)Building bonus-service...$(NC)"
	cd services/bonus-service && $(MVN_BUILD)
	@echo "$(GREEN)Bonus-service built successfully!$(NC)"

.PHONY: commission-service
commission-service:
	@echo "$(YELLOW)Building commission-service...$(NC)"
	cd services/commission-service && $(MVN_BUILD)
	@echo "$(GREEN)Commission-service built successfully!$(NC)"

.PHONY: affiliate-service
affiliate-service:
	@echo "$(YELLOW)Building affiliate-service...$(NC)"
	cd services/affiliate-service && $(MVN_BUILD)
	@echo "$(GREEN)Affiliate-service built successfully!$(NC)"

# ---- Python Services ----
.PHONY: build-python
build-python: build-kyc build-aml build-fraud
	@echo "$(GREEN)All Python services built successfully!$(NC)"

.PHONY: kyc-service
kyc-service:
	@echo "$(BLUE)Building kyc-service...$(NC)"
	cd services/kyc-service && $(PIP_INSTALL)
	@echo "$(GREEN)Kyc-service built successfully!$(NC)"

.PHONY: aml-service
aml-service:
	@echo "$(BLUE)Building aml-service...$(NC)"
	cd services/aml-service && $(PIP_INSTALL)
	@echo "$(GREEN)Aml-service built successfully!$(NC)"

.PHONY: fraud-service
fraud-service:
	@echo "$(BLUE)Building fraud-service...$(NC)"
	cd services/fraud-service && $(PIP_INSTALL)
	@echo "$(GREEN)Fraud-service built successfully!$(NC)"

# ---- Build All to Single Directory ----
.PHONY: build-all
build-all:
	@mkdir -p $(BIN_DIR)
	@$(foreach service,$(GO_SERVICES),cd services/$(service) && $(GO_BUILD) -o ../../$(BIN_DIR)/$(service) . && cd ../..;)

# ===========================================
# Run Targets (Development)
# ===========================================

.PHONY: run
run: run-go run-java run-python

.PHONY: run-go
run-go:
	@echo "$(CYAN)Starting all Go services...$(NC)"
	@echo "$(YELLOW)Use 'make run-gateway' etc. to run individual services$(NC)"

.PHONY: run-gateway
run-gateway:
	@echo "$(CYAN)Running gateway service...$(NC)"
	cd services/gateway && $(GO_RUN) .

.PHONY: run-auth
run-auth:
	@echo "$(CYAN)Running auth-service...$(NC)"
	cd services/auth-service && $(GO_RUN) .

.PHONY: run-user
run-user:
	@echo "$(CYAN)Running user-service...$(NC)"
	cd services/user-service && $(GO_RUN) .

.PHONY: run-wallet
run-wallet:
	@echo "$(CYAN)Running wallet-service...$(NC)"
	cd services/wallet-service && $(GO_RUN) .

.PHONY: run-game-registry
run-game-registry:
	@echo "$(CYAN)Running game-registry...$(NC)"
	cd services/game-registry && $(GO_RUN) .

.PHONY: run-java
run-java:
	@echo "$(YELLOW)Starting all Java services...$(NC)"

.PHONY: run-payment
run-payment:
	@echo "$(YELLOW)Running payment-service...$(NC)"
	cd services/payment-service && $(MVN_RUN)

.PHONY: run-bonus
run-bonus:
	@echo "$(YELLOW)Running bonus-service...$(NC)"
	cd services/bonus-service && $(MVN_RUN)

.PHONY: run-commission
run-commission:
	@echo "$(YELLOW)Running commission-service...$(NC)"
	cd services/commission-service && $(MVN_RUN)

.PHONY: run-affiliate
run-affiliate:
	@echo "$(YELLOW)Running affiliate-service...$(NC)"
	cd services/affiliate-service && $(MVN_RUN)

.PHONY: run-python
run-python:
	@echo "$(BLUE)Starting all Python services...$(NC)"

.PHONY: run-kyc
run-kyc:
	@echo "$(BLUE)Running kyc-service...$(NC)"
	cd services/kyc-service && $(PYTHON_RUN) main:app --reload

.PHONY: run-aml
run-aml:
	@echo "$(BLUE)Running aml-service...$(NC)"
	cd services/aml-service && $(PYTHON_RUN) main:app --reload

.PHONY: run-fraud
run-fraud:
	@echo "$(BLUE)Running fraud-service...$(NC)"
	cd services/fraud-service && $(PYTHON_RUN) main:app --reload

# ===========================================
# Test Targets
# ===========================================

.PHONY: test
test: test-go test-java test-python
	@echo "$(GREEN)All tests passed!$(NC)"

.PHONY: test-go
test-go:
	@$(foreach service,$(GO_SERVICES),cd services/$(service) && $(GO_TEST) -v ./... && cd ../..;)

.PHONY: test-java
test-java:
	@$(foreach service,$(JAVA_SERVICES),cd services/$(service) && $(MVN_TEST) && cd ../..;)

.PHONY: test-python
test-python:
	@$(foreach service,$(PYTHON_SERVICES),cd services/$(service) && $(PYTHON) -m pytest -v && cd ../..;)

.PHONY: test-coverage
test-coverage:
	@$(foreach service,$(GO_SERVICES),cd services/$(service) && $(GO_TEST) -coverprofile=coverage.out ./... && cd ../..;)

.PHONY: test-unit
test-unit:
	@$(foreach service,$(GO_SERVICES),cd services/$(service) && $(GO_TEST) -v -run 'Unit' ./... && cd ../..;)

.PHONY: test-integration
test-integration:
	@$(foreach service,$(GO_SERVICES),cd services/$(service) && $(GO_TEST) -v -run 'Integration' ./... && cd ../..;)

# ===========================================
# Lint Targets
# ===========================================

.PHONY: lint
lint: lint-go lint-java lint-python
	@echo "$(GREEN)Lint complete!$(NC)"

.PHONY: lint-go
lint-go:
	@$(foreach service,$(GO_SERVICES),cd services/$(service) && $(GO_LINT) ./... && cd ../..;)

.PHONY: lint-java
lint-java:
	@$(foreach service,$(JAVA_SERVICES),cd services/$(service) && $(MVN) validate && cd ../..;)

.PHONY: lint-python
lint-python:
	@$(foreach service,$(PYTHON_SERVICES),cd services/$(service) && $(PIP) install pylint && pylint **/*.py && cd ../..;)

.PHONY: lint-fix
lint-fix:
	@$(foreach service,$(GO_SERVICES),cd services/$(service) && $(GO_LINT) -fix ./... && cd ../..;)

# Format code
.PHONY: fmt
fmt:
	@$(foreach service,$(GO_SERVICES),cd services/$(service) && $(GO) fmt ./... && cd ../..;)

# ===========================================
# Proto Generation Targets
# ===========================================

.PHONY: proto-gen
proto-gen:
	@echo "$(YELLOW)Generating protobuf files...$(NC)"
	@mkdir -p $(PROTO_OUT_DIR)
	@$(foreach protofile,$(wildcard $(PROTO_DIR)/*.proto),$(PROTOC) $(PROTO_INCLUDES) --go_out=$(PROTO_OUT_DIR) --go-grpc_out=$(PROTO_OUT_DIR) $(protofile);)
	@echo "$(GREEN)Proto generation complete!$(NC)"

.PHONY: proto-gen-gateway
proto-gen-gateway:
	@echo "$(YELLOW)Generating gateway protobuf files...$(NC)"
	@$(PROTOC) $(PROTO_INCLUDES) --grpc-gateway_out=logtostderr=true:$(PROTO_OUT_DIR)/gateway $(PROTO_DIR)/*.proto

# ===========================================
# Docker Targets
# ===========================================

.PHONY: docker-build
docker-build: docker-build-go docker-build-java docker-build-python
	@echo "$(GREEN)Docker build complete!$(NC)"

.PHONY: docker-build-go
docker-build-go:
	@$(foreach service,$(GO_SERVICES),$(DOCKER) build -t $(IMAGE_NAME)-$(service):$(IMAGE_TAG) services/$(service);)

.PHONY: docker-build-java
docker-build-java:
	@$(foreach service,$(JAVA_SERVICES),$(DOCKER) build -t $(IMAGE_NAME)-$(service):$(IMAGE_TAG) -f services/$(service)/Dockerfile services/$(service);)

.PHONY: docker-build-python
docker-build-python:
	@$(foreach service,$(PYTHON_SERVICES),$(DOCKER) build -t $(IMAGE_NAME)-$(service):$(IMAGE_TAG) -f services/$(service)/Dockerfile services/$(service);)

.PHONY: docker-build-all
docker-build-all:
	@$(DOCKER) build -t $(IMAGE_NAME)/$(service):$(IMAGE_TAG) -f services/$(service)/Dockerfile services/$(service)

.PHONY: docker-up
docker-up:
	$(DOCKER_COMPOSE) -f docker-compose.yml up -d

.PHONY: docker-down
docker-down:
	$(DOCKER_COMPOSE) -f docker-compose.yml down

.PHONY: docker-logs
docker-logs:
	$(DOCKER_COMPOSE) -f docker-compose.yml logs -f

.PHONY: docker-logs-service
docker-logs-service:
	@echo "Usage: make docker-logs-service SERVICE=auth-service"
	@$(DOCKER_COMPOSE) -f docker-compose.yml logs -f $(SERVICE)

# ===========================================
# Docker Compose - Individual Services
# ===========================================

.PHONY: up
up: docker-up

.PHONY: down
down: docker-down

.PHONY: up-infra
up-infra:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml up -d

.PHONY: down-infra
down-infra:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml down

.PHONY: up-auth
up-auth:
	$(DOCKER_COMPOSE) -f docker-compose.yml up -d auth-service

.PHONY: up-user
up-user:
	$(DOCKER_COMPOSE) -f docker-compose.yml up -d user-service

.PHONY: up-wallet
up-wallet:
	$(DOCKER_COMPOSE) -f docker-compose.yml up -d wallet-service

.PHONY: up-payment
up-payment:
	$(DOCKER_COMPOSE) -f docker-compose.yml up -d payment-service

.PHONY: up-bonus
up-bonus:
	$(DOCKER_COMPOSE) -f docker-compose.yml up -d bonus-service

# ===========================================
# Development Targets
# ===========================================

.PHONY: dev-up
dev-up:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml up -d

.PHONY: dev-down
dev-down:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml down

.PHONY: dev-logs
dev-logs:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml logs -f

.PHONY: dev-clean
dev-clean:
	$(DOCKER_COMPOSE) -f docker-compose.dev.yml down -v
	@rm -rf $(BIN_DIR)/*

# ===========================================
# Dependency Targets
# ===========================================

.PHONY: deps
deps: deps-download deps-verify

.PHONY: deps-download
deps-download:
	@$(foreach service,$(GO_SERVICES),cd services/$(service) && $(GO_MOD) download && cd ../..;)

.PHONY: deps-verify
deps-verify:
	@$(foreach service,$(GO_SERVICES),cd services/$(service) && $(GO_MOD) verify && cd ../..;)

.PHONY: deps-tidy
deps-tidy:
	@$(foreach service,$(GO_SERVICES),cd services/$(service) && $(GO_MOD) tidy && cd ../..;)

.PHONY: deps-java
deps-java:
	@$(foreach service,$(JAVA_SERVICES),cd services/$(service) && $(MVN) dependency:resolve && cd ../..;)

.PHONY: deps-python
deps-python:
	@$(foreach service,$(PYTHON_SERVICES),cd services/$(service) && $(PIP_INSTALL) && cd ../..;)

# ===========================================
# Admin Panel Targets
# ===========================================

.PHONY: admin
admin:
	@echo "$(CYAN)Building admin panel...$(NC)"
	cd admin && $(NPM_BUILD)

.PHONY: admin-dev
admin-dev:
	@echo "$(CYAN)Starting admin panel development server...$(NC)"
	cd admin && $(NPM_DEV)

# ===========================================
# Clean Targets
# ===========================================

.PHONY: clean
clean:
	@rm -rf $(BIN_DIR)
	@rm -f services/*/coverage.out
	@rm -f services/*/*.out
	@echo "$(GREEN)Clean complete!$(NC)"

.PHONY: clean-go
clean-go:
	@$(foreach service,$(GO_SERVICES),cd services/$(service) && $(GO) clean && cd ../..;)

.PHONY: clean-java
clean-java:
	@$(foreach service,$(JAVA_SERVICES),cd services/$(service) && $(MVN) clean && cd ../..;)

.PHONY: clean-python
clean-python:
	@$(foreach service,$(PYTHON_SERVICES),cd services/$(service) && rm -rf __pycache__ .pytest_cache && cd ../..;)

# ===========================================
# Install Tools
# ===========================================

.PHONY: install-tools
install-tools:
	$(GO) install golang.org/x/lint/golint@latest
	$(GO) install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	$(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# ===========================================
# Status Check
# ===========================================

.PHONY: status
status:
	@echo "$(CYAN)=== Service Status ===$(NC)"
	@echo "$(YELLOW)Go Services:$(NC)"
	@$(foreach service,$(GO_SERVICES),@echo "  - $(service)";)
	@echo ""
	@echo "$(YELLOW)Java Services:$(NC)"
	@$(foreach service,$(JAVA_SERVICES),@echo "  - $(service)";)
	@echo ""
	@echo "$(YELLOW)Python Services:$(NC)"
	@$(foreach service,$(PYTHON_SERVICES),@echo "  - $(service)";)

# ===========================================
# Help
# ===========================================

.PHONY: help
help:
	@echo "================================================================================"
	@echo "                    GAME ENGINE CASINO - MAKEFILE COMMANDS                      "
	@echo "================================================================================"
	@echo ""
	@echo "$(CYAN)=== BUILD COMMANDS ===$(NC)"
	@echo "  make build              - Build all services (Go, Java, Python)"
	@echo "  make build-go           - Build all Go services"
	@echo "  make build-java         - Build all Java services"
	@echo "  make build-python       - Build all Python services"
	@echo "  make build-all          - Build all binaries to bin directory"
	@echo ""
	@echo "$(CYAN)=== RUN COMMANDS (Development) ===$(NC)"
	@echo "  make run                - Run all services"
	@echo "  make run-gateway        - Run gateway service"
	@echo "  make run-auth           - Run auth-service"
	@echo "  make run-user           - Run user-service"
	@echo "  make run-wallet         - Run wallet-service"
	@echo "  make run-payment        - Run payment-service (Java)"
	@echo "  make run-bonus          - Run bonus-service (Java)"
	@echo "  make run-kyc            - Run kyc-service (Python)"
	@echo ""
	@echo "$(CYAN)=== TEST COMMANDS ===$(NC)"
	@echo "  make test               - Run all tests"
	@echo "  make test-go            - Run Go service tests"
	@echo "  make test-java          - Run Java service tests"
	@echo "  make test-python        - Run Python service tests"
	@echo "  make test-coverage      - Run tests with coverage"
	@echo "  make test-unit          - Run unit tests only"
	@echo "  make test-integration   - Run integration tests only"
	@echo ""
	@echo "$(CYAN)=== LINT & FORMAT COMMANDS ===$(NC)"
	@echo "  make lint               - Run linter on all services"
	@echo "  make lint-go            - Run Go linter"
	@echo "  make lint-java          - Run Java linter"
	@echo "  make lint-python        - Run Python linter"
	@echo "  make lint-fix           - Fix linting issues"
	@echo "  make fmt                - Format code"
	@echo ""
	@echo "$(CYAN)=== PROTO COMMANDS ===$(NC)"
	@echo "  make proto-gen         - Generate protobuf files"
	@echo "  make proto-gen-gateway - Generate gRPC gateway files"
	@echo ""
	@echo "$(CYAN)=== DOCKER COMMANDS ===$(NC)"
	@echo "  make docker-build       - Build all Docker images"
	@echo "  make docker-up          - Start Docker containers"
	@echo "  make docker-down        - Stop Docker containers"
	@echo "  make docker-logs        - View Docker logs"
	@echo "  make docker-logs-service SERVICE=auth-service - View specific service logs"
	@echo ""
	@echo "$(CYAN)=== DEVELOPMENT COMMANDS ===$(NC)"
	@echo "  make dev-up             - Start development environment"
	@echo "  make dev-down           - Stop development environment"
	@echo "  make dev-logs            - View development logs"
	@echo "  make dev-clean          - Clean development environment"
	@echo ""
	@echo "$(CYAN)=== INFRASTRUCTURE COMMANDS ===$(NC)"
	@echo "  make up-infra           - Start infrastructure services (DB, Redis, etc.)"
	@echo "  make down-infra         - Stop infrastructure services"
	@echo "  make up                 - Start all services"
	@echo "  make down               - Stop all services"
	@echo ""
	@echo "$(CYAN)=== INDIVIDUAL SERVICE COMMANDS ===$(NC)"
	@echo "  make up-auth            - Start auth-service via docker-compose"
	@echo "  make up-user            - Start user-service via docker-compose"
	@echo "  make up-wallet          - Start wallet-service via docker-compose"
	@echo "  make up-payment         - Start payment-service via docker-compose"
	@echo "  make up-bonus           - Start bonus-service via docker-compose"
	@echo ""
	@echo "$(CYAN)=== ADMIN PANEL COMMANDS ===$(NC)"
	@echo "  make admin              - Build admin panel"
	@echo "  make admin-dev          - Start admin panel dev server"
	@echo ""
	@echo "$(CYAN)=== DEPENDENCY COMMANDS ===$(NC)"
	@echo "  make deps               - Download and verify dependencies"
	@echo "  make deps-download      - Download dependencies"
	@echo "  make deps-verify         - Verify dependencies"
	@echo "  make deps-tidy          - Tidy dependencies"
	@echo "  make deps-java          - Resolve Java dependencies"
	@echo "  make deps-python        - Install Python dependencies"
	@echo ""
	@echo "$(CYAN)=== CLEAN COMMANDS ===$(NC)"
	@echo "  make clean              - Clean all build artifacts"
	@echo "  make clean-go           - Clean Go service builds"
	@echo "  make clean-java         - Clean Java service builds"
	@echo "  make clean-python       - Clean Python service caches"
	@echo ""
	@echo "$(CYAN)=== UTILITY COMMANDS ===$(NC)"
	@echo "  make install-tools      - Install required tools"
	@echo "  make status             - Show all services status"
	@echo "  make help               - Show this help message"
	@echo ""
	@echo "================================================================================"
