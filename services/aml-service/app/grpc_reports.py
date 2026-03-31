"""
SAR creation and compliance report RPCs for AML service.
"""

import logging
from datetime import datetime

import grpc
from sqlalchemy import select

from app.database import async_session_factory
from app.models import AlertRecord
from app.services.rules_engine import AMLRulesEngine
from app.services.risk_model import MLRiskModel

logger = logging.getLogger(__name__)


class AMLServiceServicer:

    async def CreateSAR(self, request, context):
        async with async_session_factory() as db:
            try:
                sar_id = f"SAR-{datetime.now().strftime('%Y%m%d%H%M%S')}"
                alert = AlertRecord(
                    alert_id=sar_id,
                    user_id=request.user_id,
                    alert_type=request.suspicious_activity_type,
                    severity="high",
                    status="open",
                    description=request.description,
                    transactions=[request.transaction_id],
                    created_at=datetime.utcnow(),
                    updated_at=datetime.utcnow(),
                )
                db.add(alert)
                await db.commit()

                return {
                    "sar_id": sar_id,
                    "status": "PENDING_REVIEW",
                    "created_at": datetime.utcnow().isoformat(),
                }
            except Exception as e:
                logger.error(f"CreateSAR error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def GenerateCTRReport(self, request, context):
        async with async_session_factory() as db:
            try:
                start = datetime.fromisoformat(request.start_date)
                end = datetime.fromisoformat(request.end_date)

                from app.models import TransactionRecord
                stmt = select(TransactionRecord).where(
                    TransactionRecord.timestamp > start,
                    TransactionRecord.timestamp < end,
                    TransactionRecord.amount > AMLRulesEngine.LARGE_TRANSACTION_THRESHOLD,
                )
                result = await db.execute(stmt)
                records = result.scalars().all()

                large_transactions = [
                    {
                        "transaction_id": t.transaction_id,
                        "user_id": t.user_id,
                        "type": t.type,
                        "amount": t.amount,
                        "currency": t.currency,
                        "payment_method": t.payment_method,
                        "ip_address": t.ip_address,
                        "country": t.country,
                        "timestamp": t.timestamp.isoformat(),
                    }
                    for t in records
                ]

                return {
                    "report_type": "CTR",
                    "period": {"start": request.start_date, "end": request.end_date},
                    "transactions": large_transactions,
                    "total_count": len(large_transactions),
                }
            except Exception as e:
                logger.error(f"GenerateCTRReport error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def GenerateSARReport(self, request, context):
        async with async_session_factory() as db:
            try:
                result = await db.execute(
                    select(AlertRecord).where(AlertRecord.alert_id == request.alert_id)
                )
                record = result.scalar_one_or_none()
                if not record:
                    context.set_code(grpc.StatusCode.NOT_FOUND)
                    context.set_details("Alert not found")
                    return {}

                return {
                    "report_type": "SAR",
                    "alert": {
                        "alert_id": record.alert_id,
                        "user_id": record.user_id,
                        "alert_type": record.alert_type,
                        "severity": record.severity,
                        "status": record.status,
                        "description": record.description,
                        "transactions": record.transactions or [],
                        "assigned_to": record.assigned_to,
                        "created_at": record.created_at.isoformat(),
                        "updated_at": record.updated_at.isoformat(),
                        "resolved_at": record.resolved_at.isoformat() if record.resolved_at else None,
                        "notes": record.notes,
                    },
                    "generated_at": datetime.utcnow().isoformat(),
                    "format": "FinCEN BSA",
                }
            except Exception as e:
                logger.error(f"GenerateSARReport error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def GetRiskScore(self, request, context):
        async with async_session_factory() as db:
            try:
                score = await MLRiskModel.calculate_risk_score(db, request.user_id)
                await MLRiskModel.save_risk_score(db, score)
                return score.model_dump()
            except Exception as e:
                logger.error(f"GetRiskScore error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def BatchRiskScore(self, request, context):
        async with async_session_factory() as db:
            try:
                user_ids = list(request.user_ids)
                results = {}
                for user_id in user_ids:
                    score = await MLRiskModel.calculate_risk_score(db, user_id)
                    await MLRiskModel.save_risk_score(db, score)
                    results[user_id] = score.model_dump()
                return {"scores": results}
            except Exception as e:
                logger.error(f"BatchRiskScore error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}
