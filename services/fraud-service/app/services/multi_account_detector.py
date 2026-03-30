"""Multi-account detection via device fingerprinting and IP correlation"""

from typing import List
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select, delete

from app.models import DeviceFingerprintRecord, IpAccountRecord
from app.models.schemas import DeviceFingerprint


class MultiAccountDetector:
    """Detect multiple accounts from same device/IP"""

    @staticmethod
    async def register_fingerprint(db: AsyncSession, fingerprint: DeviceFingerprint):
        """Register a device fingerprint"""
        await db.execute(
            delete(DeviceFingerprintRecord).where(DeviceFingerprintRecord.user_id == fingerprint.user_id)
        )
        db.add(DeviceFingerprintRecord(
            user_id=fingerprint.user_id,
            canvas_hash=fingerprint.canvas_hash,
            webgl_hash=fingerprint.webgl_hash,
            audio_hash=fingerprint.audio_hash,
            fonts=fingerprint.fonts,
            screen_resolution=fingerprint.screen_resolution,
            ip_address=fingerprint.ip_address,
            user_agent=fingerprint.user_agent,
        ))

        if fingerprint.ip_address:
            result = await db.execute(
                select(IpAccountRecord).where(
                    IpAccountRecord.ip_address == fingerprint.ip_address,
                    IpAccountRecord.user_id == fingerprint.user_id,
                )
            )
            if not result.scalar_one_or_none():
                db.add(IpAccountRecord(
                    ip_address=fingerprint.ip_address,
                    user_id=fingerprint.user_id,
                ))
        await db.commit()

    @staticmethod
    async def check_multi_account(db: AsyncSession, user_id: str) -> List[str]:
        """Check if user has multiple accounts"""
        result = await db.execute(
            select(DeviceFingerprintRecord).where(DeviceFingerprintRecord.user_id == user_id)
        )
        fp = result.scalar_one_or_none()
        if not fp:
            return []

        related_accounts = set()

        # Check by canvas hash
        if fp.canvas_hash:
            result = await db.execute(
                select(DeviceFingerprintRecord).where(
                    DeviceFingerprintRecord.canvas_hash == fp.canvas_hash,
                    DeviceFingerprintRecord.user_id != user_id,
                )
            )
            for r in result.scalars().all():
                related_accounts.add(r.user_id)

        # Check by webgl hash
        if fp.webgl_hash:
            result = await db.execute(
                select(DeviceFingerprintRecord).where(
                    DeviceFingerprintRecord.webgl_hash == fp.webgl_hash,
                    DeviceFingerprintRecord.user_id != user_id,
                )
            )
            for r in result.scalars().all():
                related_accounts.add(r.user_id)

        # Check by IP
        if fp.ip_address:
            result = await db.execute(
                select(IpAccountRecord).where(
                    IpAccountRecord.ip_address == fp.ip_address,
                    IpAccountRecord.user_id != user_id,
                )
            )
            for r in result.scalars().all():
                related_accounts.add(r.user_id)

        return list(related_accounts)

    @staticmethod
    def analyze_email_patterns(email: str) -> bool:
        """Detect email variation patterns (john+1@, j.o.h.n@)"""
        if '+' in email.split('@')[0]:
            return True

        username = email.split('@')[0]
        if username.count('.') > 2:
            return True

        return False
