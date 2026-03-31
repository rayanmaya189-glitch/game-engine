# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY common/ ./common/
COPY merchant/ ./merchant/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gateway-merchant ./merchant/main.go

# Runtime stage
FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/gateway-merchant .
COPY merchant/config.yaml ./config.yaml

RUN adduser -D -u 1000 appuser
USER appuser

EXPOSE 8082

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8082/health || exit 1

ENTRYPOINT ["./gateway-merchant"]
