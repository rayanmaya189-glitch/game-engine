# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY configs/ ./configs/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o merchant-service ./cmd/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/merchant-service .
COPY configs/ ./configs/

RUN adduser -D -u 1000 appuser
USER appuser

EXPOSE 9095

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:9095/health || exit 1

ENTRYPOINT ["./merchant-service"]
