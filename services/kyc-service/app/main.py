"""
KYC Service

FastAPI health check only.
gRPC server for all business logic endpoints.
"""

import os
import logging
from contextlib import asynccontextmanager
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from app.database import engine, Base
from app.config import settings
from app.grpc_server import serve_grpc

logger = logging.getLogger(__name__)

GRPC_PORT = int(os.environ.get("KYC_GRPC_PORT", "9131"))


@asynccontextmanager
async def lifespan(app: FastAPI):
    logger.info("Starting KYC Service...")

    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)

    grpc_server = await serve_grpc(GRPC_PORT)
    logger.info(f"KYC gRPC server listening on port {GRPC_PORT}")

    logger.info("KYC Service started successfully")
    yield

    await grpc_server.stop(grace=5)
    logger.info("KYC gRPC server stopped")
    await engine.dispose()


app = FastAPI(
    title="KYC Service",
    description="Know Your Customer service for identity verification",
    version="1.0.0",
    lifespan=lifespan,
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=settings.cors_origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.get("/health")
async def health_check():
    return {"status": "healthy", "service": "kyc-service"}


@app.get("/ready")
async def readiness_check():
    return {"status": "ready", "service": "kyc-service"}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(
        "main:app",
        host="0.0.0.0",
        port=9031,
        reload=settings.debug,
    )
