from fastapi import APIRouter, HTTPException
from typing import Dict, List, Optional
from datetime import datetime

from app.models.schemas import Alert, AlertStatus, AlertSeverity, alerts_db

router = APIRouter(prefix="", tags=["alerts"])


@router.get("/alerts", response_model=List[Alert])
async def list_alerts(
    status: Optional[AlertStatus] = None,
    severity: Optional[AlertSeverity] = None,
    limit: int = 100
):
    """List alerts with optional filters"""
    alerts = list(alerts_db.values())

    if status:
        alerts = [a for a in alerts if a.status == status]
    if severity:
        alerts = [a for a in alerts if a.severity == severity]

    alerts.sort(key=lambda a: a.created_at, reverse=True)

    return alerts[:limit]


@router.get("/alerts/{alert_id}", response_model=Alert)
async def get_alert(alert_id: str):
    """Get a specific alert"""
    if alert_id not in alerts_db:
        raise HTTPException(status_code=404, detail="Alert not found")
    return alerts_db[alert_id]


@router.patch("/alerts/{alert_id}", response_model=Alert)
async def update_alert(alert_id: str, update: Dict):
    """Update alert status or assign to investigator"""
    if alert_id not in alerts_db:
        raise HTTPException(status_code=404, detail="Alert not found")

    alert = alerts_db[alert_id]

    if "status" in update:
        alert.status = AlertStatus(update["status"])
    if "assigned_to" in update:
        alert.assigned_to = update["assigned_to"]
    if "notes" in update:
        alert.notes = update["notes"]

    if alert.status == AlertStatus.RESOLVED:
        alert.resolved_at = datetime.now()

    alert.updated_at = datetime.now()

    return alert
