package com.casino.game.data.model

// Transaction and bonus related models
    val success: Boolean,
    val message: String,
    @SerializedName("bonus_amount")
    val bonusAmount: Double?
)

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

data class ApiResponse<T>(
    val success: Boolean,
    val data: T?,
    val message: String?
)

data class ChatRoom(
    val id: String,
    val name: String,
    val lastMessage: String? = null,
    val unreadCount: Int? = null
)

data class ChatMessage(
    val id: String,
    val content: String,
    val username: String,
    val timestamp: String,
    val isMine: Boolean = false
)

data class AppNotification(
    val id: String,
    val title: String,
    val message: String,
    val type: String,
    val read: Boolean = false,
    @SerializedName("created_at")
    val createdAt: String
)

data class ReferralStats(
    val code: String,
    @SerializedName("total_referrals")
    val totalReferrals: Int,
    @SerializedName("total_earnings")
    val totalEarnings: Double
)

data class ReferralEntry(
    val username: String,
    @SerializedName("joined_at")
    val joinedAt: String,
    val earned: Double
)

data class ReferralTier(
    val name: String,
    @SerializedName("min_referrals")
    val minReferrals: Int,
    @SerializedName("reward_percent")
    val rewardPercent: Double
)

data class LiveDealerTable(
    val id: String,
    val name: String,
    @SerializedName("dealer_name")
    val dealerName: String,
    @SerializedName("game_type")
    val gameType: String,
    @SerializedName("current_players")
    val currentPlayers: Int,
    @SerializedName("max_players")
    val maxPlayers: Int,
    @SerializedName("min_bet")
    val minBet: Double,
    @SerializedName("max_bet")
    val maxBet: Double,
    val status: String
)

data class DealerChatMessage(
    val sender: String,
    val content: String,
    val timestamp: String
)

data class PlayingCard(
    val suit: String,
    val value: String
) {
    val displayValue: String get() = when (value) {
        "A" -> "A"; "K" -> "K"; "Q" -> "Q"; "J" -> "J"
        else -> value
    }
}

data class BlackjackDealResponse(
    @SerializedName("player_cards")
    val playerCards: List<PlayingCard>,
    @SerializedName("dealer_cards")
    val dealerCards: List<PlayingCard>,
    @SerializedName("player_score")
    val playerScore: Int,
    @SerializedName("dealer_score")
    val dealerScore: Int,
    @SerializedName("can_double")
    val canDouble: Boolean,
    @SerializedName("can_split")
    val canSplit: Boolean,
    val result: String?
)

data class BlackjackActionResponse(
    @SerializedName("player_cards")
    val playerCards: List<PlayingCard>,
    @SerializedName("dealer_cards")
    val dealerCards: List<PlayingCard>,
    @SerializedName("player_score")
    val playerScore: Int,
    @SerializedName("dealer_score")
    val dealerScore: Int,
    val result: String?,
    @SerializedName("new_balance")
    val newBalance: Double? = null
)

data class PokerStartResponse(
    @SerializedName("player_hand")
    val playerHand: List<PlayingCard>,
    @SerializedName("community_cards")
    val communityCards: List<PlayingCard>,
    val pot: Double,
    @SerializedName("current_bet")
    val currentBet: Double,
    @SerializedName("hand_ranking")
    val handRanking: String?,
    @SerializedName("can_call")
    val canCall: Boolean,
    @SerializedName("can_raise")
    val canRaise: Boolean
)

data class PokerActionResponse(
    @SerializedName("player_hand")
    val playerHand: List<PlayingCard>,
    @SerializedName("community_cards")
    val communityCards: List<PlayingCard>,
    val pot: Double,
    @SerializedName("hand_ranking")
    val handRanking: String?,
    val result: String?,
    @SerializedName("can_call")
    val canCall: Boolean,
    @SerializedName("can_raise")
    val canRaise: Boolean,
    @SerializedName("new_balance")
    val newBalance: Double? = null,
    @SerializedName("player_chips")
    val playerChips: Double? = null
)
