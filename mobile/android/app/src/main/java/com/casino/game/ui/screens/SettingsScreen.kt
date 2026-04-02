package com.casino.game.ui.screens

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.casino.game.data.remote.ApiService
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class SettingsState(
    val username: String = "",
    val email: String = "",
    val avatarUrl: String? = null,
    val language: String = "English",
    val theme: String = "System",
    val pushNotifications: Boolean = true,
    val emailNotifications: Boolean = true,
    val smsNotifications: Boolean = false,
    val twoFactorEnabled: Boolean = false,
    val appVersion: String = "1.0.0",
    val isLoading: Boolean = false,
    val showDeleteConfirm: Boolean = false
)

@HiltViewModel
class SettingsViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(SettingsState())
    val state: StateFlow<SettingsState> = _state.asStateFlow()

    init { loadSettings() }

    private fun loadSettings() {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            try {
                val response = apiService.getProfile()
                if (response.isSuccessful) {
                    response.body()?.let { p ->
                        _state.update { it.copy(username = p.username, email = p.email, avatarUrl = p.avatarUrl, isLoading = false) }
                    }
                }
            } catch (_: Exception) { _state.update { it.copy(isLoading = false) } }
        }
    }

    fun setLanguage(lang: String) { _state.update { it.copy(language = lang) } }
    fun setTheme(theme: String) { _state.update { it.copy(theme = theme) } }
    fun togglePush() { _state.update { it.copy(pushNotifications = !it.pushNotifications) } }
    fun toggleEmail() { _state.update { it.copy(emailNotifications = !it.emailNotifications) } }
    fun toggleSms() { _state.update { it.copy(smsNotifications = !it.smsNotifications) } }
    fun toggle2FA() { _state.update { it.copy(twoFactorEnabled = !it.twoFactorEnabled) } }
    fun showDeleteDialog() { _state.update { it.copy(showDeleteConfirm = true) } }
    fun hideDeleteDialog() { _state.update { it.copy(showDeleteConfirm = false) } }

    fun deleteAccount(onSuccess: () -> Unit) {
        viewModelScope.launch {
            try { apiService.deleteAccount(); onSuccess() } catch (_: Exception) {}
        }
    }

    fun logout(onSuccess: () -> Unit) {
        viewModelScope.launch {
            try { apiService.logout(); onSuccess() } catch (_: Exception) { onSuccess() }
        }
    }
}

@Composable
fun SettingsScreen(
    onBack: () -> Unit,
    onLogout: () -> Unit,
    viewModel: SettingsViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()
    var showLanguageDialog by remember { mutableStateOf(false) }
    var showThemeDialog by remember { mutableStateOf(false) }
    val languages = listOf("English", "Spanish", "German", "French", "Japanese")
    val themes = listOf("System", "Light", "Dark")

    Column(modifier = Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
        TopAppBar(title = { Text("Settings") }, navigationIcon = { IconButton(onClick = onBack) { Icon(Icons.Default.ArrowBack, contentDescription = "Back") } })

        Card(modifier = Modifier.fillMaxWidth().padding(16.dp)) {
            Row(modifier = Modifier.padding(16.dp), verticalAlignment = Alignment.CenterVertically) {
                Icon(Icons.Default.AccountCircle, contentDescription = null, modifier = Modifier.size(56.dp), tint = MaterialTheme.colorScheme.primary)
                Spacer(modifier = Modifier.width(16.dp))
                Column { Text(state.username, style = MaterialTheme.typography.titleMedium); Text(state.email, style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant) }
            }
        }

        Text("Preferences", style = MaterialTheme.typography.titleMedium, modifier = Modifier.padding(horizontal = 16.dp))
        Spacer(modifier = Modifier.height(8.dp))

        SettingsItem(icon = Icons.Default.Language, title = "Language", subtitle = state.language, onClick = { showLanguageDialog = true })
        SettingsItem(icon = Icons.Default.Palette, title = "Theme", subtitle = state.theme, onClick = { showThemeDialog = true })

        Spacer(modifier = Modifier.height(16.dp))
        Text("Notifications", style = MaterialTheme.typography.titleMedium, modifier = Modifier.padding(horizontal = 16.dp))
        Spacer(modifier = Modifier.height(8.dp))

        SettingsToggleItem(icon = Icons.Default.Notifications, title = "Push Notifications", checked = state.pushNotifications, onToggle = viewModel::togglePush)
        SettingsToggleItem(icon = Icons.Default.Email, title = "Email Notifications", checked = state.emailNotifications, onToggle = viewModel::toggleEmail)
        SettingsToggleItem(icon = Icons.Default.Sms, title = "SMS Notifications", checked = state.smsNotifications, onToggle = viewModel::toggleSms)

        Spacer(modifier = Modifier.height(16.dp))
        Text("Security", style = MaterialTheme.typography.titleMedium, modifier = Modifier.padding(horizontal = 16.dp))
        Spacer(modifier = Modifier.height(8.dp))

        SettingsToggleItem(icon = Icons.Default.Security, title = "Two-Factor Authentication", checked = state.twoFactorEnabled, onToggle = viewModel::toggle2FA)
        SettingsItem(icon = Icons.Default.Lock, title = "Change Password", onClick = {})

        Spacer(modifier = Modifier.height(16.dp))

        SettingsItem(icon = Icons.Default.DeleteForever, title = "Delete Account", titleColor = MaterialTheme.colorScheme.error, onClick = viewModel::showDeleteDialog)
        Spacer(modifier = Modifier.height(8.dp))

        Button(onClick = { viewModel.logout(onLogout) }, modifier = Modifier.fillMaxWidth().padding(16.dp), colors = ButtonDefaults.buttonColors(containerColor = MaterialTheme.colorScheme.error)) {
            Icon(Icons.Default.Logout, contentDescription = null); Spacer(modifier = Modifier.width(8.dp)); Text("Logout")
        }

        Text("Version ${state.appVersion}", style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant, modifier = Modifier.align(Alignment.CenterHorizontally).padding(16.dp))
    }

    if (showLanguageDialog) {
        SelectionDialog(title = "Select Language", options = languages, selected = state.language, onSelect = { viewModel.setLanguage(it); showLanguageDialog = false }, onDismiss = { showLanguageDialog = false })
    }
    if (showThemeDialog) {
        SelectionDialog(title = "Select Theme", options = themes, selected = state.theme, onSelect = { viewModel.setTheme(it); showThemeDialog = false }, onDismiss = { showThemeDialog = false })
    }
    if (state.showDeleteConfirm) {
        AlertDialog(onDismissRequest = viewModel::hideDeleteDialog, title = { Text("Delete Account") }, text = { Text("This action cannot be undone. All your data will be permanently deleted.") }, confirmButton = { TextButton(onClick = { viewModel.deleteAccount(onLogout) }) { Text("Delete", color = MaterialTheme.colorScheme.error) } }, dismissButton = { TextButton(onClick = viewModel::hideDeleteDialog) { Text("Cancel") } })
    }
}

