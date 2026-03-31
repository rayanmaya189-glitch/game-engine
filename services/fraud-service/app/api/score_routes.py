"""Fraud scoring functions"""

from typing import Dict
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select

from app.database import get_db
from app.models import FraudScoreRecord
from app.models.schemas import FraudScore
from app.services.fraud_scorer import FraudScorer


async def score_transaction(
    user_id: str,
    transaction_amount: float,
    is_new_account: bool = False,
    ip_country: str = "unknown",
    payment_method_new: bool = True,
    device_matches: bool = True,
    db: AsyncSession = None,
) -> FraudScore:
    """Calculate fraud score for a transaction"""
    score = await FraudScorer.calculate_score(
        db, user_id, transaction_amount, is_new_account,
        ip_country, payment_method_new, device_matches
    )
    # Persist score
    await db.execute(
        FraudScoreRecord.__table__.delete().where(FraudScoreRecord.user_id == user_id)
    )
    db.add(FraudScoreRecord(
        user_id=score.user_id,
        score=score.score,
        category=score.category,
        signals=score.signals,
        recommendations=score.recommendations,
        last_updated=score.last_updated,
    ))
    await db.commit()
    return score


async def get_user_fraud_score(user_id: str, db: AsyncSession) -> FraudScore:
    """Get the latest fraud score for a user"""
    result = await db.execute(select(FraudScoreRecord).where(FraudScoreRecord.user_id == user_id))
    record = result.scalar_one_or_none()
    if not record:
        raise ValueError("No score found")
    return FraudScore(
        user_id=record.user_id,
        score=record.score,
        category=record.category,
        signals=record.signals or {},
        recommendations=record.recommendations or [],
        last_updated=record.last_updated,
    )
