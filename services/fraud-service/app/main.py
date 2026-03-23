"""
Fraud Detection Service

FastAPI service for detecting fraud patterns including:
- Multi-account detection (device fingerprinting, IP correlation)
- Bot detection
- Collusion detection for poker
- Real-time fraud scoring
"""

from fastapi import FastAPI, HTTPException
from pydantic import BaseModel, Field
from typing import Optional, List, Dict
from datetime import datetime, timedelta
from enum import Enum
import uuid
import hashlib

app = FastAPI(
    title="Fraud Detection Service",
    description="Real-time fraud detection and prevention",
    version="1.0.0"
)


# ============================================================================
# Data Models
# ============================================================================

class DeviceFingerprint(BaseModel):
    """Device fingerprint data for multi-account detection"""
    user_id: str
    canvas_hash: Optional[str] = None
    webgl_hash: Optional[str] = None
    audio_hash: Optional[str] = None
    fonts: Optional[List[str]] = None
    screen_resolution: Optional[str] = None
    ip_address: Optional[str] = None
    user_agent: Optional[str] = None


class FraudScore(BaseModel):
    user_id: str
    score: int = Field(ge=0, le=100)
    category: str  # low, medium, high, critical
    signals: Dict[str, float] = {}
    recommendations: List[str] = []
    last_updated: datetime = Field(default_factory=datetime.now)


class BotDetectionResult(BaseModel):
    is_bot: bool
    confidence: float = Field(ge=0, le=1)
    signals: Dict[str, float] = {}


class CollusionSignal(BaseModel):
    player_a_id: str
    player_b_id: str
    signal_type: str
    confidence: float
    evidence: Dict = {}


# ============================================================================
# In-Memory Storage
# ============================================================================

device_fingerprints: Dict[str, DeviceFingerprint] = {}
ip_accounts: Dict[str, List[str]] = {}  # IP -> list of user_ids
fraud_scores: Dict[str, FraudScore] = {}


# ============================================================================
# Multi-Account Detection
# ============================================================================

class MultiAccountDetector:
    """Detect multiple accounts from same device/IP"""
    
    @staticmethod
    def register_fingerprint(fingerprint: DeviceFingerprint):
        """Register a device fingerprint"""
        device_fingerprints[fingerprint.user_id] = fingerprint
        
        # Track IP to accounts mapping
        if fingerprint.ip_address:
            if fingerprint.ip_address not in ip_accounts:
                ip_accounts[fingerprint.ip_address] = []
            if fingerprint.user_id not in ip_accounts[fingerprint.ip_address]:
                ip_accounts[fingerprint.ip_address].append(fingerprint.user_id)
    
    @staticmethod
    def check_multi_account(user_id: str) -> List[str]:
        """Check if user has multiple accounts"""
        if user_id not in device_fingerprints:
            return []
        
        fp = device_fingerprints[user_id]
        related_accounts = []
        
        # Check canvas/webgl hashes
        for uid, other_fp in device_fingerprints.items():
            if uid == user_id:
                continue
            
            if fp.canvas_hash and other_fp.canvas_hash:
                if fp.canvas_hash == other_fp.canvas_hash:
                    related_accounts.append(uid)
            
            if fp.webgl_hash and other_fp.webgl_hash:
                if fp.webgl_hash == other_fp.webgl_hash:
                    related_accounts.append(uid)
        
        # Check IP matches
        if fp.ip_address and fp.ip_address in ip_accounts:
            related_accounts.extend([
                uid for uid in ip_accounts[fp.ip_address]
                if uid != user_id
            ])
        
        return related_accounts
    
    @staticmethod
    def analyze_email_patterns(email: str) -> bool:
        """Detect email variation patterns (john+1@, j.o.h.n@)"""
        # Simple detection for plus addressing
        if '+' in email.split('@')[0]:
            return True
        
        # Check for excessive dots in username
        username = email.split('@')[0]
        if username.count('.') > 2:
            return True
        
        return False


# ============================================================================
# Bot Detection
# ============================================================================

class BotDetector:
    """Detect automated/bot behavior"""
    
    @staticmethod
    def analyze_behavior(
        action_timestamps: List[datetime],
        mouse_movements: int,
        touch_events: int,
        session_duration: timedelta,
        perfect_play_pct: Optional[float] = None
    ) -> BotDetectionResult:
        """Analyze user behavior for bot indicators"""
        signals = {}
        
        # 1. Check timing consistency (bots have very consistent timing)
        if len(action_timestamps) > 5:
            intervals = [
                (action_timestamps[i+1] - action_timestamps[i]).total_seconds()
                for i in range(len(action_timestamps) - 1)
            ]
            variance = sum((x - sum(intervals)/len(intervals))**2 for x in intervals) / len(intervals)
            signals["timing_variance"] = variance
            
            # Very low variance indicates automation
            if variance < 0.1:
                signals["consistent_timing"] = 0.9
            else:
                signals["consistent_timing"] = 0.0
        
        # 2. Check mouse movement
        if mouse_movements < 5 and action_timestamps:
            # No mouse movement but has actions (suspicious)
            signals["no_mouse"] = 0.7
        
        # 3. Check touch events on mobile
        if touch_events < 2 and action_timestamps:
            signals["no_touch"] = 0.6
        
        # 4. Check for perfect play (chess, poker)
        if perfect_play_pct and perfect_play_pct > 0.95:
            signals["perfect_play"] = 0.8
        
        # 5. Check session duration (24/7 without breaks)
        if session_duration.total_seconds() > 8 * 3600 and not (
            datetime.now() - action_timestamps[0] < timedelta(hours=24)
        ):
            signals["no_breaks"] = 0.5
        
        # Calculate bot confidence
        total_signal = sum(signals.values())
        confidence = min(total_signal, 1.0)
        
        return BotDetectionResult(
            is_bot=confidence > 0.5,
            confidence=confidence,
            signals=signals
        )


