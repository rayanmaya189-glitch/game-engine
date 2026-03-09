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

# Binary output directory
BIN_DIR := bin

# Services to build
SERVICES := gateway auth-service user-service wallet-service game-registry

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
NC := \033[0m # No Color

# Default target
.PHONY: all
all: build

# ===========================================
# Build Targets
# ===========================================

.PHONY: build
build: build-services
	@echo "$(GREEN)Build complete!$(NC)"

.PHONY: build-services
build-services: $(SERVICES)

gateway:
	@echo "$(YELLOW)Building gateway service...$(NC)"
	cd services/gateway && $(GO_BUILD) -o ../../$(BIN_DIR)/gateway .
	@echo "$(GREEN)Gateway built successfully!$(NC)"

auth-service:
	@echo "$(YELLOW)Building auth-service...$(NC)"
	cd services/auth-service && $(GO_BUILD) -o ../../$(BIN_DIR)/auth-service .
	@echo "$(GREEN)Auth-service built successfully!$(NC)"

user-service:
	@echo "$(YELLOW)Building user-service...$(NC)"
	cd services/user-service && $(GO_BUILD) -o ../../$(BIN_DIR)/user-service .
	@echo "$(GREEN)User-service built successfully!$(NC)"

wallet-service:
	@echo "$(YELLOW)Building wallet-service...$(NC)"
	cd services/wallet-service && $(GO_BUILD) -o ../../$(BIN_DIR)/wallet-service .
	@echo "$(GREEN)Wallet-service built successfully!$(NC)"

game-registry:
	@echo "$(YELLOW)Building game-registry service...$(NC)"
	cd services/game-registry && $(GO_BUILD) -o ../../$(BIN_DIR)/game-registry .
	@echo "$(GREEN)Game-registry built successfully!$(NC)"

# Build all binaries to single directory
.PHONY: build-all
build-all:
	@mkdir -p $(BIN_DIR)
	@$(foreach service,$(SERVICES),cd services/$(service) && $(GO_BUILD) -o ../../$(BIN_DIR)/$(service) . && cd ../..;)

# ===========================================
# Test Targets
# ===========================================

.PHONY: test
test: test-services
	@echo "$(GREEN)All tests passed!$(NC)"

.PHONY: test-services
test-services:
	@$(foreach service,$(SERVICES),cd services/$(service) && $(GO_TEST) -v ./... && cd ../..;)

.PHONY: test-coverage
test-coverage:
	@$(foreach service,$(SERVICES),cd services/$(service) && $(GO_TEST) -coverprofile=coverage.out ./... && cd ../..;)

.PHONY: test-unit
test-unit:
	@$(foreach service,$(SERVICES),cd services/$(service) && $(GO_TEST) -v -run 'Unit' ./... && cd ../..;)

.PHONY: test-integration
test-integration:
	@$(foreach service,$(SERVICES),cd services/$(service) && $(GO_TEST) -v -run 'Integration' ./... && cd ../..;)

# ===========================================
# Lint Targets
# ===========================================

.PHONY: lint
lint: lint-services
	@echo "$(GREEN)Lint complete!$(NC)"

.PHONY: lint-services
lint-services:
	@$(foreach service,$(SERVICES),cd services/$(service) && $(GO_LINT) ./... && cd ../..;)

.PHONY: lint-fix
lint-fix:
	@$(foreach service,$(SERVICES),cd services/$(service) && $(GO_LINT) -fix ./... && cd ../..;)

# Format code
.PHONY: fmt
fmt:
	@$(foreach service,$(SERVICES),cd services/$(service) && $(GO) fmt ./... && cd ../..;)

# ===========================================
# Proto Generation Targets
# ===========================================

.PHONY: proto-gen
proto-gen: proto-gen-services
	@echo "$(GREEN)Proto generation complete!$(NC)"

.PHONY: proto-gen-services
proto-gen-services:
	@echo "$(YELLOW)Generating protobuf files...$(NC)"
	@mkdir -p $(PROTO_OUT_DIR)
	@$(foreach protofile,$(wildcard $(PROTO_DIR)/*.proto),$(PROTOC) $(PROTO_INCLUDES) --go_out=$(PROTO_OUT_DIR) --go-grpc_out=$(PROTO_OUT_DIR) $(protofile);)

.PHONY: proto-gen-gateway
proto-gen-gateway:
	@echo "$(YELLOW)Generating gateway protobuf files...$(NC)"
	@$(PROTOC) $(PROTO_INCLUDES) --grpc-gateway_out=logtostderr=true:$(PROTO_OUT_DIR)/gateway $(PROTO_DIR)/*.proto

# ===========================================
# Docker Targets
# ===========================================

.PHONY: docker-build
docker-build: docker-build-services
	@echo "$(GREEN)Docker build complete!$(NC)"

.PHONY: docker-build-services
docker-build-services:
	@$(foreach service,$(SERVICES),$(DOCKER) build -t $(IMAGE_NAME)-$(service):$(IMAGE_TAG) services/$(service);)

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
	@$(foreach service,$(SERVICES),cd services/$(service) && $(GO_MOD) download && cd ../..;)

.PHONY: deps-verify
deps-verify:
	@$(foreach service,$(SERVICES),cd services/$(service) && $(GO_MOD) verify && cd ../..;)

.PHONY: deps-tidy
deps-tidy:
	@$(foreach service,$(SERVICES),cd services/$(service) && $(GO_MOD) tidy && cd ../..;)

# ===========================================
# Clean Targets
# ===========================================

.PHONY: clean
clean:
	@rm -rf $(BIN_DIR)
	@rm -f services/*/coverage.out
	@rm -f services/*/*.out
	@echo "$(GREEN)Clean complete!$(NC)"

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
# Help
# ===========================================

.PHONY: help
help:
	@echo "Game Engine Casino - Makefile Commands"
	@echo ""
	@echo "Build Commands:"
	@echo "  make build              - Build all services"
	@echo "  make build-all         - Build all binaries to bin directory"
	@echo ""
	@echo "Test Commands:"
	@echo "  make test              - Run all tests"
	@echo "  make test-coverage     - Run tests with coverage"
	@echo "  make test-unit         - Run unit tests only"
	@echo "  make test-integration  - Run integration tests only"
	@echo ""
	@echo "Lint Commands:"
	@echo "  make lint              - Run linter"
	@echo "  make lint-fix          - Fix linting issues"
	@echo "  make fmt               - Format code"
	@echo ""
	@echo "Proto Commands:"
	@echo "  make proto-gen         - Generate protobuf files"
	@echo "  make proto-gen-gateway - Generate gRPC gateway files"
	@echo ""
	@echo "Docker Commands:"
	@echo "  make docker-build      - Build all Docker images"
	@echo "  make docker-up         - Start Docker containers"
	@echo "  make docker-down       - Stop Docker containers"
	@echo "  make docker-logs       - View Docker logs"
	@echo ""
	@echo "Development Commands:"
	@echo "  make dev-up            - Start development environment"
	@echo "  make dev-down          - Stop development environment"
	@echo "  make dev-logs          - View development logs"
	@echo "  make dev-clean         - Clean development environment"
	@echo ""
	@echo "Dependency Commands:"
	@echo "  make deps              - Download and verify dependencies"
	@echo "  make deps-tidy         - Tidy dependencies"
	@echo "  make install-tools     - Install required tools"
	@echo ""
	@echo "Utility Commands:"
	@echo "  make clean             - Clean build artifacts"
	@echo "  make help              - Show this help message"
