package com.casino.game.ui

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.casino.game.data.model.*
import com.casino.game.data.repository.GameRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class GamesViewModel @Inject constructor(
    private val gameRepository: GameRepository
) : ViewModel() {
    private val _state = MutableStateFlow(GamesState())
    val state: StateFlow<GamesState> = _state.asStateFlow()

    init {
        loadGames()
        loadCategories()
    }

    fun loadGames(category: String? = null) {
        viewModelScope.launch {
            _state.update { it.copy(isLoading = true) }
            gameRepository.getGames(category = category).fold(
                onSuccess = { response ->
                    _state.update { it.copy(games = response.games, isLoading = false) }
                },
                onFailure = {
                    _state.update { it.copy(isLoading = false) }
                }
            )
        }
    }

    private fun loadCategories() {
        viewModelScope.launch {
            gameRepository.getCategories().fold(
                onSuccess = { response ->
                    _state.update { it.copy(categories = response.categories) }
                },
                onFailure = { }
            )
        }
    }

    fun searchGames(query: String) {
        _state.update { it.copy(searchQuery = query) }
        loadGames(searchQuery = query)
    }

    private fun loadGames(searchQuery: String? = null) {
        viewModelScope.launch {
            _state.update { it.copy(isLoading = true) }
            gameRepository.getGames(search = searchQuery).fold(
                onSuccess = { response ->
                    _state.update { it.copy(games = response.games, isLoading = false) }
                },
                onFailure = {
                    _state.update { it.copy(isLoading = false) }
                }
            )
        }
    }
}

data class GamesState(
    val isLoading: Boolean = false,
    val games: List<Game> = emptyList(),
    val categories: List<Category> = emptyList(),
    val searchQuery: String = "",
    val error: String? = null
)
