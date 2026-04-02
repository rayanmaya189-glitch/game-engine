package com.casino.game.ui.animations

import androidx.compose.animation.core.*
import androidx.compose.foundation.Canvas
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.grid.GridCells
import androidx.compose.foundation.lazy.grid.LazyVerticalGrid
import androidx.compose.foundation.lazy.grid.items
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material3.Text
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.graphics.graphicsLayer
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import kotlin.math.*

private val RED_NUMBERS = setOf(1,3,5,7,9,12,14,16,18,19,21,23,25,27,30,32,34,36)

fun pocketColor(number: Int): Color = when {
    number == 0 -> Color(0xFF1B7A1B)
    number in RED_NUMBERS -> Color(0xFFCC0000)
    else -> Color(0xFF1A1A1A)
}

@Composable
fun RouletteWheel(
    rotation: Float,
    modifier: Modifier = Modifier
) {
    Canvas(
        modifier = modifier
            .size(280.dp)
            .graphicsLayer { rotationZ = rotation }
    ) {
        val cx = size.width / 2f
        val cy = size.height / 2f
        val outerRadius = size.minDimension / 2f
        val innerRadius = outerRadius * 0.68f
        val segmentAngle = 360f / 37f

        drawCircle(
            color = Color(0xFF3D2B1F),
            radius = outerRadius,
            center = Offset(cx, cy)
        )

        for (i in 0 until 37) {
            val startAngle = i * segmentAngle - 90f
            val color = pocketColor(i)
            drawArc(
                color = color,
                startAngle = startAngle,
                sweepAngle = segmentAngle - 1f,
                useCenter = true,
                topLeft = Offset(cx - outerRadius, cy - outerRadius),
                size = Size(outerRadius * 2, outerRadius * 2)
            )
        }

        for (i in 0 until 37) {
            val angle = Math.toRadians((i * segmentAngle - 90f).toDouble())
            val midR = (outerRadius + innerRadius) / 2f
            val sx = cx + cos(angle).toFloat() * midR
            val sy = cy + sin(angle).toFloat() * midR
            drawCircle(
                color = Color.White.copy(alpha = 0.15f),
                radius = 1.5f,
                center = Offset(sx, sy)
            )
        }

        drawCircle(
            color = Color(0xFF2D1F14),
            radius = innerRadius,
            center = Offset(cx, cy)
        )

        drawCircle(
            color = Color(0xFF8B6914),
            radius = outerRadius,
            center = Offset(cx, cy),
            style = Stroke(width = 4f)
        )

        val hubRadius = outerRadius * 0.15f
        drawCircle(
            color = Color(0xFFB8860B),
            radius = hubRadius,
            center = Offset(cx, cy)
        )
        drawCircle(
            color = Color(0xFFDAA520),
            radius = hubRadius * 0.5f,
            center = Offset(cx, cy)
        )
    }
}

@Composable
fun RouletteBall(
    angle: Float,
    radiusFraction: Float,
    bounceOffset: Float,
    modifier: Modifier = Modifier
) {
    Canvas(modifier = modifier.size(280.dp)) {
        val cx = size.width / 2f
        val cy = size.height / 2f
        val r = size.minDimension / 2f * radiusFraction
        val bx = cx + cos(Math.toRadians(angle.toDouble())).toFloat() * r
        val by = cy + sin(Math.toRadians(angle.toDouble())).toFloat() * r

        drawCircle(
            color = Color.Gray.copy(alpha = 0.4f),
            radius = 9f,
            center = Offset(bx + 2f, by + bounceOffset + 2f)
        )
        drawCircle(
            color = Color.White,
            radius = 8f,
            center = Offset(bx, by + bounceOffset)
        )
        drawCircle(
            color = Color(0xFFE0E0E0),
            radius = 4f,
            center = Offset(bx - 2f, by + bounceOffset - 2f)
        )
    }
}

