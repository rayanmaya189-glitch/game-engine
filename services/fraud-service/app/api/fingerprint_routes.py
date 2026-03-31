"""Fingerprint registration and multi-account check functions"""

from typing import Dict
from sqlalchemy.ext.asyncio import AsyncSession

from app.database import get_db
from app.models.schemas import DeviceFingerprint
from app.services.multi_account_detector import MultiAccountDetector


async def register_fingerprint(fingerprint: DeviceFingerprint, db: AsyncSession) -> Dict:
    """Register a device fingerprint"""
    await MultiAccountDetector.register_fingerprint(db, fingerprint)
    return {"status": "registered", "user_id": fingerprint.user_id}


async def check_multi_account(user_id: str, db: AsyncSession) -> Dict:
    """Check for multi-account patterns"""
    related = await MultiAccountDetector.check_multi_account(db, user_id)
    return {
        "user_id": user_id,
        "related_accounts": related,
        "is_multi_account": len(related) > 0
    }
