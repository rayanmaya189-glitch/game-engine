"""
Collusion detection RPCs for Fraud Detection Service.
"""

import logging

import grpc

from app.services.collusion_detector import CollusionDetector

logger = logging.getLogger(__name__)


class FraudServiceServicer:

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
