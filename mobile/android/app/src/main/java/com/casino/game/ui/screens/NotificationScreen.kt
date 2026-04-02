package com.casino.game.ui.screens

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
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

data class NotificationState(
    val notifications: List<AppNotification> = emptyList(),
    val isLoading: Boolean = false,
    val isRefreshing: Boolean = false,
    val error: String? = null
)

@HiltViewModel
class NotificationViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(NotificationState())
    val state: StateFlow<NotificationState> = _state.asStateFlow()

    init { loadNotifications() }

    fun loadNotifications() {
        _state.update { it.copy(isLoading = it.notifications.isEmpty()) }
        viewModelScope.launch {
            try {
                val response = apiService.getNotifications()
                if (response.isSuccessful) {
                    _state.update { it.copy(notifications = response.body() ?: emptyList(), isLoading = false, isRefreshing = false) }
                }
            } catch (_: Exception) {
                _state.update { it.copy(isLoading = false, isRefreshing = false) }
            }
        }
    }

    fun refresh() {
        _state.update { it.copy(isRefreshing = true) }
        loadNotifications()
    }

    fun markAsRead(id: String) {
        viewModelScope.launch {
            try {
                apiService.markNotificationRead(id)
                _state.update { s -> s.copy(notifications = s.notifications.map { if (it.id == id) it.copy(read = true) else it }) }
            } catch (_: Exception) {}
        }
    }

    fun deleteNotification(id: String) {
        viewModelScope.launch {
            try {
                apiService.deleteNotification(id)
                _state.update { s -> s.copy(notifications = s.notifications.filter { it.id != id }) }
            } catch (_: Exception) {}
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun NotificationScreen(
    onBack: () -> Unit,
    viewModel: NotificationViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()

    Column(modifier = Modifier.fillMaxSize()) {
        TopAppBar(
            title = { Text("Notifications") },
            navigationIcon = { IconButton(onClick = onBack) { Icon(Icons.Default.ArrowBack, contentDescription = "Back") } }
        )

        if (state.isLoading) {
            Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) { CircularProgressIndicator() }
            return
        }

        val pullRefreshState = rememberPullToRefreshState()
        if (state.isRefreshing) {
            LaunchedEffect(Unit) { viewModel.refresh() }
        }

        PullToRefreshBox(state = pullRefreshState, isRefreshing = state.isRefreshing, onRefresh = viewModel::refresh) {
            LazyColumn(modifier = Modifier.fillMaxSize(), contentPadding = PaddingValues(16.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
                if (state.notifications.isEmpty()) {
                    item {
                        Box(modifier = Modifier.fillMaxWidth().padding(32.dp), contentAlignment = Alignment.Center) {
                            Column(horizontalAlignment = Alignment.CenterHorizontally) {
                                Icon(Icons.Default.NotificationsNone, contentDescription = null, modifier = Modifier.size(64.dp), tint = MaterialTheme.colorScheme.onSurfaceVariant)
                                Spacer(modifier = Modifier.height(8.dp))
                                Text("No notifications")
                            }
                        }
                    }
                }
                items(state.notifications, key = { it.id }) { notification ->
                    SwipeToDismissBox(
                        state = rememberSwipeToDismissBoxState(),
                        backgroundContent = {
                            Surface(modifier = Modifier.fillMaxSize(), color = MaterialTheme.colorScheme.errorContainer) {
                                Box(modifier = Modifier.fillMaxSize().padding(horizontal = 20.dp), contentAlignment = Alignment.CenterEnd) {
                                    Icon(Icons.Default.Delete, contentDescription = "Delete", tint = MaterialTheme.colorScheme.onErrorContainer)
                                }
                            }
                        },
                        content = {
                            NotificationItem(notification = notification, onTap = { viewModel.markAsRead(notification.id) })
                        }
                    )
                }
            }
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
private fun NotificationItem(notification: AppNotification, onTap: () -> Unit) {
    Card(onClick = onTap, modifier = Modifier.fillMaxWidth(), colors = if (!notification.read) CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.primaryContainer.copy(alpha = 0.3f)) else CardDefaults.cardColors()) {
        Row(modifier = Modifier.padding(16.dp), verticalAlignment = Alignment.Top) {
            Icon(
                imageVector = when (notification.type) {
                    "bonus" -> Icons.Default.CardGiftcard
                    "tournament" -> Icons.Default.EmojiEvents
                    "jackpot" -> Icons.Default.Stars
                    else -> Icons.Default.Info
                },
                contentDescription = null,
                tint = when (notification.type) {
                    "bonus" -> MaterialTheme.colorScheme.tertiary
                    "tournament" -> MaterialTheme.colorScheme.primary
                    "jackpot" -> MaterialTheme.colorScheme.secondary
                    else -> MaterialTheme.colorScheme.onSurfaceVariant
                }
            )
            Spacer(modifier = Modifier.width(12.dp))
            Column(modifier = Modifier.weight(1f)) {
                Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceBetween) {
                    Text(notification.title, style = MaterialTheme.typography.bodyLarge, fontWeight = if (!notification.read) FontWeight.Bold else FontWeight.Normal)
                    if (!notification.read) {
                        Badge(modifier = Modifier.size(8.dp), containerColor = MaterialTheme.colorScheme.primary) {}
                    }
                }
                Spacer(modifier = Modifier.height(4.dp))
                Text(notification.message, style = MaterialTheme.typography.bodyMedium, color = MaterialTheme.colorScheme.onSurfaceVariant)
                Text(notification.createdAt, style = MaterialTheme.typography.labelSmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
            }
        }
    }
}
