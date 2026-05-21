#!/bin/bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Royal Platform Migration Runner${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# Function to print colored messages
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

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Check required commands
REQUIRED_COMMANDS=("psql" "mysql" "mongosh")
for cmd in "${REQUIRED_COMMANDS[@]}"; do
    if ! command_exists "$cmd"; then
        print_warning "$cmd not found. Some migrations may fail."
    fi
done

# Load environment variables
if [ -f "$ROOT_DIR/.env" ]; then
    export $(cat "$ROOT_DIR/.env" | grep -v '^#' | xargs)
    print_info "Loaded environment variables from .env"
else
    print_warning ".env file not found. Using default values."
fi

# Default values
POSTGRES_HOST="${POSTGRES_HOST:-localhost}"
POSTGRES_PORT="${POSTGRES_PORT:-5432}"
POSTGRES_USER="${POSTGRES_USER:-royal_platform}"
POSTGRES_PASSWORD="${POSTGRES_PASSWORD:-royal_platform_password}"
POSTGRES_DB="${POSTGRES_DB:-royal_platform}"

MYSQL_HOST="${MYSQL_HOST:-localhost}"
MYSQL_PORT="${MYSQL_PORT:-3306}"
MYSQL_USER="${MYSQL_USER:-royal_platform}"
MYSQL_PASSWORD="${MYSQL_PASSWORD:-royal_platform_password}"
MYSQL_DB="${MYSQL_DB:-royal_platform_games}"

MONGO_URI="${MONGO_URI:-mongodb://localhost:27017/royal_platform}"

