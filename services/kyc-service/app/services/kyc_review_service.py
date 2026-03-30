from datetime import datetime
from typing import Optional, List
import logging

from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession

from app.models.kyc import KYCRecord, KYCLevel, KYCStatus
from app.services.kyc_audit_service import KYCAuditService

logger = logging.getLogger(__name__)


class KYCReviewService:
    def __init__(self, db: AsyncSession, audit: KYCAuditService):
        self.db = db
        self.audit = audit

    async def get_kyc_by_user_id(self, user_id: str) -> Optional[KYCRecord]:
        result = await self.db.execute(
            select(KYCRecord).where(KYCRecord.user_id == user_id)
        )
        return result.scalar_one_or_none()

    async def approve_kyc(self, user_id: str, reviewer_notes: Optional[str] = None) -> KYCRecord:
        kyc = await self.get_kyc_by_user_id(user_id)
        if not kyc:
            raise ValueError("KYC record not found")

        previous_level = kyc.level
        kyc.status = KYCStatus.VERIFIED
        kyc.verified_at = datetime.utcnow()
        kyc.reviewed_at = datetime.utcnow()
        kyc.review_notes = reviewer_notes

        await self.audit.update_kyc_level(kyc)

        await self.db.commit()
        await self.db.refresh(kyc)

        await self.audit.log_audit(kyc.id, user_id, "KYC_APPROVED", previous_level.value, kyc.level.value, reviewer_notes)

        return kyc

    async def reject_kyc(self, user_id: str, reason: str) -> KYCRecord:
        kyc = await self.get_kyc_by_user_id(user_id)
        if not kyc:
            raise ValueError("KYC record not found")

        previous_status = kyc.status
        kyc.status = KYCStatus.REJECTED
        kyc.reviewed_at = datetime.utcnow()
        kyc.review_notes = reason

        await self.db.commit()
        await self.db.refresh(kyc)

        await self.audit.log_audit(kyc.id, user_id, "KYC_REJECTED", previous_status.value, KYCStatus.REJECTED.value, reason)

        return kyc

    async def get_pending_review_queue(self, limit: int = 50) -> List[KYCRecord]:
        result = await self.db.execute(
            select(KYCRecord)
            .where(KYCRecord.status == KYCStatus.NEEDS_REVIEW)
            .order_by(KYCRecord.submitted_at.desc())
            .limit(limit)
        )
        return list(result.scalars().all())

    async def check_kyc_limits(self, user_id: str, amount: float, transaction_type: str) -> bool:
        kyc = await self.get_kyc_by_user_id(user_id)
        if not kyc:
            return False

        if transaction_type == "deposit":
            return amount <= kyc.max_deposit_limit
        elif transaction_type == "withdrawal":
            return amount <= kyc.max_withdrawal_limit

        return False
