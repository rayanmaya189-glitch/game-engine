from fastapi import APIRouter
from typing import Dict, List

from app.models.schemas import RiskScore, risk_scores_db
from app.services.risk_model import MLRiskModel

router = APIRouter(prefix="", tags=["risk"])


@router.get("/risk/score/{user_id}", response_model=RiskScore)
async def get_risk_score(user_id: str):
    """Get risk score for a user"""
    score = MLRiskModel.calculate_risk_score(user_id)
    risk_scores_db[user_id] = score
    return score


@router.post("/risk/score/batch", response_model=Dict)
async def calculate_batch_risk(user_ids: List[str]):
    """Calculate risk scores for multiple users"""
    results = {}
    for user_id in user_ids:
        score = MLRiskModel.calculate_risk_score(user_id)
        risk_scores_db[user_id] = score
        results[user_id] = score

    return {"scores": results}
