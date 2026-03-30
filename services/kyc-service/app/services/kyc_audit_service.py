from datetime import datetime, timedelta
from typing import Optional
from uuid import uuid4
import logging

from sqlalchemy.ext.asyncio import AsyncSession

from app.models.kyc import KYCRecord, KYCAuditLog, KYCLevel, KYCStatus
from app.config import settings

logger = logging.getLogger(__name__)


class KYCAuditService:
    def __init__(self, db: AsyncSession):
        self.db = db

    async def log_audit(
        self,
        kyc_record_id: str,
        user_id: str,
        action: str,
        previous_value: Optional[str],
        new_value: str,
        reason: Optional[str],
    ):
        audit = KYCAuditLog(
            id=uuid4(),
            kyc_record_id=kyc_record_id,
            user_id=user_id,
            action=action,
            previous_value=previous_value,
            new_value=new_value,
            reason=reason,
        )
        self.db.add(audit)
        await self.db.commit()

    async def update_kyc_level(self, kyc: KYCRecord):
        if kyc.liveness_verified and kyc.address_verified and kyc.provider:
            kyc.level = KYCLevel.FULL
            kyc.max_deposit_limit = settings.full_max_deposit
            kyc.max_withdrawal_limit = settings.full_max_deposit
            kyc.expires_at = datetime.utcnow() + timedelta(days=365)
        elif kyc.liveness_verified and kyc.provider:
            kyc.level = KYCLevel.INTERMEDIATE
            kyc.max_deposit_limit = settings.intermediate_max_deposit
            kyc.max_withdrawal_limit = settings.intermediate_max_deposit
            kyc.expires_at = datetime.utcnow() + timedelta(days=365)
        elif kyc.provider:
            kyc.level = KYCLevel.BASIC
            kyc.max_deposit_limit = settings.basic_max_deposit
            kyc.max_withdrawal_limit = settings.basic_max_deposit
            kyc.expires_at = datetime.utcnow() + timedelta(days=365)

        if kyc.requires_manual_review:
            kyc.status = KYCStatus.NEEDS_REVIEW

    async def check_verification_complete(self, kyc: KYCRecord):
        if kyc.liveness_verified and kyc.address_verified:
            kyc.status = KYCStatus.VERIFIED
            kyc.verified_at = datetime.utcnow()
            await self.update_kyc_level(kyc)
