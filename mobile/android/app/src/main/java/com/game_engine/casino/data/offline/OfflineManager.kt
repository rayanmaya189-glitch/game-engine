package com.game_engine.casino.data.offline

import android.content.Context
import androidx.room.*
import dagger.hilt.android.qualifiers.ApplicationContext
import kotlinx.coroutines.flow.Flow
import javax.inject.Inject
import javax.inject.Singleton

/**
 * Offline Manager
 * 
 * Handles offline data caching and synchronization:
 * - Cache game catalog, user profile, transactions
 * - Queue actions when offline
 * - Sync when back online
 */

/**
 * Room Database for offline storage
 */
@Database(
    entities = [
        CachedGame::class,
        CachedTransaction::class,
        CachedUserProfile::class,
        QueuedAction::class
    ],
    version = 1,
    exportSchema = false
)
@Singleton
class OfflineDatabase @Inject constructor(
    @ApplicationContext context: Context
) : RoomDatabase() {
    val gameDao = cachedGameDao()
    val transactionDao = cachedTransactionDao()
    val userDao = cachedUserProfileDao()
    val actionDao = queuedActionDao()

    abstract fun cachedGameDao(): CachedGameDao
    abstract fun cachedTransactionDao(): CachedTransactionDao
    abstract fun cachedUserProfileDao(): CachedUserProfileDao
    abstract fun queuedActionDao(): QueuedActionDao
}

/**
 * Cached Game Entity
 */
@Entity(tableName = "cached_games")
data class CachedGame(
    @PrimaryKey val id: String,
    val name: String,
    val category: String,
    val thumbnailUrl: String,
    val rtp: Double,
    val minBet: Double,
    val maxBet: Double,
    val isFavorite: Boolean = false,
    val lastPlayedAt: Long? = null,
    val cachedAt: Long = System.currentTimeMillis()
)

/**
 * Cached Transaction Entity
 */
@Entity(tableName = "cached_transactions")
data class CachedTransaction(
    @PrimaryKey val id: String,
    val type: String, // deposit, withdrawal, bet, win
    val amount: Double,
    val currency: String,
    val status: String,
    val description: String,
    val timestamp: Long,
    val cachedAt: Long = System.currentTimeMillis()
)

/**
 * Cached User Profile
 */
@Entity(tableName = "cached_user_profile")
data class CachedUserProfile(
    @PrimaryKey val id: String,
    val username: String,
    val email: String,
    val avatarUrl: String?,
    val balance: Double,
    val bonusBalance: Double,
    val vipLevel: Int,
    val cachedAt: Long = System.currentTimeMillis()
)

/**
 * Queued Action for offline operations
 */
@Entity(tableName = "queued_actions")
data class QueuedAction(
    @PrimaryKey(autoGenerate = true) val id: Long = 0,
    val actionType: String, // deposit, withdrawal, etc.
    val payload: String, // JSON payload
    val createdAt: Long = System.currentTimeMillis(),
    val retryCount: Int = 0
)

// ========== DAOs ==========

@Dao
interface CachedGameDao {
    @Query("SELECT * FROM cached_games ORDER BY lastPlayedAt DESC")
    fun getAllGames(): Flow<List<CachedGame>>

    @Query("SELECT * FROM cached_games WHERE category = :category")
    fun getGamesByCategory(category: String): Flow<List<CachedGame>>

    @Query("SELECT * FROM cached_games WHERE isFavorite = 1")
    fun getFavoriteGames(): Flow<List<CachedGame>>

    @Query("SELECT * FROM cached_games WHERE lastPlayedAt IS NOT NULL ORDER BY lastPlayedAt DESC LIMIT 10")
    fun getRecentlyPlayed(): Flow<List<CachedGame>>

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertGames(games: List<CachedGame>)

    @Query("UPDATE cached_games SET isFavorite = :isFavorite WHERE id = :gameId")
    suspend fun updateFavorite(gameId: String, isFavorite: Boolean)

    @Query("UPDATE cached_games SET lastPlayedAt = :timestamp WHERE id = :gameId")
    suspend fun updateLastPlayed(gameId: String, timestamp: Long)

    @Query("DELETE FROM cached_games WHERE cachedAt < :timestamp")
    suspend fun deleteOldCache(timestamp: Long)
}

