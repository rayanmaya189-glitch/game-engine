"""
Fraud Detection Service

FastAPI health check only.
gRPC server for all business logic endpoints.
"""

import os
import logging
from contextlib import asynccontextmanager
from fastapi import FastAPI

from app.database import init_db
from app.grpc_server import serve_grpc

logger = logging.getLogger(__name__)

GRPC_PORT = int(os.environ.get("FRAUD_GRPC_PORT", "9115"))


@asynccontextmanager
async def lifespan(app: FastAPI):
    await init_db()
    grpc_server = await serve_grpc(GRPC_PORT)
    logger.info(f"Fraud gRPC server listening on port {GRPC_PORT}")
    yield
    await grpc_server.stop(grace=5)
    logger.info("Fraud gRPC server stopped")


app = FastAPI(
    title="Fraud Detection Service",
    description="Real-time fraud detection and prevention",
    version="1.0.0",
    lifespan=lifespan,
)


@app.get("/health")
async def health_check():
    return {"status": "healthy", "service": "fraud-service"}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=int(os.environ.get("FRAUD_PORT", "9015")))
