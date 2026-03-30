"""SQLAlchemy ORM models for AML service."""

from sqlalchemy import Column, String, Float, Integer, DateTime, Text, JSON
from datetime import datetime

from app.database import Base


class TransactionRecord(Base):
    __tablename__ = "transactions"

    transaction_id = Column(String, primary_key=True)
    user_id = Column(String, nullable=False, index=True)
    type = Column(String, nullable=False)
    amount = Column(Float, nullable=False)
    currency = Column(String, nullable=False)
    payment_method = Column(String, nullable=True)
    ip_address = Column(String, nullable=True)
    country = Column(String, nullable=True)
    timestamp = Column(DateTime, nullable=False, index=True)


class AlertRecord(Base):
    __tablename__ = "alerts"

    alert_id = Column(String, primary_key=True)
    user_id = Column(String, nullable=False, index=True)
    alert_type = Column(String, nullable=False)
    severity = Column(String, nullable=False)
    status = Column(String, nullable=False, default="new")
    description = Column(Text, nullable=False)
    transactions = Column(JSON, default=list)
    assigned_to = Column(String, nullable=True)
    created_at = Column(DateTime, default=datetime.utcnow)
    updated_at = Column(DateTime, default=datetime.utcnow)
    resolved_at = Column(DateTime, nullable=True)
    notes = Column(Text, nullable=True)


class RiskScoreRecord(Base):
    __tablename__ = "risk_scores"

    user_id = Column(String, primary_key=True)
    score = Column(Integer, nullable=False)
    category = Column(String, nullable=False)
    factors = Column(JSON, default=dict)
    last_updated = Column(DateTime, default=datetime.utcnow)
