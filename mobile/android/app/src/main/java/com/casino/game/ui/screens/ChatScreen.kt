package com.casino.game.ui.screens

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.lazy.rememberLazyListState
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.casino.game.data.model.*
import com.casino.game.data.remote.ApiService
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class ChatState(
    val messages: List<ChatMessage> = emptyList(),
    val rooms: List<ChatRoom> = emptyList(),
    val selectedRoomId: String? = null,
    val onlineCount: Int = 0,
    val messageText: String = "",
    val isLoading: Boolean = false,
    val error: String? = null
)

@HiltViewModel
class ChatViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(ChatState())
    val state: StateFlow<ChatState> = _state.asStateFlow()

    init { loadRooms() }

    fun loadRooms() {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            try {
                val response = apiService.getChatRooms()
                if (response.isSuccessful) {
                    _state.update { it.copy(rooms = response.body() ?: emptyList(), isLoading = false) }
                }
            } catch (_: Exception) {
                _state.update { it.copy(isLoading = false) }
            }
        }
    }

    fun selectRoom(roomId: String) {
        _state.update { it.copy(selectedRoomId = roomId, messages = emptyList()) }
        loadMessages(roomId)
    }

    private fun loadMessages(roomId: String) {
        viewModelScope.launch {
            try {
                val response = apiService.getChatMessages(roomId)
                if (response.isSuccessful) {
                    _state.update { it.copy(messages = response.body() ?: emptyList()) }
                }
            } catch (_: Exception) {}
        }
    }

    fun updateMessageText(text: String) { _state.update { it.copy(messageText = text) } }

    fun sendMessage() {
        val roomId = _state.value.selectedRoomId ?: return
        val text = _state.value.messageText.trim()
        if (text.isEmpty()) return
        viewModelScope.launch {
            try {
                val response = apiService.sendChatMessage(roomId, mapOf("content" to text))
                if (response.isSuccessful) {
                    response.body()?.let { msg ->
                        _state.update { it.copy(messages = it.messages + msg, messageText = "") }
                    }
                }
            } catch (_: Exception) {}
        }
    }
}

@Composable
fun ChatScreen(
    onBack: () -> Unit,
    viewModel: ChatViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()

    Column(modifier = Modifier.fillMaxSize()) {
        TopAppBar(
            title = { Text(if (state.selectedRoomId != null) "Chat" else "Chat Rooms") },
            navigationIcon = {
                IconButton(onClick = {
                    if (state.selectedRoomId != null) viewModel.loadRooms() else onBack()
                }) {
                    Icon(Icons.Default.ArrowBack, contentDescription = "Back")
                }
            },
            actions = {
                if (state.selectedRoomId != null) {
                    Row(verticalAlignment = Alignment.CenterVertically) {
                        Icon(Icons.Default.Circle, contentDescription = null, modifier = Modifier.size(8.dp), tint = MaterialTheme.colorScheme.primary)
                        Spacer(modifier = Modifier.width(4.dp))
                        Text("${state.onlineCount}", style = MaterialTheme.typography.bodySmall)
                        Spacer(modifier = Modifier.width(8.dp))
                    }
                }
            }
        )

        if (state.selectedRoomId == null) {
            ChatRoomList(rooms = state.rooms, isLoading = state.isLoading, onRoomClick = { viewModel.selectRoom(it.id) })
        } else {
            ChatMessagesView(state = state, onTextChange = viewModel::updateMessageText, onSend = viewModel::sendMessage)
        }
    }
}

