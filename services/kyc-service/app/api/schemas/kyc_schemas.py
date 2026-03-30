from typing import Optional
from pydantic import BaseModel, Field

from app.models.kyc import DocumentType, VerificationProvider


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
