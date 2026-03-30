from fastapi import APIRouter, HTTPException, Depends
from typing import Dict
from datetime import datetime
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select

from app.database import get_db
from app.models import AlertRecord, TransactionRecord

router = APIRouter(prefix="", tags=["reports"])


@router.get("/reports/ctr", response_model=Dict)
async def generate_ctr_report(start_date: str, end_date: str, db: AsyncSession = Depends(get_db)):
    """Generate Currency Transaction Report for transactions > $10K"""
    start = datetime.fromisoformat(start_date)
    end = datetime.fromisoformat(end_date)

    stmt = select(TransactionRecord).where(
        TransactionRecord.timestamp > start,
        TransactionRecord.timestamp < end,
        TransactionRecord.amount > 10000,
    )
    result = await db.execute(stmt)
    records = result.scalars().all()

    large_transactions = [
        {
            "transaction_id": t.transaction_id,
            "user_id": t.user_id,
            "type": t.type,
            "amount": t.amount,
            "currency": t.currency,
            "payment_method": t.payment_method,
            "ip_address": t.ip_address,
            "country": t.country,
            "timestamp": t.timestamp.isoformat(),
        }
        for t in records
    ]

    return {
        "report_type": "CTR",
        "period": {"start": start_date, "end": end_date},
        "transactions": large_transactions,
        "total_count": len(large_transactions)
    }


@router.get("/reports/sar/{alert_id}", response_model=Dict)
async def generate_sar_report(alert_id: str, db: AsyncSession = Depends(get_db)):
    """Generate Suspicious Activity Report for an alert"""
    result = await db.execute(select(AlertRecord).where(AlertRecord.alert_id == alert_id))
    record = result.scalar_one_or_none()
    if not record:
        raise HTTPException(status_code=404, detail="Alert not found")

    alert_data = {
        "alert_id": record.alert_id,
        "user_id": record.user_id,
        "alert_type": record.alert_type,
        "severity": record.severity,
        "status": record.status,
        "description": record.description,
        "transactions": record.transactions or [],
        "assigned_to": record.assigned_to,
        "created_at": record.created_at.isoformat(),
        "updated_at": record.updated_at.isoformat(),
        "resolved_at": record.resolved_at.isoformat() if record.resolved_at else None,
        "notes": record.notes,
    }

    return {
        "report_type": "SAR",
        "alert": alert_data,
        "generated_at": datetime.utcnow().isoformat(),
        "format": "FinCEN BSA"
    }
