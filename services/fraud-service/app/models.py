"""SQLAlchemy ORM models for fraud detection service."""

from sqlalchemy import Column, String, Float, Integer, Boolean, DateTime, JSON, Text
from datetime import datetime

from app.database import Base


class DeviceFingerprintRecord(Base):
    __tablename__ = "device_fingerprints"

    user_id = Column(String, primary_key=True)
    canvas_hash = Column(String, nullable=True)
    webgl_hash = Column(String, nullable=True)
    audio_hash = Column(String, nullable=True)
    fonts = Column(JSON, nullable=True)
    screen_resolution = Column(String, nullable=True)
    ip_address = Column(String, nullable=True, index=True)
    user_agent = Column(String, nullable=True)


class IpAccountRecord(Base):
    __tablename__ = "ip_accounts"

    id = Column(Integer, primary_key=True, autoincrement=True)
    ip_address = Column(String, nullable=False, index=True)
    user_id = Column(String, nullable=False)


class FraudScoreRecord(Base):
    __tablename__ = "fraud_scores"

    user_id = Column(String, primary_key=True)
    score = Column(Integer, nullable=False)
    category = Column(String, nullable=False)
    signals = Column(JSON, default=dict)
    recommendations = Column(JSON, default=list)
    last_updated = Column(DateTime, default=datetime.utcnow)


class FraudAlertRecord(Base):
    __tablename__ = "fraud_alerts"

    alert_id = Column(String, primary_key=True)
    user_id = Column(String, nullable=False, index=True)
    alert_type = Column(String, nullable=False)
    description = Column(Text, nullable=False)
    evidence = Column(JSON, default=dict)
    status = Column(String, default="open")
    created_at = Column(DateTime, default=datetime.utcnow)


class UserRiskProfileRecord(Base):
    __tablename__ = "user_risk_profiles"

    user_id = Column(String, primary_key=True)
    risk_score = Column(Float, default=0.0)
    is_blocked = Column(Boolean, default=False)
    flags = Column(JSON, default=list)
    transaction_count = Column(Integer, default=0)
    last_activity = Column(String, nullable=True)
    last_ip = Column(String, nullable=True)
    last_bet_time = Column(String, nullable=True)
    device_fingerprint = Column(String, nullable=True)
    block_reason = Column(String, nullable=True)
    blocked_at = Column(String, nullable=True)