@Composable
private fun SettingsItem(icon: androidx.compose.ui.graphics.vector.ImageVector, title: String, subtitle: String? = null, titleColor: androidx.compose.ui.graphics.Color = MaterialTheme.colorScheme.onSurface, onClick: () -> Unit) {
    Card(onClick = onClick, modifier = Modifier.fillMaxWidth().padding(horizontal = 16.dp, vertical = 4.dp)) {
        Row(modifier = Modifier.padding(16.dp), verticalAlignment = Alignment.CenterVertically) {
            Icon(icon, contentDescription = null, tint = if (titleColor == MaterialTheme.colorScheme.error) titleColor else MaterialTheme.colorScheme.onSurfaceVariant)
            Spacer(modifier = Modifier.width(16.dp))
            Column(modifier = Modifier.weight(1f)) {
                Text(title, style = MaterialTheme.typography.bodyLarge, color = titleColor)
                subtitle?.let { Text(it, style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant) }
            }
            Icon(Icons.Default.ChevronRight, contentDescription = null)
        }
    }
}

@Composable
private fun SettingsToggleItem(icon: androidx.compose.ui.graphics.vector.ImageVector, title: String, checked: Boolean, onToggle: () -> Unit) {
    Card(modifier = Modifier.fillMaxWidth().padding(horizontal = 16.dp, vertical = 4.dp)) {
        Row(modifier = Modifier.padding(16.dp), verticalAlignment = Alignment.CenterVertically) {
            Icon(icon, contentDescription = null, tint = MaterialTheme.colorScheme.onSurfaceVariant)
            Spacer(modifier = Modifier.width(16.dp))
            Text(title, style = MaterialTheme.typography.bodyLarge, modifier = Modifier.weight(1f))
            Switch(checked = checked, onCheckedChange = { onToggle() })
        }
    }
}

@Composable
private fun SelectionDialog(title: String, options: List<String>, selected: String, onSelect: (String) -> Unit, onDismiss: () -> Unit) {
    AlertDialog(onDismissRequest = onDismiss, title = { Text(title) }, text = {
        Column { options.forEach { option ->
            Card(onClick = { onSelect(option) }, modifier = Modifier.fillMaxWidth().padding(vertical = 4.dp)) {
                Row(modifier = Modifier.padding(16.dp), verticalAlignment = Alignment.CenterVertically) {
                    RadioButton(selected = selected == option, onClick = { onSelect(option) }); Spacer(modifier = Modifier.width(8.dp)); Text(option)
                }
            }
        }}
    }, confirmButton = { TextButton(onClick = onDismiss) { Text("Cancel") } })
}
