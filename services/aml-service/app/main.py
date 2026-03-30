"""
AML (Anti-Money Laundering) Detection Service

FastAPI service for detecting suspicious financial patterns
and generating regulatory reports.
"""

from fastapi import FastAPI

from app.api import transaction_routes, alert_routes, risk_routes, report_routes

app = FastAPI(
    title="AML Detection Service",
    description="Anti-money laundering detection and compliance",
    version="1.0.0"
)

app.include_router(transaction_routes.router)
app.include_router(alert_routes.router)
app.include_router(risk_routes.router)
app.include_router(report_routes.router)


@app.get("/health")
async def health_check():
    """Health check endpoint"""
    return {"status": "healthy", "service": "aml-service"}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=9014)
