"""
AML (Anti-Money Laundering) Detection Service

FastAPI service for detecting suspicious financial patterns
and generating regulatory reports.
"""

import os
from contextlib import asynccontextmanager
from fastapi import FastAPI

from app.database import init_db
from app.api import transaction_routes, alert_routes, risk_routes, report_routes, aml_routes


@asynccontextmanager
async def lifespan(app: FastAPI):
    await init_db()
    yield


app = FastAPI(
    title="AML Detection Service",
    description="Anti-money laundering detection and compliance",
    version="1.0.0",
    lifespan=lifespan,
)

app.include_router(transaction_routes.router)
app.include_router(alert_routes.router)
app.include_router(risk_routes.router)
app.include_router(report_routes.router)
app.include_router(aml_routes.router)


@app.get("/health")
async def health_check():
    """Health check endpoint"""
    return {"status": "healthy", "service": "aml-service"}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=int(os.environ.get("AML_PORT", "9014")))
