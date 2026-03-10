package com.casino.game.ui.auth

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.casino.game.data.model.LoginResponse
import com.casino.game.data.repository.AuthRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class AuthState(
    val isLoading: Boolean = false,
    val isLoggedIn: Boolean = false,
    val user: LoginResponse? = null,
    val error: String? = null
)

@HiltViewModel
class AuthViewModel @Inject constructor(
    private val authRepository: AuthRepository
) : ViewModel() {

    private val _state = MutableStateFlow(AuthState())
    val state: StateFlow<AuthState> = _state.asStateFlow()

    init {
        viewModelScope.launch {
            authRepository.isLoggedIn.collect { isLoggedIn ->
                _state.update { it.copy(isLoggedIn = isLoggedIn) }
            }
        }
    }

    fun login(email: String, password: String) {
        viewModelScope.launch {
            _state.update { it.copy(isLoading = true, error = null) }
            
            val result = authRepository.login(email, password)
            result.fold(
                onSuccess = { loginResponse ->
                    _state.update { 
                        it.copy(
                            isLoading = false,
                            isLoggedIn = true,
                            user = loginResponse
                        ) 
                    }
                },
                onFailure = { exception ->
                    _state.update { 
                        it.copy(
                            isLoading = false,
                            error = exception.message ?: "Login failed"
                        ) 
                    }
                }
            )
        }
    }

    fun register(email: String, password: String, username: String, phone: String?) {
        viewModelScope.launch {
            _state.update { it.copy(isLoading = true, error = null) }
            
            val result = authRepository.register(email, password, username, phone)
            result.fold(
                onSuccess = { loginResponse ->
                    _state.update { 
                        it.copy(
                            isLoading = false,
                            isLoggedIn = true,
                            user = loginResponse
                        ) 
                    }
                },
                onFailure = { exception ->
                    _state.update { 
                        it.copy(
                            isLoading = false,
                            error = exception.message ?: "Registration failed"
                        ) 
                    }
                }
            )
        }
    }

    fun logout() {
        viewModelScope.launch {
            authRepository.logout()
            _state.update { AuthState() }
        }
    }

    fun clearError() {
        _state.update { it.copy(error = null) }
    }
}
