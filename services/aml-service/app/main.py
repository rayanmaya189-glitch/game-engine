"""
AML (Anti-Money Laundering) Detection Service

FastAPI service for detecting suspicious financial patterns
and generating regulatory reports.

Features:
- Rules engine for transaction pattern detection
- ML-based risk scoring
- Alert management workflow
- CTR/SAR report generation
"""

from fastapi import FastAPI, HTTPException, BackgroundTasks
from pydantic import BaseModel, Field
from typing import Optional, List, Dict
from datetime import datetime, timedelta
from enum import Enum
import uuid

app = FastAPI(
    title="AML Detection Service",
    description="Anti-money laundering detection and compliance",
    version="1.0.0"
)


# ============================================================================
# Data Models
# ============================================================================

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


# ============================================================================
# In-Memory Storage (would be database in production)
# ============================================================================

alerts_db: Dict[str, Alert] = {}
transactions_db: Dict[str, Transaction] = {}
risk_scores_db: Dict[str, RiskScore] = {}


# ============================================================================
# AML Rules Engine
# ============================================================================

class AMLRulesEngine:
    """Rules engine for detecting suspicious transaction patterns"""
    
    # Rule thresholds
    STRUCTURING_THRESHOLD = 10000  # 3+ deposits within 24h totaling > $10K
    STRUCTURING_COUNT = 3
    RAPID_WITHDRAW_RATIO = 3  # < 3x wagering of deposit
    LARGE_TRANSACTION_THRESHOLD = 10000
    VELOCITY_THRESHOLD = 10  # > 10 transactions/hour
    
    @staticmethod
    def check_structuring(user_id: str, hours: int = 24) -> Optional[Alert]:
        """Detect structuring: multiple deposits just below reporting threshold"""
        cutoff = datetime.now() - timedelta(hours=hours)
        
        deposits = [
            t for t in transactions_db.values()
            if t.user_id == user_id 
            and t.type == "deposit"
            and t.timestamp > cutoff
            and t.amount < 10000  # Below single threshold
        ]
        
        total = sum(d.amount for d in deposits)
        
        if len(deposits) >= AMLRulesEngine.STRUCTURING_COUNT and total > AMLRulesEngine.STRUCTURING_THRESHOLD:
            return Alert(
                user_id=user_id,
                alert_type=AlertType.STRUCTURING,
                severity=AlertSeverity.HIGH,
                description=f"Structuring detected: {len(deposits)} deposits totaling ${total} within {hours}h",
                transactions=[t.transaction_id for t in deposits]
            )
        return None
    
    @staticmethod
    def check_rapid_deposit_withdraw(user_id: str) -> Optional[Alert]:
        """Detect rapid deposit-withdraw with minimal play"""
        recent_transactions = [
            t for t in transactions_db.values()
            if t.user_id == user_id
            and t.timestamp > datetime.now() - timedelta(days=7)
        ]
        
        deposits = [t for t in recent_transactions if t.type == "deposit"]
        withdrawals = [t for t in recent_transactions if t.type == "withdrawal"]
        bets = [t for t in recent_transactions if t.type == "bet"]
        
        if not deposits or not withdrawals:
            return None
            
        total_deposits = sum(d.amount for d in deposits)
        total_bets = sum(b.amount for b in bets)
        
        # Check if withdrawals exceed 3x bets (rapid withdraw pattern)
        if total_withdrawals := sum(w.amount for w in withdrawals):
            if total_bets < (total_deposits / AMLRulesEngine.RAPID_WITHDRAW_RATIO):
                return Alert(
                    user_id=user_id,
                    alert_type=AlertType.RAPID_DEPOSIT_WITHDRAW,
                    severity=AlertSeverity.MEDIUM,
                    description=f"Rapid deposit-withdraw: ${total_deposits} deposited, ${total_bets} wagered, ${total_withdrawals} withdrawn",
                    transactions=[t.transaction_id for t in recent_transactions]
                )
        return None
    
    @staticmethod
    def check_large_transaction(user_id: str, transaction: Transaction) -> Optional[Alert]:
        """Flag large transactions > $10K"""
        if transaction.amount > AMLRulesEngine.LARGE_TRANSACTION_THRESHOLD:
            return Alert(
                user_id=user_id,
                alert_type=AlertType.LARGE_TRANSACTION,
                severity=AlertSeverity.HIGH,
                description=f"Large transaction: ${transaction.amount} {transaction.type}",
                transactions=[transaction.transaction_id]
            )
        return None
    
    @staticmethod
    def check_velocity(user_id: str, hours: int = 1) -> Optional[Alert]:
        """Detect unusual transaction frequency"""
        cutoff = datetime.now() - timedelta(hours=hours)
        
        count = sum(
            1 for t in transactions_db.values()
            if t.user_id == user_id and t.timestamp > cutoff
        )
        
        if count > AMLRulesEngine.VELOCITY_THRESHOLD:
            return Alert(
                user_id=user_id,
                alert_type=AlertType.VELOCITY,
                severity=AlertSeverity.MEDIUM,
                description=f"High velocity: {count} transactions in {hours} hour(s)",
                transactions=[
                    t.transaction_id for t in transactions_db.values()
                    if t.user_id == user_id and t.timestamp > cutoff
                ]
            )
        return None
    
    @classmethod
    def run_all_rules(cls, user_id: str, transaction: Optional[Transaction] = None) -> List[Alert]:
        """Run all AML rules for a user"""
        alerts = []
        
        # Check structuring
        if alert := cls.check_structuring(user_id):
            alerts.append(alert)
        
        # Check rapid deposit-withdraw
        if alert := cls.check_rapid_deposit_withdraw(user_id):
            alerts.append(alert)
        
        # Check velocity
        if alert := cls.check_velocity(user_id):
            alerts.append(alert)
        
        # Check large transaction if provided
        if transaction and (alert := cls.check_large_transaction(user_id, transaction)):
            alerts.append(alert)
        
        return alerts