@Dao
interface CachedTransactionDao {
    @Query("SELECT * FROM cached_transactions ORDER BY timestamp DESC")
    fun getAllTransactions(): Flow<List<CachedTransaction>>

    @Query("SELECT * FROM cached_transactions WHERE type = :type ORDER BY timestamp DESC")
    fun getTransactionsByType(type: String): Flow<List<CachedTransaction>>

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertTransactions(transactions: List<CachedTransaction>)

    @Query("DELETE FROM cached_transactions WHERE cachedAt < :timestamp")
    suspend fun deleteOldCache(timestamp: Long)
}

@Dao
interface CachedUserProfileDao {
    @Query("SELECT * FROM cached_user_profile LIMIT 1")
    fun getProfile(): Flow<CachedUserProfile?>

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun saveProfile(profile: CachedUserProfile)

    @Query("UPDATE cached_user_profile SET balance = :balance WHERE id = :userId")
    suspend fun updateBalance(userId: String, balance: Double)

    @Query("DELETE FROM cached_user_profile")
    suspend fun clearProfile()
}

@Dao
interface QueuedActionDao {
    @Query("SELECT * FROM queued_actions ORDER BY createdAt ASC")
    suspend fun getPendingActions(): List<QueuedAction>

    @Insert
    suspend fun queueAction(action: QueuedAction)

    @Delete
    suspend fun deleteAction(action: QueuedAction)

    @Query("UPDATE queued_actions SET retryCount = retryCount + 1 WHERE id = :id")
    suspend fun incrementRetryCount(id: Long)

    @Query("DELETE FROM queued_actions WHERE retryCount > :maxRetries")
    suspend fun deleteFailedActions(maxRetries: Int)
}

/**
 * Offline Manager Service
 */
@Singleton
class OfflineManager @Inject constructor(
    private val database: OfflineDatabase
) {
    private val MAX_CACHE_AGE = 24 * 60 * 60 * 1000L // 24 hours
    private val MAX_RETRIES = 3

    // Game caching
    fun getCachedGames() = database.gameDao.getAllGames()
    fun getFavoriteGames() = database.gameDao.getFavoriteGames()
    fun getRecentlyPlayed() = database.gameDao.getRecentlyPlayed()

    suspend fun cacheGames(games: List<CachedGame>) {
        database.gameDao.insertGames(games)
    }

    suspend fun toggleFavorite(gameId: String, isFavorite: Boolean) {
        database.gameDao.updateFavorite(gameId, isFavorite)
    }

    suspend fun markGamePlayed(gameId: String) {
        database.gameDao.updateLastPlayed(gameId, System.currentTimeMillis())
    }

    // Transaction caching
    fun getCachedTransactions() = database.transactionDao.getAllTransactions()

    suspend fun cacheTransactions(transactions: List<CachedTransaction>) {
        database.transactionDao.insertTransactions(transactions)
    }

    // User profile caching
    fun getCachedProfile() = database.userDao.getProfile()

    suspend fun cacheProfile(profile: CachedUserProfile) {
        database.userDao.saveProfile(profile)
    }

    suspend fun updateBalance(userId: String, balance: Double) {
        database.userDao.updateBalance(userId, balance)
    }

    // Action queueing for offline
    suspend fun queueAction(actionType: String, payload: String) {
        database.actionDao.queueAction(
            QueuedAction(actionType = actionType, payload = payload)
        )
    }

    suspend fun processPendingActions(actionExecutor: suspend (QueuedAction) -> Boolean) {
        val pendingActions = database.actionDao.getPendingActions()

        for (action in pendingActions) {
            val success = actionExecutor(action)
            if (success) {
                database.actionDao.deleteAction(action)
            } else {
                database.actionDao.incrementRetryCount(action.id)
            }
        }

        // Clean up failed actions
        database.actionDao.deleteFailedActions(MAX_RETRIES)
    }

    // Cache cleanup
    suspend fun cleanOldCache() {
        val cutoff = System.currentTimeMillis() - MAX_CACHE_AGE
        database.gameDao.deleteOldCache(cutoff)
        database.transactionDao.deleteOldCache(cutoff)
    }

    suspend fun clearAllCache() {
        database.userDao.clearProfile()
    }
}
