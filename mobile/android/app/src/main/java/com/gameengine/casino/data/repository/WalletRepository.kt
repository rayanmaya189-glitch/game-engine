package com.gameengine.casino.data.repository

import com.gameengine.casino.data.api.Bonus
import com.gameengine.casino.data.api.WalletApi
import com.gameengine.casino.data.model.*
import com.gameengine.casino.util.Resource
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class WalletRepository @Inject constructor(
    private val walletApi: WalletApi
) {
    fun getWallet(): Flow<Resource<Wallet>> = flow {
        emit(Resource.Loading())
        try {
            val response = walletApi.getWallet()
            if (response.isSuccessful) {
                response.body()?.let { wallet ->
                    emit(Resource.Success(wallet))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to get wallet"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun getBalance(): Flow<Resource<WalletBalance>> = flow {
        emit(Resource.Loading())
        try {
            val response = walletApi.getBalance()
            if (response.isSuccessful) {
                response.body()?.let { balance ->
                    emit(Resource.Success(balance))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to get balance"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun deposit(amount: Double, paymentMethod: String, paymentId: String?, currency: String): Flow<Resource<DepositResponse>> = flow {
        emit(Resource.Loading())
        try {
            val request = DepositRequest(amount, paymentMethod, paymentId, currency)
            val response = walletApi.deposit(request)
            if (response.isSuccessful) {
                response.body()?.let { depositResponse ->
                    emit(Resource.Success(depositResponse))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Deposit failed"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun withdraw(amount: Double, paymentMethod: String, paymentDetails: String, currency: String): Flow<Resource<WithdrawResponse>> = flow {
        emit(Resource.Loading())
        try {
            val request = WithdrawRequest(amount, paymentMethod, paymentDetails, currency)
            val response = walletApi.withdraw(request)
            if (response.isSuccessful) {
                response.body()?.let { withdrawResponse ->
                    emit(Resource.Success(withdrawResponse))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Withdrawal failed"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun getTransactions(
        page: Int = 1,
        pageSize: Int = 20,
        type: String? = null,
        status: String? = null
    ): Flow<Resource<TransactionListResponse>> = flow {
        emit(Resource.Loading())
        try {
            val response = walletApi.getTransactions(page, pageSize, type, status)
            if (response.isSuccessful) {
                response.body()?.let { transactions ->
                    emit(Resource.Success(transactions))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to get transactions"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun getPaymentMethods(): Flow<Resource<PaymentMethodsResponse>> = flow {
        emit(Resource.Loading())
        try {
            val response = walletApi.getPaymentMethods()
            if (response.isSuccessful) {
                response.body()?.let { methods ->
                    emit(Resource.Success(methods))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to get payment methods"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun getBonuses(): Flow<Resource<List<Bonus>>> = flow {
        emit(Resource.Loading())
        try {
            val response = walletApi.getBonuses()
            if (response.isSuccessful) {
                response.body()?.let { bonuses ->
                    emit(Resource.Success(bonuses))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to get bonuses"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
}
