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
import com.casino.game.data.remote.ApiService
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class LiveDealerState(
    val tables: List<LiveDealerTable> = emptyList(),
    val selectedTable: LiveDealerTable? = null,
    val dealerChatText: String = "",
    val dealerMessages: List<DealerChatMessage> = emptyList(),
    val isLoading: Boolean = false,
    val error: String? = null
)

@HiltViewModel
class LiveDealerViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(LiveDealerState())
    val state: StateFlow<LiveDealerState> = _state.asStateFlow()

    init { loadTables() }

    fun loadTables() {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            try {
                val response = apiService.getLiveDealerTables()
                if (response.isSuccessful) {
                    _state.update { it.copy(tables = response.body() ?: emptyList(), isLoading = false) }
                }
            } catch (_: Exception) { _state.update { it.copy(isLoading = false) } }
        }
    }

    fun selectTable(table: LiveDealerTable) { _state.update { it.copy(selectedTable = table) } }
    fun clearSelection() { _state.update { it.copy(selectedTable = null) } }
    fun updateChatText(text: String) { _state.update { it.copy(dealerChatText = text) } }

    fun sendDealerMessage() {
        val table = _state.value.selectedTable ?: return
        val text = _state.value.dealerChatText.trim()
        if (text.isEmpty()) return
        viewModelScope.launch {
            try {
                val response = apiService.sendDealerChat(table.id, mapOf("content" to text))
                if (response.isSuccessful) {
                    response.body()?.let { msg ->
                        _state.update { it.copy(dealerMessages = it.dealerMessages + msg, dealerChatText = "") }
                    }
                }
            } catch (_: Exception) {}
        }
    }
}

@Composable
fun LiveDealerScreen(
    onBack: () -> Unit,
    viewModel: LiveDealerViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()

    Column(modifier = Modifier.fillMaxSize()) {
        TopAppBar(
            title = { Text(if (state.selectedTable != null) state.selectedTable!!.name else "Live Dealer") },
            navigationIcon = {
                IconButton(onClick = { if (state.selectedTable != null) viewModel.clearSelection() else onBack() }) {
                    Icon(Icons.Default.ArrowBack, contentDescription = "Back")
                }
            }
        )

        if (state.selectedTable != null) {
            LiveDealerTableView(table = state.selectedTable!!, messages = state.dealerMessages, chatText = state.dealerChatText, onTextChange = viewModel::updateChatText, onSend = viewModel::sendDealerMessage)
        } else {
            LiveDealerTableList(tables = state.tables, isLoading = state.isLoading, onTableClick = viewModel::selectTable)
        }
    }
}

