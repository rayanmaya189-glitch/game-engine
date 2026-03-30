from pydantic import BaseModel, Field
from typing import Optional, List, Dict
from datetime import datetime
from enum import Enum


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
