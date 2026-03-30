"""Repository classes for AML service database operations."""

from typing import Optional, List
from datetime import datetime
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select, delete

from app.models import TransactionRecord, AlertRecord, RiskScoreRecord


class TransactionRepository:
    @staticmethod
    async def save(db: AsyncSession, record: TransactionRecord) -> None:
        db.add(record)
        await db.commit()

    @staticmethod
    async def get_by_user(db: AsyncSession, user_id: str) -> List[TransactionRecord]:
        result = await db.execute(
            select(TransactionRecord).where(TransactionRecord.user_id == user_id)
        )
        return list(result.scalars().all())

    @staticmethod
    async def get_by_id(db: AsyncSession, transaction_id: str) -> Optional[TransactionRecord]:
        result = await db.execute(
            select(TransactionRecord).where(TransactionRecord.transaction_id == transaction_id)
        )
        return result.scalar_one_or_none()

    @staticmethod
    async def get_by_user_and_period(
        db: AsyncSession, user_id: str, start: datetime, end: datetime
    ) -> List[TransactionRecord]:
        result = await db.execute(
            select(TransactionRecord).where(
                TransactionRecord.user_id == user_id,
                TransactionRecord.timestamp >= start,
                TransactionRecord.timestamp <= end,
            )
        )
        return list(result.scalars().all())


class AlertRepository:
    @staticmethod
    async def save(db: AsyncSession, record: AlertRecord) -> None:
        db.add(record)
        await db.commit()

    @staticmethod
    async def get_by_id(db: AsyncSession, alert_id: str) -> Optional[AlertRecord]:
        result = await db.execute(
            select(AlertRecord).where(AlertRecord.alert_id == alert_id)
        )
        return result.scalar_one_or_none()

    @staticmethod
    async def get_by_user(db: AsyncSession, user_id: str, limit: int = 100) -> List[AlertRecord]:
        result = await db.execute(
            select(AlertRecord)
            .where(AlertRecord.user_id == user_id)
            .order_by(AlertRecord.created_at.desc())
            .limit(limit)
        )
        return list(result.scalars().all())

    @staticmethod
    async def list_alerts(
        db: AsyncSession,
        status: Optional[str] = None,
        severity: Optional[str] = None,
        limit: int = 100,
    ) -> List[AlertRecord]:
        stmt = select(AlertRecord)
        if status:
            stmt = stmt.where(AlertRecord.status == status)
        if severity:
            stmt = stmt.where(AlertRecord.severity == severity)
        stmt = stmt.order_by(AlertRecord.created_at.desc()).limit(limit)
        result = await db.execute(stmt)
        return list(result.scalars().all())

    @staticmethod
    async def update(db: AsyncSession, record: AlertRecord) -> None:
        await db.commit()
        await db.refresh(record)


class RiskScoreRepository:
    @staticmethod
    async def save(db: AsyncSession, record: RiskScoreRecord) -> None:
        await db.execute(
            delete(RiskScoreRecord).where(RiskScoreRecord.user_id == record.user_id)
        )
        db.add(record)
        await db.commit()

    @staticmethod
    async def get_by_user(db: AsyncSession, user_id: str) -> Optional[RiskScoreRecord]:
        result = await db.execute(
            select(RiskScoreRecord).where(RiskScoreRecord.user_id == user_id)
        )
        return result.scalar_one_or_none()
