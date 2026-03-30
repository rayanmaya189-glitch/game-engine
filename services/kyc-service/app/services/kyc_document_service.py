from datetime import datetime
from typing import Optional
from uuid import uuid4
import logging

from sqlalchemy.ext.asyncio import AsyncSession

from app.models.kyc import KYCRecord, KYCDocument, KYCLevel, KYCStatus, DocumentType
from app.services.kyc_audit_service import KYCAuditService

logger = logging.getLogger(__name__)


class KYCDocumentService:
    def __init__(self, db: AsyncSession, audit: KYCAuditService):
        self.db = db
        self.audit = audit

    async def get_kyc_by_user_id(self, user_id: str) -> Optional[KYCRecord]:
        from sqlalchemy import select
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

        await self.audit.log_audit(kyc.id, user_id, "CREATE", None, "NONE", "KYC record created")
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

        kyc.document_type = document_type
        kyc.document_number = document_number
        kyc.document_expiry = document_expiry
        kyc.document_front_url = front_image_url
        kyc.document_back_url = back_image_url
        kyc.document_selfie_url = selfie_image_url
        kyc.status = KYCStatus.IN_PROGRESS
        kyc.submitted_at = datetime.utcnow()

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

        await self.audit.log_audit(kyc.id, user_id, "DOCUMENT_SUBMITTED", None, document_type.value, "Document submitted")

        return kyc
