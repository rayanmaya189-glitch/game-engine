from datetime import datetime, timedelta

from app.models.schemas import RiskScore, transactions_db


class MLRiskModel:
    """ML-based risk scoring model"""

    @staticmethod
    def calculate_risk_score(user_id: str) -> RiskScore:
        """Calculate risk score based on user behavior patterns"""
        factors = {}

        user_transactions = [t for t in transactions_db.values() if t.user_id == user_id]

        if not user_transactions:
            return RiskScore(user_id=user_id, score=0, category="low", factors={})

        # Feature: Transaction velocity
        recent = [t for t in user_transactions if t.timestamp > datetime.now() - timedelta(days=30)]
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
