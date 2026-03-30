"""Fraud Detection Service

FastAPI service for detecting fraud patterns including:
- Multi-account detection (device fingerprinting, IP correlation)
- Bot detection
- Collusion detection for poker
- Real-time fraud scoring
"""

import os
from contextlib import asynccontextmanager
from fastapi import FastAPI

from app.database import init_db
from app.api import fingerprint_routes, bot_routes, collusion_routes, score_routes, fraud_routes


@asynccontextmanager
async def lifespan(app: FastAPI):
    await init_db()
    yield


app = FastAPI(
    title="Fraud Detection Service",
    description="Real-time fraud detection and prevention",
    version="1.0.0",
    lifespan=lifespan,
)

app.include_router(fingerprint_routes.router)
app.include_router(bot_routes.router)
app.include_router(collusion_routes.router)
app.include_router(score_routes.router)
app.include_router(fraud_routes.router)


@app.get("/health")
async def health_check():
    return {"status": "healthy", "service": "fraud-service"}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=int(os.environ.get("FRAUD_PORT", "9015")))
