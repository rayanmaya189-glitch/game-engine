package com.casino.game.ui.screens

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.input.KeyboardType
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

data class DepositState(
    val amount: String = "",
    val selectedMethod: String = "card",
    val balance: Double = 0.0,
    val isLoading: Boolean = false,
    val error: String? = null,
    val success: Boolean = false,
    val transactionId: String? = null
)

@HiltViewModel
class DepositViewModel @Inject constructor(
    private val walletRepository: WalletRepository
) : ViewModel() {
    private val _state = MutableStateFlow(DepositState())
    val state: StateFlow<DepositState> = _state.asStateFlow()

    init { loadBalance() }

    private fun loadBalance() {
        viewModelScope.launch {
            walletRepository.getBalance().fold(
                onSuccess = { _state.update { it.copy(balance = it.balance) } },
                onFailure = { }
            )
        }
    }

    fun setAmount(amount: String) { _state.update { it.copy(amount = amount) } }
    fun setMethod(method: String) { _state.update { it.copy(selectedMethod = method) } }

    fun deposit() {
        val amount = _state.value.amount.toDoubleOrNull() ?: return
        if (amount < 10) {
            _state.update { it.copy(error = "Minimum deposit is $10") }
            return
        }
        _state.update { it.copy(isLoading = true, error = null) }
        viewModelScope.launch {
            walletRepository.deposit(amount, _state.value.selectedMethod).fold(
                onSuccess = { response ->
                    _state.update {
                        it.copy(isLoading = false, success = true, transactionId = response.transactionId)
                    }
                },
                onFailure = { e ->
                    _state.update { it.copy(isLoading = false, error = e.message ?: "Deposit failed") }
                }
            )
        }
    }
}

@Composable
fun DepositScreen(
    onBack: () -> Unit,
    onDepositComplete: () -> Unit,
    viewModel: DepositViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()
    val methods = listOf("card" to "Credit Card", "crypto" to "Crypto", "bank" to "Bank Transfer")

    LaunchedEffect(state.success) { if (state.success) onDepositComplete() }

    Column(modifier = Modifier.fillMaxSize()) {
        TopAppBar(
            title = { Text("Deposit") },
            navigationIcon = {
                IconButton(onClick = onBack) {
                    Icon(Icons.Default.ArrowBack, contentDescription = "Back")
                }
            }
        )

        Column(
            modifier = Modifier
                .fillMaxSize()
                .padding(24.dp),
            verticalArrangement = Arrangement.spacedBy(16.dp)
        ) {
            Card(
                modifier = Modifier.fillMaxWidth(),
                colors = CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.primaryContainer)
            ) {
                Column(
                    modifier = Modifier.padding(20.dp),
                    horizontalAlignment = Alignment.CenterHorizontally
                ) {
                    Text("Current Balance", style = MaterialTheme.typography.bodyMedium)
                    Text(
                        "$${String.format("%,.2f", state.balance)}",
                        style = MaterialTheme.typography.headlineMedium
                    )
                }
            }

            OutlinedTextField(
                value = state.amount,
                onValueChange = viewModel::setAmount,
                label = { Text("Amount") },
                leadingIcon = { Icon(Icons.Default.AttachMoney, contentDescription = null) },
                modifier = Modifier.fillMaxWidth(),
                keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Decimal),
                singleLine = true,
                supportingText = { Text("Minimum: $10") }
            )

            Text("Payment Method", style = MaterialTheme.typography.labelLarge)
            methods.forEach { (key, label) ->
                Card(
                    onClick = { viewModel.setMethod(key) },
                    modifier = Modifier.fillMaxWidth(),
                    colors = if (state.selectedMethod == key)
                        CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.primaryContainer)
                    else CardDefaults.cardColors()
                ) {
                    Row(
                        modifier = Modifier.padding(16.dp),
                        verticalAlignment = Alignment.CenterVertically
                    ) {
                        RadioButton(
                            selected = state.selectedMethod == key,
                            onClick = { viewModel.setMethod(key) }
                        )
                        Spacer(modifier = Modifier.width(12.dp))
                        Icon(
                            imageVector = when (key) {
                                "card" -> Icons.Default.CreditCard
                                "crypto" -> Icons.Default.CurrencyBitcoin
                                else -> Icons.Default.AccountBalance
                            },
                            contentDescription = null
                        )
                        Spacer(modifier = Modifier.width(12.dp))
                        Text(label)
                    }
                }
            }

            if (state.error != null) {
                Text(state.error!!, color = MaterialTheme.colorScheme.error)
            }

            Button(
                onClick = viewModel::deposit,
                modifier = Modifier
                    .fillMaxWidth()
                    .height(56.dp),
                enabled = !state.isLoading && state.amount.isNotBlank()
            ) {
                if (state.isLoading) {
                    CircularProgressIndicator(
                        modifier = Modifier.size(24.dp),
                        color = MaterialTheme.colorScheme.onPrimary
                    )
                } else {
                    Icon(Icons.Default.Add, contentDescription = null)
                    Spacer(modifier = Modifier.width(8.dp))
                    Text("Deposit")
                }
            }
        }
    }
}
