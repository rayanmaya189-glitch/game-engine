"""
gRPC server for AML service.
Exposes anti-money laundering detection methods.
"""

import logging
from datetime import datetime
from concurrent import futures

import grpc
from sqlalchemy import select

from app.database import async_session_factory
from app.models import AlertRecord, TransactionRecord
from app.models.schemas import Transaction, RiskScore
from app.services.rules_engine import AMLRulesEngine
from app.services.risk_model import MLRiskModel

logger = logging.getLogger(__name__)

LARGE_TX_THRESHOLD = 10000
HIGH_RISK_COUNTRIES = ["KP", "IR", "SY", "CU"]


class AMLServiceServicer:

    async def CheckTransaction(self, request, context):
        async with async_session_factory() as db:
            try:
                risk_score = 0.1
                amount = getattr(request, "amount", 0)
                country = getattr(request, "country", None)

                if amount > LARGE_TX_THRESHOLD:
                    risk_score += 0.3
                if country in HIGH_RISK_COUNTRIES:
                    risk_score += 0.4

                user_id = request.user_id
                result = await db.execute(
                    select(AlertRecord).where(AlertRecord.user_id == user_id)
                )
                existing_alerts = result.scalars().all()

                return {
                    "user_id": user_id,
                    "transaction_allowed": risk_score < 0.7,
                    "risk_score": risk_score,
                    "requires_review": risk_score >= 0.5,
                    "alerts": [] if risk_score < 0.5 else [
                        "High transaction amount",
                        "Manual review required",
                    ],
                    "existing_alert_count": len(existing_alerts),
                }
            except Exception as e:
                logger.error(f"CheckTransaction error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def GetAlerts(self, request, context):
        async with async_session_factory() as db:
            try:
                status = getattr(request, "status", None) or None
                limit = getattr(request, "limit", 50)

                stmt = select(AlertRecord)
                if status:
                    stmt = stmt.where(AlertRecord.status == status)
                stmt = stmt.order_by(AlertRecord.created_at.desc()).limit(limit)
                result = await db.execute(stmt)
                records = result.scalars().all()

                return {
                    "count": len(records),
                    "alerts": [
                        {
                            "alert_id": r.alert_id,
                            "user_id": r.user_id,
                            "alert_type": r.alert_type,
                            "severity": r.severity,
                            "status": r.status,
                            "description": r.description,
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

    async def AnalyzeTransaction(self, request, context):
        async with async_session_factory() as db:
            try:
                transaction = Transaction(
                    transaction_id=request.transaction_id,
                    user_id=request.user_id,
                    type=request.type,
                    amount=request.amount,
                    currency=request.currency,
                    payment_method=getattr(request, "payment_method", None),
                    ip_address=getattr(request, "ip_address", None),
                    country=getattr(request, "country", None),
                    timestamp=datetime.utcnow(),
                )

                record = TransactionRecord(
                    transaction_id=transaction.transaction_id,
                    user_id=transaction.user_id,
                    type=transaction.type,
                    amount=transaction.amount,
                    currency=transaction.currency,
                    payment_method=transaction.payment_method,
                    ip_address=transaction.ip_address,
                    country=transaction.country,
                    timestamp=transaction.timestamp,
                )
                db.add(record)
                await db.commit()

                alerts = await AMLRulesEngine.run_all_rules(db, transaction.user_id, transaction)
                for alert in alerts:
                    db.add(AlertRecord(
                        alert_id=alert.alert_id,
                        user_id=alert.user_id,
                        alert_type=alert.alert_type.value,
                        severity=alert.severity.value,
                        status=alert.status.value,
                        description=alert.description,
                        transactions=alert.transactions,
                        assigned_to=alert.assigned_to,
                        created_at=alert.created_at,
                        updated_at=alert.updated_at,
                        resolved_at=alert.resolved_at,
                        notes=alert.notes,
                    ))
                await db.commit()

                return {
                    "transaction_id": transaction.transaction_id,
                    "alerts_generated": len(alerts),
                    "alert_ids": [a.alert_id for a in alerts],
                }
            except Exception as e:
                logger.error(f"AnalyzeTransaction error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}


async def serve_grpc(port: int) -> grpc.aio.Server:
    server = grpc.aio.server(futures.ThreadPoolExecutor(max_workers=10))
    server.add_insecure_port(f"[::]:{port}")
    await server.start()
    logger.info(f"AML gRPC server listening on port {port}")
    return server
