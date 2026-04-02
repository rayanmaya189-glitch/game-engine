package com.casino.game.ui.animations

import androidx.compose.animation.core.*
import androidx.compose.foundation.Canvas
import androidx.compose.foundation.layout.*
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.drawscope.rotate
import androidx.compose.ui.unit.dp
import kotlinx.coroutines.delay

data class ParticleState(
    val x: Float,
    val y: Float,
    val vx: Float,
    val vy: Float,
    val color: Color,
    val size: Float,
    val rotation: Float = 0f,
    val rotationSpeed: Float = 0f
)

@Composable
fun ConfettiParticle(
    isActive: Boolean,
    modifier: Modifier = Modifier
) {
    val particles = remember {
        List(40) {
            ParticleState(
                x = (0f..1f).random(),
                y = (-0.3f..-0.1f).random(),
                vx = (-0.01f..0.01f).random(),
                vy = (0.008f..0.02f).random(),
                color = listOf(
                    Color(0xFFE53935), Color(0xFF1E88E5), Color(0xFF43A047),
                    Color(0xFFFDD835), Color(0xFF8E24AA), Color(0xFFFF7043)
                ).random(),
                size = (4f..10f).random(),
                rotation = (0f..360f).random(),
                rotationSpeed = (-5f..5f).random()
            )
        }
    }
    var time by remember { mutableFloatStateOf(0f) }

    LaunchedEffect(isActive) {
        if (isActive) {
            time = 0f
            while (true) {
                delay(16)
                time += 0.016f
            }
        }
    }

    if (isActive) {
        Canvas(modifier = modifier.fillMaxSize()) {
            particles.forEach { p ->
                val px = (p.x + p.vx * time * 60f + 0.005f * time * time * 60f) * size.width
                val py = (p.y + p.vy * time * 60f + 0.5f * 0.0005f * time * time * 3600f) * size.height
                val rot = p.rotation + p.rotationSpeed * time * 60f

                if (py < size.height + 20f && px > -20f && px < size.width + 20f) {
                    rotate(rot, Offset(px, py)) {
                        drawRect(
                            color = p.color,
                            topLeft = Offset(px - p.size / 2, py - p.size / 4),
                            size = Size(p.size, p.size / 2)
                        )
                    }
                }
            }
        }
    }
}

@Composable
fun GoldParticle(
    isActive: Boolean,
    modifier: Modifier = Modifier
) {
    val particles = remember {
        List(25) {
            ParticleState(
                x = (0f..1f).random(),
                y = (0.8f..1.2f).random(),
                vx = (-0.005f..0.005f).random(),
                vy = (-0.015f..-0.005f).random(),
                color = Color(0xFFFFD700),
                size = (3f..8f).random()
            )
        }
    }
    var time by remember { mutableFloatStateOf(0f) }

    LaunchedEffect(isActive) {
        if (isActive) {
            time = 0f
            while (true) {
                delay(16)
                time += 0.016f
            }
        }
    }

    if (isActive) {
        Canvas(modifier = modifier.fillMaxSize()) {
            particles.forEach { p ->
                val px = (p.x + p.vx * time * 60f + 0.01f * kotlin.math.sin(time * 4 + p.x * 20)) * size.width
                val py = (p.y + p.vy * time * 60f) * size.height
                val alpha = (1f - (time * 0.5f)).coerceIn(0f, 1f)

                if (py > -20f) {
                    drawCircle(
                        color = p.color.copy(alpha = alpha * 0.8f),
                        radius = p.size,
                        center = Offset(px, py)
                    )
                    drawCircle(
                        color = Color.White.copy(alpha = alpha * 0.4f),
                        radius = p.size * 0.4f,
                        center = Offset(px, py)
                    )
                }
            }
        }
    }
}

@Composable
fun CoinParticle(
    isActive: Boolean,
    modifier: Modifier = Modifier
) {
    val particles = remember {
        List(15) {
            ParticleState(
                x = (0.3f..0.7f).random(),
                y = (-0.2f..0f).random(),
                vx = (-0.01f..0.01f).random(),
                vy = (0.01f..0.025f).random(),
                color = Color(0xFFFFD700),
                size = (12f..20f).random(),
                rotation = (0f..360f).random(),
                rotationSpeed = (3f..8f).random()
            )
        }
    }
    var time by remember { mutableFloatStateOf(0f) }

    LaunchedEffect(isActive) {
        if (isActive) {
            time = 0f
            while (true) {
                delay(16)
                time += 0.016f
            }
        }
    }

    if (isActive) {
        Canvas(modifier = modifier.fillMaxSize()) {
            particles.forEach { p ->
                val px = (p.x + p.vx * time * 60f) * size.width
                val py = (p.y + p.vy * time * 60f) * size.height
                val rot = p.rotation + p.rotationSpeed * time * 60f
                val scaleX = kotlin.math.cos(rot * 0.01745f)

                if (py < size.height + 20f) {
                    drawOval(
                        color = Color(0xFFB8860B),
                        topLeft = Offset(px - p.size * kotlin.math.abs(scaleX) / 2, py - p.size / 2),
                        size = Size(p.size * kotlin.math.abs(scaleX), p.size)
                    )
                    drawOval(
                        color = p.color,
                        topLeft = Offset(
                            px - p.size * kotlin.math.abs(scaleX) / 2 + 2,
                            py - p.size / 2 + 2
                        ),
                        size = Size(
                            (p.size - 4) * kotlin.math.abs(scaleX),
                            p.size - 4
                        )
                    )
                }
            }
        }
    }
}

private operator fun ClosedRange<Float>.random() =
    start + kotlin.random.Random.nextFloat() * (endInclusive - start)
