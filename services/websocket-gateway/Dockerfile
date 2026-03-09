# ============================================
# Stage 1: Build dependencies
# ============================================
FROM elixir:1.15-alpine AS deps

# Install build dependencies
RUN apk add --no-cache \
    build-base \
    git \
    make \
    cmake

# Set work directory
WORKDIR /app

# Copy mix files
COPY mix.exs mix.lock ./

# Install dependencies
RUN mix local.hex --force && \
    mix local.rebar --force && \
    mix deps.get --only prod && \
    mix deps.compile

# ============================================
# Stage 2: Build release
# ============================================
FROM deps AS builder

WORKDIR /app

# Copy source code
COPY . .

# Compile dependencies
RUN mix deps.compile

# Build the release
ENV MIX_ENV=prod
RUN mix release

# ============================================
# Stage 3: Production runtime
# ============================================
FROM alpine:3.18 AS app

# Install runtime dependencies
RUN apk add --no-cache \
    openssl \
    ncurses-libs \
    libstdc++ \
    libgcc

# Create non-root user
RUN adduser -D -u 1000 -s /bin/sh elixir

# Set work directory
WORKDIR /app

# Copy release from builder
COPY --from=builder --chown=elixir:elixir /app/_build/prod/rel/websocket_gateway ./

# Copy necessary files
RUN mkdir -p /app/etc /app/log && \
    chown -R elixir:elixir /app

# Switch to non-root user
USER elixir

# Environment variables
ENV MIX_ENV=prod
ENV PORT=8084
ENV NODE_NAME=websocket_gateway@localhost

# Expose ports
EXPOSE 8084 8085

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8084/health || exit 1

# Start command
CMD ["sh", "-c", "exec bin/websocket_gateway start"]
