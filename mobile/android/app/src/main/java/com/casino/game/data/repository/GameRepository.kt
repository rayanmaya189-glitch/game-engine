package com.casino.game.data.repository

import com.casino.game.data.model.*
import com.casino.game.data.remote.ApiService
import com.casino.game.data.remote.WebSocketService
import kotlinx.coroutines.flow.Flow
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class GameRepository @Inject constructor(
    private val apiService: ApiService,
    private val webSocketService: WebSocketService
) {
    val wsMessages: Flow<WsMessage> = webSocketService.messages

    suspend fun getGames(
        category: String? = null,
        status: String? = null,
        search: String? = null,
        page: Int = 1,
        limit: Int = 20
    ): Result<GamesResponse> {
        return try {
            val response = apiService.getGames(category, status, search, page, limit)
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load games"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getGameDetails(gameId: String): Result<GameDetails> {
        return try {
            val response = apiService.getGameDetails(gameId)
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load game details"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getCategories(): Result<CategoriesResponse> {
        return try {
            val response = apiService.getCategories()
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load categories"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getFeaturedGames(): Result<GamesResponse> {
        return try {
            val response = apiService.getFeaturedGames()
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load featured games"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getPopularGames(): Result<GamesResponse> {
        return try {
            val response = apiService.getPopularGames()
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load popular games"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    fun subscribeToGame(gameId: String) {
        webSocketService.subscribeToGame(gameId)
    }

    fun unsubscribeFromGame(gameId: String) {
        webSocketService.unsubscribeFromGame(gameId)
    }
}
