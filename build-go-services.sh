#!/usr/bin/env bash
set -euo pipefail

# Set path for Go
export PATH="/home/ali/.goenv/shims:$PATH"

BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BIN_DIR="$BASE_DIR/bin"
mkdir -p "$BIN_DIR"

echo "=========================================="
echo "  Starting Go Services Local Build"
echo "=========================================="
echo "Go version: $(go version)"

# Declare associative array of services and their entry points relative to services/
# format: "binary_name:relative_path_from_services_dir"
SERVICES=(
    "auth-service:auth-service/cmd"
    "user-service:user-service/cmd"
    "wallet-service:wallet-service/cmd"
    "game-registry:game-registry/cmd"
    "game-engine:game-engine/cmd"
    "card-games:card-games/cmd"
    "dice-games:dice-games/cmd"
    "slot-games:slot-games/cmd"
    "rng-service:rng-service/cmd"
    "betting:betting/cmd"
    "jackpot-service:jackpot-service/cmd"
    "live-dealer-service:live-dealer-service/cmd"
    "sports-betting-service:sports-betting-service/cmd"
    "leaderboard-service:leaderboard-service/cmd"
    "winners-showcase-service:winners-showcase-service/cmd"
    "multiplayer:multiplayer/cmd"
    "chat:chat/cmd"
    "notification:notification/cmd"
    "merchant-service:merchant-service/cmd"
    "agent-service:agent-service/cmd"
    "loyalty-service:loyalty-service/cmd"
    "banner-service:banner-service/cmd"
    "referral-service:referral-service/cmd"
    "gateway-admin:gateway/admin"
    "gateway-agent:gateway/agent"
    "gateway-merchant:gateway/merchant"
    "gateway-player:gateway/player"
)

SUCCESS_COUNT=0
FAILED_COUNT=0
FAILED_SERVICES=()

# 1. First, build the services using the main services module (excluding tournament)
cd "$BASE_DIR/services"

# Run go mod tidy to ensure everything is resolved
echo "Running go mod tidy..."
go mod tidy

for item in "${SERVICES[@]}"; do
    bin_name="${item%%:*}"
    rel_path="${item##*:}"
    
    echo -n "Building $bin_name... "
    if go build -o "$BIN_DIR/$bin_name" "./$rel_path" > "/tmp/build_${bin_name}.log" 2>&1; then
        echo -e "\e[32mSUCCESS\e[0m"
        SUCCESS_COUNT=$((SUCCESS_COUNT + 1))
    else
        echo -e "\e[31mFAILED\e[0m"
        FAILED_COUNT=$((FAILED_COUNT + 1))
        FAILED_SERVICES+=("$bin_name")
        cat "/tmp/build_${bin_name}.log"
    fi
done

# 2. Build tournament (which has its own go.mod)
echo -n "Building tournament... "
cd "$BASE_DIR/services/tournament"
go mod tidy
if go build -o "$BIN_DIR/tournament" ./cmd > "/tmp/build_tournament.log" 2>&1; then
    echo -e "\e[32mSUCCESS\e[0m"
    SUCCESS_COUNT=$((SUCCESS_COUNT + 1))
else
    echo -e "\e[31mFAILED\e[0m"
    FAILED_COUNT=$((FAILED_COUNT + 1))
    FAILED_SERVICES+=("tournament")
    cat "/tmp/build_tournament.log"
fi

echo "=========================================="
echo "  Build Results"
echo "=========================================="
echo "  SUCCESS: $SUCCESS_COUNT"
echo "  FAILED:  $FAILED_COUNT"
if [ ${#FAILED_SERVICES[@]} -gt 0 ]; then
    echo "  Failed services: ${FAILED_SERVICES[*]}"
fi
echo "=========================================="
