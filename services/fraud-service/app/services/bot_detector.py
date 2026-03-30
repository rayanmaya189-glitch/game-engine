"""Bot detection via behavior analysis"""

from typing import Optional, List
from datetime import datetime, timedelta

from app.models.schemas import BotDetectionResult


class BotDetector:
    """Detect automated/bot behavior"""

    @staticmethod
    def analyze_behavior(
        action_timestamps: List[datetime],
        mouse_movements: int,
        touch_events: int,
        session_duration: timedelta,
        perfect_play_pct: Optional[float] = None
    ) -> BotDetectionResult:
        """Analyze user behavior for bot indicators"""
        signals = {}

        if len(action_timestamps) > 5:
            intervals = [
                (action_timestamps[i+1] - action_timestamps[i]).total_seconds()
                for i in range(len(action_timestamps) - 1)
            ]
            variance = sum((x - sum(intervals)/len(intervals))**2 for x in intervals) / len(intervals)
            signals["timing_variance"] = variance

            if variance < 0.1:
                signals["consistent_timing"] = 0.9
            else:
                signals["consistent_timing"] = 0.0

        if mouse_movements < 5 and action_timestamps:
            signals["no_mouse"] = 0.7

        if touch_events < 2 and action_timestamps:
            signals["no_touch"] = 0.6

        if perfect_play_pct and perfect_play_pct > 0.95:
            signals["perfect_play"] = 0.8

        if session_duration.total_seconds() > 8 * 3600 and not (
            datetime.now() - action_timestamps[0] < timedelta(hours=24)
        ):
            signals["no_breaks"] = 0.5

        total_signal = sum(signals.values())
        confidence = min(total_signal, 1.0)

        return BotDetectionResult(
            is_bot=confidence > 0.5,
            confidence=confidence,
            signals=signals
        )
