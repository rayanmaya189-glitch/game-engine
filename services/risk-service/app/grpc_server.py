"""
gRPC server for risk-service.
Exposes risk scoring methods for other services to call.
"""

import asyncio
import logging
from concurrent import futures

import grpc

from app.services.risk_engine import RiskScoringEngine

logger = logging.getLogger(__name__)


class RiskServiceServicer:

    async def CalculateProfile(self, request, context):
        profile = await RiskScoringEngine.calculate_profile(
            user_id=request.user_id,
            kyc_level=request.kyc_level,
            aml_alerts=request.aml_alerts,
            fraud_score=request.fraud_score,
            transaction_risk=getattr(request, "transaction_risk", 0),
            device_risk=getattr(request, "device_risk", 0),
            velocity_risk=getattr(request, "velocity_risk", 0),
        )
        return profile

    async def AssessTransaction(self, request, context):
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

        if getattr(request, "transaction_amount", None) and getattr(request, "transaction_type", None):
            allowed, reason = RiskScoringEngine.check_transaction_allowed(
                profile,
                request.transaction_amount,
                request.transaction_type,
            )

        return {"allowed": allowed, "reason": reason, "profile": profile, "actions": profile.recommended_actions}


async def serve_grpc(port: int) -> grpc.aio.Server:
    """Start the gRPC server and return it for lifecycle management."""
    server = grpc.aio.server(futures.ThreadPoolExecutor(max_workers=10))

    # When proto-generated code is available, register the servicer:
    #   from game_engine.v1 import risk_pb2_grpc
    #   risk_pb2_grpc.add_RiskServiceServicer_to_server(RiskServiceServicer(), server)

    server.add_insecure_port(f"[::]:{port}")
    await server.start()
    return server
