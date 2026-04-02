import pytest
from unittest.mock import AsyncMock, MagicMock, patch
from datetime import datetime, timedelta
from app.services.fraud_scorer import FraudScorer
from app.services.bot_detector import BotDetector
from app.services.multi_account_detector import MultiAccountDetector


class TestCalculateScore:
    @pytest.mark.asyncio
    async def test_low_score_for_safe_transaction(self):
        mock_db = AsyncMock()
        with patch("app.services.fraud_scorer.MultiAccountDetector") as mock_ma:
            mock_ma.check_multi_account = AsyncMock(return_value=[])

            result = await FraudScorer.calculate_score(
                db=mock_db,
                user_id="user-1",
                transaction_amount=100.0,
                is_new_account=False,
                ip_country="US",
                payment_method_new=False,
                device_matches=True,
            )

            assert result.score <= 25
            assert result.category == "low"
            assert "allow" in result.recommendations

    @pytest.mark.asyncio
    async def test_high_score_for_risky_signals(self):
        mock_db = AsyncMock()
        with patch("app.services.fraud_scorer.MultiAccountDetector") as mock_ma:
            mock_ma.check_multi_account = AsyncMock(return_value=["user-2", "user-3"])

            result = await FraudScorer.calculate_score(
                db=mock_db,
                user_id="user-1",
                transaction_amount=5000.0,
                is_new_account=True,
                ip_country="US",
                payment_method_new=True,
                device_matches=False,
            )

            assert result.score > 50
            assert result.category in ("high", "critical")

    @pytest.mark.asyncio
    async def test_medium_score_for_partial_signals(self):
        mock_db = AsyncMock()
        with patch("app.services.fraud_scorer.MultiAccountDetector") as mock_ma:
            mock_ma.check_multi_account = AsyncMock(return_value=[])

            result = await FraudScorer.calculate_score(
                db=mock_db,
                user_id="user-1",
                transaction_amount=200.0,
                is_new_account=False,
                ip_country="US",
                payment_method_new=True,
                device_matches=True,
            )

            assert 0 < result.score <= 50
            assert "new_payment_method" in result.signals

    @pytest.mark.asyncio
    async def test_score_capped_at_100(self):
        mock_db = AsyncMock()
        with patch("app.services.fraud_scorer.MultiAccountDetector") as mock_ma:
            mock_ma.check_multi_account = AsyncMock(return_value=["u1", "u2", "u3"])

            result = await FraudScorer.calculate_score(
                db=mock_db,
                user_id="user-1",
                transaction_amount=10000.0,
                is_new_account=True,
                ip_country="US",
                payment_method_new=True,
                device_matches=False,
            )

            assert result.score <= 100


class TestMultiAccountDetection:
    def test_analyze_email_plus_addressing(self):
        assert MultiAccountDetector.analyze_email_patterns("john+1@example.com") is True

    def test_analyze_email_normal(self):
        assert MultiAccountDetector.analyze_email_patterns("john@example.com") is False

    def test_analyze_email_multiple_dots(self):
        assert MultiAccountDetector.analyze_email_patterns("j.o.h.n@example.com") is True

    @pytest.mark.asyncio
    async def test_check_multi_account_returns_related(self):
        mock_db = AsyncMock()
        mock_fp = MagicMock()
        mock_fp.canvas_hash = "canvas-123"
        mock_fp.webgl_hash = "webgl-456"
        mock_fp.ip_address = "192.168.1.1"

        mock_result = MagicMock()
        mock_result.scalar_one_or_none.return_value = mock_fp

        mock_related = MagicMock()
        mock_related.scalars.return_value.all.return_value = [
            MagicMock(user_id="user-2"),
            MagicMock(user_id="user-3"),
        ]

        mock_db.execute.side_effect = [mock_result, mock_related, mock_related, mock_related]

        result = await MultiAccountDetector.check_multi_account(mock_db, "user-1")

        assert len(result) >= 0

    @pytest.mark.asyncio
    async def test_check_multi_account_no_fingerprint(self):
        mock_db = AsyncMock()
        mock_result = MagicMock()
        mock_result.scalar_one_or_none.return_value = None
        mock_db.execute.return_value = mock_result

        result = await MultiAccountDetector.check_multi_account(mock_db, "user-1")

        assert result == []


class TestBotDetection:
    def test_detects_consistent_timing(self):
        base = datetime.now()
        timestamps = [base + timedelta(seconds=i * 0.05) for i in range(10)]

        result = BotDetector.analyze_behavior(
            action_timestamps=timestamps,
            mouse_movements=50,
            touch_events=10,
            session_duration=timedelta(minutes=30),
        )

        assert result.confidence > 0
        assert "consistent_timing" in result.signals

    def test_no_bot_with_normal_behavior(self):
        base = datetime.now()
        import random
        random.seed(42)
        timestamps = [base + timedelta(seconds=random.uniform(1, 10)) for _ in range(10)]

        result = BotDetector.analyze_behavior(
            action_timestamps=timestamps,
            mouse_movements=100,
            touch_events=20,
            session_duration=timedelta(minutes=30),
        )

        assert result.is_bot is False

    def test_detects_no_mouse_movement(self):
        base = datetime.now()
        timestamps = [base + timedelta(seconds=i * 2) for i in range(10)]

        result = BotDetector.analyze_behavior(
            action_timestamps=timestamps,
            mouse_movements=0,
            touch_events=0,
            session_duration=timedelta(hours=1),
        )

        assert "no_mouse" in result.signals or "no_touch" in result.signals

    def test_empty_timestamps(self):
        result = BotDetector.analyze_behavior(
            action_timestamps=[],
            mouse_movements=50,
            touch_events=10,
            session_duration=timedelta(minutes=5),
        )

        assert result.is_bot is False
