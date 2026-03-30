from pydantic import BaseModel, Field
from typing import Optional, List, Dict
from datetime import datetime
from enum import Enum
import uuid


class AlertSeverity(str, Enum):
    LOW = "low"
    MEDIUM = "medium"
    HIGH = "high"
    CRITICAL = "critical"


class AlertStatus(str, Enum):
    NEW = "new"
    ASSIGNED = "assigned"
    INVESTIGATING = "investigating"
    RESOLVED = "resolved"
    ESCALATED = "escalated"


class AlertType(str, Enum):
    STRUCTURING = "structuring"
    RAPID_DEPOSIT_WITHDRAW = "rapid_deposit_withdraw"
    LARGE_TRANSACTION = "large_transaction"
    VELOCITY = "velocity"
    ROUND_TRIPPING = "round_tripping"
    THIRD_PARTY = "third_party"
    CHIP_DUMPING = "chip_dumping"
    MINIMAL_PLAY = "minimal_play"
    PATTERN_CHANGE = "pattern_change"
    MULTI_ACCOUNT = "multi_account"
    GEOGRAPHIC_ANOMALY = "geographic_anomaly"


class Transaction(BaseModel):
    transaction_id: str
    user_id: str
    type: str  # deposit, withdrawal, bet, win
    amount: float
    currency: str
    payment_method: Optional[str] = None
    ip_address: Optional[str] = None
    country: Optional[str] = None
    timestamp: datetime


class Alert(BaseModel):
    alert_id: str = Field(default_factory=lambda: str(uuid.uuid4()))
    user_id: str
    alert_type: AlertType
    severity: AlertSeverity
    status: AlertStatus = AlertStatus.NEW
    description: str
    transactions: List[str] = []  # transaction IDs
    assigned_to: Optional[str] = None
    created_at: datetime = Field(default_factory=datetime.now)
    updated_at: datetime = Field(default_factory=datetime.now)
    resolved_at: Optional[datetime] = None
    notes: Optional[str] = None


class RiskScore(BaseModel):
    user_id: str
    score: int = Field(ge=0, le=100)
    category: str  # low, medium, high, critical
    factors: Dict[str, float] = {}
    last_updated: datetime = Field(default_factory=datetime.now)
