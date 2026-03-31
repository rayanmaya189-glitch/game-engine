"""
gRPC server for Fraud Detection Service.
Exposes fraud detection methods including multi-account, bot, and collusion detection.
"""

import logging
import random
import json
from datetime import datetime, timedelta
from concurrent import futures

import grpc
from sqlalchemy import select

from app.database import async_session_factory
from app.models import (
    DeviceFingerprintRecord,
    FraudScoreRecord,
    FraudAlertRecord,
    UserRiskProfileRecord,
)
from app.models.schemas import (
    DeviceFingerprint,
    BotDetectionResult,
    FraudScore,
    CollusionSignal,
)
from app.repositories import (
    DeviceFingerprintRepository,
    FraudAlertRepository,
    UserRiskProfileRepository,
)
from app.services.fraud_scorer import FraudScorer
from app.services.bot_detector import BotDetector
from app.services.collusion_detector import CollusionDetector
from app.services.multi_account_detector import MultiAccountDetector

logger = logging.getLogger(__name__)

HIGH_AMOUNT = 5000.0
VERY_HIGH_AMOUNT = 10000.0
BET_AMOUNT = 1000.0
RAPID_BET_SECONDS = 2


class FraudServiceServicer:

    async def CheckRisk(self, request, context):
        async with async_session_factory() as db:
            try:
                risk_score = 0.0
                risk_factors = []
                amount = getattr(request, "amount", 0)

                if amount > HIGH_AMOUNT:
                    risk_score += 0.2
                    risk_factors.append("High transaction amount")
                if amount > VERY_HIGH_AMOUNT:
                    risk_score += 0.2
                    risk_factors.append("Very high transaction amount")

                device_fingerprint = getattr(request, "device_fingerprint", None)
                if device_fingerprint:
                    fp = await DeviceFingerprintRepository.get_by_user(db, request.user_id)
                    if fp and fp.canvas_hash != device_fingerprint:
                        risk_score += 0.3
                        risk_factors.append("New device detected")

                ip_address = getattr(request, "ip_address", None)
                if ip_address:
                    if ip_address.startswith("10.") or ip_address.startswith("192."):
                        risk_score += 0.1
                        risk_factors.append("Internal IP detected")

                profile = await UserRiskProfileRepository.get_or_create(db, request.user_id)
                profile.transaction_count = (profile.transaction_count or 0) + 1
                profile.last_activity = datetime.utcnow().isoformat()
                existing_score = profile.risk_score or 0.0
                profile.risk_score = (existing_score + risk_score) / 2
                risk_score += (profile.risk_score or 0.0) * 0.3

                is_blocked = risk_score > 0.8
                allowed = risk_score < 0.7 and not is_blocked

                await UserRiskProfileRepository.save(db, profile)

                return {
                    "user_id": request.user_id,
                    "risk_score": min(risk_score, 1.0),
                    "allowed": allowed,
                    "is_blocked": is_blocked,
                    "requires_review": risk_score >= 0.5,
                    "risk_factors": risk_factors,
                    "transaction_count": profile.transaction_count,
                }
            except Exception as e:
                logger.error(f"CheckRisk error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def CheckBetRisk(self, request, context):
        async with async_session_factory() as db:
            try:
                risk_score = 0.0
                risk_factors = []
                amount = getattr(request, "amount", 0)

                if amount > BET_AMOUNT:
                    risk_score += 0.15
                    risk_factors.append("Unusually high bet amount")

                profile = await UserRiskProfileRepository.get_or_create(db, request.user_id)

                if profile.last_bet_time:
                    try:
                        last_bet = datetime.fromisoformat(profile.last_bet_time)
                        time_diff = (datetime.utcnow() - last_bet).total_seconds()
                        if time_diff < RAPID_BET_SECONDS:
                            risk_score += 0.25
                            risk_factors.append("Rapid betting pattern detected")
                    except (ValueError, TypeError):
                        pass

                profile.last_bet_time = datetime.utcnow().isoformat()
                await UserRiskProfileRepository.save(db, profile)

                return {
                    "user_id": request.user_id,
                    "game_id": getattr(request, "game_id", ""),
                    "risk_score": min(risk_score, 1.0),
                    "allowed": risk_score < 0.6,
                    "requires_review": risk_score >= 0.4,
                    "risk_factors": risk_factors,
                }
            except Exception as e:
                logger.error(f"CheckBetRisk error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def CheckAccountActivity(self, request, context):
        async with async_session_factory() as db:
            try:
                risk_score = 0.0
                risk_factors = []
                activity_type = getattr(request, "activity_type", "")

                high_risk_activities = ["withdrawal", "password_change", "email_change", "2fa_disable"]
                if activity_type in high_risk_activities:
                    risk_score += 0.3
                    risk_factors.append(f"High-risk activity: {activity_type}")

                device_fingerprint = getattr(request, "device_fingerprint", None)
                if device_fingerprint:
                    fp = await DeviceFingerprintRepository.get_by_user(db, request.user_id)
                    if fp and fp.canvas_hash != device_fingerprint:
                        risk_score += 0.4
                        risk_factors.append("Account access from new device")

                ip_address = getattr(request, "ip_address", None)
                if ip_address:
                    profile = await UserRiskProfileRepository.get_or_create(db, request.user_id)
                    if profile.last_ip and profile.last_ip != ip_address:
                        risk_score += 0.2
                        risk_factors.append("IP address changed")

                return {
                    "user_id": request.user_id,
                    "activity_type": activity_type,
                    "risk_score": min(risk_score, 1.0),
                    "allowed": risk_score < 0.7,
                    "requires_review": risk_score >= 0.5,
                    "risk_factors": risk_factors,
                    "recommendation": "block" if risk_score > 0.8 else "allow" if risk_score < 0.5 else "review",
                }
            except Exception as e:
                logger.error(f"CheckAccountActivity error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def GetUserRisk(self, request, context):
        async with async_session_factory() as db:
            try:
                profile = await UserRiskProfileRepository.get_or_create(db, request.user_id)
                profile.is_blocked = (profile.risk_score or 0) > 0.8
                await UserRiskProfileRepository.save(db, profile)

                flags = json.loads(profile.flags) if isinstance(profile.flags, str) else (profile.flags or [])

                return {
                    "user_id": profile.user_id,
                    "risk_score": profile.risk_score or 0.0,
                    "is_blocked": profile.is_blocked,
                    "flags": flags,
                    "transaction_count": profile.transaction_count or 0,
                }
            except Exception as e:
                logger.error(f"GetUserRisk error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def CreateAlert(self, request, context):
        async with async_session_factory() as db:
            try:
                alert_id = f"ALERT-{datetime.now().strftime('%Y%m%d%H%M%S')}-{random.randint(1000, 9999)}"
                evidence = getattr(request, "evidence", None)

                alert = FraudAlertRecord(
                    alert_id=alert_id,
                    user_id=request.user_id,
                    alert_type=request.alert_type,
                    description=request.description,
                    evidence=evidence or {},
                    status="open",
                    created_at=datetime.utcnow(),
                )
                await FraudAlertRepository.save(db, alert)

                profile = await UserRiskProfileRepository.get_or_create(db, request.user_id)
                flags = json.loads(profile.flags) if isinstance(profile.flags, str) else (profile.flags or [])
                flags.append(request.alert_type)
                profile.flags = json.dumps(flags)
                profile.risk_score = min((profile.risk_score or 0) + 0.1, 1.0)
                await UserRiskProfileRepository.save(db, profile)

                return {
                    "alert_id": alert_id,
                    "user_id": request.user_id,
                    "alert_type": request.alert_type,
                    "description": request.description,
                    "evidence": evidence or {},
                    "status": "open",
                    "created_at": datetime.utcnow().isoformat(),
                }
            except Exception as e:
                logger.error(f"CreateAlert error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def GetAlerts(self, request, context):
        async with async_session_factory() as db:
            try:
                status = getattr(request, "status", None) or None
                limit = getattr(request, "limit", 50)
                records = await FraudAlertRepository.list_alerts(db, status=status, limit=limit)

                return {
                    "count": len(records),
                    "alerts": [
                        {
                            "alert_id": r.alert_id,
                            "user_id": r.user_id,
                            "alert_type": r.alert_type,
                            "description": r.description,
                            "evidence": r.evidence or {},
                            "status": r.status,
                            "created_at": r.created_at.isoformat() if r.created_at else None,
                        }
                        for r in records
                    ],
                }
            except Exception as e:
                logger.error(f"GetAlerts error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def BlockUser(self, request, context):
        async with async_session_factory() as db:
            try:
                profile = await UserRiskProfileRepository.get_or_create(db, request.user_id)
                profile.is_blocked = True
                profile.block_reason = getattr(request, "reason", "")
                profile.blocked_at = datetime.utcnow().isoformat()
                await UserRiskProfileRepository.save(db, profile)
                return {"user_id": request.user_id, "blocked": True, "reason": getattr(request, "reason", "")}
            except Exception as e:
                logger.error(f"BlockUser error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def UnblockUser(self, request, context):
        async with async_session_factory() as db:
            try:
                profile = await UserRiskProfileRepository.get_or_create(db, request.user_id)
                profile.is_blocked = False
                profile.risk_score = 0.0
                profile.flags = "[]"
                await UserRiskProfileRepository.save(db, profile)
                return {"user_id": request.user_id, "unblocked": True}
            except Exception as e:
                logger.error(f"UnblockUser error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def RegisterFingerprint(self, request, context):
        async with async_session_factory() as db:
            try:
                fingerprint = DeviceFingerprint(
                    user_id=request.user_id,
                    canvas_hash=getattr(request, "canvas_hash", None),
                    webgl_hash=getattr(request, "webgl_hash", None),
                    audio_hash=getattr(request, "audio_hash", None),
                    fonts=list(getattr(request, "fonts", [])) or None,
                    screen_resolution=getattr(request, "screen_resolution", None),
                    ip_address=getattr(request, "ip_address", None),
                    user_agent=getattr(request, "user_agent", None),
                )
                await MultiAccountDetector.register_fingerprint(db, fingerprint)
                return {"status": "registered", "user_id": request.user_id}
            except Exception as e:
                logger.error(f"RegisterFingerprint error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def CheckMultiAccount(self, request, context):
        async with async_session_factory() as db:
            try:
                related = await MultiAccountDetector.check_multi_account(db, request.user_id)
                return {
                    "user_id": request.user_id,
                    "related_accounts": related,
                    "is_multi_account": len(related) > 0,
                }
            except Exception as e:
                logger.error(f"CheckMultiAccount error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def ScoreTransaction(self, request, context):
        async with async_session_factory() as db:
            try:
                score = await FraudScorer.calculate_score(
                    db,
                    user_id=request.user_id,
                    transaction_amount=getattr(request, "transaction_amount", 0),
                    is_new_account=getattr(request, "is_new_account", False),
                    ip_country=getattr(request, "ip_country", "unknown"),
                    payment_method_new=getattr(request, "payment_method_new", True),
                    device_matches=getattr(request, "device_matches", True),
                )
                await db.execute(
                    FraudScoreRecord.__table__.delete().where(FraudScoreRecord.user_id == request.user_id)
                )
                db.add(FraudScoreRecord(
                    user_id=score.user_id,
                    score=score.score,
                    category=score.category,
                    signals=score.signals,
                    recommendations=score.recommendations,
                    last_updated=score.last_updated,
                ))
                await db.commit()
                return score.model_dump()
            except Exception as e:
                logger.error(f"ScoreTransaction error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def GetUserFraudScore(self, request, context):
        async with async_session_factory() as db:
            try:
                result = await db.execute(
                    select(FraudScoreRecord).where(FraudScoreRecord.user_id == request.user_id)
                )
                record = result.scalar_one_or_none()
                if not record:
                    context.set_code(grpc.StatusCode.NOT_FOUND)
                    context.set_details("No score found")
                    return {}
                return {
                    "user_id": record.user_id,
                    "score": record.score,
                    "category": record.category,
                    "signals": record.signals or {},
                    "recommendations": record.recommendations or [],
                    "last_updated": record.last_updated.isoformat() if record.last_updated else None,
                }
            except Exception as e:
                logger.error(f"GetUserFraudScore error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def DetectBot(self, request, context):
        try:
            action_timestamps = [
                datetime.fromisoformat(ts) if isinstance(ts, str) else ts
                for ts in getattr(request, "action_timestamps", [])
            ]
            result = BotDetector.analyze_behavior(
                action_timestamps=action_timestamps,
                mouse_movements=getattr(request, "mouse_movements", 0),
                touch_events=getattr(request, "touch_events", 0),
                session_duration=timedelta(seconds=getattr(request, "session_duration_seconds", 0)),
                perfect_play_pct=getattr(request, "perfect_play_pct", None),
            )
            return result.model_dump()
        except Exception as e:
            logger.error(f"DetectBot error: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(str(e))
            return {}

    async def AnalyzeTableCollusion(self, request, context):
        try:
            table_id = getattr(request, "table_id", "")
            players = list(getattr(request, "players", []))
            signals = CollusionDetector.analyze_table(table_id, players)
            return [s.model_dump() for s in signals]
        except Exception as e:
            logger.error(f"AnalyzeTableCollusion error: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(str(e))
            return []


async def serve_grpc(port: int) -> grpc.aio.Server:
    server = grpc.aio.server(futures.ThreadPoolExecutor(max_workers=10))
    server.add_insecure_port(f"[::]:{port}")
    await server.start()
    logger.info(f"Fraud gRPC server listening on port {port}")
    return server
