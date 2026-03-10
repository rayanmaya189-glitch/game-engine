package com.casino.game.data.model

import com.google.gson.annotations.SerializedName

// Auth Models
data class LoginRequest(
    val email: String,
    val password: String
)

data class RegisterRequest(
    val email: String,
    val password: String,
    val username: String,
    val phone: String? = null
)

data class RefreshTokenRequest(
    @SerializedName("refresh_token")
    val refreshToken: String
)

data class LoginResponse(
    val user: User,
    val token: String,
    @SerializedName("refresh_token")
    val refreshToken: String,
    @SerializedName("expires_in")
    val expiresIn: Long
)

data class User(
    val id: String,
    val username: String,
    val email: String,
    val status: String,
    @SerializedName("kyc_level")
    val kycLevel: String,
    @SerializedName("created_at")
    val createdAt: String
)

data class UserProfile(
    val id: String,
    val username: String,
    val email: String,
    val phone: String?,
    val status: String,
    @SerializedName("kyc_level")
    val kycLevel: String,
    @SerializedName("created_at")
    val createdAt: String,
    @SerializedName("avatar_url")
    val avatarUrl: String?
)

data class UpdateProfileRequest(
    val username: String? = null,
    val phone: String? = null
)

// Wallet Models
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

// Game Models
data class Game(
    val id: String,
    val name: String,
    val provider: String,
    val category: String,
    val thumbnail: String?,
    val status: String,
    @SerializedName("min_bet")
    val minBet: Double,
    @SerializedName("max_bet")
    val maxBet: Double,
    @SerializedName("rtp")
    val rtp: Double,
    val volatility: String?
)

data class GameDetails(
    val id: String,
    val name: String,
    val provider: String,
    val category: String,
    val thumbnail: String?,
    val banner: String?,
    val description: String,
    val status: String,
    @SerializedName("min_bet")
    val minBet: Double,
    @SerializedName("max_bet")
    val maxBet: Double,
    @SerializedName("rtp")
    val rtp: Double,
    val volatility: String?,
    @SerializedName("paylines")
    val paylines: Int?,
    @SerializedName("reels")
    val reels: Int?,
    @SerializedName("game_features")
    val gameFeatures: List<String>?
)

data class GamesResponse(
    val games: List<Game>,
    val total: Int,
    val page: Int,
    val pages: Int
)

data class Category(
    val id: String,
    val name: String,
    val icon: String?,
    @SerializedName("game_count")
    val gameCount: Int
)

data class CategoriesResponse(
    val categories: List<Category>
)

// Tournament Models
data class Tournament(
    val id: String,
    val name: String,
    val description: String?,
    val game: String,
    @SerializedName("prize_pool")
    val prizePool: Double,
    @SerializedName("min_bet")
    val minBet: Double,
    @SerializedName("start_date")
    val startDate: String,
    @SerializedName("end_date")
    val endDate: String,
    val status: String,
    @SerializedName("player_count")
    val playerCount: Int
)

data class TournamentDetails(
    val id: String,
    val name: String,
    val description: String?,
    val game: String,
    val rules: String?,
    @SerializedName("prize_pool")
    val prizePool: Double,
    @SerializedName("min_bet")
    val minBet: Double,
    @SerializedName("start_date")
    val startDate: String,
    @SerializedName("end_date")
    val endDate: String,
    val status: String,
    @SerializedName("player_count")
    val playerCount: Int,
    val prizes: List<TournamentPrize>
)

data class TournamentPrize(
    val position: Int,
    val amount: Double,
    val type: String
)

data class TournamentsResponse(
    val tournaments: List<Tournament>,
    val total: Int,
    val page: Int,
    val pages: Int
)

data class LeaderboardEntry(
    val rank: Int,
    @SerializedName("user_id")
    val userId: String,
    val username: String,
    @SerializedName("avatar_url")
    val avatarUrl: String?,
    val score: Double,
    val prize: Double?
)

data class LeaderboardResponse(
    val entries: List<LeaderboardEntry>,
    val period: String,
    val type: String
)

// Jackpot Models
data class Jackpot(
    val id: String,
    val name: String,
    val game: String,
    @SerializedName("current_amount")
    val currentAmount: Double,
    @SerializedName("min_bet")
    val minBet: Double,
    @SerializedName("max_win")
    val maxWin: Double,
    val status: String,
    @SerializedName("hit_count")
    val hitCount: Int
)

data class JackpotDetails(
    val id: String,
    val name: String,
    val game: String,
    @SerializedName("current_amount")
    val currentAmount: Double,
    @SerializedName("min_bet")
    val minBet: Double,
    @SerializedName("max_win")
    val maxWin: Double,
    val status: String,
    @SerializedName("hit_count")
    val hitCount: Int,
    @SerializedName("recent_hits")
    val recentHits: List<JackpotHit>
)

data class JackpotHit(
    val id: String,
    val user: String,
    val amount: Double,
    @SerializedName("game_name")
    val gameName: String,
    @SerializedName("won_at")
    val wonAt: String
)

data class JackpotsResponse(
    val jackpots: List<Jackpot>
)

// Bonus Models
data class Bonus(
    val id: String,
    val name: String,
    val type: String,
    val description: String?,
    val amount: Double,
    @SerializedName("max_bonus")
    val maxBonus: Double,
    @SerializedName("min_deposit")
    val minDeposit: Double?,
    @SerializedName("wager_requirement")
    val wagerRequirement: Int?,
    @SerializedName("expires_at")
    val expiresAt: String?,
    val status: String
)

data class BonusesResponse(
    val bonuses: List<Bonus>
)

data class ClaimBonusResponse(
    val success: Boolean,
    val message: String,
    @SerializedName("bonus_amount")
    val bonusAmount: Double?
)

// Support Models
data class SupportTicket(
    val id: String,
    val subject: String,
    val status: String,
    @SerializedName("created_at")
    val createdAt: String,
    @SerializedName("last_reply")
    val lastReply: String?
)

data class TicketsResponse(
    val tickets: List<SupportTicket>
)

data class CreateTicketRequest(
    val subject: String,
    val message: String,
    val category: String
)

data class TicketResponse(
    val ticket: SupportTicket,
    val message: String
)

// WebSocket Models
data class WsMessage(
    val type: String,
    val data: Any?
)

data class GameEvent(
    val type: String,
    val gameId: String,
    val data: Map<String, Any>
)

data class JackpotUpdate(
    @SerializedName("jackpot_id")
    val jackpotId: String,
    @SerializedName("new_amount")
    val newAmount: Double,
    val timestamp: String
)

data class TournamentUpdate(
    @SerializedName("tournament_id")
    val tournamentId: String,
    @SerializedName("leaderboard")
    val leaderboard: List<LeaderboardEntry>,
    val timeRemaining: Long?
)

data class BalanceUpdate(
    @SerializedName("new_balance")
    val newBalance: Double,
    val type: String,
    val amount: Double
)

// Generic Response
data class ApiResponse<T>(
    val success: Boolean,
    val data: T?,
    val message: String?
)
