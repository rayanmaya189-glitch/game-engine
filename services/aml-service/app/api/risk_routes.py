from typing import Dict, List
from sqlalchemy.ext.asyncio import AsyncSession

from app.database import get_db
from app.models import RiskScoreRecord
from app.models.schemas import RiskScore
from app.services.risk_model import MLRiskModel


async def get_risk_score(user_id: str, db: AsyncSession) -> RiskScore:
    """Get risk score for a user"""
    score = await MLRiskModel.calculate_risk_score(db, user_id)
    await MLRiskModel.save_risk_score(db, score)
    return score


async def calculate_batch_risk(user_ids: List[str], db: AsyncSession) -> Dict:
    """Calculate risk scores for multiple users"""
    results = {}
    for user_id in user_ids:
        score = await MLRiskModel.calculate_risk_score(db, user_id)
        await MLRiskModel.save_risk_score(db, score)
        results[user_id] = score

    return {"scores": results}
