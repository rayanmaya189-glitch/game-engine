package com.casino.game.ui.screens

import androidx.compose.animation.core.*
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.itemsIndexed
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.casino.game.data.model.*
import com.casino.game.data.remote.ApiService
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.delay
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

private val SYMBOLS = listOf("\uD83C\uDFB0", "\u2B50", "\uD83C\uDF52", "\uD83C\uDF4B", "\uD83D\uDC8E", "\uD83C\uDF1F", "\uD83C\uDF47")

data class GamePlayState(
    val gameId: String = "",
    val gameName: String = "Lucky Slots",
    val balance: Double = 1000.0,
    val betAmount: Double = 10.0,
    val winAmount: Double = 0.0,
    val reels: List<List<String>> = List(3) { List(5) { SYMBOLS.random() } },
    val isSpinning: Boolean = false,
    val totalWin: Double = 0.0,
    val isLoading: Boolean = false,
    val error: String? = null
)

@HiltViewModel
class GamePlayViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(GamePlayState())
    val state: StateFlow<GamePlayState> = _state.asStateFlow()

    fun loadGame(gameId: String) {
        _state.update { it.copy(gameId = gameId, isLoading = true) }
        viewModelScope.launch {
            try {
                val response = apiService.getGameDetails(gameId)
                if (response.isSuccessful && response.body() != null) {
                    val details = response.body()!!
                    _state.update {
                        it.copy(
                            gameName = details.name,
                            isLoading = false
                        )
                    }
                }
            } catch (e: Exception) {
                _state.update { it.copy(isLoading = false, error = e.message) }
            }
        }
    }

    fun setBet(amount: Double) {
        _state.update { it.copy(betAmount = amount) }
    }

    fun spin() {
        val current = _state.value
        if (current.isSpinning || current.balance < current.betAmount) return
        _state.update {
            it.copy(
                isSpinning = true,
                balance = it.balance - it.betAmount,
                winAmount = 0.0
            )
        }
        viewModelScope.launch {
            delay(2000)
            val resultReels = List(3) { List(5) { SYMBOLS.random() } }
            val win = calculateWin(resultReels, current.betAmount)
            _state.update {
                it.copy(
                    isSpinning = false,
                    reels = resultReels,
                    winAmount = win,
                    balance = it.balance + win,
                    totalWin = it.totalWin + win
                )
            }
        }
    }

    private fun calculateWin(reels: List<List<String>>, bet: Double): Double {
        val middleRow = reels.map { it[2] }
        val uniqueSymbols = middleRow.toSet()
        return when {
            uniqueSymbols.size == 1 -> bet * 50
            uniqueSymbols.size == 2 -> bet * 10
            middleRow.groupBy { it }.values.any { it.size >= 3 } -> bet * 5
            else -> 0.0
        }
    }
}

