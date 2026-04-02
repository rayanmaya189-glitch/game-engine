package com.casino.game.ui.screens

import androidx.compose.foundation.background
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
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class LeaderboardState(
    val entries: List<LeaderboardEntry> = emptyList(),
    val period: String = "daily",
    val currentUserId: String = "",
    val isLoading: Boolean = false,
    val error: String? = null
)

@HiltViewModel
class LeaderboardViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(LeaderboardState())
    val state: StateFlow<LeaderboardState> = _state.asStateFlow()

    init { loadLeaderboard("daily") }

    fun loadLeaderboard(period: String) {
        _state.update { it.copy(isLoading = true, period = period) }
        viewModelScope.launch {
            try {
                val resp = apiService.getLeaderboard(period = period)
                if (resp.isSuccessful) {
                    resp.body()?.let { r ->
                        _state.update { it.copy(entries = r.entries, isLoading = false) }
                    }
                }
            } catch (_: Exception) { _state.update { it.copy(isLoading = false) } }
        }
    }
}

@Composable
fun LeaderboardScreen(onBack: () -> Unit, viewModel: LeaderboardViewModel = hiltViewModel()) {
    val state by viewModel.state.collectAsState()
    val periods = listOf("daily", "weekly", "monthly")

    Column(modifier = Modifier.fillMaxSize()) {
        TopAppBar(title = { Text("Leaderboard") }, navigationIcon = { IconButton(onClick = onBack) { Icon(Icons.Default.ArrowBack, contentDescription = "Back") } })

        Row(modifier = Modifier.fillMaxWidth().padding(horizontal = 16.dp), horizontalArrangement = Arrangement.spacedBy(8.dp)) {
            periods.forEach { period ->
                FilterChip(
                    selected = state.period == period,
                    onClick = { viewModel.loadLeaderboard(period) },
                    label = { Text(period.replaceFirstChar { it.uppercase() }) },
                    modifier = Modifier.weight(1f)
                )
            }
        }

        Spacer(modifier = Modifier.height(12.dp))

        if (state.isLoading) {
            Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) { CircularProgressIndicator() }
            return@Column
        }

        LazyColumn(modifier = Modifier.fillMaxSize(), contentPadding = PaddingValues(16.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
            if (state.entries.size >= 3) {
                item { PodiumSection(top3 = state.entries.take(3)) }
            }

            item {
                Text("All Rankings", style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold, modifier = Modifier.padding(top = 8.dp))
            }

            items(state.entries) { entry ->
                val isCurrentUser = entry.userId == state.currentUserId
                LeaderboardRow(entry = entry, isCurrentUser = isCurrentUser)
            }
        }
    }
}

@Composable
private fun PodiumSection(top3: List<LeaderboardEntry>) {
    Card(modifier = Modifier.fillMaxWidth(), colors = CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.primaryContainer)) {
        Column(modifier = Modifier.padding(16.dp), horizontalAlignment = Alignment.CenterHorizontally) {
            Text("Top Players", style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
            Spacer(modifier = Modifier.height(16.dp))
            Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.Center, verticalAlignment = Alignment.Bottom) {
                PodiumItem(entry = top3[1], rank = 2, height = 80.dp, color = MaterialTheme.colorScheme.secondaryContainer)
                Spacer(modifier = Modifier.width(8.dp))
                PodiumItem(entry = top3[0], rank = 1, height = 110.dp, color = MaterialTheme.colorScheme.primary)
                Spacer(modifier = Modifier.width(8.dp))
                PodiumItem(entry = top3[2], rank = 3, height = 60.dp, color = MaterialTheme.colorScheme.tertiaryContainer)
            }
        }
    }
}

@Composable
private fun PodiumItem(entry: LeaderboardEntry, rank: Int, height: androidx.compose.ui.unit.Dp, color: androidx.compose.ui.graphics.Color) {
    Column(horizontalAlignment = Alignment.CenterHorizontally) {
        Text(entry.username, style = MaterialTheme.typography.bodySmall, fontWeight = FontWeight.Bold, maxLines = 1)
        Text("${String.format("%,.0f", entry.score)}", style = MaterialTheme.typography.labelSmall)
        Spacer(modifier = Modifier.height(4.dp))
        Box(
            modifier = Modifier.width(80.dp).height(height).clip(RoundedCornerShape(topStart = 8.dp, topEnd = 8.dp)).background(color),
            contentAlignment = Alignment.Center
        ) {
            Text("#$rank", style = MaterialTheme.typography.headlineMedium, fontWeight = FontWeight.Bold, color = if (rank == 1) MaterialTheme.colorScheme.onPrimary else MaterialTheme.colorScheme.onSurface)
        }
        entry.prize?.let {
            Text("$${String.format("%,.0f", it)}", style = MaterialTheme.typography.labelSmall, color = MaterialTheme.colorScheme.primary, fontWeight = FontWeight.Bold)
        }
    }
}

@Composable
private fun LeaderboardRow(entry: LeaderboardEntry, isCurrentUser: Boolean) {
    val bgColor = if (isCurrentUser) MaterialTheme.colorScheme.secondaryContainer else MaterialTheme.colorScheme.surface
    Card(modifier = Modifier.fillMaxWidth(), colors = CardDefaults.cardColors(containerColor = bgColor)) {
        Row(modifier = Modifier.padding(12.dp), verticalAlignment = Alignment.CenterVertically) {
            Box(
                modifier = Modifier.size(36.dp).clip(CircleShape).background(
                    when (entry.rank) { 1 -> MaterialTheme.colorScheme.primary; 2 -> MaterialTheme.colorScheme.secondary; 3 -> MaterialTheme.colorScheme.tertiary; else -> MaterialTheme.colorScheme.surfaceVariant }
                ),
                contentAlignment = Alignment.Center
            ) {
                Text("#${entry.rank}", style = MaterialTheme.typography.labelMedium, fontWeight = FontWeight.Bold)
            }
            Spacer(modifier = Modifier.width(12.dp))
            Column(modifier = Modifier.weight(1f)) {
                Text(entry.username, style = MaterialTheme.typography.bodyLarge, fontWeight = if (isCurrentUser) FontWeight.Bold else FontWeight.Normal)
                if (isCurrentUser) Text("You", style = MaterialTheme.typography.labelSmall, color = MaterialTheme.colorScheme.primary)
            }
            Column(horizontalAlignment = Alignment.End) {
                Text("${String.format("%,.0f", entry.score)}", style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
                entry.prize?.let { Text("Prize: $${String.format("%,.0f", it)}", style = MaterialTheme.typography.labelSmall, color = MaterialTheme.colorScheme.primary) }
            }
        }
    }
}
