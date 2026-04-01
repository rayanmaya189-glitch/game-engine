package com.casino.game.ui.screens

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
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
import com.casino.game.data.repository.TournamentRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class PlayerTournamentState(
    val tournaments: List<Tournament> = emptyList(),
    val selectedTournament: TournamentDetails? = null,
    val leaderboard: List<LeaderboardEntry> = emptyList(),
    val isLoading: Boolean = false,
    val error: String? = null
)

@HiltViewModel
class PlayerTournamentViewModel @Inject constructor(
    private val tournamentRepository: TournamentRepository
) : ViewModel() {
    private val _state = MutableStateFlow(PlayerTournamentState())
    val state: StateFlow<PlayerTournamentState> = _state.asStateFlow()

    init { loadTournaments() }

    fun loadTournaments() {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            tournamentRepository.getTournaments(status = "active").fold(
                onSuccess = { r ->
                    _state.update { it.copy(tournaments = r.tournaments, isLoading = false) }
                },
                onFailure = { e ->
                    _state.update { it.copy(isLoading = false, error = e.message) }
                }
            )
        }
    }

    fun loadDetails(tournamentId: String) {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            tournamentRepository.getTournamentDetails(tournamentId).fold(
                onSuccess = { details ->
                    _state.update { it.copy(selectedTournament = details) }
                },
                onFailure = { }
            )
            tournamentRepository.getTournamentLeaderboard(tournamentId).fold(
                onSuccess = { lb ->
                    _state.update { it.copy(leaderboard = lb.entries, isLoading = false) }
                },
                onFailure = { _state.update { it.copy(isLoading = false) } }
            )
        }
    }

    fun clearSelection() {
        _state.update { it.copy(selectedTournament = null, leaderboard = emptyList()) }
    }
}

@Composable
fun TournamentScreen(
    viewModel: PlayerTournamentViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()

    if (state.selectedTournament != null) {
        TournamentDetailContent(
            tournament = state.selectedTournament!!,
            leaderboard = state.leaderboard,
            onBack = viewModel::clearSelection
        )
    } else {
        TournamentListContent(
            tournaments = state.tournaments,
            isLoading = state.isLoading,
            onTournamentClick = viewModel::loadDetails
        )
    }
}

@Composable
private fun TournamentListContent(
    tournaments: List<Tournament>,
    isLoading: Boolean,
    onTournamentClick: (String) -> Unit
) {
    LazyColumn(
        modifier = Modifier.fillMaxSize(),
        contentPadding = PaddingValues(16.dp),
        verticalArrangement = Arrangement.spacedBy(12.dp)
    ) {
        item {
            Text("Active Tournaments", style = MaterialTheme.typography.headlineSmall)
            Spacer(modifier = Modifier.height(4.dp))
        }
        if (isLoading) {
            item {
                Box(
                    modifier = Modifier.fillMaxWidth().padding(32.dp),
                    contentAlignment = Alignment.Center
                ) { CircularProgressIndicator() }
            }
        } else {
            items(tournaments) { tournament ->
                Card(
                    onClick = { onTournamentClick(tournament.id) },
                    modifier = Modifier.fillMaxWidth()
                ) {
                    Column(modifier = Modifier.padding(16.dp)) {
                        Row(
                            modifier = Modifier.fillMaxWidth(),
                            horizontalArrangement = Arrangement.SpaceBetween
                        ) {
                            Column(modifier = Modifier.weight(1f)) {
                                Text(
                                    tournament.name,
                                    style = MaterialTheme.typography.titleMedium,
                                    fontWeight = FontWeight.Bold
                                )
                                Text(tournament.game, style = MaterialTheme.typography.bodySmall)
                            }
                            Column(horizontalAlignment = Alignment.End) {
                                Text(
                                    "$${String.format("%,.0f", tournament.prizePool)}",
                                    style = MaterialTheme.typography.titleMedium,
                                    color = MaterialTheme.colorScheme.primary,
                                    fontWeight = FontWeight.Bold
                                )
                                Text("Prize Pool", style = MaterialTheme.typography.bodySmall)
                            }
                        }
                        Spacer(modifier = Modifier.height(8.dp))
                        Row(
                            modifier = Modifier.fillMaxWidth(),
                            horizontalArrangement = Arrangement.SpaceBetween
                        ) {
                            Text("${tournament.playerCount} players", style = MaterialTheme.typography.bodySmall)
                            Text("Min bet: $${tournament.minBet}", style = MaterialTheme.typography.bodySmall)
                        }
                    }
                }
            }
        }
    }
}

@Composable
private fun TournamentDetailContent(
    tournament: TournamentDetails,
    leaderboard: List<LeaderboardEntry>,
    onBack: () -> Unit
) {
    LazyColumn(
        modifier = Modifier.fillMaxSize(),
        contentPadding = PaddingValues(16.dp),
        verticalArrangement = Arrangement.spacedBy(12.dp)
    ) {
        item {
            Row(verticalAlignment = Alignment.CenterVertically) {
                IconButton(onClick = onBack) {
                    Icon(Icons.Default.ArrowBack, contentDescription = "Back")
                }
                Text(tournament.name, style = MaterialTheme.typography.headlineSmall)
            }
        }

        item {
            Card(
                modifier = Modifier.fillMaxWidth(),
                colors = CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.primaryContainer)
            ) {
                Column(modifier = Modifier.padding(20.dp), horizontalAlignment = Alignment.CenterHorizontally) {
                    Text("Prize Pool", style = MaterialTheme.typography.bodyMedium)
                    Text(
                        "$${String.format("%,.0f", tournament.prizePool)}",
                        style = MaterialTheme.typography.headlineLarge,
                        fontWeight = FontWeight.Bold
                    )
                    Spacer(modifier = Modifier.height(8.dp))
                    Text("${tournament.playerCount} participants", style = MaterialTheme.typography.bodySmall)
                }
            }
        }

        item {
            Button(
                onClick = { },
                modifier = Modifier.fillMaxWidth().height(48.dp)
            ) {
                Icon(Icons.Default.EmojiEvents, contentDescription = null)
                Spacer(modifier = Modifier.width(8.dp))
                Text("Join Tournament")
            }
        }

        item {
            Text("Leaderboard", style = MaterialTheme.typography.titleMedium)
        }

        items(leaderboard) { entry ->
            Card(modifier = Modifier.fillMaxWidth()) {
                Row(
                    modifier = Modifier.padding(12.dp),
                    verticalAlignment = Alignment.CenterVertically
                ) {
                    Text(
                        "#${entry.rank}",
                        style = MaterialTheme.typography.titleMedium,
                        fontWeight = FontWeight.Bold,
                        modifier = Modifier.width(40.dp)
                    )
                    Column(modifier = Modifier.weight(1f)) {
                        Text(entry.username, style = MaterialTheme.typography.bodyMedium)
                    }
                    Text(
                        "$${String.format("%,.0f", entry.score)}",
                        style = MaterialTheme.typography.titleMedium,
                        fontWeight = FontWeight.Bold
                    )
                }
            }
        }
    }
}
