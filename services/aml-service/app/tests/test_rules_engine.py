import pytest
from unittest.mock import AsyncMock, MagicMock, patch
from datetime import datetime, timedelta
from app.services.rules_engine import AMLRulesEngine
from app.models.schemas import AlertType, AlertSeverity, Transaction


class TestCheckStructuring:
    @pytest.mark.asyncio
    async def test_detects_structuring_when_multiple_deposits_below_threshold(self):
        mock_db = AsyncMock()
        deposits = []
        for i in range(5):
            record = MagicMock()
            record.amount = 9500.0
            record.transaction_id = f"tx-{i}"
            deposits.append(record)

        mock_result = MagicMock()
        mock_result.scalars.return_value.all.return_value = deposits
        mock_db.execute.return_value = mock_result

        alert = await AMLRulesEngine.check_structuring(mock_db, "user-1", hours=24)

        assert alert is not None
        assert alert.alert_type == AlertType.STRUCTURING
        assert alert.severity == AlertSeverity.HIGH
        assert alert.user_id == "user-1"

    @pytest.mark.asyncio
    async def test_no_alert_when_few_deposits(self):
        mock_db = AsyncMock()
        deposits = [MagicMock(amount=5000.0, transaction_id="tx-1")]

        mock_result = MagicMock()
        mock_result.scalars.return_value.all.return_value = deposits
        mock_db.execute.return_value = mock_result

        alert = await AMLRulesEngine.check_structuring(mock_db, "user-1")

        assert alert is None

    @pytest.mark.asyncio
    async def test_no_alert_when_total_below_threshold(self):
        mock_db = AsyncMock()
        deposits = [MagicMock(amount=100.0, transaction_id=f"tx-{i}") for i in range(5)]

        mock_result = MagicMock()
        mock_result.scalars.return_value.all.return_value = deposits
        mock_db.execute.return_value = mock_result

        alert = await AMLRulesEngine.check_structuring(mock_db, "user-1")

        assert alert is None


class TestCheckRapidDepositWithdraw:
    @pytest.mark.asyncio
    async def test_detects_rapid_deposit_withdraw_pattern(self):
        mock_db = AsyncMock()
        deposits = [MagicMock(type="deposit", amount=5000.0, transaction_id="d1")]
        bets = [MagicMock(type="bet", amount=100.0, transaction_id="b1")]
        withdrawals = [MagicMock(type="withdrawal", amount=4900.0, transaction_id="w1")]
        all_tx = deposits + bets + withdrawals

        mock_result = MagicMock()
        mock_result.scalars.return_value.all.return_value = all_tx
        mock_db.execute.return_value = mock_result

        alert = await AMLRulesEngine.check_rapid_deposit_withdraw(mock_db, "user-1")

        assert alert is not None
        assert alert.alert_type == AlertType.RAPID_DEPOSIT_WITHDRAW

    @pytest.mark.asyncio
    async def test_no_alert_when_no_deposits(self):
        mock_db = AsyncMock()
        mock_result = MagicMock()
        mock_result.scalars.return_value.all.return_value = []
        mock_db.execute.return_value = mock_result

        alert = await AMLRulesEngine.check_rapid_deposit_withdraw(mock_db, "user-1")

        assert alert is None

    @pytest.mark.asyncio
    async def test_no_alert_when_sufficient_play(self):
        mock_db = AsyncMock()
        transactions = [
            MagicMock(type="deposit", amount=1000.0, transaction_id="d1"),
            MagicMock(type="bet", amount=800.0, transaction_id="b1"),
            MagicMock(type="withdrawal", amount=200.0, transaction_id="w1"),
        ]

        mock_result = MagicMock()
        mock_result.scalars.return_value.all.return_value = transactions
        mock_db.execute.return_value = mock_result

        alert = await AMLRulesEngine.check_rapid_deposit_withdraw(mock_db, "user-1")

        assert alert is None


class TestCheckLargeTransaction:
    def test_detects_large_transaction(self):
        tx = Transaction(
            transaction_id="tx-1",
            user_id="user-1",
            type="deposit",
            amount=15000.0,
            currency="USD",
            timestamp=datetime.utcnow(),
        )
        alert = AMLRulesEngine.check_large_transaction("user-1", tx)

        assert alert is not None
        assert alert.alert_type == AlertType.LARGE_TRANSACTION
        assert alert.severity == AlertSeverity.HIGH

    def test_no_alert_for_small_transaction(self):
        tx = Transaction(
            transaction_id="tx-1",
            user_id="user-1",
            type="deposit",
            amount=500.0,
            currency="USD",
            timestamp=datetime.utcnow(),
        )
        alert = AMLRulesEngine.check_large_transaction("user-1", tx)

        assert alert is None

    def test_exact_threshold_no_alert(self):
        tx = Transaction(
            transaction_id="tx-1",
            user_id="user-1",
            type="deposit",
            amount=10000.0,
            currency="USD",
            timestamp=datetime.utcnow(),
        )
        alert = AMLRulesEngine.check_large_transaction("user-1", tx)

        assert alert is None


class TestCheckVelocity:
    @pytest.mark.asyncio
    async def test_detects_high_velocity(self):
        mock_db = AsyncMock()
        transactions = [MagicMock(transaction_id=f"tx-{i}") for i in range(15)]

        mock_result = MagicMock()
        mock_result.scalars.return_value.all.return_value = transactions
        mock_db.execute.return_value = mock_result

        alert = await AMLRulesEngine.check_velocity(mock_db, "user-1", hours=1)

        assert alert is not None
        assert alert.alert_type == AlertType.VELOCITY
        assert alert.severity == AlertSeverity.MEDIUM

    @pytest.mark.asyncio
    async def test_no_alert_when_low_velocity(self):
        mock_db = AsyncMock()
        transactions = [MagicMock(transaction_id=f"tx-{i}") for i in range(3)]

        mock_result = MagicMock()
        mock_result.scalars.return_value.all.return_value = transactions
        mock_db.execute.return_value = mock_result

        alert = await AMLRulesEngine.check_velocity(mock_db, "user-1")

        assert alert is None

    @pytest.mark.asyncio
    async def test_no_alert_when_no_transactions(self):
        mock_db = AsyncMock()
        mock_result = MagicMock()
        mock_result.scalars.return_value.all.return_value = []
        mock_db.execute.return_value = mock_result

        alert = await AMLRulesEngine.check_velocity(mock_db, "user-1")

        assert alert is None
