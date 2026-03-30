from fastapi import APIRouter, HTTPException, Depends
from pydantic import BaseModel
from typing import Optional, List
from datetime import datetime
import random
import json

from sqlalchemy.ext.asyncio import AsyncSession
from app.database import get_db
from app.models import FraudAlertRecord, UserRiskProfileRecord
from app.repositories import (
    DeviceFingerprintRepository,
    FraudAlertRepository,
    UserRiskProfileRepository,
)

router = APIRouter()


class RiskCheckRequest(BaseModel):
    user_id: str
    amount: float
    transaction_type: str
    ip_address: Optional[str] = None
    device_fingerprint: Optional[str] = None
    geo_location: Optional[str] = None


class BetCheckRequest(BaseModel):
    user_id: str
    game_id: str
    amount: float
    ip_address: Optional[str] = None
    device_fingerprint: Optional[str] = None


class AccountActivityRequest(BaseModel):
    user_id: str
    activity_type: str  # login, password_change, withdrawal, etc.
    ip_address: Optional[str] = None
    device_fingerprint: Optional[str] = None
    new_value: Optional[str] = None


class AlertRequest(BaseModel):
    user_id: str
    alert_type: str
    description: str
    evidence: Optional[dict] = None


@router.post("/check")
async def check_risk(request: RiskCheckRequest, db: AsyncSession = Depends(get_db)):
    """Check transaction for fraud risk"""
    risk_score = 0.0
    risk_factors = []

    if request.amount > 5000:
        risk_score += 0.2
        risk_factors.append("High transaction amount")
    if request.amount > 10000:
        risk_score += 0.2
        risk_factors.append("Very high transaction amount")

    if request.device_fingerprint:
        fp = await DeviceFingerprintRepository.get_by_user(db, request.user_id)
        if fp and fp.canvas_hash != request.device_fingerprint:
            risk_score += 0.3
            risk_factors.append("New device detected")

    if request.ip_address:
        if request.ip_address.startswith('10.') or request.ip_address.startswith('192.'):
            risk_score += 0.1
            risk_factors.append("Internal IP detected")

    profile = await UserRiskProfileRepository.get_or_create(db, request.user_id)
    profile.transaction_count = (profile.transaction_count or 0) + 1
    profile.last_activity = datetime.utcnow().isoformat()
    existing_score = profile.risk_score or 0.0
    profile.risk_score = (existing_score + risk_score) / 2

    risk_score += (profile.risk_score or 0.0) * 0.3

    is_blocked = risk_score > 0.8
    allowed = risk_score < 0.7 and not is_blocked

    await UserRiskProfileRepository.save(db, profile)

    return {
        "user_id": request.user_id,
        "risk_score": min(risk_score, 1.0),
        "allowed": allowed,
        "is_blocked": is_blocked,
        "requires_review": risk_score >= 0.5,
        "risk_factors": risk_factors,
        "transaction_count": profile.transaction_count,
    }


@router.post("/check/bet")
async def check_bet_risk(request: BetCheckRequest, db: AsyncSession = Depends(get_db)):
    """Check bet for fraud patterns (bonus abuse, collusion, etc.)"""
    risk_score = 0.0
    risk_factors = []

    if request.amount > 1000:
        risk_score += 0.15
        risk_factors.append("Unusually high bet amount")

    profile = await UserRiskProfileRepository.get_or_create(db, request.user_id)

    if profile.last_bet_time:
        try:
            last_bet = datetime.fromisoformat(profile.last_bet_time)
            time_diff = (datetime.utcnow() - last_bet).total_seconds()
            if time_diff < 2:
                risk_score += 0.25
                risk_factors.append("Rapid betting pattern detected")
        except (ValueError, TypeError):
            pass

    profile.last_bet_time = datetime.utcnow().isoformat()
    await UserRiskProfileRepository.save(db, profile)

    return {
        "user_id": request.user_id,
        "game_id": request.game_id,
        "risk_score": min(risk_score, 1.0),
        "allowed": risk_score < 0.6,
        "requires_review": risk_score >= 0.4,
        "risk_factors": risk_factors,
    }