# ============================================================================
# ML Risk Model (Simplified)
# ============================================================================

class MLRiskModel:
    """ML-based risk scoring model"""
    
    @staticmethod
    def calculate_risk_score(user_id: str) -> RiskScore:
        """Calculate risk score based on user behavior patterns"""
        factors = {}
        
        # Get user transactions
        user_transactions = [t for t in transactions_db.values() if t.user_id == user_id]
        
        if not user_transactions:
            return RiskScore(user_id=user_id, score=0, category="low", factors={})
        
        # Feature: Transaction velocity
        recent = [t for t in user_transactions if t.timestamp > datetime.now() - timedelta(days=30)]
        velocity = len(recent) / 30
        factors["velocity"] = min(velocity / 5, 1.0)  # Normalize to 0-1
        
        # Feature: Average transaction amount
        amounts = [t.amount for t in recent]
        avg_amount = sum(amounts) / len(amounts) if amounts else 0
        factors["avg_amount"] = min(avg_amount / 5000, 1.0)
        
        # Feature: Wagering ratio
        deposits = sum(t.amount for t in recent if t.type == "deposit")
        bets = sum(t.amount for t in recent if t.type == "bet")
        wagering_ratio = (bets / deposits) if deposits > 0 else 0
        factors["wagering_ratio"] = 1 - min(wagering_ratio, 1)
        
        # Feature: Payment method diversity
        methods = set(t.payment_method for t in recent if t.payment_method)
        factors["payment_diversity"] = min(len(methods) / 3, 1.0)
        
        # Feature: Geographic consistency
        countries = set(t.country for t in recent if t.country)
        factors["geo_consistency"] = 1 - min(len(countries) / 5, 1.0)
        
        # Calculate weighted score
        weights = {
            "velocity": 0.15,
            "avg_amount": 0.20,
            "wagering_ratio": 0.30,
            "payment_diversity": 0.15,
            "geo_consistency": 0.20
        }
        
        score = sum(factors[k] * weights[k] for k in weights)
        score = int(score * 100)
        
        # Determine category
        if score <= 25:
            category = "low"
        elif score <= 50:
            category = "medium"
        elif score <= 75:
            category = "high"
        else:
            category = "critical"
        
        return RiskScore(
            user_id=user_id,
            score=score,
            category=category,
            factors={k: round(v, 3) for k, v in factors.items()}
        )


# ============================================================================
# API Endpoints
# ============================================================================

@app.post("/transactions", response_model=Dict)
async def analyze_transaction(transaction: Transaction, background_tasks: BackgroundTasks):
    """Analyze a transaction for suspicious patterns"""
    # Store transaction
    transactions_db[transaction.transaction_id] = transaction
    
    # Run rules
    alerts = AMLRulesEngine.run_all_rules(transaction.user_id, transaction)
    
    # Store alerts
    for alert in alerts:
        alerts_db[alert.alert_id] = alert
    
    return {
        "transaction_id": transaction.transaction_id,
        "alerts_generated": len(alerts),
        "alert_ids": [a.alert_id for a in alerts]
    }


