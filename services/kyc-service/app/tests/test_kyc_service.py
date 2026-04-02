import pytest
from unittest.mock import AsyncMock, MagicMock, patch
from datetime import datetime
from uuid import uuid4
from app.services.kyc_service import KYCService
from app.services.kyc_document_service import KYCDocumentService
from app.services.kyc_review_service import KYCReviewService
from app.services.kyc_verification_service import KYCVerificationService
from app.models.kyc import KYCRecord, KYCLevel, KYCStatus, DocumentType, VerificationProvider


@pytest.fixture
def mock_db():
    return AsyncMock()


@pytest.fixture
def kyc_service(mock_db):
    return KYCService(mock_db)


class TestCreateKYCRecord:
    @pytest.mark.asyncio
    async def test_creates_new_record(self, kyc_service):
        mock_record = MagicMock()
        mock_record.id = uuid4()
        mock_record.user_id = "user-1"
        mock_record.level = KYCLevel.NONE
        mock_record.status = KYCStatus.NOT_STARTED

        kyc_service.document.create_kyc_record = AsyncMock(return_value=mock_record)

        result = await kyc_service.create_kyc_record("user-1")

        assert result.user_id == "user-1"
        assert result.level == KYCLevel.NONE
        kyc_service.document.create_kyc_record.assert_called_once_with("user-1")

    @pytest.mark.asyncio
    async def test_get_existing_record(self, kyc_service):
        mock_record = MagicMock()
        mock_record.user_id = "user-1"
        kyc_service.document.get_kyc_by_user_id = AsyncMock(return_value=mock_record)

        result = await kyc_service.get_kyc_by_user_id("user-1")

        assert result is not None
        assert result.user_id == "user-1"

    @pytest.mark.asyncio
    async def test_get_nonexistent_record(self, kyc_service):
        kyc_service.document.get_kyc_by_user_id = AsyncMock(return_value=None)

        result = await kyc_service.get_kyc_by_user_id("nonexistent")

        assert result is None


class TestVerifyDocument:
    @pytest.mark.asyncio
    async def test_verify_successful_document(self):
        mock_db = AsyncMock()
        audit = MagicMock()
        audit.log_audit = AsyncMock()
        audit.update_kyc_level = AsyncMock()

        service = KYCVerificationService(mock_db, audit)

        mock_kyc = MagicMock()
        mock_kyc.status = KYCStatus.IN_PROGRESS

        mock_result = MagicMock()
        mock_result.scalar_one_or_none.return_value = mock_kyc
        mock_db.execute.return_value = mock_result

        result = await service.verify_document(
            "user-1",
            VerificationProvider.ONFIDO,
            "ref-123",
            {"document_verified": True},
        )

        assert result.status == KYCStatus.VERIFIED

    @pytest.mark.asyncio
    async def test_verify_failed_document(self):
        mock_db = AsyncMock()
        audit = MagicMock()
        audit.log_audit = AsyncMock()

        service = KYCVerificationService(mock_db, audit)

        mock_kyc = MagicMock()
        mock_kyc.status = KYCStatus.IN_PROGRESS

        mock_result = MagicMock()
        mock_result.scalar_one_or_none.return_value = mock_kyc
        mock_db.execute.return_value = mock_result

        result = await service.verify_document(
            "user-1",
            VerificationProvider.ONFIDO,
            "ref-123",
            {"document_verified": False, "reason": "Blurry image"},
        )

        assert result.status == KYCStatus.REJECTED

    @pytest.mark.asyncio
    async def test_verify_nonexistent_record(self):
        mock_db = AsyncMock()
        audit = MagicMock()
        service = KYCVerificationService(mock_db, audit)

        mock_result = MagicMock()
        mock_result.scalar_one_or_none.return_value = None
        mock_db.execute.return_value = mock_result

        with pytest.raises(ValueError, match="KYC record not found"):
            await service.verify_document(
                "user-99",
                VerificationProvider.ONFIDO,
                "ref-123",
                {"document_verified": True},
            )


