"""
Risk Scoring Service

FastAPI service for unified risk profile management.
Aggregates signals from AML, Fraud, KYC, and transaction history
to produce a comprehensive risk score with automated actions.
"""

from fastapi import FastAPI, HTTPException
from pydantic import BaseModel, Field
from typing import Optional, List, Dict
from datetime import datetime, timedelta
from enum import Enum
import httpx

app = FastAPI(
    title="Risk Scoring Service",
    description="Unified risk scoring and automated actions",
    version="1.0.0"
)


# ============================================================================
# Data Models
# ============================================================================

class RiskCategory(str, Enum):
    LOW = "low"
    MEDIUM = "medium"
    HIGH = "high"
    CRITICAL = "critical"


class RiskProfile(BaseModel):
    user_id: str
    overall_score: int = Field(ge=0, le=100)
    category: RiskCategory
    kyc_level: int = Field(ge=0, le=3)
    aml_alerts: int = 0
    fraud_score: int = 0
    transaction_risk: int = 0
    device_risk: int = 0
    velocity_risk: int = 0
    factors: Dict[str, float] = {}
    deposit_limit: Optional[float] = None
    withdrawal_limit: Optional[float] = None
    recommended_actions: List[str] = []
    last_updated: datetime = Field(default_factory=datetime.now)


class RiskAssessmentRequest(BaseModel):
    user_id: str
    transaction_amount: Optional[float] = None
    transaction_type: Optional[str] = None  # deposit, withdrawal


class RiskAction(BaseModel):
    action: str
    reason: str
    parameters: Dict = {}


# ============================================================================
# Service URLs (would be environment variables in production)
# ============================================================================

AML_SERVICE_URL = "http://localhost:9014"
FRAUD_SERVICE_URL = "http://localhost:9015"


# ============================================================================
# Risk Scoring Engine
# ============================================================================

class RiskScoringEngine:
    """Calculate unified risk profile from multiple sources"""
    
    # Risk weights for each factor
    WEIGHTS = {
        "kyc_level": 0.20,
        "aml_alerts": 0.25,
        "fraud_score": 0.25,
        "transaction_risk": 0.15,
        "device_risk": 0.10,
        "velocity_risk": 0.05
    }
    
    # Dynamic limits by risk category
    DEPOSIT_LIMITS = {
        "low": 50000,
        "medium": 10000,
        "high": 1000,
        "critical": 0
    }
    
    WITHDRAWAL_LIMITS = {
        "low": 50000,
        "medium": 5000,
        "high": 500,
        "critical": 0
    }
    
    @staticmethod
    async def calculate_profile(
        user_id: str,
        kyc_level: int,
        aml_alerts: int,
        fraud_score: int,
        transaction_risk: int = 0,
        device_risk: int = 0,
        velocity_risk: int = 0
    ) -> RiskProfile:
        """Calculate comprehensive risk profile"""
        
        # Normalize KYC level (0-3 -> 0-100 risk, lower KYC = higher risk)
        kyc_risk = (3 - kyc_level) / 3 * 100 if kyc_level <= 3 else 100
        
        # Normalize AML alerts
        aml_risk = min(aml_alerts * 20, 100)
        
        # Combine factors with weights
        factors = {
            "kyc_risk": kyc_risk,
            "aml_risk": aml_risk,
            "fraud_risk": fraud_score,
            "transaction_risk": transaction_risk,
            "device_risk": device_risk,
            "velocity_risk": velocity_risk
        }
        
        overall_score = sum(
            factors[k] * RiskScoringEngine.WEIGHTS.get(k, 0)
            for k in factors
        )
        overall_score = int(overall_score)
        overall_score = min(100, overall_score)
        
        # Determine category
        if overall_score <= 25:
            category = RiskCategory.LOW
        elif overall_score <= 50:
            category = RiskCategory.MEDIUM
        elif overall_score <= 75:
            category = RiskCategory.HIGH
        else:
            category = RiskCategory.CRITICAL
        
        # Calculate dynamic limits
        deposit_limit = RiskScoringEngine.DEPOSIT_LIMITS[category]
        withdrawal_limit = RiskScoringEngine.WITHDRAWAL_LIMITS[category]
        
        # Generate recommended actions
        actions = RiskScoringEngine.generate_actions(category, kyc_level, aml_alerts)
        
        return RiskProfile(
            user_id=user_id,
            overall_score=overall_score,
            category=category,
            kyc_level=kyc_level,
            aml_alerts=aml_alerts,
            fraud_score=fraud_score,
            transaction_risk=transaction_risk,
            device_risk=device_risk,
            velocity_risk=velocity_risk,
            factors={k: round(v, 2) for k, v in factors.items()},
            deposit_limit=deposit_limit,
            withdrawal_limit=withdrawal_limit,
            recommended_actions=actions
        )
    
    @staticmethod
    def generate_actions(category: RiskCategory, kyc_level: int, aml_alerts: int) -> List[str]:
        """Generate recommended actions based on risk profile"""
        actions = []
        
        if category == RiskCategory.MEDIUM:
            actions.append("enhanced_monitoring")
            actions.append("lower_withdrawal_limits")
        
        elif category == RiskCategory.HIGH:
            actions.append("require_additional_kyc")
            actions.append("manual_withdrawal_review")
            actions.append("reduce_deposit_limits")
        
        elif category == RiskCategory.CRITICAL:
            actions.append("suspend_account")
            actions.append("block_all_withdrawals")
            actions.append("escalate_to_compliance")
        
        if kyc_level < 2:
            actions.append("prompt_kyc_verification")
        
        if aml_alerts > 0:
            actions.append("flag_for_aml_review")
        
        return actions
    
    @staticmethod
    def check_transaction_allowed(
        profile: RiskProfile,
        amount: float,
        transaction_type: str
    ) -> tuple[bool, Optional[str]]:
        """Check if a transaction is allowed based on risk profile"""
        
        if transaction_type == "deposit":
            limit = profile.deposit_limit
            if limit is not None and amount > limit:
                return False, f"Deposit exceeds limit of ${limit}"
        
        elif transaction_type == "withdrawal":
            limit = profile.withdrawal_limit
            if limit is not None and amount > limit:
                return False, f"Withdrawal exceeds limit of ${limit}"
            
            if profile.category == RiskCategory.CRITICAL:
                return False, "Withdrawals blocked for high-risk accounts"
        
        return True, None


