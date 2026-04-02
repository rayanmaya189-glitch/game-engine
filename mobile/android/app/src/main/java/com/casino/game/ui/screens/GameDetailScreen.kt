package com.casino.game.ui.screens

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.LazyRow
import androidx.compose.foundation.lazy.items
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
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

data class GameDetailState(
    val game: GameDetails? = null,
    val recentBigWins: List<BigWin> = emptyList(),
    val relatedGames: List<Game> = emptyList(),
    val isLoading: Boolean = false,
    val error: String? = null
)

@HiltViewModel
class GameDetailViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(GameDetailState())
    val state: StateFlow<GameDetailState> = _state.asStateFlow()

    fun loadGame(gameId: String) {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            try {
                val response = apiService.getGameDetail(gameId)
                if (response.isSuccessful) {
                    val body = response.body()
                    _state.update { it.copy(game = body?.game, recentBigWins = body?.recentBigWins ?: emptyList(), relatedGames = body?.relatedGames ?: emptyList(), isLoading = false) }
                } else {
                    _state.update { it.copy(isLoading = false, error = "Failed to load game") }
                }
            } catch (e: Exception) {
                _state.update { it.copy(isLoading = false, error = e.message) }
            }
        }
    }
}

@Composable
fun GameDetailScreen(
    gameId: String,
    onBack: () -> Unit,
    onPlayGame: (String) -> Unit,
    onGameClick: (String) -> Unit,
    viewModel: GameDetailViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()

    LaunchedEffect(gameId) { viewModel.loadGame(gameId) }

    if (state.isLoading) {
        Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) { CircularProgressIndicator() }
        return
    }

    val game = state.game ?: run {
        Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) { Text("Game not found") }
        return
    }

    LazyColumn(modifier = Modifier.fillMaxSize()) {
        item {
            Box(
                modifier = Modifier.fillMaxWidth().height(200.dp),
                contentAlignment = Alignment.Center
            ) {
                Icon(Icons.Default.Casino, contentDescription = null, modifier = Modifier.size(80.dp), tint = MaterialTheme.colorScheme.primary)
            }
        }

        item {
            Column(modifier = Modifier.padding(16.dp)) {
                Text(game.name, style = MaterialTheme.typography.headlineMedium, fontWeight = FontWeight.Bold)
                Text(game.provider, style = MaterialTheme.typography.bodyMedium, color = MaterialTheme.colorScheme.onSurfaceVariant)
                Spacer(modifier = Modifier.height(16.dp))

                Card(modifier = Modifier.fillMaxWidth()) {
                    Column(modifier = Modifier.padding(16.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
                        GameInfoRow("RTP", "${game.rtp}%")
                        GameInfoRow("Volatility", game.volatility ?: "Medium")
                        GameInfoRow("Min Bet", "$${String.format("%,.2f", game.minBet)}")
                        GameInfoRow("Max Bet", "$${String.format("%,.2f", game.maxBet)}")
                        game.paylines?.let { GameInfoRow("Paylines", "$it") }
                        game.reels?.let { GameInfoRow("Reels", "$it") }
                    }
                }

                Spacer(modifier = Modifier.height(16.dp))
                Button(onClick = { onPlayGame(game.id) }, modifier = Modifier.fillMaxWidth().height(56.dp)) {
                    Icon(Icons.Default.PlayArrow, contentDescription = null)
                    Spacer(modifier = Modifier.width(8.dp))
                    Text("Play Now")
                }
            }
        }

        if (state.recentBigWins.isNotEmpty()) {
            item {
                Column(modifier = Modifier.padding(16.dp)) {
                    Text("Recent Big Wins", style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
                    Spacer(modifier = Modifier.height(8.dp))
                }
            }
            items(state.recentBigWins.take(5)) { win ->
                Card(modifier = Modifier.fillMaxWidth().padding(horizontal = 16.dp, vertical = 4.dp)) {
                    Row(
                        modifier = Modifier.padding(12.dp),
                        horizontalArrangement = Arrangement.SpaceBetween,
                        verticalAlignment = Alignment.CenterVertically
                    ) {
                        Row(verticalAlignment = Alignment.CenterVertically) {
                            Icon(Icons.Default.EmojiEvents, contentDescription = null, tint = MaterialTheme.colorScheme.primary)
                            Spacer(modifier = Modifier.width(8.dp))
                            Column {
                                Text(win.username, style = MaterialTheme.typography.bodyMedium)
                                Text(win.date, style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
                            }
                        }
                        Text("$${String.format("%,.2f", win.amount)}", style = MaterialTheme.typography.titleMedium, color = MaterialTheme.colorScheme.primary, fontWeight = FontWeight.Bold)
                    }
                }
            }
        }

        if (state.relatedGames.isNotEmpty()) {
            item {
                Column(modifier = Modifier.padding(16.dp)) {
                    Text("Related Games", style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
                    Spacer(modifier = Modifier.height(8.dp))
                }
            }
            item {
                LazyRow(
                    contentPadding = PaddingValues(horizontal = 16.dp),
                    horizontalArrangement = Arrangement.spacedBy(12.dp)
                ) {
                    items(state.relatedGames) { relatedGame ->
                        RelatedGameCard(game = relatedGame, onClick = { onGameClick(relatedGame.id) })
                    }
                }
                Spacer(modifier = Modifier.height(16.dp))
            }
        }
    }
}

@Composable
private fun GameInfoRow(label: String, value: String) {
    Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceBetween) {
        Text(label, style = MaterialTheme.typography.bodyMedium, color = MaterialTheme.colorScheme.onSurfaceVariant)
        Text(value, style = MaterialTheme.typography.bodyMedium, fontWeight = FontWeight.Medium)
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun RelatedGameCard(game: Game, onClick: () -> Unit) {
    Card(onClick = onClick, modifier = Modifier.width(140.dp)) {
        Column {
            Box(
                modifier = Modifier.fillMaxWidth().height(80.dp),
                contentAlignment = Alignment.Center
            ) {
                Icon(Icons.Default.Casino, contentDescription = null, modifier = Modifier.size(32.dp), tint = MaterialTheme.colorScheme.primary)
            }
            Column(modifier = Modifier.padding(8.dp)) {
                Text(game.name, style = MaterialTheme.typography.bodySmall, maxLines = 1)
                Text("RTP: ${game.rtp}%", style = MaterialTheme.typography.labelSmall, color = MaterialTheme.colorScheme.primary)
            }
        }
    }
}
