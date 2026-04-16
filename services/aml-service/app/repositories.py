"""Repository classes for AML service database operations."""

from typing import Optional, List
from datetime import datetime
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select, delete

from app import db_models


class TransactionRepository:
    @staticmethod
    async def save(db: AsyncSession, record: db_models.TransactionRecord) -> None:
        db.add(record)
        await db.commit()

    @staticmethod
    async def get_by_user(db: AsyncSession, user_id: str) -> List[db_models.TransactionRecord]:
        result = await db.execute(
            select(db_models.TransactionRecord).where(db_models.TransactionRecord.user_id == user_id)
        )
        return list(result.scalars().all())

    @staticmethod
    async def get_by_id(db: AsyncSession, transaction_id: str) -> Optional[db_models.TransactionRecord]:
        result = await db.execute(
            select(db_models.TransactionRecord).where(db_models.TransactionRecord.transaction_id == transaction_id)
        )
        return result.scalar_one_or_none()

    @staticmethod
    async def get_by_user_and_period(
        db: AsyncSession, user_id: str, start: datetime, end: datetime
    ) -> List[db_models.TransactionRecord]:
        result = await db.execute(
            select(db_models.TransactionRecord).where(
                db_models.TransactionRecord.user_id == user_id,
                db_models.TransactionRecord.timestamp >= start,
                db_models.TransactionRecord.timestamp <= end,
            )
        )
        return list(result.scalars().all())


class AlertRepository:
    @staticmethod
    async def save(db: AsyncSession, record: db_models.AlertRecord) -> None:
        db.add(record)
        await db.commit()

    @staticmethod
    async def get_by_id(db: AsyncSession, alert_id: str) -> Optional[db_models.AlertRecord]:
        result = await db.execute(
            select(db_models.AlertRecord).where(db_models.AlertRecord.alert_id == alert_id)
        )
        return result.scalar_one_or_none()

    @staticmethod
    async def get_by_user(db: AsyncSession, user_id: str, limit: int = 100) -> List[db_models.AlertRecord]:
        result = await db.execute(
            select(db_models.AlertRecord)
            .where(db_models.AlertRecord.user_id == user_id)
            .order_by(db_models.AlertRecord.created_at.desc())
            .limit(limit)
        )
        return list(result.scalars().all())

    @staticmethod
    async def list_alerts(
        db: AsyncSession,
        status: Optional[str] = None,
        severity: Optional[str] = None,
        limit: int = 100,
    ) -> List[db_models.AlertRecord]:
        stmt = select(db_models.AlertRecord)
        if status:
            stmt = stmt.where(db_models.AlertRecord.status == status)
        if severity:
            stmt = stmt.where(db_models.AlertRecord.severity == severity)
        stmt = stmt.order_by(db_models.AlertRecord.created_at.desc()).limit(limit)
        result = await db.execute(stmt)
        return list(result.scalars().all())

    @staticmethod
    async def update(db: AsyncSession, record: db_models.AlertRecord) -> None:
        await db.commit()
        await db.refresh(record)


class RiskScoreRepository:
    @staticmethod
    async def save(db: AsyncSession, record: db_models.RiskScoreRecord) -> None:
        await db.execute(
            delete(db_models.RiskScoreRecord).where(db_models.RiskScoreRecord.user_id == record.user_id)
        )
        db.add(record)
        await db.commit()

    @staticmethod
    async def get_by_user(db: AsyncSession, user_id: str) -> Optional[db_models.RiskScoreRecord]:
        result = await db.execute(
            select(db_models.RiskScoreRecord).where(db_models.RiskScoreRecord.user_id == user_id)
        )
        return result.scalar_one_or_none()
