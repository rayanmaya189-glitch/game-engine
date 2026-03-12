package com.game-engine.casino.data.model

import com.google.gson.annotations.SerializedName

data class Wallet(
    @SerializedName("id")
    val id: String,
    @SerializedName("user_id")
    val userId: String,
    @SerializedName("currency")
    val currency: String,
    @SerializedName("balance")
    val balance: Double,
    @SerializedName("bonus_balance")
    val bonusBalance: Double,
    @SerializedName("pending_balance")
    val pendingBalance: Double,
    @SerializedName("total_won")
    val totalWon: Double,
    @SerializedName("total_wagered")
    val totalWagered: Double,
    @SerializedName("last_updated")
    val lastUpdated: String
)

data class WalletBalance(
    @SerializedName("balance")
    val balance: Double,
    @SerializedName("bonus_balance")
    val bonusBalance: Double,
    @SerializedName("currency")
    val currency: String
)

data class DepositRequest(
    @SerializedName("amount")
    val amount: Double,
    @SerializedName("payment_method")
    val paymentMethod: String,
    @SerializedName("payment_id")
    val paymentId: String?,
    @SerializedName("currency")
    val currency: String
)

data class DepositResponse(
    @SerializedName("transaction_id")
    val transactionId: String,
    @SerializedName("status")
    val status: String,
    @SerializedName("amount")
    val amount: Double,
    @SerializedName("currency")
    val currency: String,
    @SerializedName("payment_url")
    val paymentUrl: String?,
    @SerializedName("expires_at")
    val expiresAt: String?
)

data class WithdrawRequest(
    @SerializedName("amount")
    val amount: Double,
    @SerializedName("payment_method")
    val paymentMethod: String,
    @SerializedName("payment_details")
    val paymentDetails: String,
    @SerializedName("currency")
    val currency: String
)

data class WithdrawResponse(
    @SerializedName("transaction_id")
    val transactionId: String,
    @SerializedName("status")
    val status: String,
    @SerializedName("amount")
    val amount: Double,
    @SerializedName("currency")
    val currency: String,
    @SerializedName("fee")
    val fee: Double,
    @SerializedName("estimated_arrival")
    val estimatedArrival: String?
)

data class Transaction(
    @SerializedName("id")
    val id: String,
    @SerializedName("user_id")
    val userId: String,
    @SerializedName("type")
    val type: String,
    @SerializedName("amount")
    val amount: Double,
    @SerializedName("currency")
    val currency: String,
    @SerializedName("status")
    val status: String,
    @SerializedName("payment_method")
    val paymentMethod: String?,
    @SerializedName("reference_id")
    val referenceId: String?,
    @SerializedName("description")
    val description: String?,
    @SerializedName("created_at")
    val createdAt: String,
    @SerializedName("completed_at")
    val completedAt: String?
)

data class TransactionListResponse(
    @SerializedName("transactions")
    val transactions: List<Transaction>,
    @SerializedName("total")
    val total: Int,
    @SerializedName("page")
    val page: Int,
    @SerializedName("page_size")
    val pageSize: Int,
    @SerializedName("total_pages")
    val totalPages: Int
)

data class PaymentMethod(
    @SerializedName("id")
    val id: String,
    @SerializedName("name")
    val name: String,
    @SerializedName("type")
    val type: String,
    @SerializedName("logo_url")
    val logoUrl: String?,
    @SerializedName("min_amount")
    val minAmount: Double,
    @SerializedName("max_amount")
    val maxAmount: Double,
    @SerializedName("fee_percentage")
    val feePercentage: Double,
    @SerializedName("processing_time")
    val processingTime: String,
    @SerializedName("is_available")
    val isAvailable: Boolean
)

data class PaymentMethodsResponse(
    @SerializedName("deposit_methods")
    val depositMethods: List<PaymentMethod>,
    @SerializedName("withdraw_methods")
    val withdrawMethods: List<PaymentMethod>
)
