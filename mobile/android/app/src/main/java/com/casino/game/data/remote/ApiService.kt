package com.casino.game.data.remote

import com.casino.game.data.model.*
import retrofit2.Response
import retrofit2.http.*

interface ApiService {
    // Auth
    @POST("player/auth/login")
   Body request: Login suspend fun login(@Request): Response<LoginResponse>

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

    // Support
    @GET("player/support/tickets")
    suspend fun getSupportTickets(): Response<TicketsResponse>

    @POST("player/support/tickets")
    suspend fun createTicket(@Body request: CreateTicketRequest): Response<TicketResponse>
}
