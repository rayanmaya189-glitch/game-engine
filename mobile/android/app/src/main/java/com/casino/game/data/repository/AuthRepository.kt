package com.casino.game.data.repository

import android.content.Context
import androidx.datastore.core.DataStore
import androidx.datastore.preferences.core.Preferences
import androidx.datastore.preferences.core.edit
import androidx.datastore.preferences.core.stringPreferencesKey
import com.casino.game.data.model.*
import com.casino.game.data.remote.ApiService
import com.casino.game.data.remote.WebSocketService
import dagger.hilt.android.qualifiers.ApplicationContext
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.flow.map
import kotlinx.coroutines.runBlocking
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class AuthRepository @Inject constructor(
    private val apiService: ApiService,
    private val webSocketService: WebSocketService,
    private val dataStore: DataStore<Preferences>,
    @ApplicationContext private val context: Context
) {
    companion object {
        val AUTH_TOKEN = stringPreferencesKey("auth_token")
        val REFRESH_TOKEN = stringPreferencesKey("refresh_token")
        val USER_ID = stringPreferencesKey("user_id")
    }

    val isLoggedIn: Flow<Boolean> = dataStore.data.map { prefs ->
        prefs[AUTH_TOKEN] != null
    }

    suspend fun login(email: String, password: String): Result<LoginResponse> {
        return try {
            val response = apiService.login(LoginRequest(email, password))
            if (response.isSuccessful && response.body() != null) {
                val loginResponse = response.body()!!
                saveTokens(loginResponse)
                webSocketService.connect(loginResponse.token)
                Result.success(loginResponse)
            } else {
                Result.failure(Exception(response.errorBody()?.string() ?: "Login failed"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun register(email: String, password: String, username: String, phone: String?): Result<LoginResponse> {
        return try {
            val response = apiService.register(RegisterRequest(email, password, username, phone))
            if (response.isSuccessful && response.body() != null) {
                val loginResponse = response.body()!!
                saveTokens(loginResponse)
                webSocketService.connect(loginResponse.token)
                Result.success(loginResponse)
            } else {
                Result.failure(Exception(response.errorBody()?.string() ?: "Registration failed"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun logout() {
        try {
            apiService.logout()
        } catch (e: Exception) {
            // Ignore errors on logout
        }
        webSocketService.disconnect()
        dataStore.edit { prefs ->
            prefs.remove(AUTH_TOKEN)
            prefs.remove(REFRESH_TOKEN)
            prefs.remove(USER_ID)
        }
    }

    suspend fun refreshToken(): Result<LoginResponse> {
        return try {
            val refreshToken = dataStore.data.first()[REFRESH_TOKEN] ?: return Result.failure(Exception("No refresh token"))
            val response = apiService.refreshToken(RefreshTokenRequest(refreshToken))
            if (response.isSuccessful && response.body() != null) {
                val loginResponse = response.body()!!
                saveTokens(loginResponse)
                Result.success(loginResponse)
            } else {
                Result.failure(Exception("Token refresh failed"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    private suspend fun saveTokens(loginResponse: LoginResponse) {
        dataStore.edit { prefs ->
            prefs[AUTH_TOKEN] = loginResponse.token
            prefs[REFRESH_TOKEN] = loginResponse.refreshToken
            prefs[USER_ID] = loginResponse.user.id
        }
    }

    suspend fun getAuthToken(): String? = dataStore.data.first()[AUTH_TOKEN]

    val wsMessages = webSocketService.messages
}
