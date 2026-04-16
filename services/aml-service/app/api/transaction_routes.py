from typing import Dict, List
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select

from app.database import get_db
from app import db_models
from app.models.schemas import Transaction, Alert
from app.services.rules_engine import AMLRulesEngine


async def analyze_transaction(transaction: Transaction, db: AsyncSession) -> Dict:
    """Analyze a transaction for suspicious patterns"""
    record = db_models.TransactionRecord(
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
        db.add(db_models.AlertRecord(
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
        "alert_ids": [a.alert_id for a in alerts]
    }


async def analyze_transactions(transactions: List[Transaction], db: AsyncSession) -> Dict:
    """Batch analyze multiple transactions"""
    all_alerts = []

    for transaction in transactions:
        db.add(db_models.TransactionRecord(
            transaction_id=transaction.transaction_id,
            user_id=transaction.user_id,
            type=transaction.type,
            amount=transaction.amount,
            currency=transaction.currency,
            payment_method=transaction.payment_method,
            ip_address=transaction.ip_address,
            country=transaction.country,
            timestamp=transaction.timestamp,
        ))

        alerts = await AMLRulesEngine.run_all_rules(db, transaction.user_id, transaction)
        all_alerts.extend(alerts)

        for alert in alerts:
            db.add(db_models.AlertRecord(
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
        "transactions_processed": len(transactions),
        "total_alerts": len(all_alerts),
        "alerts": [a.model_dump() for a in all_alerts]
    }
