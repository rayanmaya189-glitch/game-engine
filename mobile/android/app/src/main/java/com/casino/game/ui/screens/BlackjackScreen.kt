package com.casino.game.ui.screens

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyRow
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.casino.game.data.model.*
import com.casino.game.data.remote.ApiService
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class BlackjackState(
    val playerCards: List<PlayingCard> = emptyList(),
    val dealerCards: List<PlayingCard> = emptyList(),
    val betAmount: Double = 10.0,
    val balance: Double = 0.0,
    val result: String? = null,
    val playerScore: Int = 0,
    val dealerScore: Int = 0,
    val canHit: Boolean = false,
    val canStand: Boolean = false,
    val canDouble: Boolean = false,
    val canSplit: Boolean = false,
    val gameActive: Boolean = false,
    val isLoading: Boolean = false
)

@HiltViewModel
class BlackjackViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(BlackjackState())
    val state: StateFlow<BlackjackState> = _state.asStateFlow()

    init { loadBalance() }

    private fun loadBalance() {
        viewModelScope.launch {
            try {
                val response = apiService.getBalance()
                if (response.isSuccessful) {
                    _state.update { it.copy(balance = response.body()?.balance ?: 0.0) }
                }
            } catch (_: Exception) {}
        }
    }

    fun setBet(amount: Double) { _state.update { it.copy(betAmount = amount) } }

    fun deal() {
        _state.update { it.copy(isLoading = true, result = null) }
        viewModelScope.launch {
            try {
                val response = apiService.blackjackDeal(mapOf("bet_amount" to _state.value.betAmount))
                if (response.isSuccessful) {
                    response.body()?.let { r ->
                        _state.update {
                            it.copy(playerCards = r.playerCards, dealerCards = r.dealerCards, playerScore = r.playerScore, dealerScore = r.dealerScore, canHit = true, canStand = true, canDouble = r.canDouble, canSplit = r.canSplit, gameActive = true, isLoading = false, result = r.result)
                        }
                    }
                }
            } catch (_: Exception) { _state.update { it.copy(isLoading = false) } }
        }
    }

    fun hit() { playAction("hit") }
    fun stand() { playAction("stand") }
    fun doubleDown() { playAction("double") }
    fun split() { playAction("split") }

    private fun playAction(action: String) {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            try {
                val response = apiService.blackjackAction(mapOf("action" to action))
                if (response.isSuccessful) {
                    response.body()?.let { r ->
                        _state.update {
                            it.copy(playerCards = r.playerCards, dealerCards = r.dealerCards, playerScore = r.playerScore, dealerScore = r.dealerScore, result = r.result, canHit = r.result == null, canStand = r.result == null, canDouble = false, canSplit = false, gameActive = r.result == null, isLoading = false, balance = r.newBalance ?: it.balance)
                        }
                    }
                }
            } catch (_: Exception) { _state.update { it.copy(isLoading = false) } }
        }
    }

    fun newGame() { _state.update { BlackjackState(balance = it.balance) } }
}

