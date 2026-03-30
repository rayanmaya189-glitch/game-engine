import os
from datetime import datetime, timedelta
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select, delete

from app.models import TransactionRecord, RiskScoreRecord
from app.models.schemas import RiskScore


class MLRiskModel:
    """ML-based risk scoring model"""

    WEIGHT_VELOCITY = float(os.environ.get("AML_WEIGHT_VELOCITY", "0.15"))
    WEIGHT_AVG_AMOUNT = float(os.environ.get("AML_WEIGHT_AVG_AMOUNT", "0.20"))
    WEIGHT_WAGERING_RATIO = float(os.environ.get("AML_WEIGHT_WAGERING_RATIO", "0.30"))
    WEIGHT_PAYMENT_DIVERSITY = float(os.environ.get("AML_WEIGHT_PAYMENT_DIVERSITY", "0.15"))
    WEIGHT_GEO_CONSISTENCY = float(os.environ.get("AML_WEIGHT_GEO_CONSISTENCY", "0.20"))

    @staticmethod
    async def calculate_risk_score(db: AsyncSession, user_id: str) -> RiskScore:
        """Calculate risk score based on user behavior patterns"""
        factors = {}

        stmt = select(TransactionRecord).where(TransactionRecord.user_id == user_id)
        result = await db.execute(stmt)
        user_transactions = result.scalars().all()

        if not user_transactions:
            return RiskScore(user_id=user_id, score=0, category="low", factors={})

        # Feature: Transaction velocity
        recent = [t for t in user_transactions if t.timestamp > datetime.utcnow() - timedelta(days=30)]
        velocity = len(recent) / 30
        factors["velocity"] = min(velocity / 5, 1.0)

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

        weights = {
            "velocity": MLRiskModel.WEIGHT_VELOCITY,
            "avg_amount": MLRiskModel.WEIGHT_AVG_AMOUNT,
            "wagering_ratio": MLRiskModel.WEIGHT_WAGERING_RATIO,
            "payment_diversity": MLRiskModel.WEIGHT_PAYMENT_DIVERSITY,
            "geo_consistency": MLRiskModel.WEIGHT_GEO_CONSISTENCY,
        }

        score = sum(factors[k] * weights[k] for k in weights)
        score = int(score * 100)

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

    @staticmethod
    async def save_risk_score(db: AsyncSession, risk_score: RiskScore) -> None:
        """Persist risk score to database"""
        await db.execute(delete(RiskScoreRecord).where(RiskScoreRecord.user_id == risk_score.user_id))
        db.add(RiskScoreRecord(
            user_id=risk_score.user_id,
            score=risk_score.score,
            category=risk_score.category,
            factors=risk_score.factors,
            last_updated=risk_score.last_updated,
        ))
        await db.commit()
