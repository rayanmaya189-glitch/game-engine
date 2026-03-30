import os
from typing import Optional, List

from app.models.schemas import RiskCategory, RiskProfile


class RiskScoringEngine:
    """Calculate unified risk profile from multiple sources"""

    WEIGHTS = {
        "kyc_level": float(os.environ.get("RISK_WEIGHT_KYC", "0.20")),
        "aml_alerts": float(os.environ.get("RISK_WEIGHT_AML", "0.25")),
        "fraud_score": float(os.environ.get("RISK_WEIGHT_FRAUD", "0.25")),
        "transaction_risk": float(os.environ.get("RISK_WEIGHT_TRANSACTION", "0.15")),
        "device_risk": float(os.environ.get("RISK_WEIGHT_DEVICE", "0.10")),
        "velocity_risk": float(os.environ.get("RISK_WEIGHT_VELOCITY", "0.05")),
    }

    DEPOSIT_LIMITS = {
        "low": int(os.environ.get("RISK_DEPOSIT_LIMIT_LOW", "50000")),
        "medium": int(os.environ.get("RISK_DEPOSIT_LIMIT_MEDIUM", "10000")),
        "high": int(os.environ.get("RISK_DEPOSIT_LIMIT_HIGH", "1000")),
        "critical": int(os.environ.get("RISK_DEPOSIT_LIMIT_CRITICAL", "0")),
    }

    WITHDRAWAL_LIMITS = {
        "low": int(os.environ.get("RISK_WITHDRAWAL_LIMIT_LOW", "50000")),
        "medium": int(os.environ.get("RISK_WITHDRAWAL_LIMIT_MEDIUM", "5000")),
        "high": int(os.environ.get("RISK_WITHDRAWAL_LIMIT_HIGH", "500")),
        "critical": int(os.environ.get("RISK_WITHDRAWAL_LIMIT_CRITICAL", "0")),
    }

    @staticmethod
    async def calculate_profile(
        user_id: str,
        kyc_level: int,
        aml_alerts: int,
        fraud_score: int,
        transaction_risk: int = 0,
        device_risk: int = 0,
        velocity_risk: int = 0
    ) -> RiskProfile:
        """Calculate comprehensive risk profile"""

        kyc_risk = (3 - kyc_level) / 3 * 100 if kyc_level <= 3 else 100

        aml_risk = min(aml_alerts * 20, 100)

        factors = {
            "kyc_risk": kyc_risk,
            "aml_risk": aml_risk,
            "fraud_risk": fraud_score,
            "transaction_risk": transaction_risk,
            "device_risk": device_risk,
            "velocity_risk": velocity_risk
        }

        overall_score = sum(
            factors[k] * RiskScoringEngine.WEIGHTS.get(k, 0)
            for k in factors
        )
        overall_score = int(overall_score)
        overall_score = min(100, overall_score)

        if overall_score <= 25:
            category = RiskCategory.LOW
        elif overall_score <= 50:
            category = RiskCategory.MEDIUM
        elif overall_score <= 75:
            category = RiskCategory.HIGH
        else:
            category = RiskCategory.CRITICAL

        deposit_limit = RiskScoringEngine.DEPOSIT_LIMITS[category]
        withdrawal_limit = RiskScoringEngine.WITHDRAWAL_LIMITS[category]

        actions = RiskScoringEngine.generate_actions(category, kyc_level, aml_alerts)

        return RiskProfile(
            user_id=user_id,
            overall_score=overall_score,
            category=category,
            kyc_level=kyc_level,
            aml_alerts=aml_alerts,
            fraud_score=fraud_score,
            transaction_risk=transaction_risk,
            device_risk=device_risk,
            velocity_risk=velocity_risk,
            factors={k: round(v, 2) for k, v in factors.items()},
            deposit_limit=deposit_limit,
            withdrawal_limit=withdrawal_limit,
            recommended_actions=actions
        )

    @staticmethod
    def generate_actions(category: RiskCategory, kyc_level: int, aml_alerts: int) -> List[str]:
        """Generate recommended actions based on risk profile"""
        actions = []

        if category == RiskCategory.MEDIUM:
            actions.append("enhanced_monitoring")
            actions.append("lower_withdrawal_limits")

        elif category == RiskCategory.HIGH:
            actions.append("require_additional_kyc")
            actions.append("manual_withdrawal_review")
            actions.append("reduce_deposit_limits")

        elif category == RiskCategory.CRITICAL:
            actions.append("suspend_account")
            actions.append("block_all_withdrawals")
            actions.append("escalate_to_compliance")

        if kyc_level < 2:
            actions.append("prompt_kyc_verification")

        if aml_alerts > 0:
            actions.append("flag_for_aml_review")

        return actions

    @staticmethod
    def check_transaction_allowed(
        profile: RiskProfile,
        amount: float,
        transaction_type: str
    ) -> tuple[bool, Optional[str]]:
        """Check if a transaction is allowed based on risk profile"""

        if transaction_type == "deposit":
            limit = profile.deposit_limit
            if limit is not None and amount > limit:
                return False, f"Deposit exceeds limit of ${limit}"

        elif transaction_type == "withdrawal":
            limit = profile.withdrawal_limit
            if limit is not None and amount > limit:
                return False, f"Withdrawal exceeds limit of ${limit}"

            if profile.category == RiskCategory.CRITICAL:
                return False, "Withdrawals blocked for high-risk accounts"

        return True, None
