package com.casino.game.ui.animations

import androidx.compose.animation.core.*
import androidx.compose.foundation.layout.*
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.unit.dp
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch

enum class ReelState { Idle, Spinning, Stopping, Stopped }

data class SlotAnimationConfig(
    val spinDurationBaseMs: Int = 1200,
    val staggerDelayMs: Int = 300,
    val symbolBounceDurationMs: Int = 400,
    val paylineGlowDurationMs: Int = 600,
    val bigWinDurationMs: Int = 2000,
    val freeSpinsGlowDurationMs: Int = 1000,
    val blurMaxAlpha: Float = 0.6f,
)

@Composable
fun SlotReelSpin(
    reelState: ReelState,
    symbolCount: Int,
    targetIndex: Int,
    modifier: Modifier = Modifier,
    config: SlotAnimationConfig = SlotAnimationConfig(),
    content: @Composable (offset: Float, blur: Float) -> Unit
) {
    val animOffset = remember { Animatable(0f) }

    LaunchedEffect(reelState) {
        when (reelState) {
            ReelState.Spinning -> {
                animOffset.snapTo(0f)
                animOffset.animateTo(
                    targetValue = symbolCount * 3f,
                    animationSpec = infiniteRepeatable(
                        animation = tween(300, easing = LinearEasing),
                        repeatMode = RepeatMode.Restart
                    )
                )
            }
            ReelState.Stopping -> {
                animOffset.stop()
                animOffset.animateTo(
                    targetValue = targetIndex.toFloat(),
                    animationSpec = tween(
                        durationMillis = config.spinDurationBaseMs,
                        easing = CubicBezierEasing(0.2f, 0.9f, 0.1f, 1f)
                    )
                )
            }
            ReelState.Stopped -> { /* already stopped */ }
            ReelState.Idle -> { animOffset.snapTo(targetIndex.toFloat()) }
        }
    }

    val blur = when (reelState) {
        ReelState.Spinning -> config.blurMaxAlpha
        ReelState.Stopping -> {
            val progress = (animOffset.value - targetIndex).coerceIn(0f, symbolCount.toFloat())
            config.blurMaxAlpha * (progress / symbolCount)
        }
        else -> 0f
    }

    content(animOffset.value, blur)
}

@Composable
fun SlotMachineSpin(
    reelStates: List<ReelState>,
    reelTargets: List<Int>,
    symbolCount: Int,
    modifier: Modifier = Modifier,
    config: SlotAnimationConfig = SlotAnimationConfig(),
    onAllStopped: () -> Unit = {},
    content: @Composable (reelIndex: Int, offset: Float, blur: Float) -> Unit
) {
    val animatedOffsets = remember { List(reelStates.size) { Animatable(0f) } }
    val animatedBlurs = remember { List(reelStates.size) { Animatable(0f) } }

    LaunchedEffect(reelStates) {
        reelStates.forEachIndexed { index, state ->
            launch {
                when (state) {
                    ReelState.Spinning -> {
                        animatedOffsets[index].snapTo(0f)
                        animatedBlurs[index].animateTo(
                            1f, tween(200, easing = FastOutLinearInEasing)
                        )
                    }
                    ReelState.Stopping -> {
                        delay((index * config.staggerDelayMs).toLong())
                        animatedOffsets[index].animateTo(
                            targetValue = reelTargets.getOrElse(index) { 0 }.toFloat(),
                            animationSpec = tween(
                                config.spinDurationBaseMs + index * config.staggerDelayMs,
                                easing = CubicBezierEasing(0.15f, 0.85f, 0.2f, 1f)
                            )
                        )
                        animatedBlurs[index].animateTo(0f, tween(150))
                    }
                    ReelState.Stopped -> {}
                    ReelState.Idle -> {
                        animatedOffsets[index].snapTo(reelTargets.getOrElse(index) { 0 }.toFloat())
                        animatedBlurs[index].snapTo(0f)
                    }
                }
            }
        }
    }

    val allStopped = reelStates.all { it == ReelState.Stopped || it == ReelState.Idle }
    LaunchedEffect(allStopped) { if (allStopped) onAllStopped() }

    Row(modifier = modifier) {
        reelStates.forEachIndexed { index, _ ->
            content(index, animatedOffsets[index].value, animatedBlurs[index].value)
        }
    }
}

