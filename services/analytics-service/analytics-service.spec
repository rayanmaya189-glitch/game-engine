name: Analytics Service
version: 1.0.0
description: Business intelligence and reporting service

services:
  - name: analytics-service
    type: Python FastAPI
    ports:
      http: 9020
      grpc: 9120
    dependencies:
      - postgresql
      - redis
      - timescaledb
    health_check: /health
