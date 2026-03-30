"""Real-time fraud scoring"""

from sqlalchemy.ext.asyncio import AsyncSession

from app.models.schemas import FraudScore
from app.services.multi_account_detector import MultiAccountDetector


class FraudScorer:
    """Calculate real-time fraud score"""

    @staticmethod
    async def calculate_score(
        db: AsyncSession,
        user_id: str,
        transaction_amount: float,
        is_new_account: bool,
        ip_country: str,
        payment_method_new: bool,
        device_matches: bool
    ) -> FraudScore:
        """Calculate fraud score for a transaction"""
        signals = {}

        if is_new_account and transaction_amount > 1000:
            signals["new_account_large_txn"] = 0.6

        if payment_method_new:
            signals["new_payment_method"] = 0.3

        if not device_matches:
            signals["device_mismatch"] = 0.4

        multi_accounts = await MultiAccountDetector.check_multi_account(db, user_id)
        if multi_accounts:
            signals["multi_account"] = 0.8

        score = int(sum(signals.values()) * 100)
        score = min(score, 100)

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
