from fastapi import APIRouter, Depends, HTTPException, status
from sqlalchemy.ext.asyncio import AsyncSession

from app.database import get_db
from app.services.kyc_service import KYCService
from app.api.schemas.kyc_schemas import KYCResponse
from app.api.document_routes import router as document_router
from app.api.verification_routes import router as verification_router
from app.api.review_routes import router as review_router
from app.api.limits_routes import router as limits_router

router = APIRouter()

router.include_router(document_router, tags=["KYC Documents"])
router.include_router(verification_router, tags=["KYC Verification"])
router.include_router(review_router, tags=["KYC Review"])
router.include_router(limits_router, tags=["KYC Limits"])


@router.post("/", response_model=KYCResponse)
async def create_kyc_record(
    user_id: str,
    db: AsyncSession = Depends(get_db)
):
    service = KYCService(db)
    kyc = await service.create_kyc_record(user_id)

    return KYCResponse(
        id=str(kyc.id),
        user_id=kyc.user_id,
        level=kyc.level.value,
        status=kyc.status.value,
        max_deposit_limit=kyc.max_deposit_limit or 0.0,
        max_withdrawal_limit=kyc.max_withdrawal_limit or 0.0,
        message="KYC record created"
    )


@router.get("/{user_id}", response_model=KYCResponse)
async def get_kyc_status(
    user_id: str,
    db: AsyncSession = Depends(get_db)
):
    service = KYCService(db)
    kyc = await service.get_kyc_by_user_id(user_id)

    if not kyc:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="KYC record not found"
        )

    return KYCResponse(
        id=str(kyc.id),
        user_id=kyc.user_id,
        level=kyc.level.value,
        status=kyc.status.value,
        max_deposit_limit=kyc.max_deposit_limit or 0.0,
        max_withdrawal_limit=kyc.max_withdrawal_limit or 0.0
    )
