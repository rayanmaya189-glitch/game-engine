package com.casino.game.data.remote

import com.casino.game.data.model.*
import com.google.gson.Gson
import kotlinx.coroutines.channels.Channel
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.receiveAsFlow
import okhttp3.*
import java.util.concurrent.TimeUnit
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class WebSocketService @Inject constructor(
    private val okHttpClient: OkHttpClient,
    private val gson: Gson
) {
    private var webSocket: WebSocket? = null
    private val _messages = Channel<WsMessage>(Channel.BUFFERED)
    val messages: Flow<WsMessage> = _messages.receiveAsFlow()

    private var isConnected = false
    private var reconnectAttempts = 0
    private val maxReconnectAttempts = 5
    private var token: String? = null
    private var baseUrl: String = ""

    fun connect(token: String, baseUrl: String = "ws://localhost:8080/ws") {
        this.token = token
        this.baseUrl = baseUrl

        val request = Request.Builder()
            .url("$baseUrl?token=$token")
            .build()

        webSocket = okHttpClient.newWebSocket(request, createWebSocketListener())
    }

    fun disconnect() {
        webSocket?.close(1000, "User disconnected")
        webSocket = null
        isConnected = false
    }

    fun sendMessage(type: String, data: Any?) {
        val message = WsMessage(type, data)
        val json = gson.toJson(message)
        webSocket?.send(json)
    }

    fun subscribeToGame(gameId: String) {
        sendMessage("subscribe_game", mapOf("game_id" to gameId))
    }

    fun unsubscribeFromGame(gameId: String) {
        sendMessage("unsubscribe_game", mapOf("game_id" to gameId))
    }

    fun subscribeToJackpot(jackpotId: String) {
        sendMessage("subscribe_jackpot", mapOf("jackpot_id" to jackpotId))
    }

    fun subscribeToTournament(tournamentId: String) {
        sendMessage("subscribe_tournament", mapOf("tournament_id" to tournamentId))
    }

    private fun reconnect() {
        if (reconnectAttempts < maxReconnectAttempts && token != null) {
            reconnectAttempts++
            Thread.sleep((reconnectAttempts * 2000).toLong())
            token?.let { connect(it, baseUrl) }
        }
    }

    private fun createWebSocketListener() = object : WebSocketListener() {
        override fun onOpen(webSocket: WebSocket, response: Response) {
            isConnected = true
            reconnectAttempts = 0
            _messages.trySend(WsMessage("connected", null))
        }

        override fun onMessage(webSocket: WebSocket, text: String) {
            try {
                val message = gson.fromJson(text, WsMessage::class.java)
                _messages.trySend(message)
            } catch (e: Exception) {
                e.printStackTrace()
            }
        }

        override fun onClosing(webSocket: WebSocket, code: Int, reason: String) {
            webSocket.close(1000, null)
        }

        override fun onClosed(webSocket: WebSocket, code: Int, reason: String) {
            isConnected = false
            _messages.trySend(WsMessage("disconnected", mapOf("code" to code, "reason" to reason)))
        }

        override fun onFailure(webSocket: WebSocket, t: Throwable, response: Response?) {
            isConnected = false
            _messages.trySend(WsMessage("error", mapOf("message" to (t.message ?: "Unknown error"))))
            reconnect()
        }
    }
}
