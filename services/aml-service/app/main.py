"""
AML (Anti-Money Laundering) Detection Service

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

GRPC_PORT = int(os.environ.get("AML_GRPC_PORT", "9114"))


@asynccontextmanager
async def lifespan(app: FastAPI):
    await init_db()
    grpc_server = await serve_grpc(GRPC_PORT)
    logger.info(f"AML gRPC server listening on port {GRPC_PORT}")
    yield
    await grpc_server.stop(grace=5)
    logger.info("AML gRPC server stopped")


app = FastAPI(
    title="AML Detection Service",
    description="Anti-money laundering detection and compliance",
    version="1.0.0",
    lifespan=lifespan,
)


@app.get("/health")
async def health_check():
    return {"status": "healthy", "service": "aml-service"}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=int(os.environ.get("AML_PORT", "9014")))
