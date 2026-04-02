import pytest
from unittest.mock import patch
from app.services.risk_engine import RiskScoringEngine
from app.models.schemas import RiskCategory, RiskProfile


class TestCalculateProfile:
    @pytest.mark.asyncio
    async def test_low_risk_profile(self):
        profile = await RiskScoringEngine.calculate_profile(
            user_id="user_1",
            kyc_level=3,
            aml_alerts=0,
            fraud_score=10,
            transaction_risk=5,
            device_risk=5,
            velocity_risk=5,
        )
        assert profile.user_id == "user_1"
        assert profile.overall_score <= 25
        assert profile.category == RiskCategory.LOW
        assert profile.deposit_limit == 50000

    @pytest.mark.asyncio
    async def test_medium_risk_profile(self):
        profile = await RiskScoringEngine.calculate_profile(
            user_id="user_2",
            kyc_level=2,
            aml_alerts=1,
            fraud_score=30,
            transaction_risk=20,
            device_risk=15,
            velocity_risk=10,
        )
        assert profile.category in (RiskCategory.LOW, RiskCategory.MEDIUM)

    @pytest.mark.asyncio
    async def test_high_risk_profile(self):
        profile = await RiskScoringEngine.calculate_profile(
            user_id="user_3",
            kyc_level=0,
            aml_alerts=5,
            fraud_score=90,
            transaction_risk=80,
            device_risk=70,
            velocity_risk=60,
        )
        assert profile.category in (RiskCategory.HIGH, RiskCategory.CRITICAL)
        assert profile.overall_score > 50

    @pytest.mark.asyncio
    async def test_critical_risk_profile(self):
        profile = await RiskScoringEngine.calculate_profile(
            user_id="user_4",
            kyc_level=0,
            aml_alerts=10,
            fraud_score=100,
            transaction_risk=100,
            device_risk=100,
            velocity_risk=100,
        )
        assert profile.category == RiskCategory.CRITICAL
        assert profile.deposit_limit == 0
        assert "suspend_account" in profile.recommended_actions

    @pytest.mark.asyncio
    async def test_profile_includes_factors(self):
        profile = await RiskScoringEngine.calculate_profile(
            user_id="user_5", kyc_level=2, aml_alerts=0, fraud_score=20
        )
        assert "kyc_risk" in profile.factors
        assert "aml_risk" in profile.factors
        assert "fraud_risk" in profile.factors


class TestGenerateActions:
    def test_low_category_no_actions(self):
        actions = RiskScoringEngine.generate_actions(RiskCategory.LOW, kyc_level=3, aml_alerts=0)
        assert actions == []

    def test_medium_category_actions(self):
        actions = RiskScoringEngine.generate_actions(RiskCategory.MEDIUM, kyc_level=3, aml_alerts=0)
        assert "enhanced_monitoring" in actions
        assert "lower_withdrawal_limits" in actions

    def test_high_category_actions(self):
        actions = RiskScoringEngine.generate_actions(RiskCategory.HIGH, kyc_level=3, aml_alerts=0)
        assert "require_additional_kyc" in actions
        assert "manual_withdrawal_review" in actions
        assert "reduce_deposit_limits" in actions

    def test_critical_category_actions(self):
        actions = RiskScoringEngine.generate_actions(RiskCategory.CRITICAL, kyc_level=3, aml_alerts=0)
        assert "suspend_account" in actions
        assert "block_all_withdrawals" in actions
        assert "escalate_to_compliance" in actions

    def test_low_kyc_prompts_verification(self):
        actions = RiskScoringEngine.generate_actions(RiskCategory.LOW, kyc_level=1, aml_alerts=0)
        assert "prompt_kyc_verification" in actions

    def test_aml_alerts_flag_for_review(self):
        actions = RiskScoringEngine.generate_actions(RiskCategory.LOW, kyc_level=3, aml_alerts=2)
        assert "flag_for_aml_review" in actions


class TestCheckTransactionAllowed:
    def _make_profile(self, **kwargs):
        defaults = {
            "user_id": "u1",
            "overall_score": 10,
            "category": RiskCategory.LOW,
            "kyc_level": 3,
            "deposit_limit": 50000,
            "withdrawal_limit": 50000,
        }
        defaults.update(kwargs)
        return RiskProfile(**defaults)

    def test_deposit_within_limit(self):
        profile = self._make_profile()
        allowed, reason = RiskScoringEngine.check_transaction_allowed(profile, 1000.0, "deposit")
        assert allowed is True
        assert reason is None

    def test_deposit_exceeds_limit(self):
        profile = self._make_profile(deposit_limit=500)
        allowed, reason = RiskScoringEngine.check_transaction_allowed(profile, 1000.0, "deposit")
        assert allowed is False
        assert "limit" in reason.lower()

    def test_withdrawal_within_limit(self):
        profile = self._make_profile()
        allowed, reason = RiskScoringEngine.check_transaction_allowed(profile, 1000.0, "withdrawal")
        assert allowed is True

    def test_withdrawal_exceeds_limit(self):
        profile = self._make_profile(withdrawal_limit=500)
        allowed, reason = RiskScoringEngine.check_transaction_allowed(profile, 1000.0, "withdrawal")
        assert allowed is False
        assert "limit" in reason.lower()

    def test_withdrawal_blocked_for_critical(self):
        profile = self._make_profile(
            category=RiskCategory.CRITICAL, withdrawal_limit=0
        )
        allowed, reason = RiskScoringEngine.check_transaction_allowed(profile, 100.0, "withdrawal")
        assert allowed is False
        assert "blocked" in reason.lower()

    def test_deposit_blocked_for_critical(self):
        profile = self._make_profile(
            category=RiskCategory.CRITICAL, deposit_limit=0
        )
        allowed, reason = RiskScoringEngine.check_transaction_allowed(profile, 100.0, "deposit")
        assert allowed is False


class TestCalculateLimits:
    @pytest.mark.asyncio
    async def test_limits_vary_by_category(self):
        low = await RiskScoringEngine.calculate_profile(
            "u", kyc_level=3, aml_alerts=0, fraud_score=5
        )
        critical = await RiskScoringEngine.calculate_profile(
            "u", kyc_level=0, aml_alerts=10, fraud_score=100,
            transaction_risk=100, device_risk=100, velocity_risk=100
        )
        assert low.deposit_limit > critical.deposit_limit
        assert low.withdrawal_limit > critical.withdrawal_limit
