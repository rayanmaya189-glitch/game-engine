package com.game-engine.casino.ui.wallet

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.game-engine.casino.data.model.*
import com.game-engine.casino.data.repository.WalletRepository
import com.game-engine.casino.util.Resource
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class WalletUiState(
    val isLoading: Boolean = false,
    val wallet: Wallet? = null,
    val balance: WalletBalance? = null,
    val transactions: List<Transaction> = emptyList(),
    val paymentMethods: PaymentMethodsResponse? = null,
    val error: String? = null
)

@HiltViewModel
class WalletViewModel @Inject constructor(
    private val walletRepository: WalletRepository
) : ViewModel() {
    
    private val _uiState = MutableStateFlow(WalletUiState())
    val uiState: StateFlow<WalletUiState> = _uiState.asStateFlow()
    
    init {
        loadWallet()
        loadTransactions()
        loadPaymentMethods()
    }
    
    fun loadWallet() {
        viewModelScope.launch {
            walletRepository.getWallet().collect { result ->
                when (result) {
                    is Resource.Loading -> {
                        _uiState.update { it.copy(isLoading = true) }
                    }
                    is Resource.Success -> {
                        _uiState.update {
                            it.copy(isLoading = false, wallet = result.data)
                        }
                    }
                    is Resource.Error -> {
                        _uiState.update {
                            it.copy(isLoading = false, error = result.message)
                        }
                    }
                }
            }
        }
        
        viewModelScope.launch {
            walletRepository.getBalance().collect { result ->
                when (result) {
                    is Resource.Success -> {
                        _uiState.update { it.copy(balance = result.data) }
                    }
                    else -> {}
                }
            }
        }
    }
    
    fun loadTransactions() {
        viewModelScope.launch {
            walletRepository.getTransactions().collect { result ->
                when (result) {
                    is Resource.Success -> {
                        _uiState.update { it.copy(transactions = result.data?.transactions ?: emptyList()) }
                    }
                    else -> {}
                }
            }
        }
    }
    
    fun loadPaymentMethods() {
        viewModelScope.launch {
            walletRepository.getPaymentMethods().collect { result ->
                when (result) {
                    is Resource.Success -> {
                        _uiState.update { it.copy(paymentMethods = result.data) }
                    }
                    else -> {}
                }
            }
        }
    }
}
