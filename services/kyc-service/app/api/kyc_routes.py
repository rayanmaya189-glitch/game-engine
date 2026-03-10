from typing import Optional
from fastapi import APIRouter, Depends, HTTPException, status
from pydantic import BaseModel, Field
from sqlalchemy.ext.asyncio import AsyncSession

from app.database import get_db
from app.services.kyc_service import KYCService
from app.models.kyc import DocumentType, VerificationProvider

router = APIRouter()


# Request/Response Models
class KYCSubmissionRequest(BaseModel):
    user_id: str = Field(..., description="User ID")
    document_type: DocumentType = Field(..., description="Type of document")
    document_number: str = Field(..., description="Document number")
    document_expiry: str = Field(..., description="Document expiry date (ISO format)")
    front_image_url: str = Field(..., description="URL to front image")
    back_image_url: Optional[str] = Field(None, description="URL to back image")
    selfie_image_url: Optional[str] = Field(None, description="URL to selfie image")


class AddressVerificationRequest(BaseModel):
    user_id: str = Field(..., description="User ID")
    address_line_1: str = Field(..., description="Address line 1")
    address_line_2: Optional[str] = Field(None, description="Address line 2")
    city: str = Field(..., description="City")
    state: Optional[str] = Field(None, description="State/Province")
    postal_code: str = Field(..., description="Postal code")
    country_of_residence: str = Field(..., description="Country of residence (ISO 3)")


class VerificationResultRequest(BaseModel):
    user_id: str = Field(..., description="User ID")
    provider: VerificationProvider = Field(..., description="Verification provider")
    provider_reference: str = Field(..., description="Provider reference")
    document_verified: bool = Field(..., description="Whether document was verified")
    reason: Optional[str] = Field(None, description="Reason if not verified")


class LivenessCheckRequest(BaseModel):
    user_id: str = Field(..., description="User ID")
    liveness_verified: bool = Field(..., description="Whether liveness was verified")
    liveness_score: float = Field(..., ge=0.0, le=1.0, description="Liveness confidence score")


class ManualReviewRequest(BaseModel):
    kyc_id: str = Field(..., description="KYC record ID")
    reviewer_notes: Optional[str] = Field(None, description="Reviewer notes")
    approved: bool = Field(..., description="Whether KYC is approved")


class KYCLimitCheckRequest(BaseModel):
    user_id: str = Field(..., description="User ID")
    amount: float = Field(..., gt=0, description="Transaction amount")
    transaction_type: str = Field(..., description="Transaction type (deposit/withdrawal)")


class KYCResponse(BaseModel):
    id: str
    user_id: str
    level: str
    status: str
    max_deposit_limit: float
    max_withdrawal_limit: float
    message: Optional[str] = None


# Routes
@router.post("/", response_model=KYCResponse)
async def create_kyc_record(
    user_id: str,
    db: AsyncSession = Depends(get_db)
):
    """Create a new KYC record for a user"""
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
    """Get KYC status for a user"""
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


@router.post("/documents", response_model=KYCResponse)
async def submit_document(
    request: KYCSubmissionRequest,
    db: AsyncSession = Depends(get_db)
):
    """Submit document for KYC verification"""
    from datetime import datetime
    
    service = KYCService(db)
    
    try:
        document_expiry = datetime.fromisoformat(request.document_expiry.replace("Z", "+00:00"))
    except ValueError:
        raise HTTPException(
            status_code=status.HTTP_400_BAD_REQUEST,
            detail="Invalid date format. Use ISO format."
        )
    
    kyc = await service.submit_document(
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


@router.post("/address", response_model=KYCResponse)
async def submit_address(
    request: AddressVerificationRequest,
    db: AsyncSession = Depends(get_db)
):
    """Submit address for verification"""
    service = KYCService(db)
    
    kyc = await service.submit_address_verification(
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


@router.post("/verify", response_model=KYCResponse)
async def verify_document(
    request: VerificationResultRequest,
    db: AsyncSession = Depends(get_db)
):
    """Process document verification result from provider"""
    service = KYCService(db)
    
    kyc = await service.verify_document(
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
    """Submit liveness check result"""
    service = KYCService(db)
    
    kyc = await service.verify_liveness(
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
    """Process address verification result"""
    service = KYCService(db)
    
    kyc = await service.verify_address(
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


@router.post("/review", response_model=KYCResponse)
async def manual_review(
    request: ManualReviewRequest,
    db: AsyncSession = Depends(get_db)
):
    """Process manual review decision"""
    service = KYCService(db)
    
    if request.approved:
        kyc = await service.approve_kyc(request.kyc_id, request.reviewer_notes)
        message = "KYC approved"
    else:
        kyc = await service.reject_kyc(request.kyc_id, request.reviewer_notes or "Rejected during manual review")
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
    """Get pending review queue"""
    service = KYCService(db)
    records = await service.get_pending_review_queue(limit)
    
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


@router.post("/limits/check")
async def check_limits(
    request: KYCLimitCheckRequest,
    db: AsyncSession = Depends(get_db)
):
    """Check if transaction is within KYC limits"""
    service = KYCService(db)
    
    within_limits = await service.check_kyc_limits(
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
