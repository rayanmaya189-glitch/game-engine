package com.casino.game.ui.animations

import androidx.compose.animation.*
import androidx.compose.animation.core.*
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import kotlinx.coroutines.delay

@Composable
fun GameEntryAnimation(
    visible: Boolean,
    content: @Composable AnimatedVisibilityScope.() -> Unit
) {
    AnimatedVisibility(
        visible = visible,
        enter = slideInVertically(
            initialOffsetY = { it },
            animationSpec = tween(500, easing = FastOutSlowInEasing)
        ) + fadeIn(animationSpec = tween(400)),
        exit = slideOutVertically(
            targetOffsetY = { it },
            animationSpec = tween(400, easing = FastOutLinearInEasing)
        ) + fadeOut(animationSpec = tween(300)),
        content = content
    )
}

@Composable
fun GameExitAnimation(
    visible: Boolean,
    content: @Composable AnimatedVisibilityScope.() -> Unit
) {
    AnimatedVisibility(
        visible = visible,
        enter = fadeIn() + scaleIn(initialScale = 0.95f),
        exit = slideOutHorizontally(
            targetOffsetX = { -it },
            animationSpec = tween(400, easing = FastOutLinearInEasing)
        ) + fadeOut(animationSpec = tween(300)),
        content = content
    )
}

@Composable
fun <T> GameCrossfade(
    currentState: T,
    modifier: Modifier = Modifier,
    content: @Composable (T) -> Unit
) {
    Crossfade(
        targetState = currentState,
        modifier = modifier,
        animationSpec = tween(400, easing = FastOutSlowInEasing),
        label = "game_crossfade",
        content = content
    )
}

@Composable
fun LoadingAnimation(
    modifier: Modifier = Modifier
) {
    val infiniteTransition = rememberInfiniteTransition(label = "loading")

    val rotation by infiniteTransition.animateFloat(
        initialValue = 0f,
        targetValue = 360f,
        animationSpec = infiniteRepeatable(
            animation = tween(1200, easing = LinearEasing),
            repeatMode = RepeatMode.Restart
        ),
        label = "chip_rotation"
    )

    val scale by infiniteTransition.animateFloat(
        initialValue = 0.8f,
        targetValue = 1.1f,
        animationSpec = infiniteRepeatable(
            animation = tween(600),
            repeatMode = RepeatMode.Reverse
        ),
        label = "chip_pulse"
    )

    val bounce by infiniteTransition.animateFloat(
        initialValue = 0f,
        targetValue = -20f,
        animationSpec = infiniteRepeatable(
            animation = tween(500, easing = FastOutSlowInEasing),
            repeatMode = RepeatMode.Reverse
        ),
        label = "chip_bounce"
    )

    Box(
        modifier = modifier.fillMaxSize(),
        contentAlignment = Alignment.Center
    ) {
        Column(
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            Box(
                modifier = Modifier
                    .offset(y = bounce.dp)
                    .graphicsLayer {
                        this.rotationZ = rotation
                        scaleX = scale
                        scaleY = scale
                    }
                    .size(56.dp)
                    .clip(CircleShape)
                    .background(Color(0xFF1565C0)),
                contentAlignment = Alignment.Center
            ) {
                repeat(8) { index ->
                    Box(
                        modifier = Modifier
                            .fillMaxSize()
                            .graphicsLayer { rotationZ = index * 45f }
                    ) {
                        Box(
                            modifier = Modifier
                                .align(Alignment.TopCenter)
                                .offset(y = 2.dp)
                                .size(7.dp, 3.dp)
                                .background(Color.White.copy(alpha = 0.6f))
                        )
                    }
                }
                Text(
                    text = "$",
                    color = Color.White,
                    fontSize = 18.sp,
                    fontWeight = FontWeight.Bold
                )
            }

            Spacer(Modifier.height(24.dp))

            val dotsAlpha by infiniteTransition.animateFloat(
                initialValue = 0.3f,
                targetValue = 1f,
                animationSpec = infiniteRepeatable(
                    animation = tween(1000),
                    repeatMode = RepeatMode.Reverse
                ),
                label = "dots_alpha"
            )

            Row(
                horizontalArrangement = Arrangement.spacedBy(8.dp),
                modifier = Modifier.graphicsLayer { alpha = dotsAlpha }
            ) {
                repeat(3) { index ->
                    val dotScale by infiniteTransition.animateFloat(
                        initialValue = 0.6f,
                        targetValue = 1f,
                        animationSpec = infiniteRepeatable(
                            animation = tween(400, delayMillis = index * 150),
                            repeatMode = RepeatMode.Reverse
                        ),
                        label = "dot_$index"
                    )
                    Box(
                        modifier = Modifier
                            .size(8.dp)
                            .graphicsLayer {
                                scaleX = dotScale
                                scaleY = dotScale
                            }
                            .clip(CircleShape)
                            .background(Color(0xFFFFD700))
                    )
                }
            }
        }
    }
}

@Composable
fun GameScreenTransition(
    targetState: Int,
    modifier: Modifier = Modifier,
    content: @Composable (Int) -> Unit
) {
    AnimatedContent(
        targetState = targetState,
        modifier = modifier,
        transitionSpec = {
            slideInHorizontally(
                initialOffsetX = { it },
                animationSpec = tween(400, easing = FastOutSlowInEasing)
            ) + fadeIn(tween(300)) togetherWith
                    slideOutHorizontally(
                        targetOffsetX = { -it },
                        animationSpec = tween(400, easing = FastOutLinearInEasing)
                    ) + fadeOut(tween(300))
        },
        label = "screen_transition",
        content = content
    )
}