class TestApproveKYC:
    @pytest.mark.asyncio
    async def test_approve_verified_record(self):
        mock_db = AsyncMock()
        audit = MagicMock()
        audit.log_audit = AsyncMock()
        audit.update_kyc_level = AsyncMock()

        service = KYCReviewService(mock_db, audit)

        mock_kyc = MagicMock()
        mock_kyc.level = KYCLevel.BASIC

        mock_result = MagicMock()
        mock_result.scalar_one_or_none.return_value = mock_kyc
        mock_db.execute.return_value = mock_result

        result = await service.approve_kyc("user-1", reviewer_notes="Looks good")

        assert result.status == KYCStatus.VERIFIED
        assert result.review_notes == "Looks good"

    @pytest.mark.asyncio
    async def test_approve_nonexistent_record(self):
        mock_db = AsyncMock()
        audit = MagicMock()
        service = KYCReviewService(mock_db, audit)

        mock_result = MagicMock()
        mock_result.scalar_one_or_none.return_value = None
        mock_db.execute.return_value = mock_result

        with pytest.raises(ValueError):
            await service.approve_kyc("user-99")


class TestRejectKYC:
    @pytest.mark.asyncio
    async def test_reject_record(self):
        mock_db = AsyncMock()
        audit = MagicMock()
        audit.log_audit = AsyncMock()

        service = KYCReviewService(mock_db, audit)

        mock_kyc = MagicMock()
        mock_kyc.status = KYCStatus.IN_PROGRESS

        mock_result = MagicMock()
        mock_result.scalar_one_or_none.return_value = mock_kyc
        mock_db.execute.return_value = mock_result

        result = await service.reject_kyc("user-1", "Document expired")

        assert result.status == KYCStatus.REJECTED
        assert result.review_notes == "Document expired"

    @pytest.mark.asyncio
    async def test_reject_nonexistent_record(self):
        mock_db = AsyncMock()
        audit = MagicMock()
        service = KYCReviewService(mock_db, audit)

        mock_result = MagicMock()
        mock_result.scalar_one_or_none.return_value = None
        mock_db.execute.return_value = mock_result

        with pytest.raises(ValueError):
            await service.reject_kyc("user-99", "reason")


class TestCheckKYCLimits:
    @pytest.mark.asyncio
    async def test_deposit_within_limit(self):
        mock_db = AsyncMock()
        audit = MagicMock()
        service = KYCReviewService(mock_db, audit)

        mock_kyc = MagicMock()
        mock_kyc.max_deposit_limit = 10000.0
        mock_kyc.max_withdrawal_limit = 5000.0

        mock_result = MagicMock()
        mock_result.scalar_one_or_none.return_value = mock_kyc
        mock_db.execute.return_value = mock_result

        allowed = await service.check_kyc_limits("user-1", 5000.0, "deposit")

        assert allowed is True

    @pytest.mark.asyncio
    async def test_deposit_exceeds_limit(self):
        mock_db = AsyncMock()
        audit = MagicMock()
        service = KYCReviewService(mock_db, audit)

        mock_kyc = MagicMock()
        mock_kyc.max_deposit_limit = 1000.0
        mock_kyc.max_withdrawal_limit = 500.0

        mock_result = MagicMock()
        mock_result.scalar_one_or_none.return_value = mock_kyc
        mock_db.execute.return_value = mock_result

        allowed = await service.check_kyc_limits("user-1", 5000.0, "deposit")

        assert allowed is False

    @pytest.mark.asyncio
    async def test_no_kyc_record(self):
        mock_db = AsyncMock()
        audit = MagicMock()
        service = KYCReviewService(mock_db, audit)

        mock_result = MagicMock()
        mock_result.scalar_one_or_none.return_value = None
        mock_db.execute.return_value = mock_result

        allowed = await service.check_kyc_limits("user-1", 100.0, "deposit")

        assert allowed is False
