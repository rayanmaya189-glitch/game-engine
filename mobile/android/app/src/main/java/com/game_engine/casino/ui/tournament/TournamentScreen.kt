package com.game_engine.casino.ui.tournament

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
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import java.text.NumberFormat
import java.util.Locale

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun TournamentScreen(
    onTournamentClick: (String) -> Unit,
    viewModel: TournamentViewModel = hiltViewModel()
) {
    val uiState by viewModel.uiState.collectAsStateWithLifecycle()
    var selectedFilter by remember { mutableStateOf(TournamentFilter.ALL) }

    Scaffold(
        topBar = {
            TopAppBar(
                title = { Text("Tournaments") },
                actions = {
                    IconButton(onClick = { /* Filter */ }) {
                        Icon(Icons.Default.FilterList, contentDescription = "Filter")
                    }
                }
            )
        }
    ) { padding ->
        Column(
            modifier = Modifier
                .fillMaxSize()
                .padding(padding)
        ) {
            Row(
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(horizontal = 16.dp, vertical = 8.dp),
                horizontalArrangement = Arrangement.spacedBy(8.dp)
            ) {
                TournamentFilter.entries.forEach { filter ->
                    FilterChip(
                        selected = selectedFilter == filter,
                        onClick = {
                            selectedFilter = filter
                            viewModel.filterTournaments(filter)
                        },
                        label = { Text(filter.name.replace("_", " ")) }
                    )
                }
            }

            when (val state = uiState) {
                is TournamentUiState.Loading -> {
                    Box(
                        modifier = Modifier.fillMaxSize(),
                        contentAlignment = Alignment.Center
                    ) {
                        CircularProgressIndicator()
                    }
                }
                is TournamentUiState.Success -> {
                    LazyColumn(
                        modifier = Modifier.fillMaxSize(),
                        contentPadding = PaddingValues(16.dp),
                        verticalArrangement = Arrangement.spacedBy(12.dp)
                    ) {
                        items(state.tournaments) { tournament ->
                            TournamentCard(
                                tournament = tournament,
                                onClick = { onTournamentClick(tournament.id) }
                            )
                        }
                    }
                }
                is TournamentUiState.Error -> {
                    Box(
                        modifier = Modifier.fillMaxSize(),
                        contentAlignment = Alignment.Center
                    ) {
                        Column(horizontalAlignment = Alignment.CenterHorizontally) {
                            Text(state.message, color = MaterialTheme.colorScheme.error)
                            Spacer(modifier = Modifier.height(8.dp))
                            Button(onClick = { viewModel.loadTournaments() }) {
                                Text("Retry")
                            }
                        }
                    }
                }
            }
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun TournamentDetailScreen(
    tournamentId: String,
    onBackClick: () -> Unit,
    onPlayClick: () -> Unit,
    viewModel: TournamentViewModel = hiltViewModel()
) {
    val tournament by viewModel.selectedTournament.collectAsStateWithLifecycle()
    val leaderboard by viewModel.leaderboard.collectAsStateWithLifecycle()

    LaunchedEffect(tournamentId) {
        viewModel.selectTournament(tournamentId)
    }

    Scaffold(
        topBar = {
            TopAppBar(
                title = { Text(tournament?.name ?: "Tournament") },
                navigationIcon = {
                    IconButton(onClick = onBackClick) {
                        Icon(Icons.Default.ArrowBack, contentDescription = "Back")
                    }
                }
            )
        }
    ) { padding ->
        tournament?.let { t ->
            LazyColumn(
                modifier = Modifier
                    .fillMaxSize()
                    .padding(padding),
                contentPadding = PaddingValues(16.dp),
                verticalArrangement = Arrangement.spacedBy(16.dp)
            ) {
                item {
                    Card(
                        modifier = Modifier.fillMaxWidth(),
                        colors = CardDefaults.cardColors(
                            containerColor = MaterialTheme.colorScheme.primaryContainer
                        )
                    ) {
                        Column(
                            modifier = Modifier
                                .fillMaxWidth()
                                .padding(24.dp),
                            horizontalAlignment = Alignment.CenterHorizontally
                        ) {
                            Text(
                                text = "Prize Pool",
                                style = MaterialTheme.typography.titleMedium
                            )
                            Text(
                                text = NumberFormat.getCurrencyInstance(Locale.US).format(t.prizePool),
                                style = MaterialTheme.typography.headlineLarge,
                                fontWeight = FontWeight.Bold
                            )
                        }
                    }
                }

                item {
                    Card(modifier = Modifier.fillMaxWidth()) {
                        Column(modifier = Modifier.padding(16.dp)) {
                            Text(
                                text = "Tournament Info",
                                style = MaterialTheme.typography.titleMedium,
                                fontWeight = FontWeight.Bold
                            )
                            Spacer(modifier = Modifier.height(12.dp))
                            InfoItem(Icons.Default.SportsPoker, "Game", t.gameType)
                            InfoItem(Icons.Default.AttachMoney, "Buy-in", NumberFormat.getCurrencyInstance(Locale.US).format(t.buyIn))
                            InfoItem(Icons.Default.People, "Players", "${t.registeredPlayers}/${t.maxPlayers}")
                        }
                    }
                }

                item {
                    Text(
                        text = "Leaderboard",
                        style = MaterialTheme.typography.titleMedium,
                        fontWeight = FontWeight.Bold
                    )
                }

                items(leaderboard) { entry ->
                    LeaderboardItem(entry = entry)
                }

                item {
                    Button(
                        onClick = {
                            if (t.isRegistered) onPlayClick()
                            else viewModel.registerForTournament(t.id)
                        },
                        modifier = Modifier.fillMaxWidth(),
                        enabled = t.status == TournamentStatus.REGISTRATION_OPEN || t.status == TournamentStatus.IN_PROGRESS
                    ) {
                        Text(
                            text = when {
                                t.isRegistered -> "Play Now"
                                t.status == TournamentStatus.REGISTRATION_OPEN -> "Register - ${NumberFormat.getCurrencyInstance(Locale.US).format(t.buyIn)}"
                                else -> "Registration Closed"
                            }
                        )
                    }
                }
            }
        }
    }
}
