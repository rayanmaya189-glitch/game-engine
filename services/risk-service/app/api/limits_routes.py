import os
from fastapi import APIRouter
from typing import Dict

router = APIRouter(tags=["limits"])

BASE_DEPOSIT = int(os.environ.get("RISK_BASE_DEPOSIT_LIMIT", "10000"))
BASE_WITHDRAWAL = int(os.environ.get("RISK_BASE_WITHDRAWAL_LIMIT", "5000"))


@router.post("/limits/calculate", response_model=Dict)
async def calculate_dynamic_limits(
    user_id: str,
    account_age_days: int,
    total_deposits: float,
    vip_level: int = 0,
    current_risk_score: int = 0
):
    """Calculate dynamic deposit/withdrawal limits"""

    if account_age_days < 30:
        age_factor = 0.5
    elif account_age_days < 90:
        age_factor = 0.75
    else:
        age_factor = 1.0

    vip_factor = 1.0 + (vip_level * 0.25)

    risk_factor = 1.0 - (current_risk_score / 200)

    deposit_limit = BASE_DEPOSIT * age_factor * vip_factor * risk_factor
    withdrawal_limit = BASE_WITHDRAWAL * age_factor * vip_factor * risk_factor

    return {
        "user_id": user_id,
        "daily_deposit_limit": round(deposit_limit, 2),
        "daily_withdrawal_limit": round(withdrawal_limit, 2),
        "calculation": {
            "base_deposit": BASE_DEPOSIT,
            "age_factor": age_factor,
            "vip_factor": vip_factor,
            "risk_factor": risk_factor
        }
    }
