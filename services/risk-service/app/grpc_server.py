"""
gRPC server for risk-service.
Exposes risk scoring methods for other services to call.
"""

import logging
import os
from concurrent import futures

import grpc

from app.services.risk_engine import RiskScoringEngine

logger = logging.getLogger(__name__)


class RiskServiceServicer:

    async def CalculateProfile(self, request, context):
        try:
            profile = await RiskScoringEngine.calculate_profile(
                user_id=request.user_id,
                kyc_level=getattr(request, "kyc_level", 0),
                aml_alerts=getattr(request, "aml_alerts", 0),
                fraud_score=getattr(request, "fraud_score", 0),
                transaction_risk=getattr(request, "transaction_risk", 0),
                device_risk=getattr(request, "device_risk", 0),
                velocity_risk=getattr(request, "velocity_risk", 0),
            )
            return profile.model_dump()
        except Exception as e:
            logger.error(f"CalculateProfile error: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(str(e))
            return {}

    async def AssessTransaction(self, request, context):
        try:
            kyc_level = getattr(request, "kyc_level", 2)
            aml_alerts = getattr(request, "aml_alerts", 0)
            fraud_score = getattr(request, "fraud_score", 0)

            profile = await RiskScoringEngine.calculate_profile(
                user_id=request.user_id,
                kyc_level=kyc_level,
                aml_alerts=aml_alerts,
                fraud_score=fraud_score,
            )

            allowed = True
            reason = None

            transaction_amount = getattr(request, "transaction_amount", None)
            transaction_type = getattr(request, "transaction_type", None)
            if transaction_amount and transaction_type:
                allowed, reason = RiskScoringEngine.check_transaction_allowed(
                    profile,
                    transaction_amount,
                    transaction_type,
                )

            return {
                "allowed": allowed,
                "reason": reason,
                "profile": profile.model_dump(),
                "actions": profile.recommended_actions,
            }
        except Exception as e:
            logger.error(f"AssessTransaction error: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(str(e))
            return {}

    async def GetUserRiskProfile(self, request, context):
        try:
            profile = await RiskScoringEngine.calculate_profile(
                user_id=request.user_id,
                kyc_level=getattr(request, "kyc_level", 1),
                aml_alerts=0,
                fraud_score=0,
            )
            return profile.model_dump()
        except Exception as e:
            logger.error(f"GetUserRiskProfile error: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(str(e))
            return {}

    async def CalculateLimits(self, request, context):
        try:
            BASE_DEPOSIT = int(os.environ.get("RISK_BASE_DEPOSIT_LIMIT", "10000"))
            BASE_WITHDRAWAL = int(os.environ.get("RISK_BASE_WITHDRAWAL_LIMIT", "5000"))

            account_age_days = getattr(request, "account_age_days", 0)
            vip_level = getattr(request, "vip_level", 0)
            current_risk_score = getattr(request, "current_risk_score", 0)

            if account_age_days < 30:
                age_factor = 0.5
            elif account_age_days < 90:
                age_factor = 0.75
            else:
                age_factor = 1.0

            vip_factor = 1.0 + (vip_level * 0.25)
            risk_factor = 1.0 - (current_risk_score / 200)

            deposit_limit = BASE_DEPOSIT * age_factor * vip_factor * risk_factor
            withdrawal_limit = BASE_WITHDRAWAL * age_factor * vip_factor * risk_factor

            return {
                "user_id": request.user_id,
                "daily_deposit_limit": round(deposit_limit, 2),
                "daily_withdrawal_limit": round(withdrawal_limit, 2),
                "calculation": {
                    "base_deposit": BASE_DEPOSIT,
                    "age_factor": age_factor,
                    "vip_factor": vip_factor,
                    "risk_factor": risk_factor,
                },
            }
        except Exception as e:
            logger.error(f"CalculateLimits error: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(str(e))
            return {}


async def serve_grpc(port: int) -> grpc.aio.Server:
    server = grpc.aio.server(futures.ThreadPoolExecutor(max_workers=10))
    server.add_insecure_port(f"[::]:{port}")
    await server.start()
    logger.info(f"Risk gRPC server listening on port {port}")
    return server
