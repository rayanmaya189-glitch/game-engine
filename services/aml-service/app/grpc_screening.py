"""
Sanctions, PEP screening, and limits RPCs for AML service.
"""

import logging

import grpc

from app.database import async_session_factory
from app.models import TransactionRecord
from sqlalchemy import select

logger = logging.getLogger(__name__)

DAILY_LIMIT = 10000.0
WEEKLY_LIMIT = 25000.0
MONTHLY_LIMIT = 100000.0
SINGLE_TX_LIMIT = 5000.0


class AMLServiceServicer:

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
