package com.casino.game.ui.screens

import androidx.compose.animation.AnimatedVisibility
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
import com.casino.game.data.remote.ApiService
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class BetHistoryState(
    val bets: List<BetEntry> = emptyList(),
    val selectedFilter: String = "all",
    val expandedBetId: String? = null,
    val isLoading: Boolean = false,
    val error: String? = null
)

@HiltViewModel
class BetHistoryViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(BetHistoryState())
    val state: StateFlow<BetHistoryState> = _state.asStateFlow()

    init { loadBets() }

    fun loadBets() {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            try {
                val response = apiService.getBetHistory()
                if (response.isSuccessful) {
                    _state.update { it.copy(bets = response.body()?.bets ?: emptyList(), isLoading = false) }
                } else {
                    _state.update { it.copy(isLoading = false, error = "Failed to load bet history") }
                }
            } catch (e: Exception) {
                _state.update { it.copy(isLoading = false, error = e.message) }
            }
        }
    }

    fun setFilter(filter: String) { _state.update { it.copy(selectedFilter = filter) } }
    fun toggleExpand(betId: String) {
        _state.update { it.copy(expandedBetId = if (it.expandedBetId == betId) null else betId) }
    }
}

@Composable
fun BetHistoryScreen(
    onBack: () -> Unit,
    viewModel: BetHistoryViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()
    val filters = listOf("all" to "All", "won" to "Won", "lost" to "Lost", "pending" to "Pending")

    val filteredBets = if (state.selectedFilter == "all") state.bets
    else state.bets.filter { it.result == state.selectedFilter }

    Column(modifier = Modifier.fillMaxSize()) {
        TopAppBar(
            title = { Text("Bet History") },
            navigationIcon = { IconButton(onClick = onBack) { Icon(Icons.Default.ArrowBack, contentDescription = "Back") } }
        )

        if (state.isLoading) {
            Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) { CircularProgressIndicator() }
        } else {
            LazyColumn(
                modifier = Modifier.fillMaxSize(),
                contentPadding = PaddingValues(16.dp),
                verticalArrangement = Arrangement.spacedBy(12.dp)
            ) {
                item {
                    Row(
                        modifier = Modifier.fillMaxWidth(),
                        horizontalArrangement = Arrangement.spacedBy(8.dp)
                    ) {
                        filters.forEach { (key, label) ->
                            FilterChip(
                                selected = state.selectedFilter == key,
                                onClick = { viewModel.setFilter(key) },
                                label = { Text(label) }
                            )
                        }
                    }
                }

                if (filteredBets.isEmpty()) {
                    item {
                        Box(
                            modifier = Modifier.fillMaxWidth().padding(48.dp),
                            contentAlignment = Alignment.Center
                        ) {
                            Column(horizontalAlignment = Alignment.CenterHorizontally) {
                                Icon(Icons.Default.Casino, contentDescription = null, modifier = Modifier.size(64.dp), tint = MaterialTheme.colorScheme.onSurfaceVariant)
                                Spacer(modifier = Modifier.height(8.dp))
                                Text("No bets found", style = MaterialTheme.typography.bodyLarge, color = MaterialTheme.colorScheme.onSurfaceVariant)
                            }
                        }
                    }
                } else {
                    items(filteredBets) { bet ->
                        BetHistoryCard(
                            bet = bet,
                            isExpanded = state.expandedBetId == bet.id,
                            onToggleExpand = { viewModel.toggleExpand(bet.id) }
                        )
                    }
                }
            }
        }
    }
}

@Composable
private fun BetHistoryCard(bet: BetEntry, isExpanded: Boolean, onToggleExpand: () -> Unit) {
    val resultColor = when (bet.result) {
        "won" -> Color(0xFF4CAF50)
        "lost" -> Color(0xFFF44336)
        "pending" -> Color(0xFFFFC107)
        else -> MaterialTheme.colorScheme.onSurface
    }

    Card(modifier = Modifier.fillMaxWidth(), onClick = onToggleExpand) {
        Column(modifier = Modifier.padding(16.dp)) {
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween,
                verticalAlignment = Alignment.CenterVertically
            ) {
                Column(modifier = Modifier.weight(1f)) {
                    Text(bet.gameName, style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
                    Text(bet.date, style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
                }
                AssistChip(
                    onClick = { },
                    label = { Text(bet.result.uppercase(), fontWeight = FontWeight.Bold) },
                    colors = AssistChipDefaults.assistChipColors(containerColor = resultColor.copy(alpha = 0.15f)),
                    labelStyle = MaterialTheme.typography.labelSmall.copy(color = resultColor)
                )
            }
            Spacer(modifier = Modifier.height(8.dp))
            Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceBetween) {
                Column { Text("Stake", style = MaterialTheme.typography.labelSmall, color = MaterialTheme.colorScheme.onSurfaceVariant); Text("$${String.format("%,.2f", bet.stake)}", style = MaterialTheme.typography.bodyMedium) }
                Column(horizontalAlignment = Alignment.End) {
                    Text("Potential Win", style = MaterialTheme.typography.labelSmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
                    Text(
                        if (bet.result == "won") "$${String.format("%,.2f", bet.actualWin ?: 0.0)}" else "$${String.format("%,.2f", bet.potentialWin)}",
                        style = MaterialTheme.typography.bodyMedium,
                        color = resultColor,
                        fontWeight = FontWeight.Bold
                    )
                }
            }
            if (bet.selections.isNotEmpty()) {
                Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.End) {
                    Icon(
                        if (isExpanded) Icons.Default.ExpandLess else Icons.Default.ExpandMore,
                        contentDescription = "Details",
                        modifier = Modifier.size(20.dp)
                    )
                }
                AnimatedVisibility(visible = isExpanded) {
                    Column(modifier = Modifier.padding(top = 8.dp), verticalArrangement = Arrangement.spacedBy(4.dp)) {
                        Divider()
                        Spacer(modifier = Modifier.height(4.dp))
                        Text("Selections", style = MaterialTheme.typography.labelMedium, fontWeight = FontWeight.Bold)
                        bet.selections.forEach { sel ->
                            Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceBetween) {
                                Text(sel.selection, style = MaterialTheme.typography.bodySmall)
                                Text("${sel.odds}x - ${sel.result}", style = MaterialTheme.typography.bodySmall,
                                    color = when (sel.result) { "won" -> Color(0xFF4CAF50); "lost" -> Color(0xFFF44336); else -> Color(0xFFFFC107) })
                            }
                        }
                    }
                }
            }
        }
    }
}
