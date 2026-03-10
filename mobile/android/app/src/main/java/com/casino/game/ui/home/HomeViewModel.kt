package com.casino.game.ui.home

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.casino.game.data.model.*
import com.casino.game.data.repository.*
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class HomeState(
    val isLoading: Boolean = false,
    val balance: BalanceResponse? = null,
    val featuredGames: List<Game> = emptyList(),
    val popularGames: List<Game> = emptyList(),
    val currentJackpots: List<Jackpot> = emptyList(),
    val activeTournaments: List<Tournament> = emptyList(),
    val error: String? = null
)

@HiltViewModel
class HomeViewModel @Inject constructor(
    private val walletRepository: WalletRepository,
    private val gameRepository: GameRepository,
    private val jackpotRepository: JackpotRepository,
    private val tournamentRepository: TournamentRepository
) : ViewModel() {

    private val _state = MutableStateFlow(HomeState())
    val state: StateFlow<HomeState> = _state.asStateFlow()

    init {
        loadHomeData()
    }

    fun loadHomeData() {
        viewModelScope.launch {
            _state.update { it.copy(isLoading = true) }

            // Load balance
            walletRepository.getBalance().fold(
                onSuccess = { balance ->
                    _state.update { it.copy(balance = balance) }
                },
                onFailure = { /* Ignore */ }
            )

            // Load featured games
            gameRepository.getFeaturedGames().fold(
                onSuccess = { response ->
                    _state.update { it.copy(featuredGames = response.games) }
                },
                onFailure = { /* Ignore */ }
            )

            // Load popular games
            gameRepository.getPopularGames().fold(
                onSuccess = { response ->
                    _state.update { it.copy(popularGames = response.games) }
                },
                onFailure = { /* Ignore */ }
            )

            // Load current jackpots
            jackpotRepository.getCurrentJackpots().fold(
                onSuccess = { response ->
                    _state.update { it.copy(currentJackpots = response.jackpots) }
                },
                onFailure = { /* Ignore */ }
            )

            // Load active tournaments
            tournamentRepository.getTournaments("active").fold(
                onSuccess = { response ->
                    _state.update { it.copy(activeTournaments = response.tournaments) }
                },
                onFailure = { /* Ignore */ }
            )

            _state.update { it.copy(isLoading = false) }
        }
    }

    fun refresh() {
        loadHomeData()
    }
}
