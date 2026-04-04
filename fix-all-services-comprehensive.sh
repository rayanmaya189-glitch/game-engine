#!/bin/bash
set -e

BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$BASE_DIR"

fix_service() {
    local svc=$1
    local dir="services/$svc"
    
    if [ ! -d "$dir" ] || [ ! -f "$dir/go.mod" ]; then
        echo "SKIP: $svc (not found)"
        return
    fi
    
    echo "=== Fixing: $svc ==="
    cd "$dir"
    
    # Fix Go version
    sed -i 's/go 1.2[0-9]/go 1.24/g' go.mod
    
    # Fix NATS dependency issue
    if grep -q "nats.go v1.5" go.mod; then
        sed -i 's/nats.go v1.5/nats.go v1.28.0/g' go.mod
    fi
    
    # Add replace directive if needed
    if ! grep -q "replace.*nats-server" go.mod; then
        echo -e "\nreplace github.com/nats-io/nats-server/v2 => github.com/nats-io/nats-server/v2 v2.9.0" >> go.mod
    fi
    
    # Copy proto files if needed
    if [ -d "../common-service/proto/gen/go/game_engine" ]; then
        mkdir -p pkg/game_engine
        cp -r ../common-service/proto/gen/go/game_engine/* pkg/game_engine/ 2>/dev/null || true
        
        # Fix proto imports
        find pkg -name "*.go" -exec sed -i "s|github.com/game_engine/common/proto/gen/go/game_engine|github.com/game_engine/$svc/pkg/game_engine|g" {} \; 2>/dev/null || true
        find pkg -name "*.go" -exec sed -i 's|"gen/go|"github.com/game_engine/$svc/pkg/game_engine/|g' {} \; 2>/dev/null || true
    fi
    
    # Fix source imports
    sed -i 's|game_engine/gen/go|github.com/game_engine/'"$svc"'/pkg/game_engine|g' cmd/main.go internal/handler/*.go internal/model/*.go 2>/dev/null || true
    sed -i 's|gen/go|github.com/game_engine/'"$svc"'/pkg/game_engine|g' internal/model/*.go 2>/dev/null || true
    
    # Fix Dockerfile Go version
    if [ -f "Dockerfile" ]; then
        sed -i 's/FROM golang:1.2[0-9]-alpine/FROM golang:1.24-alpine/g' Dockerfile
    fi
    
    # Run go mod tidy
    rm -f go.sum
    docker run --rm -v "$(pwd):/app" -w /app -e GOTOOLCHAIN=local golang:1.24-alpine sh -c "apk add --no-cache git && go mod tidy" 2>&1 | grep -v "^$" | tail -3
    
    # Build
    docker build -t "game-engine/$svc:latest" . 2>&1 | tail -5 || echo "FAILED: $svc"
    
    cd "$BASE_DIR"
}

export -f fix_service
export BASE_DIR

# Services to fix
SERVICES="
user-service
wallet-service
game-registry
"

for svc in $SERVICES; do
    fix_service "$svc"
done

echo "=== Done ==="