@router.post("/check/account")
async def check_account_activity(request: AccountActivityRequest, db: AsyncSession = Depends(get_db)):
    """Check for account takeover attempts"""
    risk_score = 0.0
    risk_factors = []

    high_risk_activities = ["withdrawal", "password_change", "email_change", "2fa_disable"]
    if request.activity_type in high_risk_activities:
        risk_score += 0.3
        risk_factors.append(f"High-risk activity: {request.activity_type}")

    if request.device_fingerprint:
        fp = await DeviceFingerprintRepository.get_by_user(db, request.user_id)
        if fp and fp.canvas_hash != request.device_fingerprint:
            risk_score += 0.4
            risk_factors.append("Account access from new device")

    if request.ip_address:
        profile = await UserRiskProfileRepository.get_or_create(db, request.user_id)
        if profile.last_ip and profile.last_ip != request.ip_address:
            risk_score += 0.2
            risk_factors.append("IP address changed")

    return {
        "user_id": request.user_id,
        "activity_type": request.activity_type,
        "risk_score": min(risk_score, 1.0),
        "allowed": risk_score < 0.7,
        "requires_review": risk_score >= 0.5,
        "risk_factors": risk_factors,
        "recommendation": "block" if risk_score > 0.8 else "allow" if risk_score < 0.5 else "review",
    }


@router.get("/user/{user_id}/risk")
async def get_user_risk(user_id: str, db: AsyncSession = Depends(get_db)):
    """Get user's risk profile"""
    profile = await UserRiskProfileRepository.get_or_create(db, user_id)
    profile.is_blocked = (profile.risk_score or 0) > 0.8
    await UserRiskProfileRepository.save(db, profile)

    flags = json.loads(profile.flags) if isinstance(profile.flags, str) else (profile.flags or [])

    return {
        "user_id": profile.user_id,
        "risk_score": profile.risk_score or 0.0,
        "is_blocked": profile.is_blocked,
        "flags": flags,
        "transaction_count": profile.transaction_count or 0,
    }


@router.post("/alert")
async def create_alert(request: AlertRequest, db: AsyncSession = Depends(get_db)):
    """Create fraud alert"""
    alert_id = f"ALERT-{datetime.now().strftime('%Y%m%d%H%M%S')}-{random.randint(1000, 9999)}"

    alert = FraudAlertRecord(
        alert_id=alert_id,
        user_id=request.user_id,
        alert_type=request.alert_type,
        description=request.description,
        evidence=request.evidence or {},
        status="open",
        created_at=datetime.utcnow(),
    )
    await FraudAlertRepository.save(db, alert)

    profile = await UserRiskProfileRepository.get_or_create(db, request.user_id)
    flags = json.loads(profile.flags) if isinstance(profile.flags, str) else (profile.flags or [])
    flags.append(request.alert_type)
    profile.flags = json.dumps(flags)
    profile.risk_score = min((profile.risk_score or 0) + 0.1, 1.0)
    await UserRiskProfileRepository.save(db, profile)

    return {
        "alert_id": alert_id,
        "user_id": request.user_id,
        "alert_type": request.alert_type,
        "description": request.description,
        "evidence": request.evidence or {},
        "status": "open",
        "created_at": datetime.utcnow().isoformat(),
    }


@router.get("/alerts")
async def get_alerts(status: Optional[str] = None, limit: int = 50, db: AsyncSession = Depends(get_db)):
    """Get fraud alerts"""
    records = await FraudAlertRepository.list_alerts(db, status=status, limit=limit)
    return {
        "count": len(records),
        "alerts": [
            {
                "alert_id": r.alert_id,
                "user_id": r.user_id,
                "alert_type": r.alert_type,
                "description": r.description,
                "evidence": r.evidence or {},
                "status": r.status,
                "created_at": r.created_at.isoformat() if r.created_at else None,
            }
            for r in records
        ],
    }


@router.post("/user/{user_id}/unblock")
async def unblock_user(user_id: str, db: AsyncSession = Depends(get_db)):
    """Unblock a user after review"""
    profile = await UserRiskProfileRepository.get_or_create(db, user_id)
    profile.is_blocked = False
    profile.risk_score = 0.0
    profile.flags = "[]"
    await UserRiskProfileRepository.save(db, profile)
    return {"user_id": user_id, "unblocked": True}


@router.post("/user/{user_id}/block")
async def block_user(user_id: str, reason: str, db: AsyncSession = Depends(get_db)):
    """Block a user"""
    profile = await UserRiskProfileRepository.get_or_create(db, user_id)
    profile.is_blocked = True
    profile.block_reason = reason
    profile.blocked_at = datetime.utcnow().isoformat()
    await UserRiskProfileRepository.save(db, profile)
    return {"user_id": user_id, "blocked": True, "reason": reason}
