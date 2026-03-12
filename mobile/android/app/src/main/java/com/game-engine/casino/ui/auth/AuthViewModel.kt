package com.game-engine.casino.ui.auth

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.game-engine.casino.data.repository.AuthRepository
import com.game-engine.casino.util.Resource
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch
import javax.inject.Inject

data class AuthUiState(
    val isLoading: Boolean = false,
    val isLoggedIn: Boolean = false,
    val error: String? = null,
    val user: com.game-engine.casino.data.model.User? = null
)

@HiltViewModel
class AuthViewModel @Inject constructor(
    private val authRepository: AuthRepository
) : ViewModel() {
    
    private val _uiState = MutableStateFlow(AuthUiState())
    val uiState: StateFlow<AuthUiState> = _uiState.asStateFlow()
    
    fun login(email: String, password: String) {
        viewModelScope.launch {
            authRepository.login(email, password, null, null).collect { result ->
                when (result) {
                    is Resource.Loading -> {
                        _uiState.update { it.copy(isLoading = true, error = null) }
                    }
                    is Resource.Success -> {
                        _uiState.update {
                            it.copy(
                                isLoading = false,
                                isLoggedIn = true,
                                user = result.data?.user,
                                error = null
                            )
                        }
                    }
                    is Resource.Error -> {
                        _uiState.update {
                            it.copy(
                                isLoading = false,
                                error = result.message
                            )
                        }
                    }
                }
            }
        }
    }
    
    fun register(
        email: String,
        username: String,
        password: String,
        phone: String?,
        currency: String,
        referralCode: String?
    ) {
        viewModelScope.launch {
            authRepository.register(email, username, password, phone, currency, referralCode).collect { result ->
                when (result) {
                    is Resource.Loading -> {
                        _uiState.update { it.copy(isLoading = true, error = null) }
                    }
                    is Resource.Success -> {
                        _uiState.update {
                            it.copy(
                                isLoading = false,
                                isLoggedIn = true,
                                user = result.data?.user,
                                error = null
                            )
                        }
                    }
                    is Resource.Error -> {
                        _uiState.update {
                            it.copy(
                                isLoading = false,
                                error = result.message
                            )
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
                    is Resource.Loading -> {
                        _uiState.update { it.copy(isLoading = true) }
                    }
                    is Resource.Success -> {
                        _uiState.update { AuthUiState() }
                    }
                    is Resource.Error -> {
                        _uiState.update { AuthUiState() }
                    }
                }
            }
        }
    }
    
    fun resetPassword(email: String) {
        viewModelScope.launch {
            authRepository.resetPassword(email).collect { result ->
                when (result) {
                    is Resource.Loading -> {
                        _uiState.update { it.copy(isLoading = true, error = null) }
                    }
                    is Resource.Success -> {
                        _uiState.update {
                            it.copy(
                                isLoading = false,
                                error = null
                            )
                        }
                    }
                    is Resource.Error -> {
                        _uiState.update {
                            it.copy(
                                isLoading = false,
                                error = result.message
                            )
                        }
                    }
                }
            }
        }
    }
    
    fun loginWithBiometric() {
        // Implementation for biometric login
        // This would typically use BiometricPrompt
    }
    
    fun clearError() {
        _uiState.update { it.copy(error = null) }
    }
}
