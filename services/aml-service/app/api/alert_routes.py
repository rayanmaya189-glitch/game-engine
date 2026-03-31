from typing import Dict, List, Optional
from datetime import datetime
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select, update

from app.database import get_db
from app.models import AlertRecord
from app.models.schemas import Alert, AlertStatus, AlertSeverity


def _record_to_alert(r: AlertRecord) -> Alert:
    return Alert(
        alert_id=r.alert_id,
        user_id=r.user_id,
        alert_type=r.alert_type,
        severity=r.severity,
        status=r.status,
        description=r.description,
        transactions=r.transactions or [],
        assigned_to=r.assigned_to,
        created_at=r.created_at,
        updated_at=r.updated_at,
        resolved_at=r.resolved_at,
        notes=r.notes,
    )


async def list_alerts(
    status: Optional[AlertStatus] = None,
    severity: Optional[AlertSeverity] = None,
    limit: int = 100,
    db: AsyncSession = None,
) -> List[Alert]:
    """List alerts with optional filters"""
    stmt = select(AlertRecord)
    if status:
        stmt = stmt.where(AlertRecord.status == status.value)
    if severity:
        stmt = stmt.where(AlertRecord.severity == severity.value)
    stmt = stmt.order_by(AlertRecord.created_at.desc()).limit(limit)

    result = await db.execute(stmt)
    records = result.scalars().all()
    return [_record_to_alert(r) for r in records]


async def get_alert(alert_id: str, db: AsyncSession) -> Alert:
    """Get a specific alert"""
    result = await db.execute(select(AlertRecord).where(AlertRecord.alert_id == alert_id))
    record = result.scalar_one_or_none()
    if not record:
        raise ValueError("Alert not found")
    return _record_to_alert(record)


async def update_alert(alert_id: str, update_data: Dict, db: AsyncSession) -> Alert:
    """Update alert status or assign to investigator"""
    result = await db.execute(select(AlertRecord).where(AlertRecord.alert_id == alert_id))
    record = result.scalar_one_or_none()
    if not record:
        raise ValueError("Alert not found")

    if "status" in update_data:
        record.status = update_data["status"]
    if "assigned_to" in update_data:
        record.assigned_to = update_data["assigned_to"]
    if "notes" in update_data:
        record.notes = update_data["notes"]

    if record.status == AlertStatus.RESOLVED.value:
        record.resolved_at = datetime.utcnow()

    record.updated_at = datetime.utcnow()
    await db.commit()
    await db.refresh(record)

    return _record_to_alert(record)
