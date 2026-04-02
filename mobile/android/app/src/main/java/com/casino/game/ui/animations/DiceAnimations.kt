package com.casino.game.ui.animations

import androidx.compose.animation.core.*
import androidx.compose.animation.*
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.unit.dp
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch

@Composable
fun DiceRollAnimation(
    finalValue: Int,
    isRolling: Boolean,
    onAnimationEnd: () -> Unit = {},
    modifier: Modifier = Modifier
) {
    var displayValue by remember { mutableIntStateOf(1) }
    val rotationX = remember { Animatable(0f) }
    val rotationY = remember { Animatable(0f) }
    val rotationZ = remember { Animatable(0f) }
    val scope = rememberCoroutineScope()

    LaunchedEffect(isRolling) {
        if (isRolling) {
            displayValue = 1
            launch {
                rotationX.animateTo(
                    targetValue = 720f + (0..360).random().toFloat(),
                    animationSpec = tween(1200, easing = FastOutSlowInEasing)
                )
            }
            launch {
                rotationY.animateTo(
                    targetValue = 540f + (0..360).random().toFloat(),
                    animationSpec = tween(1200, easing = FastOutSlowInEasing)
                )
            }
            launch {
                rotationZ.animateTo(
                    targetValue = 360f + (0..180).random().toFloat(),
                    animationSpec = tween(1200, easing = FastOutSlowInEasing)
                )
            }

            repeat(8) {
                delay(100)
                displayValue = (1..6).random()
            }
            displayValue = finalValue
            delay(200)
            onAnimationEnd()
        }
    }

    Box(
        modifier = modifier.size(80.dp),
        contentAlignment = Alignment.Center
    ) {
        DiceFace(
            value = displayValue,
            modifier = Modifier.graphicsLayer {
                this.rotationX = rotationX.value % 360f
                this.rotationY = rotationY.value % 360f
                this.rotationZ = rotationZ.value % 360f
                cameraDistance = 12f * density
            }
        )
    }
}

@Composable
fun DiceShuffleAnimation(
    diceValues: List<Int>,
    isShuffling: Boolean,
    onAnimationEnd: () -> Unit = {},
    modifier: Modifier = Modifier
) {
    var showFinal by remember { mutableStateOf(false) }

    LaunchedEffect(isShuffling) {
        if (isShuffling) {
            showFinal = false
            delay(1500)
            showFinal = true
            delay(300)
            onAnimationEnd()
        }
    }

    Row(
        modifier = modifier.padding(16.dp),
        horizontalArrangement = Arrangement.spacedBy(12.dp)
    ) {
        diceValues.forEachIndexed { index, value ->
            if (showFinal) {
                DiceRevealAnimation(value = value, delay = index * 100)
            } else {
                val infiniteTransition = rememberInfiniteTransition(label = "dice_$index")
                val rotation by infiniteTransition.animateFloat(
                    initialValue = 0f,
                    targetValue = 360f,
                    animationSpec = infiniteRepeatable(
                        animation = tween(300 + index * 50, easing = LinearEasing),
                        repeatMode = RepeatMode.Restart
                    ),
                    label = "rot_$index"
                )
                DiceFace(
                    value = (1..6).random(),
                    modifier = Modifier.graphicsLayer {
                        this.rotationZ = rotation
                        this.rotationX = rotation * 0.7f
                    }
                )
            }
        }
    }
}

@Composable
fun DiceRevealAnimation(
    value: Int,
    delay: Int = 0,
    modifier: Modifier = Modifier
) {
    var revealed by remember { mutableStateOf(false) }

    LaunchedEffect(Unit) {
        delay(delay.toLong())
        revealed = true
    }

    val scale by animateFloatAsState(
        targetValue = if (revealed) 1f else 0.3f,
        animationSpec = spring(
            dampingRatio = Spring.DampingRatioMediumBouncy,
            stiffness = Spring.StiffnessMedium
        ),
        label = "reveal_scale"
    )

    val pulseScale = if (revealed) {
        val infiniteTransition = rememberInfiniteTransition(label = "pulse")
        infiniteTransition.animateFloat(
            initialValue = 1f,
            targetValue = 1.1f,
            animationSpec = infiniteRepeatable(
                animation = tween(400),
                repeatMode = RepeatMode.Reverse
            ),
            label = "pulse_anim"
        ).value
    } else 1f

    Box(
        modifier = modifier.graphicsLayer {
            scaleX = scale * pulseScale
            scaleY = scale * pulseScale
        }
    ) {
        DiceFace(value = value)
    }
}

@Composable
fun DiceFace(
    value: Int,
    modifier: Modifier = Modifier
) {
    Box(
        modifier = modifier
            .size(64.dp)
            .clip(RoundedCornerShape(12.dp))
            .background(Color.White)
            .padding(8.dp),
        contentAlignment = Alignment.Center
    ) {
        DiceDots(value = value)
    }
}

@Composable
fun DiceDots(value: Int) {
    Box(modifier = Modifier.fillMaxSize()) {
        val dotSize = 10.dp
        val dotColor = Color.Black

        @Composable
        fun Dot(modifier: Modifier) {
            Box(
                modifier = modifier.size(dotSize).clip(CircleShape).background(dotColor)
            )
        }

        if (value in listOf(1, 3, 5)) {
            Dot(Modifier.align(Alignment.Center))
        }
        if (value in listOf(2, 3, 4, 5, 6)) {
            Dot(Modifier.align(Alignment.TopStart).offset(4.dp, 4.dp))
        }
        if (value in listOf(2, 3, 4, 5, 6)) {
            Dot(Modifier.align(Alignment.BottomEnd).offset((-4).dp, (-4).dp))
        }
        if (value in listOf(4, 5, 6)) {
            Dot(Modifier.align(Alignment.TopEnd).offset((-4).dp, 4.dp))
        }
        if (value in listOf(4, 5, 6)) {
            Dot(Modifier.align(Alignment.BottomStart).offset(4.dp, (-4).dp))
        }
        if (value == 6) {
            Dot(Modifier.align(Alignment.CenterStart).offset(4.dp, 0.dp))
            Dot(Modifier.align(Alignment.CenterEnd).offset((-4).dp, 0.dp))
        }
    }
}
