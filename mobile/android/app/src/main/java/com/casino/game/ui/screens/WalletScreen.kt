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
import com.casino.game.data.repository.WalletRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class PlayerWalletState(
    val balance: Double = 0.0,
    val bonusBalance: Double = 0.0,
    val transactions: List<Transaction> = emptyList(),
    val selectedFilter: String? = null,
    val isLoading: Boolean = false,
    val error: String? = null
)

@HiltViewModel
class PlayerWalletViewModel @Inject constructor(
    private val walletRepository: WalletRepository
) : ViewModel() {
    private val _state = MutableStateFlow(PlayerWalletState())
    val state: StateFlow<PlayerWalletState> = _state.asStateFlow()

    init { loadWallet() }

    fun loadWallet() {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            walletRepository.getBalance().fold(
                onSuccess = { b ->
                    _state.update { it.copy(balance = b.balance, bonusBalance = b.bonusBalance) }
                },
                onFailure = { }
            )
            walletRepository.getTransactions().fold(
                onSuccess = { r ->
                    _state.update { it.copy(transactions = r.transactions, isLoading = false) }
                },
                onFailure = { e ->
                    _state.update { it.copy(isLoading = false, error = e.message) }
                }
            )
        }
    }

    fun setFilter(filter: String?) { _state.update { it.copy(selectedFilter = filter) } }
}

@Composable
fun WalletScreen(
    onNavigateToDeposit: () -> Unit,
    onNavigateToWithdraw: () -> Unit,
    viewModel: PlayerWalletViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()
    val filters = listOf(null to "All", "deposit" to "Deposits", "withdraw" to "Withdrawals", "win" to "Winnings", "bet" to "Bets")

    val filteredTxns = if (state.selectedFilter != null) {
        state.transactions.filter { it.type == state.selectedFilter }
    } else state.transactions

    LazyColumn(
        modifier = Modifier.fillMaxSize(),
        contentPadding = PaddingValues(16.dp),
        verticalArrangement = Arrangement.spacedBy(12.dp)
    ) {
        item {
            Card(
                modifier = Modifier.fillMaxWidth(),
                colors = CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.primary)
            ) {
                Column(
                    modifier = Modifier.padding(24.dp).fillMaxWidth(),
                    horizontalAlignment = Alignment.CenterHorizontally
                ) {
                    Text("Total Balance", color = MaterialTheme.colorScheme.onPrimary.copy(alpha = 0.8f))
                    Text(
                        "$${String.format("%,.2f", state.balance)}",
                        style = MaterialTheme.typography.headlineLarge,
                        color = MaterialTheme.colorScheme.onPrimary
                    )
                    if (state.bonusBalance > 0) {
                        Text(
                            "Bonus: $${String.format("%,.2f", state.bonusBalance)}",
                            style = MaterialTheme.typography.bodyMedium,
                            color = MaterialTheme.colorScheme.onPrimary.copy(alpha = 0.7f)
                        )
                    }
                }
            }
        }

        item {
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.spacedBy(12.dp)
            ) {
                Button(
                    onClick = onNavigateToDeposit,
                    modifier = Modifier.weight(1f)
                ) {
                    Icon(Icons.Default.Add, contentDescription = null)
                    Spacer(modifier = Modifier.width(4.dp))
                    Text("Deposit")
                }
                OutlinedButton(
                    onClick = onNavigateToWithdraw,
                    modifier = Modifier.weight(1f)
                ) {
                    Icon(Icons.Default.Remove, contentDescription = null)
                    Spacer(modifier = Modifier.width(4.dp))
                    Text("Withdraw")
                }
            }
        }

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

        item {
            Text("Transactions", style = MaterialTheme.typography.titleMedium)
        }

        if (state.isLoading) {
            item {
                Box(
                    modifier = Modifier.fillMaxWidth().padding(32.dp),
                    contentAlignment = Alignment.Center
                ) { CircularProgressIndicator() }
            }
        } else {
            items(filteredTxns) { transaction ->
                WalletTransactionItem(transaction = transaction)
            }
        }
    }
}

@Composable
private fun WalletTransactionItem(transaction: Transaction) {
    Card(modifier = Modifier.fillMaxWidth()) {
        Row(
            modifier = Modifier.padding(12.dp),
            horizontalArrangement = Arrangement.SpaceBetween,
            verticalAlignment = Alignment.CenterVertically
        ) {
            Row(verticalAlignment = Alignment.CenterVertically) {
                Icon(
                    imageVector = when (transaction.type) {
                        "deposit" -> Icons.Default.ArrowDownward
                        "withdraw" -> Icons.Default.ArrowUpward
                        "win" -> Icons.Default.EmojiEvents
                        else -> Icons.Default.Casino
                    },
                    contentDescription = null,
                    tint = when (transaction.type) {
                        "deposit", "win" -> MaterialTheme.colorScheme.primary
                        "withdraw" -> MaterialTheme.colorScheme.error
                        else -> MaterialTheme.colorScheme.onSurface
                    }
                )
                Spacer(modifier = Modifier.width(12.dp))
                Column {
                    Text(
                        transaction.type.replaceFirstChar { it.uppercase() },
                        style = MaterialTheme.typography.bodyMedium
                    )
                    Text(transaction.createdAt, style = MaterialTheme.typography.bodySmall)
                }
            }
            Text(
                "${if (transaction.type in listOf("deposit", "win")) "+" else "-"}$${String.format("%,.2f", transaction.amount)}",
                style = MaterialTheme.typography.titleMedium,
                color = when (transaction.type) {
                    "deposit", "win" -> MaterialTheme.colorScheme.primary
                    else -> MaterialTheme.colorScheme.error
                }
            )
        }
    }
}
