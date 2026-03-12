package com.game-engine.casino.data.api

import com.game-engine.casino.data.model.*
import retrofit2.Response
import retrofit2.http.*

interface GameApi {
    
    @GET("games")
    suspend fun getGames(
        @Query("page") page: Int = 1,
        @Query("page_size") pageSize: Int = 20,
        @Query("category") category: String? = null,
        @Query("provider") provider: String? = null,
        @Query("search") search: String? = null,
        @Query("sort_by") sortBy: String? = null,
        @Query("is_featured") isFeatured: Boolean? = null,
        @Query("is_new") isNew: Boolean? = null
    ): Response<GameListResponse>
    
    @GET("games/featured")
    suspend fun getFeaturedGames(): Response<FeaturedGamesResponse>
    
    @GET("games/categories")
    suspend fun getCategories(): Response<CategoriesResponse>
    
    @GET("games/{id}")
    suspend fun getGame(@Path("id") id: String): Response<Game>
    
    @POST("games/{id}/start")
    suspend fun startGame(
        @Path("id") gameId: String,
        @Body request: StartGameRequest
    ): Response<GameSession>
    
    @POST("games/{id}/favorite")
    suspend fun addToFavorites(@Path("id") gameId: String): Response<Unit>
    
    @DELETE("games/{id}/favorite")
    suspend fun removeFromFavorites(@Path("id") gameId: String): Response<Unit>
    
    @GET("games/favorites")
    suspend fun getFavorites(
        @Query("page") page: Int = 1,
        @Query("page_size") pageSize: Int = 20
    ): Response<GameListResponse>
    
    @GET("games/recent")
    suspend fun getRecentGames(
        @Query("page") page: Int = 1,
        @Query("page_size") pageSize: Int = 20
    ): Response<GameListResponse>
    
    @GET("games/providers")
    suspend fun getProviders(): Response<List<GameProvider>>
}
