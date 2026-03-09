#!/bin/bash
# Proto code generation script
# Generates Go, Java, Python, and TypeScript code from proto definitions

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROTO_DIR="$SCRIPT_DIR/proto"
GEN_DIR="$PROTO_DIR/gen"

echo "=== Proto Code Generation ==="
echo "Proto directory: $PROTO_DIR"
echo "Generated code directory: $GEN_DIR"
echo ""

# Check if buf is installed
if ! command -v buf &> /dev/null; then
    echo "Error: buf is not installed"
    echo "Install buf: https://buf.build/docs/installation"
    exit 1
fi

# Clean generated directories
echo "Cleaning generated directories..."
rm -rf "$GEN_DIR/go" "$GEN_DIR/java" "$GEN_DIR/python" "$GEN_DIR/ts"
mkdir -p "$GEN_DIR/go" "$GEN_DIR/java" "$GEN_DIR/python" "$GEN_DIR/ts"

# Lint proto files
echo "Linting proto files..."
cd "$PROTO_DIR"
buf lint

# Generate code
echo "Generating code..."
buf generate

echo ""
echo "=== Generation Complete ==="
echo "Generated files:"
echo "  - Go:       $GEN_DIR/go"
echo "  - Java:     $GEN_DIR/java"
echo "  - Python:   $GEN_DIR/python"
echo "  - TypeScript: $GEN_DIR/ts"
echo ""

# Show generated files
echo "Generated Go files:"
find "$GEN_DIR/go" -name "*.pb.go" | head -20
echo ""
echo "Generated TypeScript files:"
find "$GEN_DIR/ts" -name "*.ts" | head -20
