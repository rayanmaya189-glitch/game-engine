"""
Risk Scoring Service

FastAPI service for unified risk profile management.
Aggregates signals from AML, Fraud, KYC, and transaction history
to produce a comprehensive risk score with automated actions.
Exposes both REST (FastAPI) and gRPC endpoints.
"""

import os
import asyncio
import logging
from contextlib import asynccontextmanager
from fastapi import FastAPI

from app.api.risk_routes import router as risk_router
from app.api.limits_routes import router as limits_router
from app.config import GRPC_PORT
from app.grpc_server import serve_grpc

logger = logging.getLogger(__name__)


@asynccontextmanager
async def lifespan(app: FastAPI):
    grpc_server = await serve_grpc(GRPC_PORT)
    logger.info(f"gRPC server listening on port {GRPC_PORT}")
    yield
    await grpc_server.stop(grace=5)
    logger.info("gRPC server stopped")


app = FastAPI(
    title="Risk Scoring Service",
    description="Unified risk scoring and automated actions",
    version="1.0.0",
    lifespan=lifespan,
)

app.include_router(risk_router)
app.include_router(limits_router)


@app.get("/health")
async def health_check():
    return {"status": "healthy", "service": "risk-service"}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=int(os.environ.get("RISK_PORT", "9016")))