@Composable
fun RouletteBettingBoard(
    highlightedBets: Map<String, Float>,
    onBetSelected: (String) -> Unit,
    modifier: Modifier = Modifier
) {
    val numbers = (0..36).toList()
    val specialBets = listOf("1-12", "13-24", "25-36", "EVEN", "ODD", "RED", "BLACK", "1-18", "19-36")

    Column(modifier = modifier.padding(4.dp)) {
        Row(modifier = Modifier.fillMaxWidth()) {
            Box(
                modifier = Modifier
                    .size(36.dp)
                    .background(pocketColor(0))
                    .clickable { onBetSelected("0") }
                    .border(1.dp, Color.White.copy(alpha = 0.3f)),
                contentAlignment = Alignment.Center
            ) {
                Text("0", color = Color.White, fontSize = 11.sp, fontWeight = FontWeight.Bold)
            }
            Spacer(Modifier.width(2.dp))
            LazyVerticalGrid(
                columns = GridCells.Fixed(12),
                modifier = Modifier.weight(1f),
                horizontalArrangement = Arrangement.spacedBy(2.dp),
                verticalArrangement = Arrangement.spacedBy(2.dp)
            ) {
                items(numbers.filter { it > 0 }) { num ->
                    val bet = num.toString()
                    val alpha = highlightedBets[bet] ?: 0f
                    Box(
                        modifier = Modifier
                            .size(28.dp)
                            .background(pocketColor(num).copy(alpha = 0.6f + alpha * 0.4f))
                            .clickable { onBetSelected(bet) }
                            .border(1.dp, Color(0xFFFFD700).copy(alpha = alpha)),
                        contentAlignment = Alignment.Center
                    ) {
                        Text(
                            num.toString(),
                            color = Color.White,
                            fontSize = 9.sp,
                            fontWeight = if (alpha > 0f) FontWeight.ExtraBold else FontWeight.Normal
                        )
                    }
                }
            }
        }

        Spacer(Modifier.height(4.dp))

        Row(
            modifier = Modifier.fillMaxWidth(),
            horizontalArrangement = Arrangement.spacedBy(2.dp)
        ) {
            specialBets.forEach { bet ->
                val alpha = highlightedBets[bet] ?: 0f
                Box(
                    modifier = Modifier
                        .weight(1f)
                        .height(28.dp)
                        .background(Color(0xFF1B5E20).copy(alpha = 0.7f + alpha * 0.3f))
                        .clickable { onBetSelected(bet) }
                        .border(1.dp, Color(0xFFFFD700).copy(alpha = alpha)),
                    contentAlignment = Alignment.Center
                ) {
                    Text(
                        bet,
                        color = Color.White,
                        fontSize = 7.sp,
                        fontWeight = if (alpha > 0f) FontWeight.Bold else FontWeight.Normal,
                        textAlign = TextAlign.Center
                    )
                }
            }
        }
    }
}

@Composable
fun RouletteChip(
    value: Int,
    isPlaced: Boolean,
    modifier: Modifier = Modifier
) {
    val scale by animateFloatAsState(
        targetValue = if (isPlaced) 1f else 0f,
        animationSpec = spring(dampingRatio = 0.5f),
        label = "chip_scale"
    )

    val colors = mapOf(
        1 to Color(0xFF1565C0),
        5 to Color(0xFFCC0000),
        25 to Color(0xFF2E7D32),
        100 to Color(0xFF1A1A1A)
    )

    Box(
        modifier = modifier
            .size(24.dp)
            .graphicsLayer { scaleX = scale; scaleY = scale }
            .shadow(if (isPlaced) 4.dp else 0.dp, CircleShape)
            .clip(CircleShape)
            .background(colors[value] ?: Color.Gray),
        contentAlignment = Alignment.Center
    ) {
        Text(
            "$$value",
            color = Color.White,
            fontSize = 8.sp,
            fontWeight = FontWeight.Bold
        )
    }
}
