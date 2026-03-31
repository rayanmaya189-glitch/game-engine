"""
Verification, address, and limits RPCs for KYC Service.
"""

import logging
from datetime import datetime

import grpc

from app.database import AsyncSessionLocal
from app.services.kyc_service import KYCService

logger = logging.getLogger(__name__)


class KYCServiceServicer:

    async def VerifyDocument(self, request, context):
        async with AsyncSessionLocal() as db:
            try:
                from app.models.kyc import VerificationProvider
                service = KYCService(db)

                kyc = await service.verification.verify_document(
                    user_id=request.user_id,
                    provider=VerificationProvider(request.provider),
                    provider_reference=request.provider_reference,
                    verification_result={
                        "document_verified": getattr(request, "document_verified", False),
                        "reason": getattr(request, "reason", None),
                    },
                )

                return {
                    "id": str(kyc.id),
                    "user_id": kyc.user_id,
                    "level": kyc.level.value,
                    "status": kyc.status.value,
                    "max_deposit_limit": kyc.max_deposit_limit or 0.0,
                    "max_withdrawal_limit": kyc.max_withdrawal_limit or 0.0,
                    "message": "Document verification completed",
                }
            except Exception as e:
                logger.error(f"VerifyDocument error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def VerifyLiveness(self, request, context):
        async with AsyncSessionLocal() as db:
            try:
                service = KYCService(db)

                kyc = await service.verification.verify_liveness(
                    user_id=request.user_id,
                    liveness_verified=getattr(request, "liveness_verified", False),
                    liveness_score=getattr(request, "liveness_score", 0.0),
                )

                return {
                    "id": str(kyc.id),
                    "user_id": kyc.user_id,
                    "level": kyc.level.value,
                    "status": kyc.status.value,
                    "max_deposit_limit": kyc.max_deposit_limit or 0.0,
                    "max_withdrawal_limit": kyc.max_withdrawal_limit or 0.0,
                    "message": "Liveness check completed",
                }
            except Exception as e:
                logger.error(f"VerifyLiveness error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def VerifyAddress(self, request, context):
        async with AsyncSessionLocal() as db:
            try:
                service = KYCService(db)

                kyc = await service.verification.verify_address(
                    user_id=request.user_id,
                    address_verified=getattr(request, "address_verified", False),
                )

                return {
                    "id": str(kyc.id),
                    "user_id": kyc.user_id,
                    "level": kyc.level.value,
                    "status": kyc.status.value,
                    "max_deposit_limit": kyc.max_deposit_limit or 0.0,
                    "max_withdrawal_limit": kyc.max_withdrawal_limit or 0.0,
                    "message": "Address verification completed",
                }
            except Exception as e:
                logger.error(f"VerifyAddress error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def SubmitAddress(self, request, context):
        async with AsyncSessionLocal() as db:
            try:
                service = KYCService(db)

                kyc = await service.verification.submit_address_verification(
                    user_id=request.user_id,
                    address_line_1=request.address_line_1,
                    address_line_2=getattr(request, "address_line_2", None),
                    city=request.city,
                    state=getattr(request, "state", None),
                    postal_code=request.postal_code,
                    country_of_residence=request.country_of_residence,
                )

                return {
                    "id": str(kyc.id),
                    "user_id": kyc.user_id,
                    "level": kyc.level.value,
                    "status": kyc.status.value,
                    "max_deposit_limit": kyc.max_deposit_limit or 0.0,
                    "max_withdrawal_limit": kyc.max_withdrawal_limit or 0.0,
                    "message": "Address submitted successfully",
                }
            except Exception as e:
                logger.error(f"SubmitAddress error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}
