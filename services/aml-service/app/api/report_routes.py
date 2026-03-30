from fastapi import APIRouter, HTTPException
from typing import Dict
from datetime import datetime

from app.models.schemas import alerts_db, transactions_db

router = APIRouter(prefix="", tags=["reports"])


@router.get("/reports/ctr", response_model=Dict)
async def generate_ctr_report(start_date: str, end_date: str):
    """Generate Currency Transaction Report for transactions > $10K"""
    start = datetime.fromisoformat(start_date)
    end = datetime.fromisoformat(end_date)

    large_transactions = [
        t for t in transactions_db.values()
        if t.timestamp > start and t.timestamp < end
        and t.amount > 10000
    ]

    return {
        "report_type": "CTR",
        "period": {"start": start_date, "end": end_date},
        "transactions": [t.dict() for t in large_transactions],
        "total_count": len(large_transactions)
    }


@router.get("/reports/sar/{alert_id}", response_model=Dict)
async def generate_sar_report(alert_id: str):
    """Generate Suspicious Activity Report for an alert"""
    if alert_id not in alerts_db:
        raise HTTPException(status_code=404, detail="Alert not found")

    alert = alerts_db[alert_id]

    return {
        "report_type": "SAR",
        "alert": alert.dict(),
        "generated_at": datetime.now().isoformat(),
        "format": "FinCEN BSA"
    }
