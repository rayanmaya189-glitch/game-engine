"""Fraud scoring endpoints"""

from fastapi import APIRouter, HTTPException
from typing import Dict

from app.models.schemas import FraudScore, fraud_scores
from app.services.fraud_scorer import FraudScorer

router = APIRouter(prefix="/score", tags=["score"])


@router.post("/transaction", response_model=FraudScore)
async def score_transaction(
    user_id: str,
    transaction_amount: float,
    is_new_account: bool = False,
    ip_country: str = "unknown",
    payment_method_new: bool = True,
    device_matches: bool = True
):
    """Calculate fraud score for a transaction"""
    score = FraudScorer.calculate_score(
        user_id,
        transaction_amount,
        is_new_account,
        ip_country,
        payment_method_new,
        device_matches
    )
    fraud_scores[user_id] = score
    return score


@router.get("/user/{user_id}", response_model=FraudScore)
async def get_user_fraud_score(user_id: str):
    """Get the latest fraud score for a user"""
    if user_id not in fraud_scores:
        raise HTTPException(status_code=404, detail="No score found")
    return fraud_scores[user_id]
