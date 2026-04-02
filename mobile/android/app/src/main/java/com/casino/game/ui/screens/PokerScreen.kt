package com.casino.game.ui.screens

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyRow
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
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

data class PokerState(
    val playerHand: List<PlayingCard> = emptyList(),
    val communityCards: List<PlayingCard> = emptyList(),
    val pot: Double = 0.0,
    val playerChips: Double = 0.0,
    val currentBet: Double = 0.0,
    val handRanking: String? = null,
    val result: String? = null,
    val betAmount: Double = 10.0,
    val canBet: Boolean = false,
    val canCall: Boolean = false,
    val canRaise: Boolean = false,
    val canFold: Boolean = false,
    val gameActive: Boolean = false,
    val isLoading: Boolean = false,
    val balance: Double = 0.0
)

@HiltViewModel
class PokerViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(PokerState())
    val state: StateFlow<PokerState> = _state.asStateFlow()

    init { loadBalance() }

    private fun loadBalance() {
        viewModelScope.launch {
            try {
                val response = apiService.getBalance()
                if (response.isSuccessful) {
                    _state.update { it.copy(balance = response.body()?.balance ?: 0.0, playerChips = response.body()?.balance ?: 0.0) }
                }
            } catch (_: Exception) {}
        }
    }

    fun setBet(amount: Double) { _state.update { it.copy(betAmount = amount) } }

    fun startGame() {
        _state.update { it.copy(isLoading = true, result = null) }
        viewModelScope.launch {
            try {
                val response = apiService.pokerStart(mapOf("bet_amount" to _state.value.betAmount))
                if (response.isSuccessful) {
                    response.body()?.let { r ->
                        _state.update {
                            it.copy(playerHand = r.playerHand, communityCards = r.communityCards, pot = r.pot, currentBet = r.currentBet, handRanking = r.handRanking, canBet = true, canCall = r.canCall, canRaise = r.canRaise, canFold = true, gameActive = true, isLoading = false, playerChips = it.playerChips - it.betAmount)
                        }
                    }
                }
            } catch (_: Exception) { _state.update { it.copy(isLoading = false) } }
        }
    }

    fun bet() { pokerAction("bet", _state.value.betAmount) }
    fun call() { pokerAction("call", 0.0) }
    fun raise() { pokerAction("raise", _state.value.currentBet * 2) }
    fun fold() { pokerAction("fold", 0.0) }

    private fun pokerAction(action: String, amount: Double) {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            try {
                val response = apiService.pokerAction(mapOf("action" to action, "amount" to amount))
                if (response.isSuccessful) {
                    response.body()?.let { r ->
                        _state.update {
                            it.copy(playerHand = r.playerHand, communityCards = r.communityCards, pot = r.pot, handRanking = r.handRanking, result = r.result, canBet = r.result == null, canCall = r.result == null && r.canCall, canRaise = r.result == null && r.canRaise, canFold = r.result == null, gameActive = r.result == null, isLoading = false, balance = r.newBalance ?: it.balance, playerChips = r.playerChips ?: it.playerChips)
                        }
                    }
                }
            } catch (_: Exception) { _state.update { it.copy(isLoading = false) } }
        }
    }

    fun newGame() { _state.update { PokerState(balance = it.balance, playerChips = it.balance) } }
}

