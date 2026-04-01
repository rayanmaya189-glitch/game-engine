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
import com.casino.game.data.repository.WalletRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class WithdrawState(
    val amount: String = "",
    val accountInfo: String = "",
    val selectedMethod: String = "bank",
    val availableBalance: Double = 0.0,
    val isLoading: Boolean = false,
    val error: String? = null,
    val success: Boolean = false
)

@HiltViewModel
class WithdrawViewModel @Inject constructor(
    private val walletRepository: WalletRepository
) : ViewModel() {
    private val _state = MutableStateFlow(WithdrawState())
    val state: StateFlow<WithdrawState> = _state.asStateFlow()

    init { loadBalance() }

    private fun loadBalance() {
        viewModelScope.launch {
            walletRepository.getBalance().fold(
                onSuccess = { _state.update { it.copy(availableBalance = it.availableBalance) } },
                onFailure = { }
            )
        }
    }

    fun setAmount(amount: String) { _state.update { it.copy(amount = amount) } }
    fun setAccountInfo(info: String) { _state.update { it.copy(accountInfo = info) } }
    fun setMethod(method: String) { _state.update { it.copy(selectedMethod = method) } }

    fun withdraw() {
        val amount = _state.value.amount.toDoubleOrNull() ?: return
        if (amount < 20) {
            _state.update { it.copy(error = "Minimum withdrawal is $20") }
            return
        }
        if (amount > _state.value.availableBalance) {
            _state.update { it.copy(error = "Insufficient balance") }
            return
        }
        if (_state.value.accountInfo.isBlank()) {
            _state.update { it.copy(error = "Please enter account information") }
            return
        }
        _state.update { it.copy(isLoading = true, error = null) }
        viewModelScope.launch {
            walletRepository.withdraw(amount, _state.value.selectedMethod, _state.value.accountInfo).fold(
                onSuccess = {
                    _state.update { it.copy(isLoading = false, success = true) }
                },
                onFailure = { e ->
                    _state.update { it.copy(isLoading = false, error = e.message ?: "Withdrawal failed") }
                }
            )
        }
    }
}

@Composable
fun WithdrawScreen(
    onBack: () -> Unit,
    onWithdrawComplete: () -> Unit,
    viewModel: WithdrawViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()
    val methods = listOf("bank" to "Bank Transfer", "crypto" to "Crypto Wallet")

    LaunchedEffect(state.success) { if (state.success) onWithdrawComplete() }

    Column(modifier = Modifier.fillMaxSize()) {
        TopAppBar(
            title = { Text("Withdraw") },
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
                    Text("Available Balance", style = MaterialTheme.typography.bodyMedium)
                    Text(
                        "$${String.format("%,.2f", state.availableBalance)}",
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
                supportingText = { Text("Min: $20 | Max: $${String.format("%,.0f", state.availableBalance)}") }
            )

            Text("Withdrawal Method", style = MaterialTheme.typography.labelLarge)
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
                        Text(label)
                    }
                }
            }

            OutlinedTextField(
                value = state.accountInfo,
                onValueChange = viewModel::setAccountInfo,
                label = { Text(if (state.selectedMethod == "bank") "Bank Account Number" else "Wallet Address") },
                leadingIcon = {
                    Icon(
                        if (state.selectedMethod == "bank") Icons.Default.AccountBalance else Icons.Default.Wallet,
                        contentDescription = null
                    )
                },
                modifier = Modifier.fillMaxWidth(),
                singleLine = true
            )

            if (state.error != null) {
                Text(state.error!!, color = MaterialTheme.colorScheme.error)
            }

            Button(
                onClick = viewModel::withdraw,
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
                    Icon(Icons.Default.Remove, contentDescription = null)
                    Spacer(modifier = Modifier.width(8.dp))
                    Text("Withdraw")
                }
            }
        }
    }
}
