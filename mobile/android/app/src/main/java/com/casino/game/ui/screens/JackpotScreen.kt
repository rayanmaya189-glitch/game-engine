package com.casino.game.ui.screens

import androidx.compose.animation.core.*
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.casino.game.data.model.*
import com.casino.game.data.repository.JackpotRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class PlayerJackpotState(
    val jackpots: List<Jackpot> = emptyList(),
    val selectedJackpot: JackpotDetails? = null,
    val selectedIndex: Int = 0,
    val isLoading: Boolean = false
)

@HiltViewModel
class PlayerJackpotViewModel @Inject constructor(
    private val jackpotRepository: JackpotRepository
) : ViewModel() {
    private val _state = MutableStateFlow(PlayerJackpotState())
    val state: StateFlow<PlayerJackpotState> = _state.asStateFlow()

    init { loadJackpots() }

    private fun loadJackpots() {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            jackpotRepository.getCurrentJackpots().fold(
                onSuccess = { r ->
                    _state.update { it.copy(jackpots = r.jackpots, isLoading = false) }
                    if (r.jackpots.isNotEmpty()) loadDetails(r.jackpots.first().id)
                },
                onFailure = { _state.update { it.copy(isLoading = false) } }
            )
        }
    }

    fun loadDetails(jackpotId: String) {
        viewModelScope.launch {
            jackpotRepository.getJackpotDetails(jackpotId).fold(
                onSuccess = { details ->
                    val idx = _state.value.jackpots.indexOfFirst { it.id == jackpotId }
                    _state.update { it.copy(selectedJackpot = details, selectedIndex = idx.coerceAtLeast(0)) }
                },
                onFailure = { }
            )
        }
    }
}

@Composable
fun JackpotScreen(viewModel: PlayerJackpotViewModel = hiltViewModel()) {
    val state by viewModel.state.collectAsState()

    LazyColumn(
        modifier = Modifier.fillMaxSize(),
        contentPadding = PaddingValues(16.dp),
        verticalArrangement = Arrangement.spacedBy(12.dp)
    ) {
        item {
            Text("Progressive Jackpots", style = MaterialTheme.typography.headlineSmall)
            Spacer(modifier = Modifier.height(4.dp))
        }

        item {
            Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
                state.jackpots.forEachIndexed { idx, jackpot ->
                    FilterChip(
                        selected = state.selectedIndex == idx,
                        onClick = { viewModel.loadDetails(jackpot.id) },
                        label = { Text(jackpot.name) }
                    )
                }
            }
        }

        state.selectedJackpot?.let { details ->
            item {
                JackpotAmountCard(details)
            }

            item {
                Text("Recent Winners", style = MaterialTheme.typography.titleMedium)
            }

            items(details.recentHits) { hit ->
                Card(modifier = Modifier.fillMaxWidth()) {
                    Row(
                        modifier = Modifier.padding(16.dp),
                        horizontalArrangement = Arrangement.SpaceBetween,
                        verticalAlignment = Alignment.CenterVertically
                    ) {
                        Column(modifier = Modifier.weight(1f)) {
                            Text(hit.user, style = MaterialTheme.typography.bodyMedium)
                            Text(hit.gameName, style = MaterialTheme.typography.bodySmall)
                        }
                        Column(horizontalAlignment = Alignment.End) {
                            Text(
                                "$${String.format("%,.0f", hit.amount)}",
                                style = MaterialTheme.typography.titleMedium,
                                color = Color(0xFFFFD700),
                                fontWeight = FontWeight.Bold
                            )
                            Text(hit.wonAt, style = MaterialTheme.typography.bodySmall)
                        }
                    }
                }
            }
        }

        if (state.isLoading) {
            item {
                Box(
                    modifier = Modifier.fillMaxWidth().padding(32.dp),
                    contentAlignment = Alignment.Center
                ) { CircularProgressIndicator() }
            }
        }
    }
}

@Composable
private fun JackpotAmountCard(details: JackpotDetails) {
    var currentAmount by remember { mutableFloatStateOf(0f) }
    val animatedAmount by animateFloatAsState(
        targetValue = currentAmount.toFloat(),
        animationSpec = tween(durationMillis = 2000, easing = FastOutSlowInEasing)
    )

    LaunchedEffect(details.currentAmount) {
        currentAmount = details.currentAmount.toFloat()
    }

    Card(
        modifier = Modifier.fillMaxWidth(),
        colors = CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.primary)
    ) {
        Column(
            modifier = Modifier.padding(24.dp).fillMaxWidth(),
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            Icon(Icons.Default.EmojiEvents, contentDescription = null, tint = Color(0xFFFFD700), modifier = Modifier.size(48.dp))
            Spacer(modifier = Modifier.height(12.dp))
            Text(
                details.name,
                style = MaterialTheme.typography.titleMedium,
                color = MaterialTheme.colorScheme.onPrimary.copy(alpha = 0.8f)
            )
            Text(
                "$${String.format("%,.2f", animatedAmount)}",
                style = MaterialTheme.typography.displaySmall,
                fontWeight = FontWeight.Bold,
                color = Color(0xFFFFD700)
            )
            Text(
                details.game,
                style = MaterialTheme.typography.bodySmall,
                color = MaterialTheme.colorScheme.onPrimary.copy(alpha = 0.7f)
            )
            Text(
                "${details.hitCount} total wins",
                style = MaterialTheme.typography.bodySmall,
                color = MaterialTheme.colorScheme.onPrimary.copy(alpha = 0.7f)
            )
        }
    }
}
