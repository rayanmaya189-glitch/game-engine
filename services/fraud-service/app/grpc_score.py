"""
Scoring, risk assessment, bot detection RPCs for Fraud Detection Service.
"""

import logging
from datetime import datetime, timedelta

import grpc
from sqlalchemy import select

from app.database import async_session_factory
from app import db_models
from app.services.fraud_scorer import FraudScorer
from app.services.bot_detector import BotDetector

logger = logging.getLogger(__name__)

BET_AMOUNT = 1000.0
RAPID_BET_SECONDS = 2


class FraudServiceServicer:

    async def CheckBetRisk(self, request, context):
        async with async_session_factory() as db:
            try:
                risk_score = 0.0
                risk_factors = []
                amount = getattr(request, "amount", 0)

                if amount > BET_AMOUNT:
                    risk_score += 0.15
                    risk_factors.append("Unusually high bet amount")

                from app.repositories import UserRiskProfileRepository
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
                    from app.repositories import DeviceFingerprintRepository
                    fp = await DeviceFingerprintRepository.get_by_user(db, request.user_id)
                    if fp and fp.canvas_hash != device_fingerprint:
                        risk_score += 0.4
                        risk_factors.append("Account access from new device")

                ip_address = getattr(request, "ip_address", None)
                if ip_address:
                    from app.repositories import UserRiskProfileRepository
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
                    db_models.FraudScoreRecord.__table__.delete().where(db_models.FraudScoreRecord.user_id == request.user_id)
                )
                db.add(db_models.FraudScoreRecord(
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
                    select(db_models.FraudScoreRecord).where(db_models.FraudScoreRecord.user_id == request.user_id)
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
