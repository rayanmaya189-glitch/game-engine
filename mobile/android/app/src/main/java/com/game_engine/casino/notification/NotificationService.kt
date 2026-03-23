package com.game_engine.casino.notification

import android.app.NotificationChannel
import android.app.NotificationManager
import android.app.PendingIntent
import android.content.Context
import android.content.Intent
import android.os.Build
import androidx.core.app.NotificationCompat
import com.game_engine.casino.R
import com.game_engine.casino.ui.MainActivity
import com.google.firebase.messaging.FirebaseMessagingService
import com.google.firebase.messaging.RemoteMessage
import dagger.hilt.android.AndroidEntryPoint

/**
 * Firebase Cloud Messaging Service
 * 
 * Handles push notifications for:
 * - Game alerts (your turn, tournament starting)
 * - Financial (deposits, withdrawals, bonuses)
 * - Promotions (new games, special offers)
 * - System messages
 */
@AndroidEntryPoint
class FcmService : FirebaseMessagingService() {

    override fun onNewToken(token: String) {
        super.onNewToken(token)
        // Send token to server
        sendTokenToServer(token)
    }

    override fun onMessageReceived(remoteMessage: RemoteMessage) {
        super.onMessageReceived(remoteMessage)

        // Get notification data
        val data = remoteMessage.data
        val notificationType = data["type"] ?: "system"

        // Build and show notification
        when (notificationType) {
            "game" -> showGameNotification(data)
            "financial" -> showFinancialNotification(data)
            "promotion" -> showPromotionNotification(data)
            "tournament" -> showTournamentNotification(data)
            else -> showSystemNotification(data)
        }
    }

    private fun showGameNotification(data: Map<String, String>) {
        val title = data["title"] ?: "Game Update"
        val body = data["body"] ?: ""
        val gameId = data["game_id"]
        
        val intent = Intent(this, MainActivity::class.java).apply {
            putExtra("screen", "game")
            putExtra("game_id", gameId)
            flags = Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TASK
        }

        showNotification(
            channelId = CHANNEL_GAME,
            channelName = "Game Alerts",
            title = title,
            body = body,
            intent = intent
        )
    }

    private fun showFinancialNotification(data: Map<String, String>) {
        val title = data["title"] ?: "Transaction Update"
        val body = data["body"] ?: ""
        
        val intent = Intent(this, MainActivity::class.java).apply {
            putExtra("screen", "wallet")
            flags = Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TASK
        }

        showNotification(
            channelId = CHANNEL_FINANCIAL,
            channelName = "Financial",
            title = title,
            body = body,
            intent = intent
        )
    }

    private fun showPromotionNotification(data: Map<String, String>) {
        val title = data["title"] ?: "Special Offer"
        val body = data["body"] ?: ""
        val bonusId = data["bonus_id"]
        
        val intent = Intent(this, MainActivity::class.java).apply {
            putExtra("screen", "bonus")
            putExtra("bonus_id", bonusId)
            flags = Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TASK
        }

        showNotification(
            channelId = CHANNEL_PROMOTION,
            channelName = "Promotions",
            title = title,
            body = body,
            intent = intent
        )
    }

    private fun showTournamentNotification(data: Map<String, String>) {
        val title = data["title"] ?: "Tournament"
        val body = data["body"] ?: ""
        val tournamentId = data["tournament_id"]
        
        val intent = Intent(this, MainActivity::class.java).apply {
            putExtra("screen", "tournament")
            putExtra("tournament_id", tournamentId)
            flags = Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TASK
        }

        showNotification(
            channelId = CHANNEL_TOURNAMENT,
            channelName = "Tournaments",
            title = title,
            body = body,
            intent = intent
        )
    }

    private fun showSystemNotification(data: Map<String, String>) {
        val title = data["title"] ?: "System Message"
        val body = data["body"] ?: ""
        
        val intent = Intent(this, MainActivity::class.java).apply {
            flags = Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TASK
        }

        showNotification(
            channelId = CHANNEL_SYSTEM,
            channelName = "System",
            title = title,
            body = body,
            intent = intent
        )
    }

    private fun showNotification(
        channelId: String,
        channelName: String,
        title: String,
        body: String,
        intent: Intent
    ) {
        createNotificationChannel(channelId, channelName)

        val pendingIntent = PendingIntent.getActivity(
            this,
            0,
            intent,
            PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE
        )

        val notification = NotificationCompat.Builder(this, channelId)
            .setSmallIcon(R.drawable.ic_notification)
            .setContentTitle(title)
            .setContentText(body)
            .setPriority(NotificationCompat.PRIORITY_HIGH)
            .setAutoCancel(true)
            .setContentIntent(pendingIntent)
            .build()

        val notificationManager = getSystemService(Context.NOTIFICATION_SERVICE) as NotificationManager
        notificationManager.notify(System.currentTimeMillis().toInt(), notification)
    }

    private fun createNotificationChannel(channelId: String, channelName: String) {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            val channel = NotificationChannel(
                channelId,
                channelName,
                NotificationManager.IMPORTANCE_HIGH
            ).apply {
                description = "Notifications for $channelName"
                enableVibration(true)
            }

            val notificationManager = getSystemService(Context.NOTIFICATION_SERVICE) as NotificationManager
            notificationManager.createNotificationChannel(channel)
        }
    }

    private fun sendTokenToServer(token: String) {
        // API call to register token with backend
    }

    companion object {
        const val CHANNEL_GAME = "game_alerts"
        const val CHANNEL_FINANCIAL = "financial"
        const val CHANNEL_PROMOTION = "promotions"
        const val CHANNEL_TOURNAMENT = "tournaments"
        const val CHANNEL_SYSTEM = "system"
    }
}
