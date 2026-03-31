"""Bot detection functions"""

from typing import Optional, List
from datetime import datetime, timedelta

from app.models.schemas import BotDetectionResult
from app.services.bot_detector import BotDetector


async def detect_bot(
    action_timestamps: List[datetime],
    mouse_movements: int = 0,
    touch_events: int = 0,
    session_duration_seconds: int = 0,
    perfect_play_pct: Optional[float] = None
) -> BotDetectionResult:
    """Detect if behavior indicates a bot"""
    return BotDetector.analyze_behavior(
        action_timestamps,
        mouse_movements,
        touch_events,
        timedelta(seconds=session_duration_seconds),
        perfect_play_pct
    )