@Composable
fun SlotSymbolLand(
    isLanding: Boolean,
    modifier: Modifier = Modifier,
    config: SlotAnimationConfig = SlotAnimationConfig(),
    content: @Composable (scale: Float, rotation: Float) -> Unit
) {
    val transition = updateTransition(isLanding, label = "symbol_land")

    val scale by transition.animateFloat(
        transitionSpec = {
            if (targetState) spring(dampingRatio = 0.4f, stiffness = Spring.StiffnessMedium)
            else snap()
        },
        label = "land_scale"
    ) { landing -> if (landing) 1.1f else 1f }

    val rotation by transition.animateFloat(
        transitionSpec = {
            if (targetState) keyframes {
                durationMillis = config.symbolBounceDurationMs
                0f at 0
                8f at durationMillis / 4
                -5f at durationMillis / 2
                3f at durationMillis * 3 / 4
                0f at durationMillis
            }
            else snap()
        },
        label = "land_rotation"
    ) { landing -> if (landing) 1f else 0f }

    Box(modifier.graphicsLayer { scaleX = scale; scaleY = scale; rotationZ = rotation }) {
        content(scale, rotation)
    }
}

@Composable
fun SlotPaylineWin(
    isActive: Boolean,
    modifier: Modifier = Modifier,
    config: SlotAnimationConfig = SlotAnimationConfig(),
    content: @Composable (glowIntensity: Float) -> Unit
) {
    val glow by animateFloatAsState(
        targetValue = if (isActive) 1f else 0f,
        animationSpec = if (isActive) infiniteRepeatable(
            animation = keyframes {
                durationMillis = config.paylineGlowDurationMs
                0f at 0
                1f at durationMillis / 2
                0.3f at durationMillis
            },
            repeatMode = RepeatMode.Restart
        ) else snap(),
        label = "payline_glow"
    )

    content(glow)
}

@Composable
fun SlotBigWin(
    isActive: Boolean,
    modifier: Modifier = Modifier,
    config: SlotAnimationConfig = SlotAnimationConfig(),
    onCoinsBurst: () -> Unit = {},
    content: @Composable (shakeX: Float, shakeY: Float, coinCount: Int) -> Unit
) {
    var coinCount by remember { mutableIntStateOf(0) }

    val shakeX by animateFloatAsState(
        targetValue = if (isActive) 1f else 0f,
        animationSpec = if (isActive) infiniteRepeatable(
            animation = keyframes {
                durationMillis = 100
                -4f at 0; 4f at 25; -3f at 50; 3f at 75; 0f at 100
            },
            repeatMode = RepeatMode.Restart
        ) else snap(),
        label = "bigwin_shake_x"
    )

    val shakeY by animateFloatAsState(
        targetValue = if (isActive) 1f else 0f,
        animationSpec = if (isActive) infiniteRepeatable(
            animation = keyframes {
                durationMillis = 130
                3f at 0; -3f at 33; 2f at 66; 0f at 130
            },
            repeatMode = RepeatMode.Restart
        ) else snap(),
        label = "bigwin_shake_y"
    )

    LaunchedEffect(isActive) {
        if (isActive) {
            repeat(20) {
                coinCount++
                delay(config.bigWinDurationMs / 20)
            }
            onCoinsBurst()
        } else {
            coinCount = 0
        }
    }

    content(shakeX * 4f, shakeY * 3f, coinCount)
}

@Composable
fun SlotFreeSpins(
    isActive: Boolean,
    remainingSpins: Int,
    modifier: Modifier = Modifier,
    config: SlotAnimationConfig = SlotAnimationConfig(),
    content: @Composable (glowAlpha: Float, pulseScale: Float, count: Int) -> Unit
) {
    val transition = updateTransition(isActive, label = "free_spins")

    val glowAlpha by transition.animateFloat(
        transitionSpec = {
            if (targetState) infiniteRepeatable(
                animation = tween(config.freeSpinsGlowDurationMs, easing = FastOutSlowInEasing),
                repeatMode = RepeatMode.Reverse
            )
            else snap()
        },
        label = "fs_glow"
    ) { active -> if (active) 1f else 0f }

    val pulseScale by transition.animateFloat(
        transitionSpec = {
            if (targetState) infiniteRepeatable(
                animation = tween(800, easing = FastOutSlowInEasing),
                repeatMode = RepeatMode.Reverse
            )
            else snap()
        },
        label = "fs_pulse"
    ) { active -> if (active) 1.05f else 1f }

    content(glowAlpha, pulseScale, remainingSpins)
}
