import os
from typing import Optional, List
from datetime import datetime, timedelta
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select

from app.models import TransactionRecord
from app.models.schemas import Alert, AlertType, AlertSeverity, Transaction


class AMLRulesEngine:
    """Rules engine for detecting suspicious transaction patterns"""

    STRUCTURING_THRESHOLD = int(os.environ.get("AML_STRUCTURING_THRESHOLD", "10000"))
    STRUCTURING_COUNT = int(os.environ.get("AML_STRUCTURING_COUNT", "3"))
    RAPID_WITHDRAW_RATIO = int(os.environ.get("AML_RAPID_WITHDRAW_RATIO", "3"))
    LARGE_TRANSACTION_THRESHOLD = int(os.environ.get("AML_LARGE_TRANSACTION_THRESHOLD", "10000"))
    VELOCITY_THRESHOLD = int(os.environ.get("AML_VELOCITY_THRESHOLD", "10"))

    @staticmethod
    async def check_structuring(db: AsyncSession, user_id: str, hours: int = 24) -> Optional[Alert]:
        """Detect structuring: multiple deposits just below reporting threshold"""
        cutoff = datetime.utcnow() - timedelta(hours=hours)

        stmt = select(TransactionRecord).where(
            TransactionRecord.user_id == user_id,
            TransactionRecord.type == "deposit",
            TransactionRecord.timestamp > cutoff,
            TransactionRecord.amount < 10000,
        )
        result = await db.execute(stmt)
        deposits = result.scalars().all()

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
    async def check_rapid_deposit_withdraw(db: AsyncSession, user_id: str) -> Optional[Alert]:
        """Detect rapid deposit-withdraw with minimal play"""
        cutoff = datetime.utcnow() - timedelta(days=7)
        stmt = select(TransactionRecord).where(
            TransactionRecord.user_id == user_id,
            TransactionRecord.timestamp > cutoff,
        )
        result = await db.execute(stmt)
        recent_transactions = result.scalars().all()

        deposits = [t for t in recent_transactions if t.type == "deposit"]
        withdrawals = [t for t in recent_transactions if t.type == "withdrawal"]
        bets = [t for t in recent_transactions if t.type == "bet"]

        if not deposits or not withdrawals:
            return None

        total_deposits = sum(d.amount for d in deposits)
        total_bets = sum(b.amount for b in bets)
        total_withdrawals = sum(w.amount for w in withdrawals)

        if total_withdrawals:
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
    async def check_velocity(db: AsyncSession, user_id: str, hours: int = 1) -> Optional[Alert]:
        """Detect unusual transaction frequency"""
        cutoff = datetime.utcnow() - timedelta(hours=hours)

        stmt = select(TransactionRecord).where(
            TransactionRecord.user_id == user_id,
            TransactionRecord.timestamp > cutoff,
        )
        result = await db.execute(stmt)
        recent = result.scalars().all()
        count = len(recent)

        if count > AMLRulesEngine.VELOCITY_THRESHOLD:
            return Alert(
                user_id=user_id,
                alert_type=AlertType.VELOCITY,
                severity=AlertSeverity.MEDIUM,
                description=f"High velocity: {count} transactions in {hours} hour(s)",
                transactions=[t.transaction_id for t in recent]
            )
        return None

    @classmethod
    async def run_all_rules(cls, db: AsyncSession, user_id: str, transaction: Optional[Transaction] = None) -> List[Alert]:
        """Run all AML rules for a user"""
        alerts = []

        if alert := await cls.check_structuring(db, user_id):
            alerts.append(alert)

        if alert := await cls.check_rapid_deposit_withdraw(db, user_id):
            alerts.append(alert)

        if alert := await cls.check_velocity(db, user_id):
            alerts.append(alert)

        if transaction and (alert := cls.check_large_transaction(user_id, transaction)):
            alerts.append(alert)

        return alerts
