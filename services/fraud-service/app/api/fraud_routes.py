from fastapi import APIRouter, HTTPException
from pydantic import BaseModel
from typing import Optional, List
from datetime import datetime
import random

router = APIRouter()

# Models
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


# In-memory storage for demo
alerts_db = []
user_risk_profiles = {}
device_fingerprints = {}


@router.post("/check")
async def check_risk(request: RiskCheckRequest):
    """Check transaction for fraud risk"""
    risk_score = 0.0
    risk_factors = []
    
    # Amount-based risk
    if request.amount > 5000:
        risk_score += 0.2
        risk_factors.append("High transaction amount")
    if request.amount > 10000:
        risk_score += 0.2
        risk_factors.append("Very high transaction amount")
    
    # Device fingerprint check
    if request.device_fingerprint:
        if request.user_id in device_fingerprints:
            if device_fingerprints[request.user_id] != request.device_fingerprint:
                risk_score += 0.3
                risk_factors.append("New device detected")
        else:
            device_fingerprints[request.user_id] = request.device_fingerprint
    
    # IP-based risk
    if request.ip_address:
        # Check for VPN/proxy (simplified)
        if request.ip_address.startswith('10.') or request.ip_address.startswith('192.'):
            risk_score += 0.1
            risk_factors.append("Internal IP detected")
    
    # Initialize user profile if not exists
    if request.user_id not in user_risk_profiles:
        user_risk_profiles[request.user_id] = {
            "user_id": request.user_id,
            "risk_score": 0.0,
            "is_blocked": False,
            "flags": [],
            "transaction_count": 0,
            "last_activity": None
        }
    
    # Update user profile
    profile = user_risk_profiles[request.user_id]
    profile["transaction_count"] += 1
    profile["last_activity"] = datetime.now().isoformat()
    profile["risk_score"] = (profile["risk_score"] + risk_score) / 2
    
    # Add accumulated risk
    risk_score += profile["risk_score"] * 0.3
    
    is_blocked = risk_score > 0.8
    allowed = risk_score < 0.7 and not is_blocked
    
    return {
        "user_id": request.user_id,
        "risk_score": min(risk_score, 1.0),
        "allowed": allowed,
        "is_blocked": is_blocked,
        "requires_review": risk_score >= 0.5,
        "risk_factors": risk_factors,
        "transaction_count": profile["transaction_count"]
    }


@router.post("/check/bet")
async def check_bet_risk(request: BetCheckRequest):
    """Check bet for fraud patterns (bonus abuse, collusion, etc.)"""
    risk_score = 0.0
    risk_factors = []
    
    # Pattern: Unusual bet size for user's history
    if request.amount > 1000:
        risk_score += 0.15
        risk_factors.append("Unusually high bet amount")
    
    # Pattern: Rapid betting (bots)
    if request.user_id in user_risk_profiles:
        profile = user_risk_profiles[request.user_id]
        if profile.get("last_bet_time"):
            last_bet = datetime.fromisoformat(profile["last_bet_time"])
            time_diff = (datetime.now() - last_bet).total_seconds()
            if time_diff < 2:  # Less than 2 seconds between bets
                risk_score += 0.25
                risk_factors.append("Rapid betting pattern detected")
    
    # Update last bet time
    if request.user_id not in user_risk_profiles:
        user_risk_profiles[request.user_id] = {"user_id": request.user_id}
    user_risk_profiles[request.user_id]["last_bet_time"] = datetime.now().isoformat()
    
    return {
        "user_id": request.user_id,
        "game_id": request.game_id,
        "risk_score": min(risk_score, 1.0),
        "allowed": risk_score < 0.6,
        "requires_review": risk_score >= 0.4,
        "risk_factors": risk_factors
    }


