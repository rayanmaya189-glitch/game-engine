package com.casino.game.data.remote

import com.casino.game.data.model.*
import retrofit2.Response
import retrofit2.http.*

interface ApiService {
    // Auth
    @POST("player/auth/login")
    suspend fun login(@Body request: LoginRequest): Response<LoginResponse>

    @POST("player/auth/register")
    suspend fun register(@Body request: RegisterRequest): Response<LoginResponse>

    @POST("player/auth/refresh")
    suspend fun refreshToken(@Body request: RefreshTokenRequest): Response<LoginResponse>

    @POST("player/auth/logout")
    suspend fun logout(): Response<Unit>

    // User Profile
    @GET("player/profile")
    suspend fun getProfile(): Response<UserProfile>

    @PUT("player/profile")
    suspend fun updateProfile(@Body request: UpdateProfileRequest): Response<UserProfile>

    // Wallet
    @GET("player/wallet/balance")
    suspend fun getBalance(): Response<BalanceResponse>

    @GET("player/wallet/transactions")
    suspend fun getTransactions(
        @Query("page") page: Int = 1,
        @Query("limit") limit: Int = 20
    ): Response<TransactionsResponse>

    @POST("player/wallet/deposit")
    suspend fun deposit(@Body request: DepositRequest): Response<DepositResponse>

    @POST("player/wallet/withdraw")
    suspend fun withdraw(@Body request: WithdrawRequest): Response<WithdrawResponse>

    // Games
    @GET("player/games")
    suspend fun getGames(
        @Query("category") category: String? = null,
        @Query("status") status: String? = null,
        @Query("search") search: String? = null,
        @Query("page") page: Int = 1,
        @Query("limit") limit: Int = 20
    ): Response<GamesResponse>

    @GET("player/games/{id}")
    suspend fun getGameDetails(@Path("id") gameId: String): Response<GameDetails>

    @GET("player/games/categories")
    suspend fun getCategories(): Response<CategoriesResponse>

    @GET("player/games/featured")
    suspend fun getFeaturedGames(): Response<GamesResponse>

    @GET("player/games/popular")
    suspend fun getPopularGames(): Response<GamesResponse>

    // Tournaments
    @GET("player/tournaments")
    suspend fun getTournaments(
        @Query("status") status: String? = null,
        @Query("page") page: Int = 1,
        @Query("limit") limit: Int = 20
    ): Response<TournamentsResponse>

    @GET("player/tournaments/{id}")
    suspend fun getTournamentDetails(@Path("id") tournamentId: String): Response<TournamentDetails>

    @GET("player/tournaments/{id}/leaderboard")
    suspend fun getTournamentLeaderboard(@Path("id") tournamentId: String): Response<LeaderboardResponse>

    // Jackpots
    @GET("player/jackpots")
    suspend fun getJackpots(): Response<JackpotsResponse>

    @GET("player/jackpots/{id}")
    suspend fun getJackpotDetails(@Path("id") jackpotId: String): Response<JackpotDetails>

    @GET("player/jackpots/current")
    suspend fun getCurrentJackpots(): Response<JackpotsResponse>

    // Bonuses
    @GET("player/bonuses")
    suspend fun getBonuses(): Response<BonusesResponse>

    @GET("player/bonuses/available")
    suspend fun getAvailableBonuses(): Response<BonusesResponse>

    @POST("player/bonuses/{id}/claim")
    suspend fun claimBonus(@Path("id") bonusId: String): Response<ClaimBonusResponse>

    // Leaderboard
    @GET("player/leaderboard")
    suspend fun getLeaderboard(
        @Query("type") type: String = "all",
        @Query("period") period: String = "daily"
    ): Response<LeaderboardResponse>

    // Game Play
    @POST("player/games/{id}/spin")
    suspend fun spinGame(@Path("id") gameId: String, @Body request: SpinRequest): Response<SpinResponse>

    @POST("player/tournaments/{id}/join")
    suspend fun joinTournament(@Path("id") tournamentId: String): Response<Unit>

    // Support
    @GET("player/support/tickets")
    suspend fun getSupportTickets(): Response<TicketsResponse>

    @POST("player/support/tickets")
    suspend fun createTicket(@Body request: CreateTicketRequest): Response<TicketResponse>

    // Chat
    @GET("player/chat/rooms")
    suspend fun getChatRooms(): Response<List<ChatRoom>>

    @GET("player/chat/rooms/{id}/messages")
    suspend fun getChatMessages(@Path("id") roomId: String): Response<List<ChatMessage>>

    @POST("player/chat/rooms/{id}/messages")
    suspend fun sendChatMessage(@Path("id") roomId: String, @Body request: Map<String, String>): Response<ChatMessage>

    // Notifications
    @GET("player/notifications")
    suspend fun getNotifications(): Response<List<AppNotification>>

    @PUT("player/notifications/{id}/read")
    suspend fun markNotificationRead(@Path("id") notificationId: String): Response<Unit>

    @DELETE("player/notifications/{id}")
    suspend fun deleteNotification(@Path("id") notificationId: String): Response<Unit>

    // Referral
    @GET("player/referral/stats")
    suspend fun getReferralStats(): Response<ReferralStats>

    @GET("player/referral/history")
    suspend fun getReferralHistory(): Response<List<ReferralEntry>>

    @GET("player/referral/tiers")
    suspend fun getReferralTiers(): Response<List<ReferralTier>>

    // Live Dealer
    @GET("player/live-dealer/tables")
    suspend fun getLiveDealerTables(): Response<List<LiveDealerTable>>

    @POST("player/live-dealer/tables/{id}/chat")
    suspend fun sendDealerChat(@Path("id") tableId: String, @Body request: Map<String, String>): Response<DealerChatMessage>

    // Blackjack
    @POST("player/games/blackjack/deal")
    suspend fun blackjackDeal(@Body request: Map<String, Double>): Response<BlackjackDealResponse>

    @POST("player/games/blackjack/action")
    suspend fun blackjackAction(@Body request: Map<String, String>): Response<BlackjackActionResponse>

    // Poker
    @POST("player/games/poker/start")
    suspend fun pokerStart(@Body request: Map<String, Double>): Response<PokerStartResponse>

    @POST("player/games/poker/action")
    suspend fun pokerAction(@Body request: Map<String, Any>): Response<PokerActionResponse>

    // Account
    @DELETE("player/account")
    suspend fun deleteAccount(): Response<Unit>
}