@Composable
private fun LiveDealerTableList(tables: List<LiveDealerTable>, isLoading: Boolean, onTableClick: (LiveDealerTable) -> Unit) {
    if (isLoading) {
        Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) { CircularProgressIndicator() }
        return
    }
    LazyColumn(modifier = Modifier.fillMaxSize(), contentPadding = PaddingValues(16.dp), verticalArrangement = Arrangement.spacedBy(12.dp)) {
        if (tables.isEmpty()) {
            item {
                Box(modifier = Modifier.fillMaxWidth().padding(32.dp), contentAlignment = Alignment.Center) {
                    Column(horizontalAlignment = Alignment.CenterHorizontally) {
                        Icon(Icons.Default.Videocam, contentDescription = null, modifier = Modifier.size(64.dp), tint = MaterialTheme.colorScheme.onSurfaceVariant)
                        Spacer(modifier = Modifier.height(8.dp))
                        Text("No live tables available")
                    }
                }
            }
        }
        items(tables) { table -> LiveDealerTableCard(table = table, onClick = { onTableClick(table) }) }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun LiveDealerTableCard(table: LiveDealerTable, onClick: () -> Unit) {
    Card(onClick = onClick, modifier = Modifier.fillMaxWidth()) {
        Column {
            Surface(modifier = Modifier.fillMaxWidth().height(140.dp), color = MaterialTheme.colorScheme.surfaceVariant) {
                Box(contentAlignment = Alignment.Center) {
                    Column(horizontalAlignment = Alignment.CenterHorizontally) {
                        Icon(Icons.Default.Videocam, contentDescription = null, modifier = Modifier.size(48.dp), tint = MaterialTheme.colorScheme.primary)
                        Text("Live", style = MaterialTheme.typography.labelMedium, color = MaterialTheme.colorScheme.error)
                    }
                }
            }
            Column(modifier = Modifier.padding(16.dp)) {
                Text(table.name, style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
                Spacer(modifier = Modifier.height(4.dp))
                Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceBetween) {
                    InfoChip(icon = Icons.Default.Person, text = table.dealerName)
                    InfoChip(icon = Icons.Default.Group, text = "${table.currentPlayers}/${table.maxPlayers}")
                }
                Spacer(modifier = Modifier.height(8.dp))
                Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceBetween) {
                    Text("Bet: $${table.minBet} - $${table.maxBet}", style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
                    AssistChip(onClick = {}, label = { Text(table.gameType) })
                }
            }
        }
    }
}

@Composable
private fun LiveDealerTableView(table: LiveDealerTable, messages: List<DealerChatMessage>, chatText: String, onTextChange: (String) -> Unit, onSend: () -> Unit) {
    Column(modifier = Modifier.fillMaxSize()) {
        Surface(modifier = Modifier.fillMaxWidth().height(220.dp), color = MaterialTheme.colorScheme.surfaceVariant) {
            Box(contentAlignment = Alignment.Center) {
                Column(horizontalAlignment = Alignment.CenterHorizontally) {
                    Icon(Icons.Default.VideocamOff, contentDescription = null, modifier = Modifier.size(48.dp))
                    Text("Camera Feed", style = MaterialTheme.typography.bodyMedium)
                }
            }
        }

        Row(modifier = Modifier.fillMaxWidth().padding(12.dp), horizontalArrangement = Arrangement.SpaceBetween, verticalAlignment = Alignment.CenterVertically) {
            Column {
                Text(table.dealerName, style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
                Text(table.gameType, style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
            }
            Button(onClick = { }) {
                Icon(Icons.Default.PlayArrow, contentDescription = null); Spacer(modifier = Modifier.width(4.dp)); Text("Join Table")
            }
        }

        Divider()
        Text("Dealer Chat", style = MaterialTheme.typography.labelMedium, modifier = Modifier.padding(horizontal = 16.dp, vertical = 8.dp))

        LazyColumn(modifier = Modifier.weight(1f).padding(horizontal = 16.dp), verticalArrangement = Arrangement.spacedBy(4.dp)) {
            items(messages) { msg ->
                Text("${msg.sender}: ${msg.content}", style = MaterialTheme.typography.bodySmall)
            }
        }

        Surface(tonalElevation = 3.dp) {
            Row(modifier = Modifier.fillMaxWidth().padding(8.dp), verticalAlignment = Alignment.CenterVertically) {
                OutlinedTextField(value = chatText, onValueChange = onTextChange, modifier = Modifier.weight(1f), placeholder = { Text("Chat with dealer...") }, singleLine = true)
                Spacer(modifier = Modifier.width(8.dp))
                IconButton(onClick = onSend, enabled = chatText.isNotBlank()) {
                    Icon(Icons.Default.Send, contentDescription = "Send")
                }
            }
        }
    }
}

@Composable
private fun InfoChip(icon: androidx.compose.ui.graphics.vector.ImageVector, text: String) {
    Row(verticalAlignment = Alignment.CenterVertically) {
        Icon(icon, contentDescription = null, modifier = Modifier.size(14.dp), tint = MaterialTheme.colorScheme.onSurfaceVariant)
        Spacer(modifier = Modifier.width(4.dp))
        Text(text, style = MaterialTheme.typography.bodySmall)
    }
}
