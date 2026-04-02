package com.casino.game.ui.animations

import androidx.compose.animation.core.*
import androidx.compose.animation.*
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.scale
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch

@Composable
fun WinAnimation(
    amount: String,
    isVisible: Boolean,
    onAnimationEnd: () -> Unit = {},
    modifier: Modifier = Modifier
) {
    var showConfetti by remember { mutableStateOf(false) }
    var counterValue by remember { mutableFloatStateOf(0f) }

    val glowAlpha by animateFloatAsState(
        targetValue = if (isVisible) 0.6f else 0f,
        animationSpec = infiniteRepeatable(tween(800), repeatMode = RepeatMode.Reverse),
        label = "glow"
    )

    val amountFloat = amount.filter { it.isDigit() }.toFloatOrNull() ?: 0f

    LaunchedEffect(isVisible) {
        if (isVisible) {
            showConfetti = true
            Animatable(0f).animateTo(amountFloat, tween(1500, easing = FastOutSlowInEasing)) {
                counterValue = this.value
            }
            delay(2000)
            showConfetti = false
            onAnimationEnd()
        }
    }

    if (isVisible) {
        Box(modifier = modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
            if (showConfetti) ConfettiParticle(isActive = true, Modifier.fillMaxSize())
            Column(
                horizontalAlignment = Alignment.CenterHorizontally,
                modifier = Modifier
                    .clip(RoundedCornerShape(16.dp))
                    .background(Color(0xFF1B5E20).copy(alpha = glowAlpha + 0.4f))
                    .padding(32.dp)
            ) {
                Text("YOU WIN!", color = Color(0xFFFFD700), fontSize = 28.sp, fontWeight = FontWeight.Bold)
                Spacer(Modifier.height(8.dp))
                Text("$${counterValue.toInt()}", color = Color.White, fontSize = 36.sp, fontWeight = FontWeight.Bold)
            }
        }
    }
}

@Composable
fun LoseAnimation(
    isVisible: Boolean,
    onAnimationEnd: () -> Unit = {},
    modifier: Modifier = Modifier
) {
    val shakeOffset = remember { Animatable(0f) }
    val flashAlpha = remember { Animatable(0f) }

    LaunchedEffect(isVisible) {
        if (isVisible) {
            launch {
                repeat(6) {
                    shakeOffset.animateTo(10f, tween(50))
                    shakeOffset.animateTo(-10f, tween(50))
                }
                shakeOffset.animateTo(0f, tween(50))
            }
            launch {
                flashAlpha.animateTo(0.3f, tween(100))
                flashAlpha.animateTo(0f, tween(400))
            }
            delay(1000)
            onAnimationEnd()
        }
    }

    if (isVisible) {
        Box(
            modifier = modifier
                .fillMaxSize()
                .graphicsLayer { translationX = shakeOffset.value }
                .background(Color.Red.copy(alpha = flashAlpha.value)),
            contentAlignment = Alignment.Center
        ) {
            val textAlpha by animateFloatAsState(1f, tween(500), label = "tf")
            Text(
                "BETTER LUCK NEXT TIME",
                color = Color.White.copy(alpha = textAlpha),
                fontSize = 20.sp,
                fontWeight = FontWeight.Bold
            )
        }
    }
}

@Composable
fun PushAnimation(
    isVisible: Boolean,
    onAnimationEnd: () -> Unit = {},
    modifier: Modifier = Modifier
) {
    val neutralFlash by animateColorAsState(
        targetValue = if (isVisible) Color(0xFF424242) else Color.Transparent,
        animationSpec = infiniteRepeatable(tween(600), repeatMode = RepeatMode.Reverse),
        label = "nf"
    )

    LaunchedEffect(isVisible) {
        if (isVisible) { delay(1500); onAnimationEnd() }
    }

    AnimatedVisibility(visible = isVisible, enter = fadeIn() + scaleIn(), exit = fadeOut() + scaleOut()) {
        Box(modifier = modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
            Box(
                modifier = Modifier
                    .clip(RoundedCornerShape(12.dp))
                    .background(neutralFlash.copy(alpha = 0.3f))
                    .padding(horizontal = 40.dp, vertical = 20.dp),
                contentAlignment = Alignment.Center
            ) {
                Text("PUSH", color = Color.White, fontSize = 24.sp, fontWeight = FontWeight.Bold)
            }
        }
    }
}

@Composable
fun JackpotAnimation(
    amount: String,
    isVisible: Boolean,
    onAnimationEnd: () -> Unit = {},
    modifier: Modifier = Modifier
) {
    var showParticles by remember { mutableStateOf(false) }

    LaunchedEffect(isVisible) {
        if (isVisible) {
            showParticles = true
            delay(3000)
            showParticles = false
            onAnimationEnd()
        }
    }

    val bannerScale by animateFloatAsState(
        targetValue = if (isVisible) 1f else 0f,
        animationSpec = spring(Spring.DampingRatioMediumBouncy, Spring.StiffnessLow),
        label = "bs"
    )

    if (isVisible) {
        Box(modifier = modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
            if (showParticles) GoldParticle(isActive = true, Modifier.fillMaxSize())
            val flashColor by animateColorAsState(
                Color(0xFFFFD700).copy(alpha = 0.15f),
                infiniteRepeatable(tween(300), repeatMode = RepeatMode.Reverse),
                label = "fc"
            )
            Box(Modifier.fillMaxSize().background(flashColor))
            Column(horizontalAlignment = Alignment.CenterHorizontally, modifier = Modifier.scale(bannerScale)) {
                Text("★ JACKPOT ★", color = Color(0xFFFFD700), fontSize = 40.sp, fontWeight = FontWeight.Black)
                Spacer(Modifier.height(12.dp))
                Text(amount, color = Color.White, fontSize = 48.sp, fontWeight = FontWeight.Bold)
            }
        }
    }
}
