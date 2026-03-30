"""Data models for fraud detection service"""

from pydantic import BaseModel, Field
from typing import Optional, List, Dict
from datetime import datetime


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


# In-Memory Storage
device_fingerprints: Dict[str, DeviceFingerprint] = {}
ip_accounts: Dict[str, List[str]] = {}
fraud_scores: Dict[str, FraudScore] = {}
