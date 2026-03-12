package com.game_engine.casino.ui.games

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.game_engine.casino.data.model.Game
import com.game_engine.casino.data.model.GameCategory
import com.game_engine.casino.data.repository.GameRepository
import com.game_engine.casino.util.Resource
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class GamesUiState(
    val isLoading: Boolean = false,
    val games: List<Game> = emptyList(),
    val categories: List<GameCategory> = emptyList(),
    val selectedCategory: String? = null,
    val searchQuery: String = "",
    val error: String? = null,
    val currentPage: Int = 1,
    val totalPages: Int = 1
)

@HiltViewModel
class GamesViewModel @Inject constructor(
    private val gameRepository: GameRepository
) : ViewModel() {
    
    private val _uiState = MutableStateFlow(GamesUiState())
    val uiState: StateFlow<GamesUiState> = _uiState.asStateFlow()
    
    init {
        loadCategories()
        loadGames()
    }
    
    fun loadGames(category: String? = null, search: String? = null) {
        viewModelScope.launch {
            gameRepository.getGames(
                page = 1,
                category = category,
                search = search
            ).collect { result ->
                when (result) {
                    is Resource.Loading -> {
                        _uiState.update { it.copy(isLoading = true) }
                    }
                    is Resource.Success -> {
                        result.data?.let { response ->
                            _uiState.update {
                                it.copy(
                                    isLoading = false,
                                    games = response.games,
                                    currentPage = response.page,
                                    totalPages = response.totalPages,
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
    
    fun selectCategory(category: String?) {
        _uiState.update { it.copy(selectedCategory = category) }
        loadGames(category = category)
    }
    
    fun search(query: String) {
        _uiState.update { it.copy(searchQuery = query) }
        loadGames(search = query.ifEmpty { null })
    }
    
    fun loadMore() {
        val currentState = _uiState.value
        if (currentState.currentPage < currentState.totalPages && !currentState.isLoading) {
            viewModelScope.launch {
                gameRepository.getGames(
                    page = currentState.currentPage + 1,
                    category = currentState.selectedCategory,
                    search = currentState.searchQuery.ifEmpty { null }
                ).collect { result ->
                    when (result) {
                        is Resource.Success -> {
                            result.data?.let { response ->
                                _uiState.update {
                                    it.copy(
                                        games = it.games + response.games,
                                        currentPage = response.page,
                                        totalPages = response.totalPages
                                    )
                                }
                            }
                        }
                        else -> {}
                    }
                }
            }
        }
    }
}
