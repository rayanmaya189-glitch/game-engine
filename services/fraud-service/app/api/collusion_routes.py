"""Collusion detection endpoint"""

from fastapi import APIRouter
from typing import List

from app.models.schemas import CollusionSignal
from app.services.collusion_detector import CollusionDetector

router = APIRouter(prefix="/collusion", tags=["collusion"])


@router.get("/table/{table_id}", response_model=List[CollusionSignal])
async def analyze_table_collusion(table_id: str, players: List[str]):
    """Analyze a table for collusion patterns"""
    return CollusionDetector.analyze_table(table_id, players)
