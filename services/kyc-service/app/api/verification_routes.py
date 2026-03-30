from fastapi import APIRouter, Depends
from sqlalchemy.ext.asyncio import AsyncSession

from app.database import get_db
from app.services.kyc_service import KYCService
from app.api.schemas.kyc_schemas import (
    VerificationResultRequest,
    LivenessCheckRequest,
    AddressVerificationRequest,
    KYCResponse,
)

router = APIRouter()


@router.post("/verify", response_model=KYCResponse)
async def verify_document(
    request: VerificationResultRequest,
    db: AsyncSession = Depends(get_db)
):
    service = KYCService(db)

    kyc = await service.verification.verify_document(
        user_id=request.user_id,
        provider=request.provider,
        provider_reference=request.provider_reference,
        verification_result={
            "document_verified": request.document_verified,
            "reason": request.reason
        },
    )

    return KYCResponse(
        id=str(kyc.id),
        user_id=kyc.user_id,
        level=kyc.level.value,
        status=kyc.status.value,
        max_deposit_limit=kyc.max_deposit_limit or 0.0,
        max_withdrawal_limit=kyc.max_withdrawal_limit or 0.0,
        message="Document verification completed"
    )


@router.post("/liveness", response_model=KYCResponse)
async def verify_liveness(
    request: LivenessCheckRequest,
    db: AsyncSession = Depends(get_db)
):
    service = KYCService(db)

    kyc = await service.verification.verify_liveness(
        user_id=request.user_id,
        liveness_verified=request.liveness_verified,
        liveness_score=request.liveness_score,
    )

    return KYCResponse(
        id=str(kyc.id),
        user_id=kyc.user_id,
        level=kyc.level.value,
        status=kyc.status.value,
        max_deposit_limit=kyc.max_deposit_limit or 0.0,
        max_withdrawal_limit=kyc.max_withdrawal_limit or 0.0,
        message="Liveness check completed"
    )


@router.post("/address/verify", response_model=KYCResponse)
async def verify_address(
    user_id: str,
    address_verified: bool,
    db: AsyncSession = Depends(get_db)
):
    service = KYCService(db)

    kyc = await service.verification.verify_address(
        user_id=user_id,
        address_verified=address_verified,
    )

    return KYCResponse(
        id=str(kyc.id),
        user_id=kyc.user_id,
        level=kyc.level.value,
        status=kyc.status.value,
        max_deposit_limit=kyc.max_deposit_limit or 0.0,
        max_withdrawal_limit=kyc.max_withdrawal_limit or 0.0,
        message="Address verification completed"
    )


@router.post("/address", response_model=KYCResponse)
async def submit_address(
    request: AddressVerificationRequest,
    db: AsyncSession = Depends(get_db)
):
    service = KYCService(db)

    kyc = await service.verification.submit_address_verification(
        user_id=request.user_id,
        address_line_1=request.address_line_1,
        address_line_2=request.address_line_2,
        city=request.city,
        state=request.state,
        postal_code=request.postal_code,
        country_of_residence=request.country_of_residence,
    )

    return KYCResponse(
        id=str(kyc.id),
        user_id=kyc.user_id,
        level=kyc.level.value,
        status=kyc.status.value,
        max_deposit_limit=kyc.max_deposit_limit or 0.0,
        max_withdrawal_limit=kyc.max_withdrawal_limit or 0.0,
        message="Address submitted successfully"
    )
