# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY common/ ./common/
COPY agent/ ./agent/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gateway-agent ./agent/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/gateway-agent .
COPY agent/config.yaml ./config.yaml

RUN adduser -D -u 1000 appuser
USER appuser

EXPOSE 8083

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8083/health || exit 1

ENTRYPOINT ["./gateway-agent"]
