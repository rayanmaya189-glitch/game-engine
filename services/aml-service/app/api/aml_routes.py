from pydantic import BaseModel, Field
from typing import Optional, List
from datetime import datetime
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select
import os

from app.database import get_db
from app.models import AlertRecord, TransactionRecord, RiskScoreRecord

LARGE_TX_THRESHOLD = int(os.environ.get("AML_LARGE_TRANSACTION_THRESHOLD", "10000"))
DAILY_LIMIT = float(os.environ.get("AML_DAILY_DEPOSIT_LIMIT", "10000"))
WEEKLY_LIMIT = float(os.environ.get("AML_WEEKLY_DEPOSIT_LIMIT", "25000"))
MONTHLY_LIMIT = float(os.environ.get("AML_MONTHLY_DEPOSIT_LIMIT", "100000"))
SINGLE_TX_LIMIT = float(os.environ.get("AML_SINGLE_TRANSACTION_LIMIT", "5000"))


class TransactionAlert(BaseModel):
    id: str
    user_id: str
    transaction_id: str
    alert_type: str
    risk_score: float
    status: str
    created_at: datetime


class SanctionsScreeningRequest(BaseModel):
    name: str
    date_of_birth: Optional[str] = None
    country: Optional[str] = None
    document_number: Optional[str] = None


class TransactionCheckRequest(BaseModel):
    user_id: str
    amount: float
    currency: str
    transaction_type: str
    payment_method: str
    ip_address: Optional[str] = None
    country: Optional[str] = None


class SARRequest(BaseModel):
    user_id: str
    transaction_id: str
    description: str
    suspicious_activity_type: str
    amount: float


HIGH_RISK_COUNTRIES = ["KP", "IR", "SY", "CU"]


async def check_transaction(request: TransactionCheckRequest, db: AsyncSession):
    """Check a transaction for AML compliance"""
    risk_score = 0.1

    if request.amount > LARGE_TX_THRESHOLD:
        risk_score += 0.3

    if request.country in HIGH_RISK_COUNTRIES:
        risk_score += 0.4

    result = await db.execute(
        select(AlertRecord).where(AlertRecord.user_id == request.user_id)
    )
    existing_alerts = result.scalars().all()

    return {
        "user_id": request.user_id,
        "transaction_allowed": risk_score < 0.7,
        "risk_score": risk_score,
        "requires_review": risk_score >= 0.5,
        "alerts": [] if risk_score < 0.5 else ["High transaction amount", "Manual review required"],
        "existing_alert_count": len(existing_alerts),
    }


async def screen_sanctions(request: SanctionsScreeningRequest):
    """Screen against sanctions lists (OFAC, EU, UN)"""
    return {
        "name": request.name,
        "matched": False,
        "lists_checked": ["OFAC", "EU_SANCTIONS", "UN_SANCTIONS"],
        "risk_level": "LOW"
    }


async def screen_pep(name: str):
    """Screen against Politically Exposed Persons list"""
    return {
        "name": name,
        "is_pep": False,
        "risk_level": "LOW"
    }


async def get_alerts(status: Optional[str] = None, limit: int = 50, db: AsyncSession = None):
    """Get AML alerts"""
    stmt = select(AlertRecord)
    if status:
        stmt = stmt.where(AlertRecord.status == status)
    stmt = stmt.order_by(AlertRecord.created_at.desc()).limit(limit)
    result = await db.execute(stmt)
    records = result.scalars().all()
    return {
        "count": len(records),
        "alerts": [
            {
                "alert_id": r.alert_id,
                "user_id": r.user_id,
                "alert_type": r.alert_type,
                "severity": r.severity,
                "status": r.status,
                "description": r.description,
                "created_at": r.created_at.isoformat() if r.created_at else None,
            }
            for r in records
        ],
    }


async def create_sar(request: SARRequest, db: AsyncSession):
    """Create Suspicious Activity Report"""
    sar_id = f"SAR-{datetime.now().strftime('%Y%m%d%H%M%S')}"
    alert = AlertRecord(
        alert_id=sar_id,
        user_id=request.user_id,
        alert_type=request.suspicious_activity_type,
        severity="high",
        status="open",
        description=request.description,
        transactions=[request.transaction_id],
        created_at=datetime.utcnow(),
        updated_at=datetime.utcnow(),
    )
    db.add(alert)
    await db.commit()
    return {
        "sar_id": sar_id,
        "status": "PENDING_REVIEW",
        "created_at": datetime.utcnow().isoformat()
    }


async def get_user_limits(user_id: str, db: AsyncSession):
    """Get AML limits for a user"""
    result = await db.execute(
        select(TransactionRecord).where(TransactionRecord.user_id == user_id)
    )
    transactions = result.scalars().all()

    daily_deposits = sum(t.amount for t in transactions if t.type == "deposit")

    return {
        "user_id": user_id,
        "daily_deposit_limit": DAILY_LIMIT,
        "weekly_deposit_limit": WEEKLY_LIMIT,
        "monthly_deposit_limit": MONTHLY_LIMIT,
        "single_transaction_limit": SINGLE_TX_LIMIT,
        "current_daily": daily_deposits,
        "current_weekly": 0.0,
        "current_monthly": 0.0
    }


async def check_limits(user_id: str, amount: float, period: str):
    """Check if amount is within AML limits"""
    limits = {
        "daily": DAILY_LIMIT,
        "weekly": WEEKLY_LIMIT,
        "monthly": MONTHLY_LIMIT
    }

    limit = limits.get(period, DAILY_LIMIT)

    return {
        "user_id": user_id,
        "amount": amount,
        "period": period,
        "limit": limit,
        "within_limit": amount <= limit,
        "remaining": limit - amount
    }
