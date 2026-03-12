package com.game-engine.casino.ui.home

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.game-engine.casino.data.model.Game
import com.game-engine.casino.data.model.GameCategory
import com.game-engine.casino.data.model.WalletBalance
import com.game-engine.casino.data.repository.AuthRepository
import com.game-engine.casino.data.repository.GameRepository
import com.game-engine.casino.data.repository.WalletRepository
import com.game-engine.casino.util.Resource
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class HomeUiState(
    val isLoading: Boolean = false,
    val user: com.game-engine.casino.data.model.User? = null,
    val balance: WalletBalance? = null,
    val featuredGames: List<Game> = emptyList(),
    val popularGames: List<Game> = emptyList(),
    val jackpotGames: List<Game> = emptyList(),
    val categories: List<GameCategory> = emptyList(),
    val error: String? = null
)

@HiltViewModel
class HomeViewModel @Inject constructor(
    private val authRepository: AuthRepository,
    private val gameRepository: GameRepository,
    private val walletRepository: WalletRepository
) : ViewModel() {
    
    private val _uiState = MutableStateFlow(HomeUiState())
    val uiState: StateFlow<HomeUiState> = _uiState.asStateFlow()
    
    init {
        loadData()
    }
    
    fun loadData() {
        loadUser()
        loadBalance()
        loadFeaturedGames()
        loadCategories()
    }
    
    private fun loadUser() {
        viewModelScope.launch {
            authRepository.getCurrentUser().collect { result ->
                when (result) {
                    is Resource.Success -> {
                        _uiState.update { it.copy(user = result.data) }
                    }
                    else -> {}
                }
            }
        }
    }
    
    private fun loadBalance() {
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
    
    private fun loadFeaturedGames() {
        viewModelScope.launch {
            gameRepository.getFeaturedGames().collect { result ->
                when (result) {
                    is Resource.Loading -> {
                        _uiState.update { it.copy(isLoading = true) }
                    }
                    is Resource.Success -> {
                        result.data?.let { response ->
                            _uiState.update {
                                it.copy(
                                    isLoading = false,
                                    featuredGames = response.featured,
                                    popularGames = response.popular,
                                    jackpotGames = response.jackpot,
                                    error = null
                                )
                            }
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
    }
    
    private fun loadCategories() {
        viewModelScope.launch {
            gameRepository.getCategories().collect { result ->
                when (result) {
                    is Resource.Success -> {
                        _uiState.update { it.copy(categories = result.data ?: emptyList()) }
                    }
                    else -> {}
                }
            }
        }
    }
    
    fun refresh() {
        loadData()
    }
}
