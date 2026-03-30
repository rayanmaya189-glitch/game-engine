"""Multi-account detection via device fingerprinting and IP correlation"""

from typing import List

from app.models.schemas import DeviceFingerprint, device_fingerprints, ip_accounts


class MultiAccountDetector:
    """Detect multiple accounts from same device/IP"""

    @staticmethod
    def register_fingerprint(fingerprint: DeviceFingerprint):
        """Register a device fingerprint"""
        device_fingerprints[fingerprint.user_id] = fingerprint

        if fingerprint.ip_address:
            if fingerprint.ip_address not in ip_accounts:
                ip_accounts[fingerprint.ip_address] = []
            if fingerprint.user_id not in ip_accounts[fingerprint.ip_address]:
                ip_accounts[fingerprint.ip_address].append(fingerprint.user_id)

    @staticmethod
    def check_multi_account(user_id: str) -> List[str]:
        """Check if user has multiple accounts"""
        if user_id not in device_fingerprints:
            return []

        fp = device_fingerprints[user_id]
        related_accounts = []

        for uid, other_fp in device_fingerprints.items():
            if uid == user_id:
                continue

            if fp.canvas_hash and other_fp.canvas_hash:
                if fp.canvas_hash == other_fp.canvas_hash:
                    related_accounts.append(uid)

            if fp.webgl_hash and other_fp.webgl_hash:
                if fp.webgl_hash == other_fp.webgl_hash:
                    related_accounts.append(uid)

        if fp.ip_address and fp.ip_address in ip_accounts:
            related_accounts.extend([
                uid for uid in ip_accounts[fp.ip_address]
                if uid != user_id
            ])

        return related_accounts

    @staticmethod
    def analyze_email_patterns(email: str) -> bool:
        """Detect email variation patterns (john+1@, j.o.h.n@)"""
        if '+' in email.split('@')[0]:
            return True

        username = email.split('@')[0]
        if username.count('.') > 2:
            return True

        return False
