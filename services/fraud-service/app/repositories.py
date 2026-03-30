"""Repository classes for fraud service database operations."""

from typing import Optional, List
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select, delete

from app.models import (
    DeviceFingerprintRecord,
    IpAccountRecord,
    FraudScoreRecord,
    FraudAlertRecord,
    UserRiskProfileRecord,
)


class DeviceFingerprintRepository:
    @staticmethod
    async def save(db: AsyncSession, record: DeviceFingerprintRecord) -> None:
        await db.execute(
            delete(DeviceFingerprintRecord).where(
                DeviceFingerprintRecord.user_id == record.user_id
            )
        )
        db.add(record)
        await db.commit()

    @staticmethod
    async def get_by_user(db: AsyncSession, user_id: str) -> Optional[DeviceFingerprintRecord]:
        result = await db.execute(
            select(DeviceFingerprintRecord).where(
                DeviceFingerprintRecord.user_id == user_id
            )
        )
        return result.scalar_one_or_none()


class IpAccountRepository:
    @staticmethod
    async def save(db: AsyncSession, record: IpAccountRecord) -> None:
        db.add(record)
        await db.commit()

    @staticmethod
    async def exists(db: AsyncSession, ip_address: str, user_id: str) -> bool:
        result = await db.execute(
            select(IpAccountRecord).where(
                IpAccountRecord.ip_address == ip_address,
                IpAccountRecord.user_id == user_id,
            )
        )
        return result.scalar_one_or_none() is not None


class FraudScoreRepository:
    @staticmethod
    async def save(db: AsyncSession, record: FraudScoreRecord) -> None:
        await db.execute(
            delete(FraudScoreRecord).where(FraudScoreRecord.user_id == record.user_id)
        )
        db.add(record)
        await db.commit()

    @staticmethod
    async def get_by_user(db: AsyncSession, user_id: str) -> Optional[FraudScoreRecord]:
        result = await db.execute(
            select(FraudScoreRecord).where(FraudScoreRecord.user_id == user_id)
        )
        return result.scalar_one_or_none()


class FraudAlertRepository:
    @staticmethod
    async def save(db: AsyncSession, record: FraudAlertRecord) -> None:
        db.add(record)
        await db.commit()

    @staticmethod
    async def list_alerts(
        db: AsyncSession,
        status: Optional[str] = None,
        limit: int = 50,
    ) -> List[FraudAlertRecord]:
        stmt = select(FraudAlertRecord)
        if status:
            stmt = stmt.where(FraudAlertRecord.status == status)
        stmt = stmt.order_by(FraudAlertRecord.created_at.desc()).limit(limit)
        result = await db.execute(stmt)
        return list(result.scalars().all())


class UserRiskProfileRepository:
    @staticmethod
    async def get_or_create(db: AsyncSession, user_id: str) -> UserRiskProfileRecord:
        result = await db.execute(
            select(UserRiskProfileRecord).where(
                UserRiskProfileRecord.user_id == user_id
            )
        )
        record = result.scalar_one_or_none()
        if not record:
            record = UserRiskProfileRecord(
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
    async def save(db: AsyncSession, record: UserRiskProfileRecord) -> None:
        await db.commit()
        await db.refresh(record)
