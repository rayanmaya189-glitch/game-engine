package com.casino.game.ui.animations

import androidx.compose.animation.core.*
import androidx.compose.animation.*
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.*
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

data class CardData(
    val suit: String,
    val value: String,
    val color: Color = if (suit in listOf("♥", "♦")) Color.Red else Color.Black
)

@Composable
fun CardShuffleAnimation(
    cards: List<CardData>,
    onShuffleComplete: () -> Unit = {}
) {
    var shuffleCount by remember { mutableIntStateOf(0) }
    var isAnimating by remember { mutableStateOf(true) }
    val maxShuffles = 3

    val fanAngle by animateFloatAsState(
        targetValue = if (isAnimating) 15f else 0f,
        animationSpec = tween(durationMillis = 500, easing = FastOutSlowInEasing),
        label = "fan"
    )

    LaunchedEffect(shuffleCount) {
        if (shuffleCount < maxShuffles) {
            isAnimating = true
            delay(600)
            isAnimating = false
            delay(400)
            shuffleCount++
        } else {
            onShuffleComplete()
        }
    }

    Box(
        modifier = Modifier.fillMaxWidth().height(200.dp),
        contentAlignment = Alignment.Center
    ) {
        cards.take(5).forEachIndexed { index, card ->
            val offset by animateFloatAsState(
                targetValue = if (isAnimating) (index - 2) * fanAngle else 0f,
                animationSpec = spring(dampingRatio = Spring.DampingRatioMediumBouncy),
                label = "offset_$index"
            )

            CardFace(
                card = card,
                modifier = Modifier
                    .graphicsLayer {
                        rotationZ = offset
                        translationX = (index - 2) * if (isAnimating) 30f else 0f
                    }
                    .offset(x = (index - 2).dp * 4)
            )
        }
    }
}

@Composable
fun CardDealAnimation(
    card: CardData,
    targetIndex: Int,
    isDealing: Boolean,
    onAnimationEnd: () -> Unit = {}
) {
    var isVisible by remember { mutableStateOf(false) }

    LaunchedEffect(isDealing) {
        if (isDealing) {
            delay(targetIndex * 200L)
            isVisible = true
            delay(600)
            onAnimationEnd()
        }
    }

    AnimatedVisibility(
        visible = isVisible,
        enter = slideInHorizontally(
            initialOffsetX = { -it * 2 },
            animationSpec = tween(500, easing = FastOutSlowInEasing)
        ) + fadeIn(),
        exit = fadeOut()
    ) {
        var flipped by remember { mutableStateOf(false) }
        val rotation by animateFloatAsState(
            targetValue = if (flipped) 180f else 0f,
            animationSpec = tween(400, easing = LinearEasing),
            label = "flip"
        )

        LaunchedEffect(Unit) {
            delay(200)
            flipped = true
        }

        CardFace(
            card = card,
            modifier = Modifier.graphicsLayer {
                rotationY = rotation
                cameraDistance = 12f * density
            }
        )
    }
}

@Composable
fun CardFlipAnimation(
    card: CardData,
    isFlipped: Boolean,
    modifier: Modifier = Modifier
) {
    val rotation by animateFloatAsState(
        targetValue = if (isFlipped) 180f else 0f,
        animationSpec = tween(600, easing = FastOutSlowInEasing),
        label = "cardFlip"
    )

    val showFront = rotation > 90f

    Box(
        modifier = modifier
            .size(width = 80.dp, height = 120.dp)
            .graphicsLayer {
                rotationX = rotation
                cameraDistance = 12f * density
            },
        contentAlignment = Alignment.Center
    ) {
        if (showFront) {
            CardFace(card = card)
        } else {
            CardBack()
        }
    }
}

@Composable
fun CardFanAnimation(
    cards: List<CardData>,
    isFanned: Boolean,
    modifier: Modifier = Modifier
) {
    val maxAngle = 60f
    val cardCount = cards.size.coerceAtMost(7)
    val spread = if (cardCount > 1) maxAngle / (cardCount - 1) else 0f

    Box(
        modifier = modifier.fillMaxWidth().height(200.dp),
        contentAlignment = Alignment.BottomCenter
    ) {
        cards.take(cardCount).forEachIndexed { index, card ->
            val angle by animateFloatAsState(
                targetValue = if (isFanned) (index - cardCount / 2) * spread else 0f,
                animationSpec = spring(
                    dampingRatio = Spring.DampingRatioLowBouncy,
                    stiffness = Spring.StiffnessLow
                ),
                label = "fan_angle_$index"
            )

            val elevation by animateDpAsState(
                targetValue = if (isFanned) (cardCount - index).dp else 0.dp,
                animationSpec = tween(300),
                label = "elevation_$index"
            )

            CardFace(
                card = card,
                modifier = Modifier
                    .graphicsLayer {
                        rotationZ = angle
                        transformOrigin = androidx.compose.ui.graphics.TransformOrigin(0.5f, 1f)
                    }
                    .offset(y = -elevation)
            )
        }
    }
}

@Composable
fun CardFace(
    card: CardData,
    modifier: Modifier = Modifier
) {
    Box(
        modifier = modifier
            .size(width = 70.dp, height = 100.dp)
            .clip(RoundedCornerShape(8.dp))
            .background(Color.White)
            .border(1.dp, Color.Gray, RoundedCornerShape(8.dp))
            .padding(6.dp),
        contentAlignment = Alignment.TopStart
    ) {
        Column(horizontalAlignment = Alignment.CenterHorizontally) {
            Text(
                text = card.value,
                color = card.color,
                fontSize = 14.sp,
                fontWeight = FontWeight.Bold
            )
            Text(
                text = card.suit,
                color = card.color,
                fontSize = 16.sp
            )
        }

        Text(
            text = card.suit,
            color = card.color.copy(alpha = 0.3f),
            fontSize = 36.sp,
            modifier = Modifier.align(Alignment.Center)
        )
    }
}

@Composable
fun CardBack(modifier: Modifier = Modifier) {
    Box(
        modifier = modifier
            .size(width = 70.dp, height = 100.dp)
            .clip(RoundedCornerShape(8.dp))
            .background(Color(0xFF1A237E))
            .border(2.dp, Color(0xFFFFD700), RoundedCornerShape(8.dp)),
        contentAlignment = Alignment.Center
    ) {
        Text(
            text = "♠♣",
            color = Color(0xFFFFD700).copy(alpha = 0.5f),
            fontSize = 24.sp
        )
    }
}
