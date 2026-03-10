from fastapi import APIRouter
from pydantic import BaseModel
from typing import Optional

router = APIRouter()

class RiskCheckRequest(BaseModel):
    user_id: str
    amount: float
    transaction_type: str
    ip_address: Optional[str] = None
    device_fingerprint: Optional[str] = None

@router.post("/check")
async def check_risk(request: RiskCheckRequest):
    """Check transaction for fraud risk"""
    risk_score = 0.1
    if request.amount > 5000:
        risk_score += 0.3
    return {
        "user_id": request.user_id,
        "risk_score": risk_score,
        "allowed": risk_score < 0.7,
        "requires_review": risk_score >= 0.5
    }

@router.post("/device/fingerprint")
async def register_device(user_id: str, fingerprint: str):
    """Register device fingerprint"""
    return {"user_id": user_id, "fingerprint": fingerprint, "registered": True}

@router.get("/user/{user_id}/risk")
async def get_user_risk(user_id: str):
    """Get user's risk profile"""
    return {
        "user_id": user_id,
        "risk_score": 0.2,
        "is_blocked": False,
        "flags": []
    }

@router.post("/alert")
async def create_alert(user_id: str, alert_type: str, description: str):
    """Create fraud alert"""
    return {"alert_id": "ALERT-001", "status": "created"}
