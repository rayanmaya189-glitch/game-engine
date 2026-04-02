package com.casino.game.ui.animations

import androidx.compose.animation.core.*
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material3.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.DpOffset
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import kotlinx.coroutines.delay

enum class ChipColor(val color: Color, val label: String) {
    RED(Color(0xFFD32F2F), "5"),
    BLUE(Color(0xFF1565C0), "10"),
    GREEN(Color(0xFF2E7D32), "25"),
    BLACK(Color(0xFF212121), "100")
}

@Composable
fun ChipBetAnimation(
    chipColor: ChipColor,
    startPosition: DpOffset,
    targetPosition: DpOffset,
    isActive: Boolean,
    onAnimationEnd: () -> Unit = {},
    modifier: Modifier = Modifier
) {
    var animating by remember { mutableStateOf(false) }

    LaunchedEffect(isActive) {
        if (isActive) {
            animating = true
            delay(600)
            onAnimationEnd()
        } else {
            animating = false
        }
    }

    val offsetX by animateDpAsState(
        targetValue = if (animating) targetPosition.x else startPosition.x,
        animationSpec = tween(500, easing = FastOutSlowInEasing),
        label = "chip_bet_x"
    )

    val offsetY by animateDpAsState(
        targetValue = if (animating) targetPosition.y else startPosition.y,
        animationSpec = tween(500, easing = FastOutSlowInEasing),
        label = "chip_bet_y"
    )

    val elevation by animateDpAsState(
        targetValue = if (animating) 16.dp else 4.dp,
        animationSpec = tween(250),
        label = "chip_elevation"
    )

    val scale by animateFloatAsState(
        targetValue = if (animating) 1.2f else 1f,
        animationSpec = tween(300),
        label = "chip_scale"
    )

    CasinoChip(
        chipColor = chipColor,
        modifier = modifier
            .offset(x = offsetX, y = offsetY)
            .graphicsLayer {
                scaleX = scale
                scaleY = scale
            }
            .shadow(elevation, CircleShape)
    )
}

@Composable
fun ChipCollectAnimation(
    chipColor: ChipColor,
    startPosition: DpOffset,
    targetPosition: DpOffset,
    isActive: Boolean,
    delayMs: Long = 0,
    onAnimationEnd: () -> Unit = {},
    modifier: Modifier = Modifier
) {
    var animating by remember { mutableStateOf(false) }

    LaunchedEffect(isActive) {
        if (isActive) {
            delay(delayMs)
            animating = true
            delay(700)
            onAnimationEnd()
        }
    }

    val offsetX by animateDpAsState(
        targetValue = if (animating) targetPosition.x else startPosition.x,
        animationSpec = spring(
            dampingRatio = Spring.DampingRatioNoBouncy,
            stiffness = Spring.StiffnessMedium
        ),
        label = "collect_x"
    )

    val offsetY by animateDpAsState(
        targetValue = if (animating) targetPosition.y else startPosition.y,
        animationSpec = spring(
            dampingRatio = Spring.DampingRatioNoBouncy,
            stiffness = Spring.StiffnessMedium
        ),
        label = "collect_y"
    )

    val alpha by animateFloatAsState(
        targetValue = if (animating) 0.7f else 1f,
        animationSpec = tween(500),
        label = "collect_alpha"
    )

    CasinoChip(
        chipColor = chipColor,
        modifier = modifier
            .offset(x = offsetX, y = offsetY)
            .graphicsLayer { this.alpha = alpha }
    )
}

@Composable
fun ChipStackAnimation(
    chipColors: List<ChipColor>,
    isStacking: Boolean,
    modifier: Modifier = Modifier
) {
    Box(modifier = modifier.size(60.dp, (chipColors.size * 8 + 40).dp)) {
        chipColors.forEachIndexed { index, color ->
            var visible by remember { mutableStateOf(false) }

            LaunchedEffect(isStacking) {
                if (isStacking) {
                    delay(index * 150L)
                    visible = true
                }
            }

            val bounceOffset by animateDpAsState(
                targetValue = if (visible) 0.dp else (-30).dp,
                animationSpec = spring(
                    dampingRatio = Spring.DampingRatioMediumBouncy,
                    stiffness = Spring.StiffnessMedium
                ),
                label = "stack_bounce_$index"
            )

            if (visible || isStacking) {
                CasinoChip(
                    chipColor = color,
                    modifier = Modifier
                        .align(Alignment.BottomCenter)
                        .offset(y = -(index * 8).dp + bounceOffset)
                        .graphicsLayer {
                            shadowElevation = 4f
                        }
                )
            }
        }
    }
}

@Composable
fun CasinoChip(
    chipColor: ChipColor,
    modifier: Modifier = Modifier,
    size: Dp = 48.dp
) {
    Box(
        modifier = modifier
            .size(size)
            .clip(CircleShape)
            .background(chipColor.color)
            .border(3.dp, Color.White, CircleShape)
            .border(1.dp, chipColor.color.copy(alpha = 0.5f), CircleShape),
        contentAlignment = Alignment.Center
    ) {
        repeat(8) { index ->
            val angle = index * 45f
            Box(
                modifier = Modifier
                    .fillMaxSize()
                    .graphicsLayer {
                        rotationZ = angle
                    }
            ) {
                Box(
                    modifier = Modifier
                        .align(Alignment.TopCenter)
                        .offset(y = 2.dp)
                        .size(6.dp, 3.dp)
                        .background(Color.White.copy(alpha = 0.6f))
                )
            }
        }
        Text(
            text = chipColor.label,
            color = Color.White,
            fontSize = 12.sp,
            fontWeight = FontWeight.Bold
        )
    }
}
