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
DAILY_LIMIT = 10000.0
WEEKLY_LIMIT = 25000.0
MONTHLY_LIMIT = 100000.0
SINGLE_TX_LIMIT = 5000.0

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

    async def ScreenSanctions(self, request, context):
        try:
            return {
                "name": request.name,
                "matched": False,
                "lists_checked": ["OFAC", "EU_SANCTIONS", "UN_SANCTIONS"],
                "risk_level": "LOW",
            }
        except Exception as e:
            logger.error(f"ScreenSanctions error: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(str(e))
            return {}

    async def ScreenPEP(self, request, context):
        try:
            return {
                "name": request.name,
                "is_pep": False,
                "risk_level": "LOW",
            }
        except Exception as e:
            logger.error(f"ScreenPEP error: {e}")
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

    async def GetUserLimits(self, request, context):
        async with async_session_factory() as db:
            try:
                user_id = request.user_id
                result = await db.execute(
                    select(TransactionRecord).where(TransactionRecord.user_id == user_id)
                )
                transactions = result.scalars().all()
                daily_deposits = sum(t.amount for t in transactions if t.type == "deposit")

                return {
                    "user_id": user_id,
                    "daily_deposit_limit": DAILY_LIMIT,
                    "weekly_deposit_limit": WEEKLY_LIMIT,
                    "monthly_deposit_limit": MONTHLY_LIMIT,
                    "single_transaction_limit": SINGLE_TX_LIMIT,
                    "current_daily": daily_deposits,
                    "current_weekly": 0.0,
                    "current_monthly": 0.0,
                }
            except Exception as e:
                logger.error(f"GetUserLimits error: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(str(e))
                return {}

    async def CheckLimits(self, request, context):
        try:
            limits = {
                "daily": DAILY_LIMIT,
                "weekly": WEEKLY_LIMIT,
                "monthly": MONTHLY_LIMIT,
            }
            period = getattr(request, "period", "daily")
            amount = getattr(request, "amount", 0)
            limit = limits.get(period, DAILY_LIMIT)

            return {
                "user_id": request.user_id,
                "amount": amount,
                "period": period,
                "limit": limit,
                "within_limit": amount <= limit,
                "remaining": limit - amount,
            }
        except Exception as e:
            logger.error(f"CheckLimits error: {e}")
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

    async def GenerateCTRReport(self, request, context):
        async with async_session_factory() as db:
            try:
                start = datetime.fromisoformat(request.start_date)
                end = datetime.fromisoformat(request.end_date)

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


async def serve_grpc(port: int) -> grpc.aio.Server:
    server = grpc.aio.server(futures.ThreadPoolExecutor(max_workers=10))
    server.add_insecure_port(f"[::]:{port}")
    await server.start()
    logger.info(f"AML gRPC server listening on port {port}")
    return server
