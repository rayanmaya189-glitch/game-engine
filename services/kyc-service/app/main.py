from contextlib import asynccontextmanager
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
import logging

from app.api import kyc_routes
from app.database import engine, Base
from app.config import settings

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(name)s - %(levelname)s - %(message)s"
)
logger = logging.getLogger(__name__)


@asynccontextmanager
async def lifespan(app: FastAPI):
    # Startup
    logger.info("Starting KYC Service...")
    
    # Create tables
    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)
    
    logger.info("KYC Service started successfully")
    
    yield
    
    # Shutdown
    logger.info("Shutting down KYC Service...")
    await engine.dispose()


app = FastAPI(
    title="KYC Service",
    description="Know Your Customer service for identity verification",
    version="1.0.0",
    lifespan=lifespan
)

# CORS middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=settings.cors_origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Include routers
app.include_router(kyc_routes.router, prefix="/api/v1/kyc", tags=["KYC"])


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
        reload=settings.debug
    )
