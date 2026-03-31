from datetime import datetime
from sqlalchemy.ext.asyncio import AsyncSession

from app.database import get_db
from app.services.kyc_service import KYCService
from app.api.schemas.kyc_schemas import KYCSubmissionRequest, KYCResponse


async def submit_document(
    request: KYCSubmissionRequest,
    db: AsyncSession
):
    service = KYCService(db)

    try:
        document_expiry = datetime.fromisoformat(request.document_expiry.replace("Z", "+00:00"))
    except ValueError:
        raise ValueError("Invalid date format. Use ISO format.")

    kyc = await service.document.submit_document(
        user_id=request.user_id,
        document_type=request.document_type,
        document_number=request.document_number,
        document_expiry=document_expiry,
        front_image_url=request.front_image_url,
        back_image_url=request.back_image_url,
        selfie_image_url=request.selfie_image_url,
    )

    return KYCResponse(
        id=str(kyc.id),
        user_id=kyc.user_id,
        level=kyc.level.value,
        status=kyc.status.value,
        max_deposit_limit=kyc.max_deposit_limit or 0.0,
        max_withdrawal_limit=kyc.max_withdrawal_limit or 0.0,
        message="Document submitted successfully"
    )
