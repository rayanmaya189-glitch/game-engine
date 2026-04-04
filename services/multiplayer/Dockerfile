# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make protoc

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY configs/ ./configs/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o multiplayer ./cmd/main.go

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/multiplayer .

COPY configs/ ./configs/

RUN adduser -D -u 1000 appuser
USER appuser

EXPOSE 9021 8021

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8021/health || exit 1

ENTRYPOINT ["./multiplayer"]
