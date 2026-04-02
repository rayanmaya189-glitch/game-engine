package com.casino.game.data.model

import com.google.gson.annotations.SerializedName

data class PaymentTransaction(
    val id: String,
    val type: String,
    val amount: Double,
    val status: String,
    val date: String,
    val description: String
)

data class PaymentHistoryResponse(
    val transactions: List<PaymentTransaction>,
    val total: Int
)

data class KycDocument(
    val type: String,
    val status: String,
    @SerializedName("rejection_reason")
    val rejectionReason: String? = null,
    @SerializedName("uploaded_at")
    val uploadedAt: String? = null
)

data class KycStatusResponse(
    val level: Int,
    val documents: List<KycDocument>
)

data class BigWin(
    val username: String,
    val amount: Double,
    val date: String
)

data class GameDetailResponse(
    val game: GameDetails,
    @SerializedName("recent_big_wins")
    val recentBigWins: List<BigWin>,
    @SerializedName("related_games")
    val relatedGames: List<Game>
)

data class BetSelection(
    val selection: String,
    val odds: Double,
    val result: String
)

data class BetEntry(
    val id: String,
    @SerializedName("game_name")
    val gameName: String,
    val stake: Double,
    @SerializedName("potential_win")
    val potentialWin: Double,
    @SerializedName("actual_win")
    val actualWin: Double?,
    val result: String,
    val date: String,
    val selections: List<BetSelection>
)

data class BetHistoryResponse(
    val bets: List<BetEntry>,
    val total: Int
)