# Function to run PostgreSQL migrations
run_postgres_migrations() {
    local service=$1
    local migration_dir="$ROOT_DIR/go/services/$service/migrations"
    
    if [ ! -d "$migration_dir" ]; then
        print_warning "No migrations directory for $service"
        return 0
    fi
    
    print_info "Running PostgreSQL migrations for $service..."
    
    # Create migrations table if not exists
    PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
        -U "$POSTGRES_USER" -d "${POSTGRES_DB}_${service}" <<EOF
CREATE TABLE IF NOT EXISTS schema_migrations (
    version VARCHAR(255) PRIMARY KEY,
    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
EOF
    
    # Run pending migrations
    for migration_file in "$migration_dir"/*.up.sql; do
        if [ -f "$migration_file" ]; then
            version=$(basename "$migration_file" | cut -d'_' -f1)
            
            # Check if already applied
            applied=$(PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
                -U "$POSTGRES_USER" -d "${POSTGRES_DB}_${service}" -t -c \
                "SELECT COUNT(*) FROM schema_migrations WHERE version = '$version';")
            
            if [ "$(echo $applied | tr -d ' ')" -eq 0 ]; then
                print_info "Applying migration: $migration_file"
                PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
                    -U "$POSTGRES_USER" -d "${POSTGRES_DB}_${service}" -f "$migration_file"
                
                # Record migration
                PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
                    -U "$POSTGRES_USER" -d "${POSTGRES_DB}_${service}" -c \
                    "INSERT INTO schema_migrations (version) VALUES ('$version');"
                
                print_success "Applied: $version"
            else
                print_info "Skipping (already applied): $version"
            fi
        fi
    done
}

# Function to rollback PostgreSQL migrations
rollback_postgres_migrations() {
    local service=$1
    local migration_dir="$ROOT_DIR/go/services/$service/migrations"
    
    if [ ! -d "$migration_dir" ]; then
        return 0
    fi
    
    print_info "Rolling back migrations for $service..."
    
    # Get last applied migration
    last_version=$(PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
        -U "$POSTGRES_USER" -d "${POSTGRES_DB}_${service}" -t -c \
        "SELECT version FROM schema_migrations ORDER BY applied_at DESC LIMIT 1;" | tr -d ' ')
    
    if [ -z "$last_version" ]; then
        print_warning "No migrations to rollback"
        return 0
    fi
    
    down_file=$(find "$migration_dir" -name "${last_version}*.down.sql" | head -n 1)
    
    if [ -f "$down_file" ]; then
        print_info "Rolling back: $down_file"
        PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
            -U "$POSTGRES_USER" -d "${POSTGRES_DB}_${service}" -f "$down_file"
        
        # Remove migration record
        PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
            -U "$POSTGRES_USER" -d "${POSTGRES_DB}_${service}" -c \
            "DELETE FROM schema_migrations WHERE version = '$last_version';"
        
        print_success "Rolled back: $last_version"
    else
        print_error "No down migration found for version $last_version"
        return 1
    fi
}

# Function to show migration status
show_migration_status() {
    local service=$1
    local migration_dir="$ROOT_DIR/go/services/$service/migrations"
    
    if [ ! -d "$migration_dir" ]; then
        return 0
    fi
    
    print_info "Migration status for $service:"
    echo "----------------------------------------"
    
    PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
        -U "$POSTGRES_USER" -d "${POSTGRES_DB}_${service}" -c \
        "SELECT version, applied_at FROM schema_migrations ORDER BY applied_at;"
}

# Function to create a new migration
create_migration() {
    local service=$1
    local name=$2
    local migration_dir="$ROOT_DIR/go/services/$service/migrations"
    local timestamp=$(date +%Y%m%d%H%M%S)
    
    if [ ! -d "$migration_dir" ]; then
        mkdir -p "$migration_dir"
    fi
    
    local up_file="$migration_dir/${timestamp}_${name}.up.sql"
    local down_file="$migration_dir/${timestamp}_${name}.down.sql"
    
    cat > "$up_file" <<EOF
-- Migration: ${name}
-- Created: $(date)
-- Description: Add your SQL here

-- Example:
-- CREATE TABLE IF NOT EXISTS example_table (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(255) NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );
EOF
    
    cat > "$down_file" <<EOF
-- Rollback: ${name}
-- Description: Add rollback SQL here

-- Example:
-- DROP TABLE IF EXISTS example_table;
EOF
    
    print_success "Created migration files:"
    echo "  Up:   $up_file"
    echo "  Down: $down_file"
}

# Main command handler
case "${1:-help}" in
    up)
        print_info "Running all pending migrations..."
        if [ -n "$2" ]; then
            # Migrate specific service
            run_postgres_migrations "$2"
        else
            # Migrate all services
            for service_dir in "$ROOT_DIR"/go/services/*/; do
                service=$(basename "$service_dir")
                run_postgres_migrations "$service"
            done
        fi
        print_success "All migrations completed!"
        ;;
    
    down)
        print_info "Rolling back last migration..."
        if [ -n "$2" ]; then
            rollback_postgres_migrations "$2"
        else
            # Rollback last service that was migrated
            for service_dir in "$ROOT_DIR"/go/services/*/; do
                service=$(basename "$service_dir")
                rollback_postgres_migrations "$service"
            done
        fi
        ;;
    
    status)
        print_info "Checking migration status..."
        if [ -n "$2" ]; then
            show_migration_status "$2"
        else
            for service_dir in "$ROOT_DIR"/go/services/*/; do
                service=$(basename "$service_dir")
                show_migration_status "$service"
                echo ""
            done
        fi
        ;;
    
    create)
        if [ -z "$2" ] || [ -z "$3" ]; then
            print_error "Usage: $0 create <service> <migration_name>"
            exit 1
        fi
        create_migration "$2" "$3"
        ;;
    
    reset)
        print_warning "This will drop and recreate all databases!"
        read -p "Are you sure? (yes/no): " confirm
        if [ "$confirm" != "yes" ]; then
            print_info "Aborted"
            exit 0
        fi
        
        for service_dir in "$ROOT_DIR"/go/services/*/; do
            service=$(basename "$service_dir")
            print_info "Resetting database for $service..."
            PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
                -U "$POSTGRES_USER" -d postgres -c \
                "DROP DATABASE IF EXISTS ${POSTGRES_DB}_${service};"
            PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
                -U "$POSTGRES_USER" -d postgres -c \
                "CREATE DATABASE ${POSTGRES_DB}_${service} OWNER $POSTGRES_USER;"
        done
        
        print_success "Databases reset. Run 'migrate.sh up' to apply migrations."
        ;;
    
    help|*)
        echo "Usage: $0 <command> [options]"
        echo ""
        echo "Commands:"
        echo "  up [service]              Run all pending migrations"
        echo "  down [service]            Rollback last migration"
        echo "  status [service]          Show migration status"
        echo "  create <service> <name>   Create new migration"
        echo "  reset                     Reset all databases (DANGER!)"
        echo "  help                      Show this help message"
        echo ""
        echo "Examples:"
        echo "  $0 up                     # Run all migrations"
        echo "  $0 up auth-service        # Run migrations for auth-service only"
        echo "  $0 create auth-service add_users_table"
        echo "  $0 status                 # Show status for all services"
        ;;
esac

exit 0
