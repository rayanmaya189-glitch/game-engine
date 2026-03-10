from datetime import datetime, timedelta
from typing import Optional, List
from uuid import uuid4
import logging

from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession

from app.models.kyc import KYCRecord, KYCDocument, KYCQueueItem, KYCAuditLog, KYCLevel, KYCStatus, DocumentType, VerificationProvider
from app.config import settings

logger = logging.getLogger(__name__)


class KYCService:
    def __init__(self, db: AsyncSession):
        self.db = db

    async def get_kyc_by_user_id(self, user_id: str) -> Optional[KYCRecord]:
        result = await self.db.execute(
            select(KYCRecord).where(KYCRecord.user_id == user_id)
        )
        return result.scalar_one_or_none()

    async def create_kyc_record(self, user_id: str) -> KYCRecord:
        kyc = KYCRecord(
            id=uuid4(),
            user_id=user_id,
            level=KYCLevel.NONE,
            status=KYCStatus.NOT_STARTED,
        )
        self.db.add(kyc)
        await self.db.commit()
        await self.db.refresh(kyc)
        
        await self._log_audit(kyc.id, user_id, "CREATE", None, "NONE", "KYC record created")
        return kyc

    async def submit_document(
        self,
        user_id: str,
        document_type: DocumentType,
        document_number: str,
        document_expiry: datetime,
        front_image_url: str,
        back_image_url: Optional[str] = None,
        selfie_image_url: Optional[str] = None,
    ) -> KYCRecord:
        kyc = await self.get_kyc_by_user_id(user_id)
        if not kyc:
            kyc = await self.create_kyc_record(user_id)

        # Update KYC record
        kyc.document_type = document_type
        kyc.document_number = document_number
        kyc.document_expiry = document_expiry
        kyc.document_front_url = front_image_url
        kyc.document_back_url = back_image_url
        kyc.document_selfie_url = selfie_image_url
        kyc.status = KYCStatus.IN_PROGRESS
        kyc.submitted_at = datetime.utcnow()

        # Create document record
        document = KYCDocument(
            id=uuid4(),
            kyc_record_id=kyc.id,
            user_id=user_id,
            document_type=document_type,
            document_number=document_number,
            front_image_url=front_image_url,
            back_image_url=back_image_url,
            selfie_image_url=selfie_image_url,
            is_primary=True,
        )
        self.db.add(document)

        await self.db.commit()
        await self.db.refresh(kyc)
        
        await self._log_audit(kyc.id, user_id, "DOCUMENT_SUBMITTED", None, document_type.value, "Document submitted")
        
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
        
        await self._log_audit(kyc.id, user_id, "ADDRESS_SUBMITTED", None, city, "Address verification submitted")
        
        return kyc

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

        # Update verification details
        kyc.provider = provider
        kyc.provider_reference = provider_reference
        
        if verification_result.get("document_verified"):
            kyc.status = KYCStatus.VERIFIED
            kyc.verified_at = datetime.utcnow()
            await self._update_kyc_level(kyc)
        else:
            kyc.status = KYCStatus.REJECTED
            kyc.review_notes = verification_result.get("reason", "Document verification failed")

        await self.db.commit()
        await self.db.refresh(kyc)
        
        await self._log_audit(kyc.id, user_id, "DOCUMENT_VERIFIED", None, str(verification_result.get("document_verified")), "Document verification completed")
        
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
            await self._check_verification_complete(kyc)

        await self.db.commit()
        await self.db.refresh(kyc)
        
        await self._log_audit(kyc.id, user_id, "LIVENESS_CHECK", None, str(liveness_verified), f"Liveness score: {liveness_score}")
        
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
            await self._check_verification_complete(kyc)

        await self.db.commit()
        await self.db.refresh(kyc)
        
        await self._log_audit(kyc.id, user_id, "ADDRESS_VERIFIED", None, str(address_verified), "Address verification completed")
        
        return kyc

    async def approve_kyc(self, user_id: str, reviewer_notes: Optional[str] = None) -> KYCRecord:
        kyc = await self.get_kyc_by_user_id(user_id)
        if not kyc:
            raise ValueError("KYC record not found")

        previous_level = kyc.level
        kyc.status = KYCStatus.VERIFIED
        kyc.verified_at = datetime.utcnow()
        kyc.reviewed_at = datetime.utcnow()
        kyc.review_notes = reviewer_notes
        
        await self._update_kyc_level(kyc)
        
        await self.db.commit()
        await self.db.refresh(kyc)
        
        await self._log_audit(kyc.id, user_id, "KYC_APPROVED", previous_level.value, kyc.level.value, reviewer_notes)
        
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
        
        await self._log_audit(kyc.id, user_id, "KYC_REJECTED", previous_status.value, KYCStatus.REJECTED.value, reason)
        
        return kyc

    async def get_pending_review_queue(self, limit: int = 50) -> List[KYCRecord]:
        result = await self.db.execute(
            select(KYCRecord)
            .where(KYCRecord.status == KYCStatus.NEEDS_REVIEW)
            .order_by(KYCRecord.submitted_at.desc())
            .limit(limit)
        )
        return list(result.scalars().all())

    async def get_kyc_by_id(self, kyc_id: str) -> Optional[KYCRecord]:
        result = await self.db.execute(
            select(KYCRecord).where(KYCRecord.id == kyc_id)
        )
        return result.scalar_one_or_none()

    async def check_kyc_limits(self, user_id: str, amount: float, transaction_type: str) -> bool:
        kyc = await self.get_kyc_by_user_id(user_id)
        if not kyc:
            return False

        if transaction_type == "deposit":
            return amount <= kyc.max_deposit_limit
        elif transaction_type == "withdrawal":
            return amount <= kyc.max_withdrawal_limit
        
        return False

    async def _update_kyc_level(self, kyc: KYCRecord):
        """Update KYC level based on verification status"""
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

    async def _check_verification_complete(self, kyc: KYCRecord):
        """Check if all verifications are complete"""
        if kyc.liveness_verified and kyc.address_verified:
            kyc.status = KYCStatus.VERIFIED
            kyc.verified_at = datetime.utcnow()
            await self._update_kyc_level(kyc)

    async def _log_audit(
        self,
        kyc_record_id: str,
        user_id: str,
        action: str,
        previous_value: Optional[str],
        new_value: str,
        reason: Optional[str],
    ):
        """Log KYC audit event"""
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
