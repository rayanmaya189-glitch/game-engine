package com.game_engine.casino.ui.tournament

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.CircleShape
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
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import java.text.NumberFormat
import java.time.format.DateTimeFormatter
import java.util.Locale

/**
 * Tournament Screen
 * 
 * Displays tournament lobby with browse, register, and play functionality.
 */
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
            // Filter Chips
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

@Composable
fun TournamentCard(
    tournament: Tournament,
    onClick: () -> Unit
) {
    val currencyFormat = remember { NumberFormat.getCurrencyInstance(Locale.US) }
    val dateFormatter = remember { DateTimeFormatter.ofPattern("MMM dd, HH:mm") }

    Card(
        modifier = Modifier
            .fillMaxWidth()
            .clickable(onClick = onClick),
        shape = RoundedCornerShape(12.dp)
    ) {
        Column(
            modifier = Modifier.padding(16.dp)
        ) {
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween,
                verticalAlignment = Alignment.CenterVertically
            ) {
                Text(
                    text = tournament.name,
                    style = MaterialTheme.typography.titleMedium,
                    fontWeight = FontWeight.Bold
                )
                StatusChip(tournament.status)
            }

            Spacer(modifier = Modifier.height(8.dp))

            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.spacedBy(16.dp)
            ) {
                InfoItem(
                    icon = Icons.Default.SportsPoker,
                    label = "Game",
                    value = tournament.gameType
                )
                InfoItem(
                    icon = Icons.Default.AttachMoney,
                    label = "Buy-in",
                    value = currencyFormat.format(tournament.buyIn)
                )
            }

            Spacer(modifier = Modifier.height(8.dp))

            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.spacedBy(16.dp)
            ) {
                InfoItem(
                    icon = Icons.Default.EmojiEvents,
                    label = "Prize Pool",
                    value = currencyFormat.format(tournament.prizePool)
                )
                InfoItem(
                    icon = Icons.Default.People,
                    label = "Players",
                    value = "${tournament.registeredPlayers}/${tournament.maxPlayers}"
                )
            }

            Spacer(modifier = Modifier.height(12.dp))

            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween,
                verticalAlignment = Alignment.CenterVertically
            ) {
                Text(
                    text = "${tournament.startTime.format(dateFormatter)} - ${tournament.endTime.format(dateFormatter)}",
                    style = MaterialTheme.typography.bodySmall,
                    color = MaterialTheme.colorScheme.onSurfaceVariant
                )

                if (tournament.isRegistered) {
                    Button(
                        onClick = { /* Play */ },
                        colors = ButtonDefaults.buttonColors(
                            containerColor = MaterialTheme.colorScheme.primary
                        )
                    ) {
                        Text("Play Now")
                    }
                } else if (tournament.status == TournamentStatus.REGISTRATION_OPEN) {
                    Button(onClick = { /* Register */ }) {
                        Text("Register")
                    }
                }
            }
        }
    }
}

@Composable
fun StatusChip(status: TournamentStatus) {
    val (color, text) = when (status) {
        TournamentStatus.UPCOMING -> MaterialTheme.colorScheme.secondary to "Upcoming"
        TournamentStatus.REGISTRATION_OPEN -> MaterialTheme.colorScheme.primary to "Register"
        TournamentStatus.REGISTRATION_CLOSED -> MaterialTheme.colorScheme.tertiary to "Closed"
        TournamentStatus.IN_PROGRESS -> MaterialTheme.colorScheme.primary to "Live"
        TournamentStatus.COMPLETED -> MaterialTheme.colorScheme.surfaceVariant to "Completed"
        TournamentStatus.CANCELLED -> MaterialTheme.colorScheme.error to "Cancelled"
    }

    Surface(
        shape = RoundedCornerShape(16.dp),
        color = color.copy(alpha = 0.1f)
    ) {
        Text(
            text = text,
            modifier = Modifier.padding(horizontal = 12.dp, vertical = 4.dp),
            style = MaterialTheme.typography.labelSmall,
            color = color
        )
    }
}

@Composable
fun InfoItem(
    icon: androidx.compose.ui.graphics.vector.ImageVector,
    label: String,
    value: String
) {
    Row(
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(4.dp)
    ) {
        Icon(
            imageVector = icon,
            contentDescription = null,
            modifier = Modifier.size(16.dp),
            tint = MaterialTheme.colorScheme.onSurfaceVariant
        )
        Column {
            Text(
                text = label,
                style = MaterialTheme.typography.labelSmall,
                color = MaterialTheme.colorScheme.onSurfaceVariant
            )
            Text(
                text = value,
                style = MaterialTheme.typography.bodyMedium,
                fontWeight = FontWeight.Medium
            )
        }
    }
}

/**
 * Tournament Detail Screen
 */
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
                // Prize Pool Card
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

                // Info Section
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

                // Leaderboard
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

                // Action Button
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

@Composable
fun LeaderboardItem(entry: LeaderboardEntry) {
    Card(
        modifier = Modifier.fillMaxWidth(),
        colors = if (entry.isCurrentPlayer) {
            CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.primaryContainer)
        } else {
            CardDefaults.cardColors()
        }
    ) {
        Row(
            modifier = Modifier
                .fillMaxWidth()
                .padding(12.dp),
            verticalAlignment = Alignment.CenterVertically
        ) {
            // Rank
            Box(
                modifier = Modifier
                    .size(32.dp)
                    .clip(CircleShape)
                    .background(
                        when (entry.rank) {
                            1 -> Color(0xFFFFD700)
                            2 -> Color(0xFFC0C0C0)
                            3 -> Color(0xFFCD7F32)
                            else -> MaterialTheme.colorScheme.surfaceVariant
                        }
                    ),
                contentAlignment = Alignment.Center
            ) {
                Text(
                    text = "${entry.rank}",
                    fontWeight = FontWeight.Bold
                )
            }

            Spacer(modifier = Modifier.width(12.dp))

            Column(modifier = Modifier.weight(1f)) {
                Text(
                    text = entry.playerName,
                    fontWeight = if (entry.isCurrentPlayer) FontWeight.Bold else FontWeight.Normal
                )
            }

            Text(
                text = NumberFormat.getNumberInstance(Locale.US).format(entry.chipCount),
                fontWeight = FontWeight.Bold
            )
        }
    }
}
