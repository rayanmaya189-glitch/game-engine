"""Repository classes for fraud service database operations."""

from typing import Optional, List
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select, delete

from app import db_models


class DeviceFingerprintRepository:
    @staticmethod
    async def save(db: AsyncSession, record: db_models.DeviceFingerprintRecord) -> None:
        await db.execute(
            delete(db_models.DeviceFingerprintRecord).where(
                db_models.DeviceFingerprintRecord.user_id == record.user_id
            )
        )
        db.add(record)
        await db.commit()

    @staticmethod
    async def get_by_user(db: AsyncSession, user_id: str) -> Optional[db_models.DeviceFingerprintRecord]:
        result = await db.execute(
            select(db_models.DeviceFingerprintRecord).where(
                db_models.DeviceFingerprintRecord.user_id == user_id
            )
        )
        return result.scalar_one_or_none()


class IpAccountRepository:
    @staticmethod
    async def save(db: AsyncSession, record: db_models.IpAccountRecord) -> None:
        db.add(record)
        await db.commit()

    @staticmethod
    async def exists(db: AsyncSession, ip_address: str, user_id: str) -> bool:
        result = await db.execute(
            select(db_models.IpAccountRecord).where(
                db_models.IpAccountRecord.ip_address == ip_address,
                db_models.IpAccountRecord.user_id == user_id,
            )
        )
        return result.scalar_one_or_none() is not None


class FraudScoreRepository:
    @staticmethod
    async def save(db: AsyncSession, record: db_models.FraudScoreRecord) -> None:
        await db.execute(
            delete(db_models.FraudScoreRecord).where(db_models.FraudScoreRecord.user_id == record.user_id)
        )
        db.add(record)
        await db.commit()

    @staticmethod
    async def get_by_user(db: AsyncSession, user_id: str) -> Optional[db_models.FraudScoreRecord]:
        result = await db.execute(
            select(db_models.FraudScoreRecord).where(db_models.FraudScoreRecord.user_id == user_id)
        )
        return result.scalar_one_or_none()


class FraudAlertRepository:
    @staticmethod
    async def save(db: AsyncSession, record: db_models.FraudAlertRecord) -> None:
        db.add(record)
        await db.commit()

    @staticmethod
    async def list_alerts(
        db: AsyncSession,
        status: Optional[str] = None,
        limit: int = 50,
    ) -> List[db_models.FraudAlertRecord]:
        stmt = select(db_models.FraudAlertRecord)
        if status:
            stmt = stmt.where(db_models.FraudAlertRecord.status == status)
        stmt = stmt.order_by(db_models.FraudAlertRecord.created_at.desc()).limit(limit)
        result = await db.execute(stmt)
        return list(result.scalars().all())


class UserRiskProfileRepository:
    @staticmethod
    async def get_or_create(db: AsyncSession, user_id: str) -> db_models.UserRiskProfileRecord:
        result = await db.execute(
            select(db_models.UserRiskProfileRecord).where(
                db_models.UserRiskProfileRecord.user_id == user_id
            )
        )
        record = result.scalar_one_or_none()
        if not record:
            record = db_models.UserRiskProfileRecord(
                user_id=user_id,
                risk_score=0.0,
                is_blocked=False,
                flags="[]",
                transaction_count=0,
            )
            db.add(record)
            await db.commit()
            await db.refresh(record)
        return record

    @staticmethod
    async def save(db: AsyncSession, record: db_models.UserRiskProfileRecord) -> None:
        await db.commit()
        await db.refresh(record)
