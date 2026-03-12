package com.game-engine.casino.websocket

import com.game-engine.casino.util.PreferencesManager
import kotlinx.coroutines.channels.Channel
import kotlinx.coroutines.flow.receiveAsFlow
import okhttp3.*
import org.json.JSONObject
import java.util.concurrent.TimeUnit
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class GameSocketManager @Inject constructor(
    private val preferencesManager: PreferencesManager
) {
    private var webSocket: WebSocket? = null
    private val client = OkHttpClient.Builder()
        .readTimeout(30, TimeUnit.SECONDS)
        .pingInterval(15, TimeUnit.SECONDS)
        .build()
    
    private val _events = Channel<SocketEvent>(Channel.BUFFERED)
    val events = _events.receiveAsFlow()
    
    private var isConnected = false
    private var reconnectAttempts = 0
    private val maxReconnectAttempts = 5
    
    companion object {
        private const val WS_URL = "wss://api.casino-game.engine/v1/ws"
    }
    
    fun connect() {
        if (isConnected) return
        
        val token = preferencesManager.getAccessToken()
        if (token.isNullOrEmpty()) {
            return
        }
        
        val request = Request.Builder()
            .url("$WS_URL?token=$token")
            .build()
        
        webSocket = client.newWebSocket(request, createListener())
    }
    
    fun disconnect() {
        webSocket?.close(1000, "User disconnected")
        webSocket = null
        isConnected = false
        reconnectAttempts = 0
    }
    
    fun sendMessage(type: String, payload: Map<String, Any> = emptyMap()) {
        if (!isConnected) {
            return
        }
        
        val message = JSONObject().apply {
            put("type", type)
            put("payload", JSONObject(payload))
        }
        
        webSocket?.send(message.toString())
    }
    
    fun joinGame(gameId: String) {
        sendMessage("join_game", mapOf("game_id" to gameId))
    }
    
    fun leaveGame(gameId: String) {
        sendMessage("leave_game", mapOf("game_id" to gameId))
    }
    
    fun sendGameAction(gameId: String, action: String, data: Map<String, Any> = emptyMap()) {
        sendMessage("game_action", mapOf(
            "game_id" to gameId,
            "action" to action,
            "data" to data
        ))
    }
    
    private fun createListener(): WebSocketListener {
        return object : WebSocketListener() {
            override fun onOpen(webSocket: WebSocket, response: Response) {
                isConnected = true
                reconnectAttempts = 0
                _events.trySend(SocketEvent.Connected)
            }
            
            override fun onMessage(webSocket: WebSocket, text: String) {
                try {
                    val json = JSONObject(text)
                    val type = json.getString("type")
                    val payload = json.optJSONObject("payload")
                    
                    when (type) {
                        "game_update" -> {
                            _events.trySend(SocketEvent.GameUpdate(payload))
                        }
                        "jackpot_update" -> {
                            _events.trySend(SocketEvent.JackpotUpdate(payload))
                        }
                        "leaderboard_update" -> {
                            _events.trySend(SocketEvent.LeaderboardUpdate(payload))
                        }
                        "chat_message" -> {
Send(SocketEvent                            _events.try.ChatMessage(payload))
                        }
                        "tournament_update" -> {
                            _events.trySend(SocketEvent.TournamentUpdate(payload))
                        }
                        "error" -> {
                            val message = payload?.optString("message") ?: "Unknown error"
                            _events.trySend(SocketEvent.Error(message))
                        }
                        "ping" -> {
                            webSocket.send("{\"type\":\"pong\"}")
                        }
                    }
                } catch (e: Exception) {
                    _events.trySend(SocketEvent.Error("Failed to parse message"))
                }
            }
            
            override fun onClosing(webSocket: WebSocket, code: Int, reason: String) {
                webSocket.close(1000, null)
            }
            
            override fun onClosed(webSocket: WebSocket, code: Int, reason: String) {
                isConnected = false
                _events.trySend(SocketEvent.Disconnected(reason))
            }
            
            override fun onFailure(webSocket: WebSocket, t: Throwable, response: Response?) {
                isConnected = false
                _events.trySend(SocketEvent.Error(t.message ?: "Connection failed"))
                attemptReconnect()
            }
        }
    }
    
    private fun attemptReconnect() {
        if (reconnectAttempts < maxReconnectAttempts) {
            reconnectAttempts++
            Thread.sleep((reconnectAttempts * 2000).toLong())
            connect()
        }
    }
}

sealed class SocketEvent {
    object Connected : SocketEvent()
    data class Disconnected(val reason: String) : SocketEvent()
    data class GameUpdate(val payload: JSONObject) : SocketEvent()
    data class JackpotUpdate(val payload: JSONObject) : SocketEvent()
    data class LeaderboardUpdate(val payload: JSONObject) : SocketEvent()
    data class ChatMessage(val payload: JSONObject) : SocketEvent()
    data class TournamentUpdate(val payload: JSONObject) : SocketEvent()
    data class Error(val message: String) : SocketEvent()
}
