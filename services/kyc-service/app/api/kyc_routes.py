from sqlalchemy.ext.asyncio import AsyncSession

from app.database import get_db
from app.services.kyc_service import KYCService
from app.api.schemas.kyc_schemas import KYCResponse


async def create_kyc_record(
    user_id: str,
    db: AsyncSession
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


async def get_kyc_status(
    user_id: str,
    db: AsyncSession
):
    service = KYCService(db)
    kyc = await service.get_kyc_by_user_id(user_id)

    if not kyc:
        raise ValueError("KYC record not found")

    return KYCResponse(
        id=str(kyc.id),
        user_id=kyc.user_id,
        level=kyc.level.value,
        status=kyc.status.value,
        max_deposit_limit=kyc.max_deposit_limit or 0.0,
        max_withdrawal_limit=kyc.max_withdrawal_limit or 0.0
    )
