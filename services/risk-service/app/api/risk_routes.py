from fastapi import APIRouter
from typing import Dict

from app.models.schemas import RiskProfile, RiskAssessmentRequest
from app.services.risk_engine import RiskScoringEngine

router = APIRouter(prefix="/risk", tags=["risk"])


@router.post("/profile", response_model=RiskProfile)
async def calculate_risk_profile(
    user_id: str,
    kyc_level: int = 0,
    aml_alerts: int = 0,
    fraud_score: int = 0,
    transaction_risk: int = 0,
    device_risk: int = 0,
    velocity_risk: int = 0
):
    """Calculate risk profile from provided signals"""
    return await RiskScoringEngine.calculate_profile(
        user_id, kyc_level, aml_alerts, fraud_score,
        transaction_risk, device_risk, velocity_risk
    )


@router.post("/assess", response_model=Dict)
async def assess_transaction(request: RiskAssessmentRequest):
    """Assess if a transaction should be allowed"""

    kyc_level = 2
    aml_alerts = 0
    fraud_score = 0

    profile = await RiskScoringEngine.calculate_profile(
        request.user_id,
        kyc_level,
        aml_alerts,
        fraud_score
    )

    allowed = True
    reason = None

    if request.transaction_amount and request.transaction_type:
        allowed, reason = RiskScoringEngine.check_transaction_allowed(
            profile,
            request.transaction_amount,
            request.transaction_type
        )

    return {
        "allowed": allowed,
        "reason": reason,
        "profile": profile.dict(),
        "actions": profile.recommended_actions
    }


@router.get("/profile/{user_id}", response_model=RiskProfile)
async def get_user_risk_profile(user_id: str):
    """Get existing risk profile (would fetch from cache/database)"""
    return await RiskScoringEngine.calculate_profile(
        user_id,
        kyc_level=1,
        aml_alerts=0,
        fraud_score=0
    )
