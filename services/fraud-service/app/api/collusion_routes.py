"""Collusion detection functions"""

from typing import List

from app.models.schemas import CollusionSignal
from app.services.collusion_detector import CollusionDetector


async def analyze_table_collusion(table_id: str, players: List[str]) -> List[CollusionSignal]:
    """Analyze a table for collusion patterns"""
    return CollusionDetector.analyze_table(table_id, players)
