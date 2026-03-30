from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession

from app.database import get_db
from app.services.kyc_service import KYCService
from app.api.schemas.kyc_schemas import ManualReviewRequest, KYCResponse

router = APIRouter()


@router.post("/review", response_model=KYCResponse)
async def manual_review(
    request: ManualReviewRequest,
    db: AsyncSession = Depends(get_db)
):
    service = KYCService(db)

    if request.approved:
        kyc = await service.review.approve_kyc(request.kyc_id, request.reviewer_notes)
        message = "KYC approved"
    else:
        kyc = await service.review.reject_kyc(request.kyc_id, request.reviewer_notes or "Rejected during manual review")
        message = "KYC rejected"

    return KYCResponse(
        id=str(kyc.id),
        user_id=kyc.user_id,
        level=kyc.level.value,
        status=kyc.status.value,
        max_deposit_limit=kyc.max_deposit_limit or 0.0,
        max_withdrawal_limit=kyc.max_withdrawal_limit or 0.0,
        message=message
    )


@router.get("/queue/pending")
async def get_pending_review_queue(
    limit: int = 50,
    db: AsyncSession = Depends(get_db)
):
    service = KYCService(db)
    records = await service.review.get_pending_review_queue(limit)

    return {
        "count": len(records),
        "items": [
            {
                "id": str(r.id),
                "user_id": r.user_id,
                "status": r.status.value,
                "submitted_at": r.submitted_at.isoformat() if r.submitted_at else None
            }
            for r in records
        ]
    }