@app.post("/batch/analyze", response_model=Dict)
async def analyze_transactions(transactions: List[Transaction]):
    """Batch analyze multiple transactions"""
    all_alerts = []
    
    for transaction in transactions:
        transactions_db[transaction.transaction_id] = transaction
        
        # Run rules per user
        alerts = AMLRulesEngine.run_all_rules(transaction.user_id, transaction)
        all_alerts.extend(alerts)
        
        # Store alerts
        for alert in alerts:
            alerts_db[alert.alert_id] = alert
    
    return {
        "transactions_processed": len(transactions),
        "total_alerts": len(all_alerts),
        "alerts": [a.dict() for a in all_alerts]
    }


@app.get("/risk/score/{user_id}", response_model=RiskScore)
async def get_risk_score(user_id: str):
    """Get risk score for a user"""
    score = MLRiskModel.calculate_risk_score(user_id)
    risk_scores_db[user_id] = score
    return score


@app.post("/risk/score/batch", response_model=Dict)
async def calculate_batch_risk(user_ids: List[str]):
    """Calculate risk scores for multiple users"""
    results = {}
    for user_id in user_ids:
        score = MLRiskModel.calculate_risk_score(user_id)
        risk_scores_db[user_id] = score
        results[user_id] = score
    
    return {"scores": results}


@app.get("/alerts", response_model=List[Alert])
async def list_alerts(
    status: Optional[AlertStatus] = None,
    severity: Optional[AlertSeverity] = None,
    limit: int = 100
):
    """List alerts with optional filters"""
    alerts = list(alerts_db.values())
    
    if status:
        alerts = [a for a in alerts if a.status == status]
    if severity:
        alerts = [a for a in alerts if a.severity == severity]
    
    # Sort by created_at descending
    alerts.sort(key=lambda a: a.created_at, reverse=True)
    
    return alerts[:limit]


@app.get("/alerts/{alert_id}", response_model=Alert)
async def get_alert(alert_id: str):
    """Get a specific alert"""
    if alert_id not in alerts_db:
        raise HTTPException(status_code=404, detail="Alert not found")
    return alerts_db[alert_id]


@app.patch("/alerts/{alert_id}", response_model=Alert)
async def update_alert(alert_id: str, update: Dict):
    """Update alert status or assign to investigator"""
    if alert_id not in alerts_db:
        raise HTTPException(status_code=404, detail="Alert not found")
    
    alert = alerts_db[alert_id]
    
    if "status" in update:
        alert.status = AlertStatus(update["status"])
    if "assigned_to" in update:
        alert.assigned_to = update["assigned_to"]
    if "notes" in update:
        alert.notes = update["notes"]
    
    if alert.status == AlertStatus.RESOLVED:
        alert.resolved_at = datetime.now()
    
    alert.updated_at = datetime.now()
    
    return alert


@app.get("/health")
async def health_check():
    """Health check endpoint"""
    return {"status": "healthy", "service": "aml-service"}


# ============================================================================
# Regulatory Reports (CTR/SAR)
# ============================================================================

@app.get("/reports/ctr", response_model=Dict)
async def generate_ctr_report(start_date: str, end_date: str):
    """Generate Currency Transaction Report for transactions > $10K"""
    start = datetime.fromisoformat(start_date)
    end = datetime.fromisoformat(end_date)
    
    large_transactions = [
        t for t in transactions_db.values()
        if t.timestamp > start and t.timestamp < end
        and t.amount > 10000
    ]
    
    return {
        "report_type": "CTR",
        "period": {"start": start_date, "end": end_date},
        "transactions": [t.dict() for t in large_transactions],
        "total_count": len(large_transactions)
    }


@app.get("/reports/sar/{alert_id}", response_model=Dict)
async def generate_sar_report(alert_id: str):
    """Generate Suspicious Activity Report for an alert"""
    if alert_id not in alerts_db:
        raise HTTPException(status_code=404, detail="Alert not found")
    
    alert = alerts_db[alert_id]
    
    return {
        "report_type": "SAR",
        "alert": alert.dict(),
        "generated_at": datetime.now().isoformat(),
        "format": "FinCEN BSA"
    }


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=9014)
