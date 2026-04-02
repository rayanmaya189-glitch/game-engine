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

data class PaymentHistoryState(
    val transactions: List<PaymentTransaction> = emptyList(),
    val selectedFilter: String = "all",
    val isLoading: Boolean = false,
    val isRefreshing: Boolean = false,
    val error: String? = null
)

@HiltViewModel
class PaymentHistoryViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(PaymentHistoryState())
    val state: StateFlow<PaymentHistoryState> = _state.asStateFlow()

    init { loadHistory() }

    fun loadHistory(refresh: Boolean = false) {
        _state.update { if (refresh) it.copy(isRefreshing = true) else it.copy(isLoading = true) }
        viewModelScope.launch {
            try {
                val response = apiService.getPaymentHistory()
                if (response.isSuccessful) {
                    _state.update { it.copy(transactions = response.body()?.transactions ?: emptyList(), isLoading = false, isRefreshing = false, error = null) }
                } else {
                    _state.update { it.copy(isLoading = false, isRefreshing = false, error = "Failed to load history") }
                }
            } catch (e: Exception) {
                _state.update { it.copy(isLoading = false, isRefreshing = false, error = e.message) }
            }
        }
    }

    fun setFilter(filter: String) { _state.update { it.copy(selectedFilter = filter) } }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun PaymentHistoryScreen(
    onBack: () -> Unit,
    viewModel: PaymentHistoryViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()
    val filters = listOf("all" to "All", "deposit" to "Deposits", "withdrawal" to "Withdrawals", "bet" to "Bets", "win" to "Wins", "bonus" to "Bonuses")

    val filteredTxns = if (state.selectedFilter == "all") state.transactions
    else state.transactions.filter { it.type == state.selectedFilter }

    Column(modifier = Modifier.fillMaxSize()) {
        TopAppBar(
            title = { Text("Payment History") },
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

                if (filteredTxns.isEmpty()) {
                    item {
                        Box(
                            modifier = Modifier.fillMaxWidth().padding(48.dp),
                            contentAlignment = Alignment.Center
                        ) {
                            Column(horizontalAlignment = Alignment.CenterHorizontally) {
                                Icon(Icons.Default.Receipt, contentDescription = null, modifier = Modifier.size(64.dp), tint = MaterialTheme.colorScheme.onSurfaceVariant)
                                Spacer(modifier = Modifier.height(8.dp))
                                Text("No transactions found", style = MaterialTheme.typography.bodyLarge, color = MaterialTheme.colorScheme.onSurfaceVariant)
                            }
                        }
                    }
                } else {
                    items(filteredTxns) { txn -> PaymentTransactionItem(txn) }
                }
            }
        }
    }
}

@Composable
private fun PaymentTransactionItem(txn: PaymentTransaction) {
    Card(modifier = Modifier.fillMaxWidth()) {
        Row(
            modifier = Modifier.padding(12.dp),
            horizontalArrangement = Arrangement.SpaceBetween,
            verticalAlignment = Alignment.CenterVertically
        ) {
            Row(verticalAlignment = Alignment.CenterVertically, modifier = Modifier.weight(1f)) {
                Icon(
                    imageVector = when (txn.type) {
                        "deposit" -> Icons.Default.ArrowDownward
                        "withdrawal" -> Icons.Default.ArrowUpward
                        "win" -> Icons.Default.EmojiEvents
                        "bonus" -> Icons.Default.CardGiftcard
                        else -> Icons.Default.Casino
                    },
                    contentDescription = null,
                    tint = when (txn.type) {
                        "deposit", "win", "bonus" -> MaterialTheme.colorScheme.primary
                        "withdrawal" -> MaterialTheme.colorScheme.error
                        else -> MaterialTheme.colorScheme.onSurface
                    }
                )
                Spacer(modifier = Modifier.width(12.dp))
                Column {
                    Text(txn.description, style = MaterialTheme.typography.bodyMedium, maxLines = 1)
                    Text(txn.date, style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
                }
            }
            Column(horizontalAlignment = Alignment.End) {
                Text(
                    "${if (txn.type in listOf("deposit", "win", "bonus")) "+" else "-"}$${String.format("%,.2f", txn.amount)}",
                    style = MaterialTheme.typography.titleMedium,
                    color = when (txn.type) {
                        "deposit", "win", "bonus" -> MaterialTheme.colorScheme.primary
                        else -> MaterialTheme.colorScheme.error
                    }
                )
                AssistChip(
                    onClick = { },
                    label = { Text(txn.status.replaceFirstChar { it.uppercase() }, style = MaterialTheme.typography.labelSmall) },
                    colors = AssistChipDefaults.assistChipColors(
                        containerColor = when (txn.status) {
                            "completed" -> MaterialTheme.colorScheme.primaryContainer
                            "pending" -> MaterialTheme.colorScheme.tertiaryContainer
                            else -> MaterialTheme.colorScheme.errorContainer
                        }
                    )
                )
            }
        }
    }
}
