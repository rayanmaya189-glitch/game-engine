package com.game-engine.casino.data.model

import com.google.gson.annotations.SerializedName

data class Game(
    @SerializedName("id")
    val id: String,
    @SerializedName("name")
    val name: String,
    @SerializedName("slug")
    val slug: String,
    @SerializedName("description")
    val description: String?,
    @SerializedName("provider")
    val provider: GameProvider?,
    @SerializedName("category")
    val category: GameCategory?,
    @SerializedName("thumbnail_url")
    val thumbnailUrl: String?,
    @SerializedName("background_url")
    val backgroundUrl: String?,
    @SerializedName("game_url")
    val gameUrl: String?,
    @SerializedName("rtp")
    val rtp: Double?,
    @SerializedName("min_bet")
    val minBet: Double?,
    @SerializedName("max_bet")
    val maxBet: Double?,
    @SerializedName("volatility")
    val volatility: String?,
    @SerializedName("is_featured")
    val isFeatured: Boolean,
    @SerializedName("is_new")
    val isNew: Boolean,
    @SerializedName("is_hot")
    val isHot: Boolean,
    @SerializedName("is_favorite")
    val isFavorite: Boolean,
    @SerializedName("is_available")
    val isAvailable: Boolean,
    @SerializedName("jackpot_amount")
    val jackpotAmount: Double?,
    @SerializedName("play_count")
    val playCount: Int,
    @SerializedName("rating")
    val rating: Double?
)

data class GameProvider(
    @SerializedName("id")
    val id: String,
    @SerializedName("name")
    val name: String,
    @SerializedName("logo_url")
    val logoUrl: String?
)

data class GameCategory(
    @SerializedName("id")
    val id: String,
    @SerializedName("name")
    val name: String,
    @SerializedName("slug")
    val slug: String,
    @SerializedName("icon")
    val icon: String?,
    @SerializedName("game_count")
    val gameCount: Int
)

data class GameListResponse(
    @SerializedName("games")
    val games: List<Game>,
    @SerializedName("total")
    val total: Int,
    @SerializedName("page")
    val page: Int,
    @SerializedName("page_size")
    val pageSize: Int,
    @SerializedName("total_pages")
    val totalPages: Int
)

data class FeaturedGamesResponse(
    @SerializedName("featured")
    val featured: List<Game>,
    @SerializedName("popular")
    val popular: List<Game>,
    @SerializedName("new")
    val newGames: List<Game>,
    @SerializedName("jackpot")
    val jackpot: List<Game>
)

data class CategoriesResponse(
    @SerializedName("categories")
    val categories: List<GameCategory>
)

data class GameSession(
    @SerializedName("session_id")
    val sessionId: String,
    @SerializedName("game_id")
    val gameId: String,
    @SerializedName("play_url")
    val playUrl: String,
    @SerializedName("fun_play_url")
    val funPlayUrl: String?,
    @SerializedName("expires_at")
    val expiresAt: String
)

data class StartGameRequest(
    @SerializedName("game_id")
    val gameId: String,
    @SerializedName("is_fun_play")
    val isFunPlay: Boolean = false
)
