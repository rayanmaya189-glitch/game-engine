from contextlib import asynccontextmanager
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
import logging

from app.api import aml_routes

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


@asynccontextmanager
async def lifespan(app: FastAPI):
    logger.info("Starting AML Service...")
    yield
    logger.info("Shutting down AML Service...")


app = FastAPI(
    title="AML Service",
    description="Anti-Money Laundering service for transaction monitoring and sanctions screening",
    version="1.0.0",
    lifespan=lifespan
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(aml_routes.router, prefix="/api/v1/aml", tags=["AML"])


@app.get("/health")
async def health_check():
    return {"status": "healthy", "service": "aml-service"}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=9032)
