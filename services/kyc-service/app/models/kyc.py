from datetime import datetime
from enum import Enum
from typing import Optional
from sqlalchemy import Column, String, DateTime, Float, Text, Boolean, Enum as SQLEnum, Index
from sqlalchemy.dialects.postgresql import UUID
import uuid

from app.database import Base


class KYCLevel(str, Enum):
    NONE = "NONE"
    BASIC = "BASIC"
    INTERMEDIATE = "INTERMEDIATE"
    FULL = "FULL"


class KYCStatus(str, Enum):
    NOT_STARTED = "NOT_STARTED"
    PENDING = "PENDING"
    IN_PROGRESS = "IN_PROGRESS"
    VERIFIED = "VERIFIED"
    REJECTED = "REJECTED"
    EXPIRED = "EXPIRED"
    NEEDS_REVIEW = "NEEDS_REVIEW"


class DocumentType(str, Enum):
    PASSPORT = "PASSPORT"
    DRIVERS_LICENSE = "DRIVERS_LICENSE"
    NATIONAL_ID = "NATIONAL_ID"
    RESIDENCE_PERMIT = "RESIDENCE_PERMIT"


class VerificationProvider(str, Enum):
    ONFIDO = "ONFIDO"
    JUMIO = "JUMIO"
    SUMSUB = "SUMSUB"
    INTERNAL = "INTERNAL"


class KYCRecord(Base):
    __tablename__ = "kyc_records"

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    user_id = Column(String(255), nullable=False, unique=True, index=True)
    
    # KYC Level
    level = Column(SQLEnum(KYCLevel), nullable=False, default=KYCLevel.NONE)
    status = Column(SQLEnum(KYCStatus), nullable=False, default=KYCStatus.NOT_STARTED)
    
    # Personal Information
    first_name = Column(String(255))
    last_name = Column(String(255))
    date_of_birth = Column(DateTime)
    nationality = Column(String(3))  # ISO 3166-1 alpha-3
    country = Column(String(3))  # ISO 3166-1 alpha-3
    
    # Address
    address_line_1 = Column(String(500))
    address_line_2 = Column(String(500))
    city = Column(String(255))
    state = Column(String(255))
    postal_code = Column(String(20))
    country_of_residence = Column(String(3))
    
    # Document Verification
    document_type = Column(SQLEnum(DocumentType))
    document_number = Column(String(255))
    document_expiry = Column(DateTime)
    document_front_url = Column(Text)
    document_back_url = Column(Text)
    document_selfie_url = Column(Text)
    
    # Verification Details
    provider = Column(SQLEnum(VerificationProvider))
    provider_reference = Column(String(255))
    provider_verification_id = Column(String(255))
    
    # Liveness Check
    liveness_verified = Column(Boolean, default=False)
    liveness_score = Column(Float)
    liveness_check_at = Column(DateTime)
    
    # Address Verification
    address_verified = Column(Boolean, default=False)
    address_verification_at = Column(DateTime)
    
    # Review Process
    requires_manual_review = Column(Boolean, default=False)
    reviewed_by = Column(String(255))
    review_notes = Column(Text)
    reviewed_at = Column(DateTime)
    
    # Limits based on KYC level
    max_deposit_limit = Column(Float, default=0.0)
    max_withdrawal_limit = Column(Float, default=0.0)
    
    # Timestamps
    submitted_at = Column(DateTime)
    verified_at = Column(DateTime)
    expires_at = Column(DateTime)
    created_at = Column(DateTime, default=datetime.utcnow)
    updated_at = Column(DateTime, default=datetime.utcnow, onupdate=datetime.utcnow)

    __table_args__ = (
        Index("idx_kyc_status", "status"),
        Index("idx_kyc_level", "level"),
        Index("idx_kyc_user_status", "user_id", "status"),
    )


class KYCDocument(Base):
    __tablename__ = "kyc_documents"

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    kyc_record_id = Column(UUID(as_uuid=True), nullable=False, index=True)
    user_id = Column(String(255), nullable=False, index=True)
    
    document_type = Column(SQLEnum(DocumentType), nullable=False)
    document_number = Column(String(255))
    
    # Document Files
    front_image_url = Column(Text)
    back_image_url = Column(Text)
    selfie_image_url = Column(Text)
    video_url = Column(Text)
    
    # Verification Results
    document_verified = Column(Boolean, default=False)
    document_match_score = Column(Float)
    ocr_data = Column(Text)  # JSON string
    
    # Liveness
    liveness_verified = Column(Boolean, default=False)
    liveness_score = Column(Float)
    
    # Provider info
    provider = Column(SQLEnum(VerificationProvider))
    provider_reference = Column(String(255))
    provider_response = Column(Text)  # JSON string
    
    # Status
    is_primary = Column(Boolean, default=False)
    verified_at = Column(DateTime)
    expires_at = Column(DateTime)
    created_at = Column(DateTime, default=datetime.utcnow)
    updated_at = Column(DateTime, default=datetime.utcnow, onupdate=datetime.utcnow)


class KYCQueueItem(Base):
    __tablename__ = "kyc_queue"

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    kyc_record_id = Column(UUID(as_uuid=True), nullable=False, index=True)
    user_id = Column(String(255), nullable=False, index=True)
    
    priority = Column(String(20), default="NORMAL")  # HIGH, NORMAL, LOW
    reason = Column(String(500))
    assigned_to = Column(String(255))
    
    status = Column(String(20), default="PENDING")  # PENDING, IN_PROGRESS, COMPLETED, CANCELLED
    started_at = Column(DateTime)
    completed_at = Column(DateTime)
    
    created_at = Column(DateTime, default=datetime.utcnow)
    updated_at = Column(DateTime, default=datetime.utcnow, onupdate=datetime.utcnow)


class KYCAuditLog(Base):
    __tablename__ = "kyc_audit_log"

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    kyc_record_id = Column(UUID(as_uuid=True), nullable=False, index=True)
    user_id = Column(String(255), nullable=False, index=True)
    
    action = Column(String(100), nullable=False)
    previous_value = Column(Text)
    new_value = Column(Text)
    reason = Column(Text)
    
    ip_address = Column(String(45))
    user_agent = Column(String(500))
    
    created_at = Column(DateTime, default=datetime.utcnow)