@router.post("/check/account")
async def check_account_activity(request: AccountActivityRequest):
    """Check for account takeover attempts"""
    risk_score = 0.0
    risk_factors = []
    
    # High-risk activities
    high_risk_activities = ["withdrawal", "password_change", "email_change", "2fa_disable"]
    if request.activity_type in high_risk_activities:
        risk_score += 0.3
        risk_factors.append(f"High-risk activity: {request.activity_type}")
    
    # New device check
    if request.device_fingerprint:
        if request.user_id in device_fingerprints:
            if device_fingerprints[request.user_id] != request.device_fingerprint:
                risk_score += 0.4
                risk_factors.append("Account access from new device")
    
    # IP change detection
    if request.ip_address:
        profile = user_risk_profiles.get(request.user_id, {})
        if profile.get("last_ip") and profile["last_ip"] != request.ip_address:
            risk_score += 0.2
            risk_factors.append("IP address changed")
    
    return {
        "user_id": request.user_id,
        "activity_type": request.activity_type,
        "risk_score": min(risk_score, 1.0),
        "allowed": risk_score < 0.7,
        "requires_review": risk_score >= 0.5,
        "risk_factors": risk_factors,
        "recommendation": "block" if risk_score > 0.8 else "allow" if risk_score < 0.5 else "review"
    }


@router.post("/device/fingerprint")
async def register_device(user_id: str, fingerprint: str):
    """Register device fingerprint"""
    device_fingerprints[user_id] = fingerprint
    return {"user_id": user_id, "fingerprint": fingerprint, "registered": True}


@router.get("/user/{user_id}/risk")
async def get_user_risk(user_id: str):
    """Get user's risk profile"""
    if user_id not in user_risk_profiles:
        user_risk_profiles[user_id] = {
            "user_id": user_id,
            "risk_score": 0.0,
            "is_blocked": False,
            "flags": [],
            "transaction_count": 0
        }
    
    profile = user_risk_profiles[user_id]
    profile["is_blocked"] = profile.get("risk_score", 0) > 0.8
    
    return profile


@router.post("/alert")
async def create_alert(request: AlertRequest):
    """Create fraud alert"""
    alert_id = f"ALERT-{datetime.now().strftime('%Y%m%d%H%M%S')}-{random.randint(1000, 9999)}"
    
    alert = {
        "alert_id": alert_id,
        "user_id": request.user_id,
        "alert_type": request.alert_type,
        "description": request.description,
        "evidence": request.evidence or {},
        "status": "open",
        "created_at": datetime.now().isoformat()
    }
    
    alerts_db.append(alert)
    
    # Update user risk profile
    if request.user_id not in user_risk_profiles:
        user_risk_profiles[request.user_id] = {"user_id": request.user_id, "flags": []}
    user_risk_profiles[request.user_id]["flags"].append(request.alert_type)
    user_risk_profiles[request.user_id]["risk_score"] = min(
        user_risk_profiles[request.user_id].get("risk_score", 0) + 0.1,
        1.0
    )
    
    return alert


@router.get("/alerts")
async def get_alerts(status: Optional[str] = None, limit: int = 50):
    """Get fraud alerts"""
    filtered = alerts_db
    if status:
        filtered = [a for a in filtered if a.get("status") == status]
    
    return {
        "count": len(filtered),
        "alerts": filtered[:limit]
    }


@router.post("/user/{user_id}/unblock")
async def unblock_user(user_id: str):
    """Unblock a user after review"""
    if user_id in user_risk_profiles:
        user_risk_profiles[user_id]["is_blocked"] = False
        user_risk_profiles[user_id]["risk_score"] = 0.0
        user_risk_profiles[user_id]["flags"] = []
        return {"user_id": user_id, "unblocked": True}
    raise HTTPException(status_code=404, detail="User not found")


@router.post("/user/{user_id}/block")
async def block_user(user_id: str, reason: str):
    """Block a user"""
    if user_id not in user_risk_profiles:
        user_risk_profiles[user_id] = {"user_id": user_id}
    
    user_risk_profiles[user_id]["is_blocked"] = True
    user_risk_profiles[user_id]["block_reason"] = reason
    user_risk_profiles[user_id]["blocked_at"] = datetime.now().isoformat()
    
    return {"user_id": user_id, "blocked": True, "reason": reason}
