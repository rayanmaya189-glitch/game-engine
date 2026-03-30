from datetime import datetime
from typing import Optional
import logging

from sqlalchemy import select
from sqlalchemy.ext.asyncio import AsyncSession

from app.models.kyc import KYCRecord, KYCStatus, VerificationProvider
from app.services.kyc_audit_service import KYCAuditService

logger = logging.getLogger(__name__)


class KYCVerificationService:
    def __init__(self, db: AsyncSession, audit: KYCAuditService):
        self.db = db
        self.audit = audit

    async def get_kyc_by_user_id(self, user_id: str) -> Optional[KYCRecord]:
        result = await self.db.execute(
            select(KYCRecord).where(KYCRecord.user_id == user_id)
        )
        return result.scalar_one_or_none()

    async def verify_document(
        self,
        user_id: str,
        provider: VerificationProvider,
        provider_reference: str,
        verification_result: dict,
    ) -> KYCRecord:
        kyc = await self.get_kyc_by_user_id(user_id)
        if not kyc:
            raise ValueError("KYC record not found")

        kyc.provider = provider
        kyc.provider_reference = provider_reference

        if verification_result.get("document_verified"):
            kyc.status = KYCStatus.VERIFIED
            kyc.verified_at = datetime.utcnow()
            await self.audit.update_kyc_level(kyc)
        else:
            kyc.status = KYCStatus.REJECTED
            kyc.review_notes = verification_result.get("reason", "Document verification failed")

        await self.db.commit()
        await self.db.refresh(kyc)

        await self.audit.log_audit(kyc.id, user_id, "DOCUMENT_VERIFIED", None, str(verification_result.get("document_verified")), "Document verification completed")

        return kyc

    async def verify_liveness(
        self,
        user_id: str,
        liveness_verified: bool,
        liveness_score: float,
    ) -> KYCRecord:
        kyc = await self.get_kyc_by_user_id(user_id)
        if not kyc:
            raise ValueError("KYC record not found")

        kyc.liveness_verified = liveness_verified
        kyc.liveness_score = liveness_score
        kyc.liveness_check_at = datetime.utcnow()

        if liveness_verified and kyc.status == KYCStatus.IN_PROGRESS:
            await self.audit.check_verification_complete(kyc)

        await self.db.commit()
        await self.db.refresh(kyc)

        await self.audit.log_audit(kyc.id, user_id, "LIVENESS_CHECK", None, str(liveness_verified), f"Liveness score: {liveness_score}")

        return kyc

    async def verify_address(
        self,
        user_id: str,
        address_verified: bool,
    ) -> KYCRecord:
        kyc = await self.get_kyc_by_user_id(user_id)
        if not kyc:
            raise ValueError("KYC record not found")

        kyc.address_verified = address_verified
        kyc.address_verification_at = datetime.utcnow()

        if address_verified and kyc.status == KYCStatus.IN_PROGRESS:
            await self.audit.check_verification_complete(kyc)

        await self.db.commit()
        await self.db.refresh(kyc)

        await self.audit.log_audit(kyc.id, user_id, "ADDRESS_VERIFIED", None, str(address_verified), "Address verification completed")

        return kyc

    async def submit_address_verification(
        self,
        user_id: str,
        address_line_1: str,
        address_line_2: Optional[str],
        city: str,
        state: Optional[str],
        postal_code: str,
        country_of_residence: str,
    ) -> KYCRecord:
        kyc = await self.get_kyc_by_user_id(user_id)
        if not kyc:
            raise ValueError("KYC record not found")

        kyc.address_line_1 = address_line_1
        kyc.address_line_2 = address_line_2
        kyc.city = city
        kyc.state = state
        kyc.postal_code = postal_code
        kyc.country_of_residence = country_of_residence

        await self.db.commit()
        await self.db.refresh(kyc)

        await self.audit.log_audit(kyc.id, user_id, "ADDRESS_SUBMITTED", None, city, "Address verification submitted")

        return kyc
