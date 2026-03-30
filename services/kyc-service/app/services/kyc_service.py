from typing import Optional
import logging

from sqlalchemy.ext.asyncio import AsyncSession

from app.models.kyc import KYCRecord
from app.services.kyc_audit_service import KYCAuditService
from app.services.kyc_document_service import KYCDocumentService
from app.services.kyc_verification_service import KYCVerificationService
from app.services.kyc_review_service import KYCReviewService

logger = logging.getLogger(__name__)


class KYCService:
    def __init__(self, db: AsyncSession):
        self.db = db
        self.audit = KYCAuditService(db)
        self.document = KYCDocumentService(db, self.audit)
        self.verification = KYCVerificationService(db, self.audit)
        self.review = KYCReviewService(db, self.audit)

    async def get_kyc_by_user_id(self, user_id: str) -> Optional[KYCRecord]:
        return await self.document.get_kyc_by_user_id(user_id)

    async def create_kyc_record(self, user_id: str) -> KYCRecord:
        return await self.document.create_kyc_record(user_id)

    async def get_kyc_by_id(self, kyc_id: str) -> Optional[KYCRecord]:
        from sqlalchemy import select
        result = await self.db.execute(
            select(KYCRecord).where(KYCRecord.id == kyc_id)
        )
        return result.scalar_one_or_none()
