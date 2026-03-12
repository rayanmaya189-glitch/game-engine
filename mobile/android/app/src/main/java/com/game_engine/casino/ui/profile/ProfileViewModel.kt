package com.game_engine.casino.ui.profile

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.game_engine.casino.data.model.User
import com.game_engine.casino.data.repository.AuthRepository
import com.game_engine.casino.util.Resource
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class ProfileUiState(
    val isLoading: Boolean = false,
    val user: User? = null,
    val isLoggedOut: Boolean = false,
    val error: String? = null
)

@HiltViewModel
class ProfileViewModel @Inject constructor(
    private val authRepository: AuthRepository
) : ViewModel() {
    
    private val _uiState = MutableStateFlow(ProfileUiState())
    val uiState: StateFlow<ProfileUiState> = _uiState.asStateFlow()
    
    init {
        loadUser()
    }
    
    private fun loadUser() {
        viewModelScope.launch {
            authRepository.getCurrentUser().collect { result ->
                when (result) {
                    is Resource.Loading -> {
                        _uiState.update { it.copy(isLoading = true) }
                    }
                    is Resource.Success -> {
                        _uiState.update {
                            it.copy(isLoading = false, user = result.data)
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
    
    fun logout() {
        viewModelScope.launch {
            authRepository.logout().collect { result ->
                when (result) {
                    is Resource.Success -> {
                        _uiState.update { it.copy(isLoggedOut = true) }
                    }
                    else -> {
                        _uiState.update { it.copy(isLoggedOut = true) }
                    }
                }
            }
        }
    }
    
    fun updateProfile(
        firstName: String?,
        lastName: String?,
        phone: String?,
        language: String?,
        currency: String?
    ) {
        viewModelScope.launch {
            authRepository.updateProfile(firstName, lastName, phone, language, currency).collect { result ->
                when (result) {
                    is Resource.Loading -> {
                        _uiState.update { it.copy(isLoading = true) }
                    }
                    is Resource.Success -> {
                        _uiState.update {
                            it.copy(isLoading = false, user = result.data)
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
