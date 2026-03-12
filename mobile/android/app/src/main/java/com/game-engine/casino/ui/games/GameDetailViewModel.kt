package com.game-engine.casino.ui.games

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.game-engine.casino.data.model.Game
import com.game-engine.casino.data.repository.GameRepository
import com.game-engine.casino.util.Resource
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class GameDetailUiState(
    val isLoading: Boolean = false,
    val game: Game? = null,
    val isFavorite: Boolean = false,
    val gameSession: com.game-engine.casino.data.model.GameSession? = null,
    val error: String? = null
)

@HiltViewModel
class GameDetailViewModel @Inject constructor(
    private val gameRepository: GameRepository
) : ViewModel() {
    
    private val _uiState = MutableStateFlow(GameDetailUiState())
    val uiState: StateFlow<GameDetailUiState> = _uiState.asStateFlow()
    
    fun loadGame(gameId: String) {
        viewModelScope.launch {
            gameRepository.getGame(gameId).collect { result ->
                when (result) {
                    is Resource.Loading -> {
                        _uiState.update { it.copy(isLoading = true) }
                    }
                    is Resource.Success -> {
                        result.data?.let { game ->
                            _uiState.update {
                                it.copy(
                                    isLoading = false,
                                    game = game,
                                    isFavorite = game.isFavorite,
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
    
    fun toggleFavorite(gameId: String) {
        val currentFavorite = _uiState.value.isFavorite
        viewModelScope.launch {
            val flow = if (currentFavorite) {
                gameRepository.removeFromFavorites(gameId)
            } else {
                gameRepository.addToFavorites(gameId)
            }
            
            flow.collect { result ->
                when (result) {
                    is Resource.Success -> {
                        _uiState.update { it.copy(isFavorite = !currentFavorite) }
                    }
                    else -> {}
                }
            }
        }
    }
    
    fun startGame(gameId: String, isFunPlay: Boolean) {
        viewModelScope.launch {
            gameRepository.startGame(gameId, isFunPlay).collect { result ->
                when (result) {
                    is Resource.Loading -> {
                        _uiState.update { it.copy(isLoading = true) }
                    }
                    is Resource.Success -> {
                        result.data?.let { session ->
                            _uiState.update {
                                it.copy(
                                    isLoading = false,
                                    gameSession = session
                                )
                            }
                            // Here you would typically open a WebView with the playUrl
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
}
