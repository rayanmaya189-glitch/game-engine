package com.casino.game.ui.animations

import androidx.compose.animation.core.*
import androidx.compose.foundation.Canvas
import androidx.compose.foundation.layout.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.unit.dp
import kotlinx.coroutines.delay
import kotlin.math.*

enum class WheelSpinState { Spinning, Decelerating, Stopped }

data class RouletteAnimationConfig(
    val totalRotations: Int = 4,
    val spinDurationMs: Int = 4000,
    val decelerationDurationMs: Int = 2500,
    val ballBounceDurationMs: Int = 1500,
    val revealGlowDurationMs: Int = 800,
)

@Composable
fun RouletteWheelSpin(
    targetPocket: Int,
    spinState: WheelSpinState,
    modifier: Modifier = Modifier,
    config: RouletteAnimationConfig = RouletteAnimationConfig(),
    content: @Composable (rotation: Float) -> Unit
) {
    val totalDegrees = remember(targetPocket) {
        config.totalRotations * 360f + (360f - targetPocket * (360f / 37f))
    }

    val rotation by animateFloatAsState(
        targetValue = when (spinState) {
            WheelSpinState.Spinning -> totalDegrees * 0.6f
            WheelSpinState.Decelerating -> totalDegrees
            WheelSpinState.Stopped -> totalDegrees
        },
        animationSpec = when (spinState) {
            WheelSpinState.Spinning -> tween(
                durationMillis = config.spinDurationMs,
                easing = LinearEasing
            )
            WheelSpinState.Decelerating -> tween(
                durationMillis = config.decelerationDurationMs,
                easing = CubicBezierEasing(0.1f, 0.8f, 0.2f, 1f)
            )
            WheelSpinState.Stopped -> snap()
        },
        label = "wheel_rotation"
    )

    Box(modifier = modifier, contentAlignment = Alignment.Center) {
        content(rotation)
    }
}

@Composable
fun RouletteBallBounce(
    targetPocket: Int,
    isActive: Boolean,
    modifier: Modifier = Modifier,
    config: RouletteAnimationConfig = RouletteAnimationConfig(),
    onAnimationEnd: () -> Unit = {}
) {
    val transition = updateTransition(targetState = isActive, label = "ball_bounce")

    val rimProgress by transition.animateFloat(
        transitionSpec = {
            if (targetState) tween(durationMillis = config.ballBounceDurationMs, easing = FastOutSlowInEasing)
            else snap()
        },
        label = "rim_progress"
    ) { active -> if (active) 1f else 0f }

    val bounceOffset by transition.animateFloat(
        transitionSpec = {
            if (targetState) keyframes {
                durationMillis = config.ballBounceDurationMs
                0f at 0
                -12f at (durationMillis * 0.15).toInt()
                0f at (durationMillis * 0.3).toInt()
                -8f at (durationMillis * 0.45).toInt()
                0f at (durationMillis * 0.6).toInt()
                -4f at (durationMillis * 0.75).toInt()
                0f at durationMillis
            }
            else snap()
        },
        label = "bounce_offset"
    ) { active -> if (active) 1f else 0f }

    LaunchedEffect(isActive) {
        if (isActive) {
            delay(config.ballBounceDurationMs.toLong())
            onAnimationEnd()
        }
    }

    val pocketAngle = targetPocket * (360f / 37f)
    val angle = remember(rimProgress, pocketAngle) {
        lerp(0f, pocketAngle + 720f, rimProgress)
    }

    Canvas(modifier = modifier.size(280.dp)) {
        val radius = size.minDimension / 2f * (0.82f - 0.12f * rimProgress)
        val cx = size.width / 2f + cos(Math.toRadians(angle.toDouble())).toFloat() * radius
        val cy = size.height / 2f + sin(Math.toRadians(angle.toDouble())).toFloat() * radius

        drawCircle(
            color = Color.White,
            radius = 8f,
            center = Offset(cx, cy + bounceOffset * bounceOffset)
        )
        drawCircle(
            color = Color.LightGray,
            radius = 8f,
            center = Offset(cx, cy + bounceOffset * bounceOffset),
            style = Stroke(width = 1.5f)
        )
    }
}

@Composable
fun RouletteNumberReveal(
    winningNumber: Int,
    isRevealed: Boolean,
    modifier: Modifier = Modifier,
    config: RouletteAnimationConfig = RouletteAnimationConfig(),
    content: @Composable (glowAlpha: Float, scale: Float) -> Unit
) {
    val transition = updateTransition(isRevealed, label = "number_reveal")

    val glowAlpha by transition.animateFloat(
        transitionSpec = {
            if (targetState) infiniteRepeatable(
                animation = keyframes {
                    durationMillis = config.revealGlowDurationMs
                    0f at 0
                    1f at durationMillis / 2
                    0f at durationMillis
                },
                repeatMode = RepeatMode.Restart
            )
            else snap()
        },
        label = "glow_alpha"
    ) { revealed -> if (revealed) 1f else 0f }

    val scale by transition.animateFloat(
        transitionSpec = {
            if (targetState) spring(dampingRatio = Spring.DampingRatioMediumBouncy)
            else snap()
        },
        label = "reveal_scale"
    ) { revealed -> if (revealed) 1.15f else 1f }

    content(glowAlpha, scale)
}

@Composable
fun RouletteBoardHighlight(
    winningBets: List<String>,
    isHighlighted: Boolean,
    modifier: Modifier = Modifier,
    content: @Composable (betAlphas: Map<String, Float>) -> Unit
) {
    val alphas = remember(winningBets) {
        winningBets.associateWith { Animatable(0f) }
    }

    LaunchedEffect(isHighlighted) {
        if (isHighlighted) {
            alphas.values.forEachIndexed { index, animatable ->
                delay(index * 80L)
                animatable.animateTo(
                    targetValue = 1f,
                    animationSpec = tween(400, easing = FastOutSlowInEasing)
                )
            }
        } else {
            alphas.values.forEach { it.snapTo(0f) }
        }
    }

    val betAlphas = alphas.mapValues { it.value.value }
    content(betAlphas)
}

private fun lerp(start: Float, stop: Float, fraction: Float): Float {
    return start + (stop - start) * fraction
}
