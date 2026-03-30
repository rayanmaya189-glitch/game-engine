from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession

from app.database import get_db
from app.services.kyc_service import KYCService
from app.api.schemas.kyc_schemas import KYCLimitCheckRequest

router = APIRouter()


@router.post("/limits/check")
async def check_limits(
    request: KYCLimitCheckRequest,
    db: AsyncSession = Depends(get_db)
):
    service = KYCService(db)

    within_limits = await service.review.check_kyc_limits(
        user_id=request.user_id,
        amount=request.amount,
        transaction_type=request.transaction_type
    )

    return {
        "user_id": request.user_id,
        "amount": request.amount,
        "transaction_type": request.transaction_type,
        "within_limits": within_limits
    }