# ============================================================================
# Collusion Detection (Poker)
# ============================================================================

class CollusionDetector:
    """Detect collusion in poker games"""
    
    # In production, would query game history
    game_history: Dict[str, List[Dict]] = {}
    
    @staticmethod
    def analyze_table(table_id: str, players: List[str]) -> List[CollusionSignal]:
        """Analyze players at a table for collusion patterns"""
        signals = []
        
        # Would fetch game history for these players
        # Simplified for now
        
        return signals
    
    @staticmethod
    def check_chip_dumping(player_a: str, player_b: str, games: List[Dict]) -> Optional[CollusionSignal]:
        """Detect chip dumping between two players"""
        total_transfers_a_to_b = 0
        total_transfers_b_to_a = 0
        
        for game in games:
            # Simplified - would analyze actual hand history
            pass
        
        if total_transfers_a_to_b > 5 or total_transfers_b_to_a > 5:
            return CollusionSignal(
                player_a_id=player_a,
                player_b_id=player_b,
                signal_type="chip_dumping",
                confidence=0.7,
                evidence={"transfers": total_transfers_a_to_b + total_transfers_b_to_a}
            )
        
        return None


# ============================================================================
# Real-Time Fraud Scoring
# ============================================================================

class FraudScorer:
    """Calculate real-time fraud score"""
    
    @staticmethod
    def calculate_score(
        user_id: str,
        transaction_amount: float,
        is_new_account: bool,
        ip_country: str,
        payment_method_new: bool,
        device_matches: bool
    ) -> FraudScore:
        """Calculate fraud score for a transaction"""
        signals = {}
        
        # Factor: New account + large transaction
        if is_new_account and transaction_amount > 1000:
            signals["new_account_large_txn"] = 0.6
        
        # Factor: New payment method
        if payment_method_new:
            signals["new_payment_method"] = 0.3
        
        # Factor: IP country mismatch (would check against profile)
        # Simplified
        
        # Factor: Device fingerprint mismatch
        if not device_matches:
            signals["device_mismatch"] = 0.4
        
        # Factor: Multi-account detection
        multi_accounts = MultiAccountDetector.check_multi_account(user_id)
        if multi_accounts:
            signals["multi_account"] = 0.8
        
        # Calculate weighted score
        score = int(sum(signals.values()) * 100)
        score = min(score, 100)
        
        # Determine category
        if score <= 25:
            category = "low"
            recommendations = ["allow"]
        elif score <= 50:
            category = "medium"
            recommendations = ["allow", "enhanced_monitoring"]
        elif score <= 75:
            category = "high"
            recommendations = ["allow_with_verification"]
        else:
            category = "critical"
            recommendations = ["block"]
        
        return FraudScore(
            user_id=user_id,
            score=score,
            category=category,
            signals={k: round(v, 2) for k, v in signals.items()},
            recommendations=recommendations
        )


# ============================================================================
# API Endpoints
# ============================================================================

@app.post("/fingerprint/register", response_model=Dict)
async def register_fingerprint(fingerprint: DeviceFingerprint):
    """Register a device fingerprint"""
    MultiAccountDetector.register_fingerprint(fingerprint)
    return {"status": "registered", "user_id": fingerprint.user_id}


@app.get("/fingerprint/check/{user_id}", response_model=Dict)
async def check_multi_account(user_id: str):
    """Check for multi-account patterns"""
    related = MultiAccountDetector.check_multi_account(user_id)
    return {
        "user_id": user_id,
        "related_accounts": related,
        "is_multi_account": len(related) > 0
    }


@app.post("/bot/detect", response_model=BotDetectionResult)
async def detect_bot(
    action_timestamps: List[datetime],
    mouse_movements: int = 0,
    touch_events: int = 0,
    session_duration_seconds: int = 0,
    perfect_play_pct: Optional[float] = None
):
    """Detect if behavior indicates a bot"""
    return BotDetector.analyze_behavior(
        action_timestamps,
        mouse_movements,
        touch_events,
        timedelta(seconds=session_duration_seconds),
        perfect_play_pct
    )


@app.get("/collusion/table/{table_id}", response_model=List[CollusionSignal])
async def analyze_table_collusion(table_id: str, players: List[str]):
    """Analyze a table for collusion patterns"""
    return CollusionDetector.analyze_table(table_id, players)


@app.post("/score/transaction", response_model=FraudScore)
async def score_transaction(
    user_id: str,
    transaction_amount: float,
    is_new_account: bool = False,
    ip_country: str = "unknown",
    payment_method_new: bool = True,
    device_matches: bool = True
):
    """Calculate fraud score for a transaction"""
    score = FraudScorer.calculate_score(
        user_id,
        transaction_amount,
        is_new_account,
        ip_country,
        payment_method_new,
        device_matches
    )
    fraud_scores[user_id] = score
    return score


@app.get("/score/user/{user_id}", response_model=FraudScore)
async def get_user_fraud_score(user_id: str):
    """Get the latest fraud score for a user"""
    if user_id not in fraud_scores:
        raise HTTPException(status_code=404, detail="No score found")
    return fraud_scores[user_id]


@app.get("/health")
async def health_check():
    return {"status": "healthy", "service": "fraud-service"}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=9015)