# ============================================================================
# API Endpoints
# ============================================================================

@app.post("/risk/profile", response_model=RiskProfile)
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


@app.post("/risk/assess", response_model=Dict)
async def assess_transaction(request: RiskAssessmentRequest):
    """Assess if a transaction should be allowed"""
    
    # In production, would fetch actual data from services
    # For now, use provided or default values
    kyc_level = 2  # Would fetch from KYC service
    aml_alerts = 0  # Would fetch from AML service
    fraud_score = 0  # Would fetch from Fraud service
    
    profile = await RiskScoringEngine.calculate_profile(
        request.user_id,
        kyc_level,
        aml_alerts,
        fraud_score
    )
    
    # Check transaction if provided
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


@app.get("/risk/profile/{user_id}", response_model=RiskProfile)
async def get_user_risk_profile(user_id: str):
    """Get existing risk profile (would fetch from cache/database)"""
    # Would fetch from Redis or database
    # For demo, calculate fresh
    return await RiskScoringEngine.calculate_profile(
        user_id,
        kyc_level=1,
        aml_alerts=0,
        fraud_score=0
    )


@app.get("/health")
async def health_check():
    return {"status": "healthy", "service": "risk-service"}


# ============================================================================
# Dynamic Limits Endpoints
# ============================================================================

@app.post("/limits/calculate", response_model=Dict)
async def calculate_dynamic_limits(
    user_id: str,
    account_age_days: int,
    total_deposits: float,
    vip_level: int = 0,
    current_risk_score: int = 0
):
    """Calculate dynamic deposit/withdrawal limits"""
    
    # Base limits
    base_deposit = 10000
    base_withdrawal = 5000
    
    # Adjust for account age
    if account_age_days < 30:
        age_factor = 0.5
    elif account_age_days < 90:
        age_factor = 0.75
    else:
        age_factor = 1.0
    
    # Adjust for VIP level
    vip_factor = 1.0 + (vip_level * 0.25)
    
    # Adjust for risk
    risk_factor = 1.0 - (current_risk_score / 200)  # 0.5 at score 100
    
    deposit_limit = base_deposit * age_factor * vip_factor * risk_factor
    withdrawal_limit = base_withdrawal * age_factor * vip_factor * risk_factor
    
    return {
        "user_id": user_id,
        "daily_deposit_limit": round(deposit_limit, 2),
        "daily_withdrawal_limit": round(withdrawal_limit, 2),
        "calculation": {
            "base_deposit": base_deposit,
            "age_factor": age_factor,
            "vip_factor": vip_factor,
            "risk_factor": risk_factor
        }
    }


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=9016)
