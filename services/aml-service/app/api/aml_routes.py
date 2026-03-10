from fastapi import APIRouter, HTTPException
from pydantic import BaseModel, Field
from typing import Optional, List
from datetime import datetime

router = APIRouter()

# Models
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


# Routes
@router.post("/transaction/check")
async def check_transaction(request: TransactionCheckRequest):
    """Check a transaction for AML compliance"""
    # Mock implementation - would integrate with actual monitoring systems
    risk_score = 0.1  # Low risk by default
    
    # Check for suspicious patterns
    if request.amount > 10000:
        risk_score += 0.3
    
    # Check for high-risk countries
    high_risk_countries = ["KP", "IR", "SY", "CU"]
    if request.country in high_risk_countries:
        risk_score += 0.4
    
    return {
        "user_id": request.user_id,
        "transaction_allowed": risk_score < 0.7,
        "risk_score": risk_score,
        "requires_review": risk_score >= 0.5,
        "alerts": [] if risk_score < 0.5 else ["High transaction amount", "Manual review required"]
    }


@router.post("/sanctions/screen")
async def screen_sanctions(request: SanctionsScreeningRequest):
    """Screen against sanctions lists (OFAC, EU, UN)"""
    # Mock implementation - would integrate with actual sanctions list providers
    return {
        "name": request.name,
        "matched": False,
        "lists_checked": ["OFAC", "EU_SANCTIONS", "UN_SANCTIONS"],
        "risk_level": "LOW"
    }


@router.post("/pep/screen")
async def screen_pep(name: str):
    """Screen against Politically Exposed Persons list"""
    # Mock implementation
    return {
        "name": name,
        "is_pep": False,
        "risk_level": "LOW"
    }


@router.get("/alerts")
async def get_alerts(status: Optional[str] = None, limit: int = 50):
    """Get AML alerts"""
    return {
        "count": 0,
        "alerts": []
    }


@router.post("/sar")
async def create_sar(request: SARRequest):
    """Create Suspicious Activity Report"""
    sar_id = f"SAR-{datetime.now().strftime('%Y%m%d')}-001"
    return {
        "sar_id": sar_id,
        "status": "PENDING_REVIEW",
        "created_at": datetime.now().isoformat()
    }


@router.get("/limits/user/{user_id}")
async def get_user_limits(user_id: str):
    """Get AML limits for a user"""
    return {
        "user_id": user_id,
        "daily_deposit_limit": 10000.0,
        "weekly_deposit_limit": 25000.0,
        "monthly_deposit_limit": 100000.0,
        "single_transaction_limit": 5000.0,
        "current_daily": 1500.0,
        "current_weekly": 5000.0,
        "current_monthly": 15000.0
    }


@router.post("/limits/check")
async def check_limits(user_id: str, amount: float, period: str):
    """Check if amount is within AML limits"""
    limits = {
        "daily": 10000.0,
        "weekly": 25000.0,
        "monthly": 100000.0
    }
    
    limit = limits.get(period, 10000.0)
    
    return {
        "user_id": user_id,
        "amount": amount,
        "period": period,
        "limit": limit,
        "within_limit": amount <= limit,
        "remaining": limit - amount
    }
