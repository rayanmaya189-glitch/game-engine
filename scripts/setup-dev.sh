#!/bin/bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Royal Platform Development Setup${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check required commands
check_requirements() {
    print_info "Checking system requirements..."
    
    local missing=()
    
    # Required tools
    if ! command -v go >/dev/null 2>&1; then missing+=("go"); fi
    if ! command -v docker >/dev/null 2>&1; then missing+=("docker"); fi
    if ! command -v docker-compose >/dev/null 2>&1; then missing+=("docker-compose"); fi
    if ! command -v psql >/dev/null 2>&1; then missing+=("postgresql-client"); fi
    if ! command -v node >/dev/null 2>&1; then missing+=("nodejs"); fi
    
    # Optional but recommended
    if ! command -v kubectl >/dev/null 2>&1; then print_warning "kubectl not found (optional)"; fi
    if ! command -v helm >/dev/null 2>&1; then print_warning "helm not found (optional)"; fi
    if ! command -v buf >/dev/null 2>&1; then print_warning "buf not found (will install)"; fi
    
    if [ ${#missing[@]} -ne 0 ]; then
        print_error "Missing required tools: ${missing[*]}"
        print_info "Please install them before continuing"
        exit 1
    fi
    
    print_success "All required tools are installed"
}

# Install Go tools
install_go_tools() {
    print_info "Installing Go development tools..."
    
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
    go install github.com/bufbuild/buf/cmd/buf@latest
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    go install github.com/mockitools/mockify/cmd/mockify@latest || true
    
    print_success "Go tools installed"
}

# Install Node.js dependencies
install_node_deps() {
    print_info "Installing Node.js dependencies for admin panel..."
    
    if [ -d "$ROOT_DIR/frontend/admin-panel" ]; then
        cd "$ROOT_DIR/frontend/admin-panel"
        npm install
        print_success "Admin panel dependencies installed"
    else
        print_warning "Admin panel directory not found"
    fi
}

# Create .env from example
setup_env() {
    print_info "Setting up environment configuration..."
    
    if [ -f "$ROOT_DIR/.env" ]; then
        print_warning ".env already exists, skipping..."
    elif [ -f "$ROOT_DIR/.env.example" ]; then
        cp "$ROOT_DIR/.env.example" "$ROOT_DIR/.env"
        print_success "Created .env from .env.example"
        print_warning "Please review and update .env with your configuration"
    else
        print_warning "No .env.example found, you'll need to create .env manually"
    fi
}

# Initialize databases
init_databases() {
    print_info "Initializing databases..."
    
    # This will be done via Docker Compose in dev mode
    print_info "Databases will be initialized when you run 'docker-compose up'"
}

# Generate protobuf files
generate_proto() {
    print_info "Generating protobuf files..."
    
    if command -v buf >/dev/null 2>&1; then
        cd "$ROOT_DIR/proto" && buf generate
        print_success "Protobuf files generated"
    else
        print_warning "buf not found, skipping proto generation"
    fi
}

# Run migrations
run_migrations() {
    print_info "Running database migrations..."
    
    # Wait for databases to be ready
    print_info "Waiting for databases to be ready..."
    sleep 10
    
    "$SCRIPT_DIR/migrate.sh" up || print_warning "Migrations failed (databases may not be running yet)"
}

# Seed test data
seed_data() {
    print_info "Seeding test data..."
    
    "$SCRIPT_DIR/seed-data.sh" || print_warning "Seeding failed (databases may not be running yet)"
}

# Main setup flow
main() {
    check_requirements
    echo ""
    
    install_go_tools
    echo ""
    
    setup_env
    echo ""
    
    generate_proto
    echo ""
    
    install_node_deps
    echo ""
    
    print_success "=========================================="
    print_success "  Development setup completed!"
    print_success "=========================================="
    print_success ""
    print_info "Next steps:"
    print_info "  1. Review and update .env file"
    print_info "  2. Run: docker-compose -f docker/docker-compose.dev.yml up -d"
    print_info "  3. Run: make migrate-up"
    print_info "  4. Run: make seed"
    print_info "  5. Run: make dev (to start all services)"
    print_success ""
    print_info "For help, run: make help"
    print_success ""
}

main "$@"
