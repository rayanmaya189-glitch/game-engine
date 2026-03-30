"""Fingerprint registration and multi-account check endpoints"""

from fastapi import APIRouter
from typing import Dict

from app.models.schemas import DeviceFingerprint
from app.services.multi_account_detector import MultiAccountDetector

router = APIRouter(prefix="/fingerprint", tags=["fingerprint"])


@router.post("/register", response_model=Dict)
async def register_fingerprint(fingerprint: DeviceFingerprint):
    """Register a device fingerprint"""
    MultiAccountDetector.register_fingerprint(fingerprint)
    return {"status": "registered", "user_id": fingerprint.user_id}


@router.get("/check/{user_id}", response_model=Dict)
async def check_multi_account(user_id: str):
    """Check for multi-account patterns"""
    related = MultiAccountDetector.check_multi_account(user_id)
    return {
        "user_id": user_id,
        "related_accounts": related,
        "is_multi_account": len(related) > 0
    }
