#!/bin/bash
# =============================================================================
# Complete Service Fixer - Fixes all Go services for deployment
# =============================================================================
set -e

BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$BASE_DIR"

echo "=========================================="
echo "  Starting Complete Service Build"
echo "=========================================="

fix_and_build_service() {
    local svc=$1
    local dir="services/$svc"
    
    echo ""
    echo "=== Processing: $svc ==="
    
    if [ ! -d "$dir" ]; then
        echo "SKIP: Directory not found"
        return 1
    fi
    
    if [ ! -f "$dir/go.mod" ]; then
        echo "SKIP: No go.mod"
        return 1
    fi
    
    cd "$dir"
    
    # 1. Fix go.mod
    sed -i 's/go 1.2[0-9]/go 1.24/g' go.mod 2>/dev/null || true
    sed -i 's/go 1.24.0/go 1.24/g' go.mod 2>/dev/null || true
    
    # 2. Fix NATS version
    if grep -q "nats.go v1.5" go.mod; then
        sed -i 's/nats.go v1.5/nats.go v1.28.0/g' go.mod
    fi
    if grep -q "nats.go v1.6" go.mod; then
        sed -i 's/nats.go v1.6/nats.go v1.28.0/g' go.mod
    fi
    if ! grep -q "nats.go v1.2" go.mod; then
        if ! grep -q "nats.go v1.28" go.mod && ! grep -q "github.com/nats-io/nats.go" go.mod; then
            sed -i '/^require (/a\	github.com/nats-io/nats.go v1.28.0' go.mod 2>/dev/null || true
        fi
    fi
    
    # 3. Add replace directive
    if ! grep -q "replace.*nats-server" go.mod; then
        echo "" >> go.mod
        echo "replace github.com/nats-io/nats-server/v2 => github.com/nats-io/nats-server/v2 v2.9.0" >> go.mod
    fi
    
    # 4. Fix grpc version
    if grep -q "grpc v1.6" go.mod || grep -q "grpc v1.6" go.mod; then
        sed -i 's/grpc v1.6[0-9]*/grpc v1.68.1/g' go.mod
    fi
    
    # 5. Fix protobuf version
    if grep -q "protobuf v1.3" go.mod || grep -q "protobuf v1.3" go.mod; then
        sed -i 's/protobuf v1.3[0-9]*/protobuf v1.35.1/g' go.mod
    fi
    
    # 6. Fix validator version
    if grep -q "validator/v10 v10.22" go.mod; then
        sed -i 's/validator\/v10 v10.22/validator\/v10 v10.16.0/g' go.mod
    fi
    
    # 7. Fix Dockerfile
    if [ -f "Dockerfile" ]; then
        sed -i 's/FROM golang:1.2[0-9]-alpine/FROM golang:1.24-alpine/g' Dockerfile
        sed -i 's/FROM golang:1.24.0-alpine/FROM golang:1.24-alpine/g' Dockerfile
    fi
    
    # 8. Copy proto files
    if [ -d "$BASE_DIR/services/common-service/proto/gen/go/game_engine" ]; then
        mkdir -p pkg/game_engine
        cp -r "$BASE_DIR/services/common-service/proto/gen/go/game_engine"/* pkg/game_engine/ 2>/dev/null || true
    fi
    
    # 9. Fix proto imports in pkg
    if [ -d "pkg" ]; then
        find pkg -name "*.go" -exec sed -i 's|github.com/game_engine/common/proto/gen/go/game_engine|github.com/game_engine/'"$svc"'/pkg/game_engine|g' {} \; 2>/dev/null || true
        find pkg -name "*.go" -exec sed -i 's|"gen/go|"github.com/game_engine/'"$svc"'/pkg/game_engine/|g' {} \; 2>/dev/null || true
        find pkg -name "*.go" -exec sed -i 's|gen/go/|github.com/game_engine/'"$svc"'/pkg/game_engine/|g' {} \; 2>/dev/null || true
        find pkg -name "*.go" -exec sed -i 's|"gen/|"github.com/game_engine/'"$svc"'/pkg/game_engine/|g' {} \; 2>/dev/null || true
    fi
    
    # 10. Fix source imports
    for f in $(find . -name "*.go" -path "*/cmd/*" -o -name "*.go" -path "*/internal/*" 2>/dev/null); do
        sed -i 's|game_engine/gen/go|github.com/game_engine/'"$svc"'/pkg/game_engine|g' "$f" 2>/dev/null || true
        sed -i 's|game_engine/common/proto|github.com/game_engine/'"$svc"'/pkg/game_engine|g' "$f" 2>/dev/null || true
        sed -i 's|gen/go|github.com/game_engine/'"$svc"'/pkg/game_engine|g' "$f" 2>/dev/null || true
    done
    
    # 11. Run go mod tidy
    rm -f go.sum go.mod.bak
    docker run --rm -v "$(pwd):/app" -w /app -e GOTOOLCHAIN=local golang:1.24-alpine sh -c "apk add --no-cache git && go mod tidy" 2>&1 | grep -E "(downloading|found|requires|error)" | tail -5 || true
    
    # 12. Build
    if docker build -t "game-engine/$svc:latest" . 2>&1 | tail -3 | grep -q "Successfully"; then
        echo "SUCCESS: $svc built"
        return 0
    else
        echo "FAILED: $svc build failed"
        return 1
    fi
}

export -f fix_and_build_service
export BASE_DIR

# List of all Go services to fix
SERVICES="
auth-service
user-service
wallet-service
game-registry
card-games
dice-games
slot-games
rng-service
betting
tournament
jackpot-service
live-dealer-service
sports-betting-service
game-engine
multiplayer
chat
notification
leaderboard-service
winners-showcase-service
merchant-service
agent-service
loyalty-service
banner-service
referral-service
"

SUCCESS_COUNT=0
FAILED_COUNT=0

for svc in $SERVICES; do
    if fix_and_build_service "$svc"; then
        ((SUCCESS_COUNT++))
    else
        ((FAILED_COUNT++))
    fi
    cd "$BASE_DIR"
done

echo ""
echo "=========================================="
echo "  Build Complete"
echo "=========================================="
echo "  SUCCESS: $SUCCESS_COUNT services"
echo "  FAILED:  $FAILED_COUNT services"
echo "=========================================="

# Show built images
echo ""
echo "Built images:"
docker images --format "table {{.Repository}}\t{{.Tag}}" | grep game-engine | head -20
