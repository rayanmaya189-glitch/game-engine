"""
gRPC server for KYC Service.
Exposes identity verification and KYC management methods.
"""

import logging
from datetime import datetime
from concurrent import futures

import grpc
from sqlalchemy import select

from app.database import AsyncSessionLocal
from app.models.kyc import KYCRecord, KYCStatus, KYCLevel
from app.services.kyc_service import KYCService

logger = logging.getLogger(__name__)


class KYCServiceServicer:

    async def CreateKYCRecord(self, request, context):
        async with AsyncSessionLocal() as db:
            try:
                service = KYCService(db)
                kyc = await service.create_kyc_record(request.user_id)
                return {
                    "id": str(kyc.id),
                    "user_id": kyc.user_id,
                    "level": kyc.level.value,
                    "status": kyc.status.value,
                    "max_deposit_limit": kyc.max_deposit_limit or 0.0,
                    "max_withdrawal_limit": kyc.max_withdrawal_limit or 0.0,
                    "message": "KYC record created",
                }
            except Exception as e:
                logger.error(f"CreateKYCRecord error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def GetKYCStatus(self, request, context):
        async with AsyncSessionLocal() as db:
            try:
                service = KYCService(db)
                kyc = await service.get_kyc_by_user_id(request.user_id)
                if not kyc:
                    context.set_code(grpc.StatusCode.NOT_FOUND)
                    context.set_details("KYC record not found")
                    return {}
                return {
                    "id": str(kyc.id),
                    "user_id": kyc.user_id,
                    "level": kyc.level.value,
                    "status": kyc.status.value,
                    "max_deposit_limit": kyc.max_deposit_limit or 0.0,
                    "max_withdrawal_limit": kyc.max_withdrawal_limit or 0.0,
                }
            except Exception as e:
                logger.error(f"GetKYCStatus error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def SubmitDocument(self, request, context):
        async with AsyncSessionLocal() as db:
            try:
                from app.models.kyc import DocumentType
                service = KYCService(db)

                try:
                    document_expiry = datetime.fromisoformat(
                        request.document_expiry.replace("Z", "+00:00")
                    )
                except ValueError:
                    context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
                    context.set_details("Invalid date format. Use ISO format.")
                    return {}

                kyc = await service.document.submit_document(
                    user_id=request.user_id,
                    document_type=DocumentType(request.document_type),
                    document_number=request.document_number,
                    document_expiry=document_expiry,
                    front_image_url=request.front_image_url,
                    back_image_url=getattr(request, "back_image_url", None),
                    selfie_image_url=getattr(request, "selfie_image_url", None),
                )

                return {
                    "id": str(kyc.id),
                    "user_id": kyc.user_id,
                    "level": kyc.level.value,
                    "status": kyc.status.value,
                    "max_deposit_limit": kyc.max_deposit_limit or 0.0,
                    "max_withdrawal_limit": kyc.max_withdrawal_limit or 0.0,
                    "message": "Document submitted successfully",
                }
            except Exception as e:
                logger.error(f"SubmitDocument error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

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

    async def ManualReview(self, request, context):
        async with AsyncSessionLocal() as db:
            try:
                service = KYCService(db)
                approved = getattr(request, "approved", False)
                reviewer_notes = getattr(request, "reviewer_notes", None)

                if approved:
                    kyc = await service.review.approve_kyc(request.user_id, reviewer_notes)
                    message = "KYC approved"
                else:
                    kyc = await service.review.reject_kyc(
                        request.user_id,
                        reviewer_notes or "Rejected during manual review"
                    )
                    message = "KYC rejected"

                return {
                    "id": str(kyc.id),
                    "user_id": kyc.user_id,
                    "level": kyc.level.value,
                    "status": kyc.status.value,
                    "max_deposit_limit": kyc.max_deposit_limit or 0.0,
                    "max_withdrawal_limit": kyc.max_withdrawal_limit or 0.0,
                    "message": message,
                }
            except Exception as e:
                logger.error(f"ManualReview error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def GetPendingReviewQueue(self, request, context):
        async with AsyncSessionLocal() as db:
            try:
                service = KYCService(db)
                limit = getattr(request, "limit", 50)
                records = await service.review.get_pending_review_queue(limit)

                return {
                    "count": len(records),
                    "items": [
                        {
                            "id": str(r.id),
                            "user_id": r.user_id,
                            "status": r.status.value,
                            "submitted_at": r.submitted_at.isoformat() if r.submitted_at else None,
                        }
                        for r in records
                    ],
                }
            except Exception as e:
                logger.error(f"GetPendingReviewQueue error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def CheckLimits(self, request, context):
        async with AsyncSessionLocal() as db:
            try:
                service = KYCService(db)

                within_limits = await service.review.check_kyc_limits(
                    user_id=request.user_id,
                    amount=getattr(request, "amount", 0),
                    transaction_type=getattr(request, "transaction_type", "deposit"),
                )

                return {
                    "user_id": request.user_id,
                    "amount": getattr(request, "amount", 0),
                    "transaction_type": getattr(request, "transaction_type", "deposit"),
                    "within_limits": within_limits,
                }
            except Exception as e:
                logger.error(f"CheckLimits error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}


async def serve_grpc(port: int) -> grpc.aio.Server:
    server = grpc.aio.server(futures.ThreadPoolExecutor(max_workers=10))
    server.add_insecure_port(f"[::]:{port}")
    await server.start()
    logger.info(f"KYC gRPC server listening on port {port}")
    return server
