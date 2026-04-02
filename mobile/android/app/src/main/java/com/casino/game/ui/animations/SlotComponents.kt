package com.casino.game.ui.animations

import androidx.compose.animation.core.*
import androidx.compose.foundation.Canvas
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.itemsIndexed
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.Shadow
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import kotlin.math.abs

val SLOT_SYMBOLS = listOf("🍒", "🍋", "7️⃣", "💎", "⭐", "🍊")

@Composable
fun SlotSymbol(
    symbol: String,
    modifier: Modifier = Modifier,
    scale: Float = 1f,
    alpha: Float = 1f,
    rotation: Float = 0f
) {
    Box(
        modifier = modifier
            .size(64.dp)
            .graphicsLayer {
                scaleX = scale; scaleY = scale
                rotationZ = rotation
                this.alpha = alpha
            },
        contentAlignment = Alignment.Center
    ) {
        Text(
            text = symbol,
            fontSize = 36.sp,
            textAlign = TextAlign.Center,
            style = TextStyle(
                shadow = Shadow(
                    color = Color.Black.copy(alpha = 0.3f),
                    offset = Offset(2f, 2f),
                    blurRadius = 4f
                )
            )
        )
    }
}

@Composable
fun SlotReel(
    symbols: List<String>,
    currentOffset: Float,
    blurAlpha: Float,
    modifier: Modifier = Modifier
) {
    val visibleCount = 3
    val symbolHeight = 72

    Box(
        modifier = modifier
            .width(72.dp)
            .height((symbolHeight * visibleCount).dp)
            .clip(RoundedCornerShape(8.dp))
            .background(Color(0xFF1A1A2E))
            .border(2.dp, Color(0xFF3D3D5C), RoundedCornerShape(8.dp))
    ) {
        LazyColumn(
            modifier = Modifier.fillMaxSize(),
            horizontalAlignment = Alignment.CenterHorizontally,
            userScrollEnabled = false
        ) {
            val totalSymbols = symbols.size
            val baseIndex = ((currentOffset % totalSymbols + totalSymbols) % totalSymbols).toInt()
            val fraction = currentOffset - currentOffset.toInt()

            items(visibleCount + 2) { displayIndex ->
                val symbolIndex = (baseIndex + displayIndex - 1 + totalSymbols) % totalSymbols
                val offsetY = if (displayIndex == 1) fraction else 0f

                Box(
                    modifier = Modifier
                        .height(symbolHeight.dp)
                        .fillMaxWidth()
                        .graphicsLayer { translationY = offsetY * symbolHeight },
                    contentAlignment = Alignment.Center
                ) {
                    SlotSymbol(
                        symbol = symbols[symbolIndex],
                        alpha = when {
                            abs(blurAlpha) > 0.5f -> 0.5f + (1f - abs(blurAlpha)) * 0.5f
                            else -> 1f
                        },
                        scale = if (blurAlpha > 0.5f) 0.9f else 1f
                    )
                }
            }
        }

        Box(
            modifier = Modifier
                .fillMaxSize()
                .background(
                    Brush.verticalGradient(
                        0f to Color(0xFF1A1A2E).copy(alpha = 0.7f),
                        0.2f to Color.Transparent,
                        0.8f to Color.Transparent,
                        1f to Color(0xFF1A1A2E).copy(alpha = 0.7f)
                    )
                )
        )
    }
}

@Composable
fun SlotPayline(
    isVisible: Boolean,
    glowIntensity: Float,
    modifier: Modifier = Modifier,
    color: Color = Color(0xFFFFD700)
) {
    Canvas(
        modifier = modifier
            .fillMaxWidth()
            .height(4.dp)
    ) {
        drawLine(
            color = color.copy(alpha = 0.3f + glowIntensity * 0.7f),
            start = Offset(0f, size.height / 2),
            end = Offset(size.width, size.height / 2),
            strokeWidth = size.height * (1f + glowIntensity * 0.5f)
        )

        if (isVisible) {
            drawLine(
                color = Color.White.copy(alpha = glowIntensity * 0.6f),
                start = Offset(0f, size.height / 2),
                end = Offset(size.width, size.height / 2),
                strokeWidth = size.height * 0.5f
            )
        }
    }
}

@Composable
fun SlotMachine(
    reelContents: List<List<String>>,
    reelOffsets: List<Float>,
    reelBlurs: List<Float>,
    activePaylines: List<Pair<Int, Float>>,
    modifier: Modifier = Modifier
) {
    Column(
        modifier = modifier
            .width(IntrinsicSize.Max)
            .background(Color(0xFF0D0D1A), RoundedCornerShape(16.dp))
            .border(3.dp, Color(0xFFB8860B), RoundedCornerShape(16.dp))
            .padding(12.dp),
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        Text(
            "★ SLOT MACHINE ★",
            color = Color(0xFFFFD700),
            fontSize = 18.sp,
            fontWeight = FontWeight.ExtraBold,
            modifier = Modifier.padding(bottom = 8.dp)
        )

        Box {
            Row(
                horizontalArrangement = Arrangement.spacedBy(6.dp),
                modifier = Modifier.padding(vertical = 4.dp)
            ) {
                reelContents.forEachIndexed { index, symbols ->
                    SlotReel(
                        symbols = symbols,
                        currentOffset = reelOffsets.getOrElse(index) { 0f },
                        blurAlpha = reelBlurs.getOrElse(index) { 0f }
                    )
                }
            }

            activePaylines.forEach { (lineIndex, glow) ->
                val yOffset = when (lineIndex) {
                    0 -> 0.33f; 1 -> 0.5f; 2 -> 0.67f
                    else -> 0.5f
                }
                val lineColor = when (lineIndex) {
                    0 -> Color(0xFFFFD700)
                    1 -> Color(0xFFFF4444)
                    2 -> Color(0xFF44FF44)
                    else -> Color.White
                }

                Box(
                    modifier = Modifier
                        .fillMaxWidth()
                        .align(Alignment.CenterStart)
                        .offset(y = ((yOffset - 0.5f) * 216).dp)
                ) {
                    SlotPayline(
                        isVisible = glow > 0f,
                        glowIntensity = glow,
                        color = lineColor
                    )
                }
            }
        }

        Spacer(Modifier.height(8.dp))

        Row(
            modifier = Modifier.fillMaxWidth(),
            horizontalArrangement = Arrangement.SpaceEvenly
        ) {
            listOf("BET", "WIN", "BALANCE").forEach { label ->
                Column(horizontalAlignment = Alignment.CenterHorizontally) {
                    Text(label, color = Color.Gray, fontSize = 10.sp)
                    Text("0", color = Color.White, fontSize = 14.sp, fontWeight = FontWeight.Bold)
                }
            }
        }
    }
}
