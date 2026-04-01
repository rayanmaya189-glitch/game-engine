package com.casino.game.data.model

import com.google.gson.annotations.SerializedName

data class BalanceResponse(
    val balance: Double,
    @SerializedName("bonus_balance")
    val bonusBalance: Double,
    @SerializedName("pending_balance")
    val pendingBalance: Double,
    val currency: String
)

data class Transaction(
    val id: String,
    val type: String,
    val amount: Double,
    val status: String,
    val method: String?,
    @SerializedName("created_at")
    val createdAt: String,
    @SerializedName("transaction_id")
    val transactionId: String?
)

data class TransactionsResponse(
    val transactions: List<Transaction>,
    val total: Int,
    val page: Int,
    val pages: Int
)

data class DepositRequest(
    val amount: Double,
    val method: String,
    @SerializedName("payment_id")
    val paymentId: String? = null
)

data class DepositResponse(
    @SerializedName("transaction_id")
    val transactionId: String,
    val status: String,
    val message: String
)

data class WithdrawRequest(
    val amount: Double,
    val method: String,
    @SerializedName("account_info")
    val accountInfo: String
)

data class WithdrawResponse(
    @SerializedName("transaction_id")
    val transactionId: String,
    val status: String,
    val message: String
)
