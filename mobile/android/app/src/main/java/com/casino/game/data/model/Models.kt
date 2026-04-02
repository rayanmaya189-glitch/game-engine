package com.casino.game.data.model

import com.google.gson.annotations.SerializedName

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