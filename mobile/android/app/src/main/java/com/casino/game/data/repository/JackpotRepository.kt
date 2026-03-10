package com.casino.game.data.repository

import com.casino.game.data.model.*
import com.casino.game.data.remote.ApiService
import com.casino.game.data.remote.WebSocketService
import com.casino.game.data.repository.AuthRepository
import kotlinx.coroutines.flow.Flow
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class JackpotRepository @Inject constructor(
    private val apiService: ApiService,
    private val webSocketService: WebSocketService,
    private val authRepository: AuthRepository
) {
    val wsMessages: Flow<WsMessage> = authRepository.wsMessages

    suspend fun getJackpots(): Result<JackpotsResponse> {
        return try {
            val response = apiService.getJackpots()
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load jackpots"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getJackpotDetails(jackpotId: String): Result<JackpotDetails> {
        return try {
            val response = apiService.getJackpotDetails(jackpotId)
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load jackpot details"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getCurrentJackpots(): Result<JackpotsResponse> {
        return try {
            val response = apiService.getCurrentJackpots()
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load current jackpots"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    fun subscribeToJackpot(jackpotId: String) {
        webSocketService.subscribeToJackpot(jackpotId)
    }
}
