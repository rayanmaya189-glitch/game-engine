package com.game-engine.casino.data.repository

import com.game-engine.casino.data.api.GameApi
import com.game-engine.casino.data.model.*
import com.game-engine.casino.util.Resource
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.flow
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class GameRepository @Inject constructor(
    private val gameApi: GameApi
) {
    fun getGames(
        page: Int = 1,
        pageSize: Int = 20,
        category: String? = null,
        provider: String? = null,
        search: String? = null,
        sortBy: String? = null,
        isFeatured: Boolean? = null,
        isNew: Boolean? = null
    ): Flow<Resource<GameListResponse>> = flow {
        emit(Resource.Loading())
        try {
            val response = gameApi.getGames(page, pageSize, category, provider, search, sortBy, isFeatured, isNew)
            if (response.isSuccessful) {
                response.body()?.let { games ->
                    emit(Resource.Success(games))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to get games"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun getFeaturedGames(): Flow<Resource<FeaturedGamesResponse>> = flow {
        emit(Resource.Loading())
        try {
            val response = gameApi.getFeaturedGames()
            if (response.isSuccessful) {
                response.body()?.let { featured ->
                    emit(Resource.Success(featured))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to get featured games"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun getCategories(): Flow<Resource<List<GameCategory>>> = flow {
        emit(Resource.Loading())
        try {
            val response = gameApi.getCategories()
            if (response.isSuccessful) {
                response.body()?.let { categories ->
                    emit(Resource.Success(categories.categories))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to get categories"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun getGame(gameId: String): Flow<Resource<Game>> = flow {
        emit(Resource.Loading())
        try {
            val response = gameApi.getGame(gameId)
            if (response.isSuccessful) {
                response.body()?.let { game ->
                    emit(Resource.Success(game))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to get game"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun startGame(gameId: String, isFunPlay: Boolean = false): Flow<Resource<GameSession>> = flow {
        emit(Resource.Loading())
        try {
            val request = StartGameRequest(gameId, isFunPlay)
            val response = gameApi.startGame(gameId, request)
            if (response.isSuccessful) {
                response.body()?.let { session ->
                    emit(Resource.Success(session))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to start game"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun addToFavorites(gameId: String): Flow<Resource<Unit>> = flow {
        emit(Resource.Loading())
        try {
            val response = gameApi.addToFavorites(gameId)
            if (response.isSuccessful) {
                emit(Resource.Success(Unit))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to add to favorites"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun removeFromFavorites(gameId: String): Flow<Resource<Unit>> = flow {
        emit(Resource.Loading())
        try {
            val response = gameApi.removeFromFavorites(gameId)
            if (response.isSuccessful) {
                emit(Resource.Success(Unit))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to remove from favorites"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun getFavorites(page: Int = 1, pageSize: Int = 20): Flow<Resource<GameListResponse>> = flow {
        emit(Resource.Loading())
        try {
            val response = gameApi.getFavorites(page, pageSize)
            if (response.isSuccessful) {
                response.body()?.let { games ->
                    emit(Resource.Success(games))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to get favorites"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun getRecentGames(page: Int = 1, pageSize: Int = 20): Flow<Resource<GameListResponse>> = flow {
        emit(Resource.Loading())
        try {
            val response = gameApi.getRecentGames(page, pageSize)
            if (response.isSuccessful) {
                response.body()?.let { games ->
                    emit(Resource.Success(games))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to get recent games"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
    
    fun getProviders(): Flow<Resource<List<GameProvider>>> = flow {
        emit(Resource.Loading())
        try {
            val response = gameApi.getProviders()
            if (response.isSuccessful) {
                response.body()?.let { providers ->
                    emit(Resource.Success(providers))
                } ?: emit(Resource.Error("Empty response"))
            } else {
                emit(Resource.Error(response.message() ?: "Failed to get providers"))
            }
        } catch (e: Exception) {
            emit(Resource.Error(e.message ?: "Network error"))
        }
    }
}
