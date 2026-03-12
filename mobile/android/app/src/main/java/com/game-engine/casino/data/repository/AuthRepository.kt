package com.game-engine.casino.data.repository

import com.game-engine.casino.data.api.AuthApi
import com.game-engine.casino.data.model.*
import com.game-engine.casino.util.PreferencesManager
import com.game-engine.casino.util.Resource
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class AuthRepository @Inject constructor(
    private val authApi: AuthApi,
    private val preferencesManager: PreferencesManager
) {
    fun login(email: String, password: String, deviceId: String?, deviceName: String?): Flow<Resource<LoginResponse>> = flow {
        emit(Resource.Loading())
        try {
            val response = authApi.login(LoginRequest(email, password, deviceId, deviceName))
            if (response.isSuccessful) {
                response.body()?.let { loginResponse ->
                    preferencesManager.saveAccessToken(loginResponse.accessToken)
                    preferencesManager.saveRefreshToken(loginResponse.refreshToken)
                    preferencesManager.saveUserId(loginResponse.user.id)
                    emit(Resource.Success(loginResponse))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Login failed"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun register(
        email: String,
        username: String,
        password: String,
        phone: String?,
        currency: String,
        referralCode: String?
    ): Flow<Resource<RegisterResponse>> = flow {
        emit(Resource.Loading())
        try {
            val request = RegisterRequest(email, username, password, phone, currency, referralCode)
            val response = authApi.register(request)
            if (response.isSuccessful) {
                response.body()?.let { registerResponse ->
                    preferencesManager.saveAccessToken(registerResponse.accessToken)
                    preferencesManager.saveRefreshToken(registerResponse.refreshToken)
                    preferencesManager.saveUserId(registerResponse.user.id)
                    emit(Resource.Success(registerResponse))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Registration failed"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun logout(): Flow<Resource<Unit>> = flow {
        emit(Resource.Loading())
        try {
            authApi.logout()
            preferencesManager.clearTokens()
            emit(Resource.Success(Unit))
        } catch (e: Exception) {
            preferencesManager.clearTokens()
            emit(Resource.Success(Unit))
        }
    }
    
    fun refreshToken(): Flow<Resource<Unit>> = flow {
        emit(Resource.Loading())
        try {
            val refreshToken = preferencesManager.getRefreshToken()
            if (refreshToken.isNullOrEmpty()) {
                emit(Resource.Error("No refresh token"))
                return@flow
            }
            val response = authApi.refreshToken(RefreshTokenRequest(refreshToken))
            if (response.isSuccessful) {
                response.body()?.let { refreshResponse ->
                    preferencesManager.saveAccessToken(refreshResponse.accessToken)
                    preferencesManager.saveRefreshToken(refreshResponse.refreshToken)
                    emit(Resource.Success(Unit))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error("Token refresh failed"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun getCurrentUser(): Flow<Resource<User>> = flow {
        emit(Resource.Loading())
        try {
            val response = authApi.getCurrentUser()
            if (response.isSuccessful) {
                response.body()?.let { user ->
                    emit(Resource.Success(user))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to get user"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun updateProfile(
        firstName: String?,
        lastName: String?,
        phone: String?,
        language: String?,
        currency: String?
    ): Flow<Resource<User>> = flow {
        emit(Resource.Loading())
        try {
            val request = ProfileUpdateRequest(firstName, lastName, phone, language, currency)
            val response = authApi.updateProfile(request)
            if (response.isSuccessful) {
                response.body()?.let { user ->
                    emit(Resource.Success(user))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to update profile"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun changePassword(currentPassword: String, newPassword: String): Flow<Resource<Unit>> = flow {
        emit(Resource.Loading())
        try {
            val request = ChangePasswordRequest(currentPassword, newPassword)
            val response = authApi.changePassword(request)
            if (response.isSuccessful) {
                emit(Resource.Success(Unit))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to change password"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun resetPassword(email: String): Flow<Resource<Unit>> = flow {
        emit(Resource.Loading())
        try {
            val request = ResetPasswordRequest(email)
            val response = authApi.resetPassword(request)
            if (response.isSuccessful) {
                emit(Resource.Success(Unit))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to reset password"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun isLoggedIn(): Boolean = !preferencesManager.getAccessToken().isNullOrEmpty()
    
    fun getAccessToken(): String? = preferencesManager.getAccessToken()
}
