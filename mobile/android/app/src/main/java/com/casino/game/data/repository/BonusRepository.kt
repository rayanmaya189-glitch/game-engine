package com.casino.game.data.repository

import com.casino.game.data.model.*
import com.casino.game.data.remote.ApiService
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class BonusRepository @Inject constructor(
    private val apiService: ApiService
) {
    suspend fun getBonuses(): Result<BonusesResponse> {
        return try {
            val response = apiService.getBonuses()
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load bonuses"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getAvailableBonuses(): Result<BonusesResponse> {
        return try {
            val response = apiService.getAvailableBonuses()
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load available bonuses"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun claimBonus(bonusId: String): Result<ClaimBonusResponse> {
        return try {
            val response = apiService.claimBonus(bonusId)
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to claim bonus"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }
}
