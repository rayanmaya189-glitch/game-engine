"""
Risk Scoring Service

FastAPI health check only.
gRPC server for all business logic endpoints.
"""

import os
import logging
from contextlib import asynccontextmanager
from fastapi import FastAPI

from app.grpc_server import serve_grpc

logger = logging.getLogger(__name__)

GRPC_PORT = int(os.environ.get("RISK_GRPC_PORT", "9116"))


@asynccontextmanager
async def lifespan(app: FastAPI):
    grpc_server = await serve_grpc(GRPC_PORT)
    logger.info(f"Risk gRPC server listening on port {GRPC_PORT}")
    yield
    await grpc_server.stop(grace=5)
    logger.info("Risk gRPC server stopped")


app = FastAPI(
    title="Risk Scoring Service",
    description="Unified risk scoring and automated actions",
    version="1.0.0",
    lifespan=lifespan,
)


@app.get("/health")
async def health_check():
    return {"status": "healthy", "service": "risk-service"}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=int(os.environ.get("RISK_PORT", "9016")))
