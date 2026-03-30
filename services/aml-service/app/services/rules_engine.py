from typing import Optional, List
from datetime import datetime, timedelta

from app.models.schemas import (
    Alert, AlertType, AlertSeverity, Transaction, transactions_db
)


class AMLRulesEngine:
    """Rules engine for detecting suspicious transaction patterns"""

    # Rule thresholds
    STRUCTURING_THRESHOLD = 10000  # 3+ deposits within 24h totaling > $10K
    STRUCTURING_COUNT = 3
    RAPID_WITHDRAW_RATIO = 3  # < 3x wagering of deposit
    LARGE_TRANSACTION_THRESHOLD = 10000
    VELOCITY_THRESHOLD = 10  # > 10 transactions/hour

    @staticmethod
    def check_structuring(user_id: str, hours: int = 24) -> Optional[Alert]:
        """Detect structuring: multiple deposits just below reporting threshold"""
        cutoff = datetime.now() - timedelta(hours=hours)

        deposits = [
            t for t in transactions_db.values()
            if t.user_id == user_id
            and t.type == "deposit"
            and t.timestamp > cutoff
            and t.amount < 10000  # Below single threshold
        ]

        total = sum(d.amount for d in deposits)

        if len(deposits) >= AMLRulesEngine.STRUCTURING_COUNT and total > AMLRulesEngine.STRUCTURING_THRESHOLD:
            return Alert(
                user_id=user_id,
                alert_type=AlertType.STRUCTURING,
                severity=AlertSeverity.HIGH,
                description=f"Structuring detected: {len(deposits)} deposits totaling ${total} within {hours}h",
                transactions=[t.transaction_id for t in deposits]
            )
        return None

    @staticmethod
    def check_rapid_deposit_withdraw(user_id: str) -> Optional[Alert]:
        """Detect rapid deposit-withdraw with minimal play"""
        recent_transactions = [
            t for t in transactions_db.values()
            if t.user_id == user_id
            and t.timestamp > datetime.now() - timedelta(days=7)
        ]

        deposits = [t for t in recent_transactions if t.type == "deposit"]
        withdrawals = [t for t in recent_transactions if t.type == "withdrawal"]
        bets = [t for t in recent_transactions if t.type == "bet"]

        if not deposits or not withdrawals:
            return None

        total_deposits = sum(d.amount for d in deposits)
        total_bets = sum(b.amount for b in bets)

        # Check if withdrawals exceed 3x bets (rapid withdraw pattern)
        if total_withdrawals := sum(w.amount for w in withdrawals):
            if total_bets < (total_deposits / AMLRulesEngine.RAPID_WITHDRAW_RATIO):
                return Alert(
                    user_id=user_id,
                    alert_type=AlertType.RAPID_DEPOSIT_WITHDRAW,
                    severity=AlertSeverity.MEDIUM,
                    description=f"Rapid deposit-withdraw: ${total_deposits} deposited, ${total_bets} wagered, ${total_withdrawals} withdrawn",
                    transactions=[t.transaction_id for t in recent_transactions]
                )
        return None

    @staticmethod
    def check_large_transaction(user_id: str, transaction: Transaction) -> Optional[Alert]:
        """Flag large transactions > $10K"""
        if transaction.amount > AMLRulesEngine.LARGE_TRANSACTION_THRESHOLD:
            return Alert(
                user_id=user_id,
                alert_type=AlertType.LARGE_TRANSACTION,
                severity=AlertSeverity.HIGH,
                description=f"Large transaction: ${transaction.amount} {transaction.type}",
                transactions=[transaction.transaction_id]
            )
        return None

    @staticmethod
    def check_velocity(user_id: str, hours: int = 1) -> Optional[Alert]:
        """Detect unusual transaction frequency"""
        cutoff = datetime.now() - timedelta(hours=hours)

        count = sum(
            1 for t in transactions_db.values()
            if t.user_id == user_id and t.timestamp > cutoff
        )

        if count > AMLRulesEngine.VELOCITY_THRESHOLD:
            return Alert(
                user_id=user_id,
                alert_type=AlertType.VELOCITY,
                severity=AlertSeverity.MEDIUM,
                description=f"High velocity: {count} transactions in {hours} hour(s)",
                transactions=[
                    t.transaction_id for t in transactions_db.values()
                    if t.user_id == user_id and t.timestamp > cutoff
                ]
            )
        return None

    @classmethod
    def run_all_rules(cls, user_id: str, transaction: Optional[Transaction] = None) -> List[Alert]:
        """Run all AML rules for a user"""
        alerts = []

        if alert := cls.check_structuring(user_id):
            alerts.append(alert)

        if alert := cls.check_rapid_deposit_withdraw(user_id):
            alerts.append(alert)

        if alert := cls.check_velocity(user_id):
            alerts.append(alert)

        if transaction and (alert := cls.check_large_transaction(user_id, transaction)):
            alerts.append(alert)

        return alerts
