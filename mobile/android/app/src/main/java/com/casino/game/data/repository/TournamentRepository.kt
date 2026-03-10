package com.casino.game.data.repository

import com.casino.game.data.model.*
import com.casino.game.data.remote.ApiService
import com.casino.game.data.remote.WebSocketService
import com.casino.game.data.repository.AuthRepository
import kotlinx.coroutines.flow.Flow
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class TournamentRepository @Inject constructor(
    private val apiService: ApiService,
    private val webSocketService: WebSocketService,
    private val authRepository: AuthRepository
) {
    val wsMessages: Flow<WsMessage> = authRepository.wsMessages

    suspend fun getTournaments(status: String? = null, page: Int = 1, limit: Int = 20): Result<TournamentsResponse> {
        return try {
            val response = apiService.getTournaments(status, page, limit)
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load tournaments"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getTournamentDetails(tournamentId: String): Result<TournamentDetails> {
        return try {
            val response = apiService.getTournamentDetails(tournamentId)
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load tournament details"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getTournamentLeaderboard(tournamentId: String): Result<LeaderboardResponse> {
        return try {
            val response = apiService.getTournamentLeaderboard(tournamentId)
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load leaderboard"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getLeaderboard(type: String = "all", period: String = "daily"): Result<LeaderboardResponse> {
        return try {
            val response = apiService.getLeaderboard(type, period)
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load leaderboard"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    fun subscribeToTournament(tournamentId: String) {
        webSocketService.subscribeToTournament(tournamentId)
    }
}
