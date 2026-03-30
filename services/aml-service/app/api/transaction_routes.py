from fastapi import APIRouter
from typing import Dict, List

from app.models.schemas import Transaction, alerts_db, transactions_db
from app.services.rules_engine import AMLRulesEngine

router = APIRouter(prefix="", tags=["transactions"])


@router.post("/transactions", response_model=Dict)
async def analyze_transaction(transaction: Transaction):
    """Analyze a transaction for suspicious patterns"""
    transactions_db[transaction.transaction_id] = transaction

    alerts = AMLRulesEngine.run_all_rules(transaction.user_id, transaction)

    for alert in alerts:
        alerts_db[alert.alert_id] = alert

    return {
        "transaction_id": transaction.transaction_id,
        "alerts_generated": len(alerts),
        "alert_ids": [a.alert_id for a in alerts]
    }


@router.post("/batch/analyze", response_model=Dict)
async def analyze_transactions(transactions: List[Transaction]):
    """Batch analyze multiple transactions"""
    all_alerts = []

    for transaction in transactions:
        transactions_db[transaction.transaction_id] = transaction

        alerts = AMLRulesEngine.run_all_rules(transaction.user_id, transaction)
        all_alerts.extend(alerts)

        for alert in alerts:
            alerts_db[alert.alert_id] = alert

    return {
        "transactions_processed": len(transactions),
        "total_alerts": len(all_alerts),
        "alerts": [a.dict() for a in all_alerts]
    }