@Composable
fun BlackjackScreen(onBack: () -> Unit, viewModel: BlackjackViewModel = hiltViewModel()) {
    val state by viewModel.state.collectAsState()

    Column(modifier = Modifier.fillMaxSize()) {
        TopAppBar(title = { Text("Blackjack") }, navigationIcon = { IconButton(onClick = onBack) { Icon(Icons.Default.ArrowBack, contentDescription = "Back") } })

        Column(modifier = Modifier.fillMaxSize().padding(16.dp), horizontalAlignment = Alignment.CenterHorizontally) {
            Text("Balance: $${String.format("%.2f", state.balance)}", style = MaterialTheme.typography.titleMedium, color = MaterialTheme.colorScheme.primary)
            Spacer(modifier = Modifier.height(16.dp))

            Text("Dealer", style = MaterialTheme.typography.labelMedium)
            if (state.dealerCards.isNotEmpty()) {
                CardRow(cards = state.dealerCards, hideFirst = state.gameActive && state.dealerCards.size == 2)
                Text(if (state.gameActive && state.dealerCards.size == 2) "? + ${state.dealerCards.last().value}" else "Score: ${state.dealerScore}", style = MaterialTheme.typography.bodyMedium)
            } else {
                Spacer(modifier = Modifier.height(60.dp))
            }

            Spacer(modifier = Modifier.weight(1f))

            state.result?.let {
                val color = when (it) {
                    "win" -> MaterialTheme.colorScheme.primary
                    "blackjack" -> MaterialTheme.colorScheme.tertiary
                    "push" -> MaterialTheme.colorScheme.onSurface
                    else -> MaterialTheme.colorScheme.error
                }
                Text(it.uppercase(), style = MaterialTheme.typography.headlineMedium, fontWeight = FontWeight.Bold, color = color)
                Spacer(modifier = Modifier.height(8.dp))
            }

            Spacer(modifier = Modifier.weight(1f))

            Text("Your Hand", style = MaterialTheme.typography.labelMedium)
            if (state.playerCards.isNotEmpty()) {
                CardRow(cards = state.playerCards)
                Text("Score: ${state.playerScore}", style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
            } else {
                Spacer(modifier = Modifier.height(60.dp))
            }

            Spacer(modifier = Modifier.height(24.dp))

            if (!state.gameActive && state.result == null) {
                Row(verticalAlignment = Alignment.CenterVertically) {
                    Text("Bet: $", style = MaterialTheme.typography.bodyLarge)
                    listOf(10.0, 25.0, 50.0, 100.0).forEach { amount ->
                        FilterChip(selected = state.betAmount == amount, onClick = { viewModel.setBet(amount) }, label = { Text("$${amount.toInt()}") }, modifier = Modifier.padding(horizontal = 4.dp))
                    }
                }
                Spacer(modifier = Modifier.height(12.dp))
                Button(onClick = viewModel::deal, enabled = !state.isLoading && state.balance >= state.betAmount, modifier = Modifier.fillMaxWidth()) {
                    if (state.isLoading) CircularProgressIndicator(modifier = Modifier.size(20.dp)) else Text("Deal")
                }
            } else if (state.gameActive) {
                Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceEvenly) {
                    Button(onClick = viewModel::hit, enabled = state.canHit && !state.isLoading) { Text("Hit") }
                    Button(onClick = viewModel::stand, enabled = state.canStand && !state.isLoading) { Text("Stand") }
                    OutlinedButton(onClick = viewModel::doubleDown, enabled = state.canDouble && !state.isLoading) { Text("Double") }
                    if (state.canSplit) OutlinedButton(onClick = viewModel::split, enabled = !state.isLoading) { Text("Split") }
                }
            } else {
                Button(onClick = viewModel::newGame, modifier = Modifier.fillMaxWidth()) { Text("New Game") }
            }
        }
    }
}

@Composable
private fun CardRow(cards: List<PlayingCard>, hideFirst: Boolean = false) {
    LazyRow(horizontalArrangement = Arrangement.spacedBy((-20).dp), contentPadding = PaddingValues(horizontal = 20.dp)) {
        items(cards.indices.toList()) { index ->
            val card = cards[index]
            val hidden = hideFirst && index == 0
            PlayingCardView(card = if (hidden) null else card, zIndex = index.toFloat())
        }
    }
}

@Composable
private fun PlayingCardView(card: PlayingCard?, zIndex: Float = 0f) {
    val isRed = card?.suit in listOf("hearts", "diamonds")
    Surface(
        modifier = Modifier.width(55.dp).height(80.dp).offset(x = (zIndex * 2).dp),
        shape = RoundedCornerShape(6.dp),
        color = if (card == null) MaterialTheme.colorScheme.primary else MaterialTheme.colorScheme.surface,
        shadowElevation = 4.dp
    ) {
        Box(contentAlignment = Alignment.Center, modifier = Modifier.clip(RoundedCornerShape(6.dp)).background(if (card == null) MaterialTheme.colorScheme.primary else MaterialTheme.colorScheme.surface)) {
            if (card == null) {
                Icon(Icons.Default.Casino, contentDescription = null, tint = MaterialTheme.colorScheme.onPrimary, modifier = Modifier.size(24.dp))
            } else {
                Column(horizontalAlignment = Alignment.CenterHorizontally) {
                    Text(card.displayValue, fontSize = 18.sp, fontWeight = FontWeight.Bold, color = if (isRed) MaterialTheme.colorScheme.error else MaterialTheme.colorScheme.onSurface)
                    Text(card.suitSymbol, fontSize = 14.sp, color = if (isRed) MaterialTheme.colorScheme.error else MaterialTheme.colorScheme.onSurface)
                }
            }
        }
    }
}

private val PlayingCard.suitSymbol get() = when (suit) { "hearts" -> "♥"; "diamonds" -> "♦"; "clubs" -> "♣"; "spades" -> "♠"; else -> "?" }
