"""
Risk Scoring Service

FastAPI service for unified risk profile management.
Aggregates signals from AML, Fraud, KYC, and transaction history
to produce a comprehensive risk score with automated actions.
"""

from fastapi import FastAPI

from app.api.risk_routes import router as risk_router
from app.api.limits_routes import router as limits_router

app = FastAPI(
    title="Risk Scoring Service",
    description="Unified risk scoring and automated actions",
    version="1.0.0"
)

app.include_router(risk_router)
app.include_router(limits_router)


@app.get("/health")
async def health_check():
    return {"status": "healthy", "service": "risk-service"}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=9016)