@Composable
fun PokerScreen(onBack: () -> Unit, viewModel: PokerViewModel = hiltViewModel()) {
    val state by viewModel.state.collectAsState()

    Column(modifier = Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
        TopAppBar(title = { Text("Poker") }, navigationIcon = { IconButton(onClick = onBack) { Icon(Icons.Default.ArrowBack, contentDescription = "Back") } })

        Column(modifier = Modifier.fillMaxSize().padding(16.dp), horizontalAlignment = Alignment.CenterHorizontally) {
            Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceBetween) {
                Text("Chips: $${String.format("%.2f", state.playerChips)}", style = MaterialTheme.typography.bodyMedium)
                Text("Pot: $${String.format("%.2f", state.pot)}", style = MaterialTheme.typography.bodyMedium, fontWeight = FontWeight.Bold, color = MaterialTheme.colorScheme.primary)
            }
            Spacer(modifier = Modifier.height(16.dp))

            Text("Community Cards", style = MaterialTheme.typography.labelMedium)
            Spacer(modifier = Modifier.height(8.dp))
            if (state.communityCards.isNotEmpty()) {
                CardRow(cards = state.communityCards)
            } else {
                Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) { repeat(5) { PlaceholderCard() } }
            }

            state.handRanking?.let {
                Spacer(modifier = Modifier.height(8.dp))
                AssistChip(onClick = {}, label = { Text(it, fontWeight = FontWeight.Bold) })
            }

            Spacer(modifier = Modifier.height(24.dp))

            state.result?.let {
                val color = when (it) { "win" -> MaterialTheme.colorScheme.primary; "lose" -> MaterialTheme.colorScheme.error; else -> MaterialTheme.colorScheme.onSurface }
                Text(it.uppercase(), style = MaterialTheme.typography.headlineMedium, fontWeight = FontWeight.Bold, color = color)
                Spacer(modifier = Modifier.height(16.dp))
            }

            Text("Your Hand", style = MaterialTheme.typography.labelMedium)
            Spacer(modifier = Modifier.height(8.dp))
            if (state.playerHand.isNotEmpty()) { CardRow(cards = state.playerHand) } else { Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) { repeat(2) { PlaceholderCard() } } }

            Spacer(modifier = Modifier.height(24.dp))

            if (!state.gameActive && state.result == null) {
                Row(verticalAlignment = Alignment.CenterVertically) {
                    Text("Bet: $", style = MaterialTheme.typography.bodyLarge)
                    listOf(10.0, 25.0, 50.0, 100.0).forEach { amount ->
                        FilterChip(selected = state.betAmount == amount, onClick = { viewModel.setBet(amount) }, label = { Text("$${amount.toInt()}") }, modifier = Modifier.padding(horizontal = 4.dp))
                    }
                }
                Spacer(modifier = Modifier.height(12.dp))
                Button(onClick = viewModel::startGame, enabled = !state.isLoading && state.balance >= state.betAmount, modifier = Modifier.fillMaxWidth()) {
                    if (state.isLoading) CircularProgressIndicator(modifier = Modifier.size(20.dp)) else Text("Start Game")
                }
            } else if (state.gameActive) {
                Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
                    Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.spacedBy(8.dp)) {
                        Button(onClick = viewModel::bet, enabled = state.canBet && !state.isLoading, modifier = Modifier.weight(1f)) { Text("Bet") }
                        Button(onClick = viewModel::call, enabled = state.canCall && !state.isLoading, modifier = Modifier.weight(1f)) { Text("Call") }
                    }
                    Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.spacedBy(8.dp)) {
                        OutlinedButton(onClick = viewModel::raise, enabled = state.canRaise && !state.isLoading, modifier = Modifier.weight(1f)) { Text("Raise") }
                        OutlinedButton(onClick = viewModel::fold, enabled = state.canFold && !state.isLoading, modifier = Modifier.weight(1f), colors = ButtonDefaults.outlinedButtonColors(contentColor = MaterialTheme.colorScheme.error)) { Text("Fold") }
                    }
                }
            } else {
                Button(onClick = viewModel::newGame, modifier = Modifier.fillMaxWidth()) { Text("New Game") }
            }
        }
    }
}

@Composable
private fun CardRow(cards: List<PlayingCard>) {
    LazyRow(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
        items(cards) { card -> PokerCardView(card = card) }
    }
}

@Composable
private fun PokerCardView(card: PlayingCard) {
    val isRed = card.suit in listOf("hearts", "diamonds")
    Surface(modifier = Modifier.width(50.dp).height(72.dp), shape = RoundedCornerShape(6.dp), color = MaterialTheme.colorScheme.surface, shadowElevation = 4.dp) {
        Box(contentAlignment = Alignment.Center) {
            Column(horizontalAlignment = Alignment.CenterHorizontally) {
                Text(card.displayValue, fontSize = 16.sp, fontWeight = FontWeight.Bold, color = if (isRed) MaterialTheme.colorScheme.error else MaterialTheme.colorScheme.onSurface)
                Text(when (card.suit) { "hearts" -> "♥"; "diamonds" -> "♦"; "clubs" -> "♣"; "spades" -> "♠"; else -> "?" }, fontSize = 12.sp, color = if (isRed) MaterialTheme.colorScheme.error else MaterialTheme.colorScheme.onSurface)
            }
        }
    }
}

@Composable
private fun PlaceholderCard() {
    Surface(modifier = Modifier.width(50.dp).height(72.dp), shape = RoundedCornerShape(6.dp), color = MaterialTheme.colorScheme.surfaceVariant.copy(alpha = 0.5f)) {
        Box(contentAlignment = Alignment.Center) { Icon(Icons.Default.Casino, contentDescription = null, modifier = Modifier.size(20.dp), tint = MaterialTheme.colorScheme.onSurfaceVariant.copy(alpha = 0.3f)) }
    }
}
