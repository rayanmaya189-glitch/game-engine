package com.casino.game.data.model

import com.google.gson.annotations.SerializedName

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

data class SpinRequest(
    @SerializedName("game_id")
    val gameId: String,
    @SerializedName("bet_amount")
    val betAmount: Double
)

data class SpinResponse(
    val reels: List<List<String>>,
    @SerializedName("win_amount")
    val winAmount: Double,
    @SerializedName("new_balance")
    val newBalance: Double,
    val multiplier: Double
)