@Composable
private fun ChatRoomList(rooms: List<ChatRoom>, isLoading: Boolean, onRoomClick: (ChatRoom) -> Unit) {
    if (isLoading) {
        Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) { CircularProgressIndicator() }
        return
    }
    LazyColumn(modifier = Modifier.fillMaxSize(), contentPadding = PaddingValues(16.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
        if (rooms.isEmpty()) {
            item {
                Box(modifier = Modifier.fillMaxWidth().padding(32.dp), contentAlignment = Alignment.Center) {
                    Column(horizontalAlignment = Alignment.CenterHorizontally) {
                        Icon(Icons.Default.Chat, contentDescription = null, modifier = Modifier.size(64.dp), tint = MaterialTheme.colorScheme.onSurfaceVariant)
                        Spacer(modifier = Modifier.height(8.dp))
                        Text("No chat rooms available")
                    }
                }
            }
        }
        items(rooms) { room ->
            Card(onClick = { onRoomClick(room) }, modifier = Modifier.fillMaxWidth()) {
                Row(modifier = Modifier.padding(16.dp), verticalAlignment = Alignment.CenterVertically) {
                    Surface(shape = CircleShape, color = MaterialTheme.colorScheme.primaryContainer, modifier = Modifier.size(40.dp)) {
                        Box(contentAlignment = Alignment.Center) {
                            Icon(Icons.Default.Chat, contentDescription = null, modifier = Modifier.size(20.dp))
                        }
                    }
                    Spacer(modifier = Modifier.width(12.dp))
                    Column(modifier = Modifier.weight(1f)) {
                        Text(room.name, style = MaterialTheme.typography.bodyLarge, fontWeight = FontWeight.Medium)
                        Text(room.lastMessage ?: "No messages yet", style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant, maxLines = 1)
                    }
                    room.unreadCount?.takeIf { it > 0 }?.let {
                        Badge { Text(it.toString()) }
                    }
                }
            }
        }
    }
}

@Composable
private fun ChatMessagesView(state: ChatState, onTextChange: (String) -> Unit, onSend: () -> Unit) {
    val listState = rememberLazyListState()
    LaunchedEffect(state.messages.size) { if (state.messages.isNotEmpty()) listState.animateScrollToItem(state.messages.size - 1) }

    Column(modifier = Modifier.fillMaxSize()) {
        LazyColumn(state = listState, modifier = Modifier.weight(1f).padding(horizontal = 16.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
            items(state.messages) { message -> ChatMessageItem(message = message) }
        }
        Surface(tonalElevation = 3.dp) {
            Row(modifier = Modifier.fillMaxWidth().padding(8.dp), verticalAlignment = Alignment.CenterVertically) {
                OutlinedTextField(value = state.messageText, onValueChange = onTextChange, modifier = Modifier.weight(1f), placeholder = { Text("Type a message...") }, singleLine = true)
                Spacer(modifier = Modifier.width(8.dp))
                IconButton(onClick = onSend, enabled = state.messageText.isNotBlank()) {
                    Icon(Icons.Default.Send, contentDescription = "Send", tint = if (state.messageText.isNotBlank()) MaterialTheme.colorScheme.primary else MaterialTheme.colorScheme.onSurfaceVariant)
                }
            }
        }
    }
}

@Composable
private fun ChatMessageItem(message: ChatMessage) {
    Column(modifier = Modifier.fillMaxWidth()) {
        Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = if (message.isMine) Arrangement.End else Arrangement.Start) {
            Card(modifier = Modifier.widthIn(max = 280.dp), colors = CardDefaults.cardColors(containerColor = if (message.isMine) MaterialTheme.colorScheme.primaryContainer else MaterialTheme.colorScheme.surfaceVariant)) {
                Column(modifier = Modifier.padding(12.dp)) {
                    if (!message.isMine) {
                        Text(message.username, style = MaterialTheme.typography.labelSmall, fontWeight = FontWeight.Bold, color = MaterialTheme.colorScheme.primary)
                        Spacer(modifier = Modifier.height(2.dp))
                    }
                    Text(message.content, style = MaterialTheme.typography.bodyMedium)
                    Text(message.timestamp, style = MaterialTheme.typography.labelSmall, color = MaterialTheme.colorScheme.onSurfaceVariant, modifier = Modifier.align(Alignment.End))
                }
            }
        }
    }
}