@Composable
fun GamePlayScreen(
    gameId: String,
    onBack: () -> Unit,
    viewModel: GamePlayViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()
    val betOptions = listOf(5.0, 10.0, 25.0, 50.0, 100.0)

    LaunchedEffect(gameId) { viewModel.loadGame(gameId) }

    Column(
        modifier = Modifier
            .fillMaxSize()
            .background(MaterialTheme.colorScheme.background)
    ) {
        TopAppBar(
            title = { Text(state.gameName) },
            navigationIcon = {
                IconButton(onClick = onBack) {
                    Icon(Icons.Default.ArrowBack, contentDescription = "Back")
                }
            },
            actions = {
                Text(
                    "$${String.format("%,.2f", state.balance)}",
                    modifier = Modifier.padding(end = 16.dp),
                    fontWeight = FontWeight.Bold
                )
            }
        )

        Spacer(modifier = Modifier.height(8.dp))

        Card(
            modifier = Modifier
                .fillMaxWidth()
                .padding(horizontal = 16.dp),
            colors = CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.primaryContainer)
        ) {
            Column(
                modifier = Modifier.padding(16.dp),
                horizontalAlignment = Alignment.CenterHorizontally
            ) {
                Text("Win", style = MaterialTheme.typography.bodySmall)
                Text(
                    "$${String.format("%,.2f", state.winAmount)}",
                    style = MaterialTheme.typography.headlineMedium,
                    fontWeight = FontWeight.Bold,
                    color = if (state.winAmount > 0) Color(0xFF4CAF50) else MaterialTheme.colorScheme.onPrimaryContainer
                )
            }
        }

        Spacer(modifier = Modifier.height(16.dp))

        SlotMachineReels(reels = state.reels, isSpinning = state.isSpinning)

        Spacer(modifier = Modifier.height(16.dp))

        Text(
            "Bet Amount",
            modifier = Modifier.padding(horizontal = 16.dp),
            style = MaterialTheme.typography.labelMedium
        )
        Spacer(modifier = Modifier.height(8.dp))

        Row(
            modifier = Modifier
                .fillMaxWidth()
                .padding(horizontal = 16.dp),
            horizontalArrangement = Arrangement.spacedBy(8.dp)
        ) {
            betOptions.forEach { bet ->
                FilterChip(
                    selected = state.betAmount == bet,
                    onClick = { viewModel.setBet(bet) },
                    label = { Text("$${bet.toInt()}") },
                    modifier = Modifier.weight(1f)
                )
            }
        }

        Spacer(modifier = Modifier.height(16.dp))

        Button(
            onClick = { viewModel.spin() },
            modifier = Modifier
                .fillMaxWidth()
                .padding(horizontal = 16.dp)
                .height(56.dp),
            enabled = !state.isSpinning && state.balance >= state.betAmount,
            shape = RoundedCornerShape(12.dp)
        ) {
            if (state.isSpinning) {
                CircularProgressIndicator(
                    modifier = Modifier.size(24.dp),
                    color = MaterialTheme.colorScheme.onPrimary
                )
            } else {
                Icon(Icons.Default.Casino, contentDescription = null)
                Spacer(modifier = Modifier.width(8.dp))
                Text("SPIN - $${state.betAmount.toInt()}")
            }
        }

        if (state.error != null) {
            Spacer(modifier = Modifier.height(8.dp))
            Text(
                state.error!!,
                color = MaterialTheme.colorScheme.error,
                modifier = Modifier.padding(horizontal = 16.dp)
            )
        }
    }
}

@Composable
private fun SlotMachineReels(reels: List<List<String>>, isSpinning: Boolean) {
    Card(
        modifier = Modifier
            .fillMaxWidth()
            .padding(horizontal = 16.dp),
        shape = RoundedCornerShape(12.dp)
    ) {
        Column(
            modifier = Modifier.padding(8.dp),
            verticalArrangement = Arrangement.spacedBy(4.dp)
        ) {
            for (row in 0 until 3) {
                Row(
                    modifier = Modifier.fillMaxWidth(),
                    horizontalArrangement = Arrangement.SpaceEvenly
                ) {
                    for (col in 0 until 5) {
                        val symbol = reels[col][row]
                        val animatedOffset by animateFloatAsState(
                            targetValue = if (isSpinning) 1000f else 0f,
                            animationSpec = if (isSpinning) infiniteRepeatable(
                                animation = tween(200, easing = LinearEasing),
                                repeatMode = RepeatMode.Restart
                            ) else tween(500)
                        )
                        Box(
                            modifier = Modifier
                                .size(52.dp)
                                .clip(RoundedCornerShape(8.dp))
                                .background(MaterialTheme.colorScheme.surface)
                                .border(
                                    1.dp,
                                    MaterialTheme.colorScheme.outline,
                                    RoundedCornerShape(8.dp)
                                ),
                            contentAlignment = Alignment.Center
                        ) {
                            Text(
                                text = symbol,
                                fontSize = if (isSpinning) (20 + animatedOffset % 4).sp else 24.sp,
                                textAlign = TextAlign.Center
                            )
                        }
                    }
                }
            }
        }
    }
}
