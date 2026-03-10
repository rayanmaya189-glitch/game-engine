package com.casino.game.data.repository

import com.casino.game.data.model.*
import com.casino.game.data.remote.ApiService
import com.casino.game.data.remote.WebSocketService
import com.casino.game.data.repository.AuthRepository.Companion.wsMessages
import kotlinx.coroutines.flow.Flow
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class WalletRepository @Inject constructor(
    private val apiService: ApiService,
    private val webSocketService: WebSocketService,
    private val authRepository: AuthRepository
) {
    val wsMessages: Flow<WsMessage> = authRepository.wsMessages

    suspend fun getBalance(): Result<BalanceResponse> {
        return try {
            val response = apiService.getBalance()
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load balance"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun getTransactions(page: Int = 1, limit: Int = 20): Result<TransactionsResponse> {
        return try {
            val response = apiService.getTransactions(page, limit)
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Failed to load transactions"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun deposit(amount: Double, method: String): Result<DepositResponse> {
        return try {
            val response = apiService.deposit(DepositRequest(amount, method))
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Deposit failed"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }

    suspend fun withdraw(amount: Double, method: String, accountInfo: String): Result<WithdrawResponse> {
        return try {
            val response = apiService.withdraw(WithdrawRequest(amount, method, accountInfo))
            if (response.isSuccessful && response.body() != null) {
                Result.success(response.body()!!)
            } else {
                Result.failure(Exception("Withdrawal failed"))
            }
        } catch (e: Exception) {
            Result.failure(e)
        }
    }
}
