# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY gen/ ./gen/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o winners-showcase-service ./cmd/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/winners-showcase-service .

RUN adduser -D -u 1000 appuser
USER appuser

EXPOSE 9025

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:9025/health || exit 1

ENTRYPOINT ["./winners-showcase-service"]
