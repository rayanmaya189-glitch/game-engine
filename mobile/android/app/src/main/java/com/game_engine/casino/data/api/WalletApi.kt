package com.game_engine.casino.data.api

import com.game_engine.casino.data.model.*
import retrofit2.Response
import retrofit2.http.*

interface WalletApi {
    
    @GET("wallet")
    suspend fun getWallet(): Response<Wallet>
    
    @GET("wallet/balance")
    suspend fun getBalance(): Response<WalletBalance>
    
    @POST("wallet/deposit")
    suspend fun deposit(@Body request: DepositRequest): Response<DepositResponse>
    
    @POST("wallet/withdraw")
    suspend fun withdraw(@Body request: WithdrawRequest): Response<WithdrawResponse>
    
    @GET("wallet/transactions")
    suspend fun getTransactions(
        @Query("page") page: Int = 1,
        @Query("page_size") pageSize: Int = 20,
        @Query("type") type: String? = null,
        @Query("status") status: String? = null
    ): Response<TransactionListResponse>
    
    @GET("wallet/payment-methods")
    suspend fun getPaymentMethods(): Response<PaymentMethodsResponse>
    
    @GET("wallet/bonus")
    suspend fun getBonuses(): Response<List<Bonus>>
    
    @GET("wallet/transaction/{id}")
    suspend fun getTransaction(@Path("id") id: String): Response<Transaction>
}

data class Bonus(
    val id: String,
    val name: String,
    val description: String?,
    val amount: Double,
    val currency: String,
    val type: String,
    val wageringRequirement: Double,
    val remainingWagering: Double,
    val expiresAt: String?,
    val isActive: Boolean
)
