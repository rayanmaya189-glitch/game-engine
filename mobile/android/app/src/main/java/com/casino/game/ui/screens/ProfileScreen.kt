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
import com.casino.game.data.model.UserProfile
import com.casino.game.data.repository.AuthRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class PlayerProfileState(
    val username: String = "",
    val email: String = "",
    val phone: String = "",
    val kycLevel: String = "none",
    val avatarUrl: String? = null,
    val referralCode: String = "",
    val isLoading: Boolean = false,
    val error: String? = null,
    val notificationsEnabled: Boolean = true,
    val twoFactorEnabled: Boolean = false,
    val language: String = "English"
)

@HiltViewModel
class PlayerProfileViewModel @Inject constructor(
    private val authRepository: AuthRepository
) : ViewModel() {
    private val _state = MutableStateFlow(PlayerProfileState())
    val state: StateFlow<PlayerProfileState> = _state.asStateFlow()

    init { loadProfile() }

    private fun loadProfile() {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            try {
                val token = authRepository.getAuthToken()
                if (token != null) {
                    _state.update {
                        it.copy(
                            username = "Player",
                            email = "player@casino.com",
                            referralCode = "REF_${System.currentTimeMillis().toString().takeLast(6)}",
                            isLoading = false
                        )
                    }
                }
            } catch (e: Exception) {
                _state.update { it.copy(isLoading = false, error = e.message) }
            }
        }
    }

    fun toggleNotifications() {
        _state.update { it.copy(notificationsEnabled = !it.notificationsEnabled) }
    }

    fun toggle2FA() {
        _state.update { it.copy(twoFactorEnabled = !it.twoFactorEnabled) }
    }

    fun setLanguage(lang: String) {
        _state.update { it.copy(language = lang) }
    }

    fun logout(onSuccess: () -> Unit) {
        viewModelScope.launch {
            authRepository.logout()
            onSuccess()
        }
    }
}

@Composable
fun ProfileScreen(
    onLogout: () -> Unit,
    onNavigateToDeposit: () -> Unit,
    onNavigateToWithdraw: () -> Unit,
    viewModel: PlayerProfileViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()
    val languages = listOf("English", "Spanish", "German", "French", "Japanese")
    var showLanguageDialog by remember { mutableStateOf(false) }

    Column(
        modifier = Modifier
            .fillMaxSize()
            .verticalScroll(rememberScrollState())
            .padding(16.dp)
    ) {
        Card(modifier = Modifier.fillMaxWidth()) {
            Column(
                modifier = Modifier.padding(24.dp),
                horizontalAlignment = Alignment.CenterHorizontally
            ) {
                Icon(
                    imageVector = Icons.Default.AccountCircle,
                    contentDescription = null,
                    modifier = Modifier.size(80.dp),
                    tint = MaterialTheme.colorScheme.primary
                )
                Spacer(modifier = Modifier.height(12.dp))
                Text(state.username, style = MaterialTheme.typography.headlineSmall)
                Text(state.email, style = MaterialTheme.typography.bodyMedium)
                Text(state.phone, style = MaterialTheme.typography.bodySmall)
                Spacer(modifier = Modifier.height(8.dp))
                AssistChip(
                    onClick = { },
                    label = { Text("KYC: ${state.kycLevel}") },
                    leadingIcon = {
                        Icon(
                            if (state.kycLevel == "verified") Icons.Default.VerifiedUser else Icons.Default.Warning,
                            contentDescription = null,
                            modifier = Modifier.size(16.dp)
                        )
                    }
                )
            }
        }

        Spacer(modifier = Modifier.height(16.dp))

        Text("Account", style = MaterialTheme.typography.titleMedium)
        Spacer(modifier = Modifier.height(8.dp))

        ProfileMenuItem(icon = Icons.Default.Edit, title = "Edit Profile") { }
        ProfileMenuItem(icon = Icons.Default.Upload, title = "Upload KYC Documents") { }
        ProfileMenuItem(icon = Icons.Default.Share, title = "Referral Code: ${state.referralCode}") { }

        Spacer(modifier = Modifier.height(16.dp))
        Text("Settings", style = MaterialTheme.typography.titleMedium)
        Spacer(modifier = Modifier.height(8.dp))

        Card(modifier = Modifier.fillMaxWidth().padding(vertical = 4.dp)) {
            Row(
                modifier = Modifier.padding(16.dp),
                verticalAlignment = Alignment.CenterVertically
            ) {
                Icon(Icons.Default.Language, contentDescription = null)
                Spacer(modifier = Modifier.width(16.dp))
                Text(state.language, style = MaterialTheme.typography.bodyLarge)
                Spacer(modifier = Modifier.weight(1f))
                TextButton(onClick = { showLanguageDialog = true }) { Text("Change") }
            }
        }

        Card(modifier = Modifier.fillMaxWidth().padding(vertical = 4.dp)) {
            Row(
                modifier = Modifier.padding(16.dp),
                verticalAlignment = Alignment.CenterVertically
            ) {
                Icon(Icons.Default.Notifications, contentDescription = null)
                Spacer(modifier = Modifier.width(16.dp))
                Text("Notifications", style = MaterialTheme.typography.bodyLarge)
                Spacer(modifier = Modifier.weight(1f))
                Switch(
                    checked = state.notificationsEnabled,
                    onCheckedChange = { viewModel.toggleNotifications() }
                )
            }
        }

        Card(modifier = Modifier.fillMaxWidth().padding(vertical = 4.dp)) {
            Row(
                modifier = Modifier.padding(16.dp),
                verticalAlignment = Alignment.CenterVertically
            ) {
                Icon(Icons.Default.Security, contentDescription = null)
                Spacer(modifier = Modifier.width(16.dp))
                Text("Two-Factor Auth", style = MaterialTheme.typography.bodyLarge)
                Spacer(modifier = Modifier.weight(1f))
                Switch(
                    checked = state.twoFactorEnabled,
                    onCheckedChange = { viewModel.toggle2FA() }
                )
            }
        }

        Spacer(modifier = Modifier.weight(1f))

        Button(
            onClick = { viewModel.logout(onLogout) },
            modifier = Modifier.fillMaxWidth(),
            colors = ButtonDefaults.buttonColors(containerColor = MaterialTheme.colorScheme.error)
        ) {
            Icon(Icons.Default.Logout, contentDescription = null)
            Spacer(modifier = Modifier.width(8.dp))
            Text("Logout")
        }
    }

    if (showLanguageDialog) {
        AlertDialog(
            onDismissRequest = { showLanguageDialog = false },
            title = { Text("Select Language") },
            text = {
                Column {
                    languages.forEach { lang ->
                        Card(
                            onClick = {
                                viewModel.setLanguage(lang)
                                showLanguageDialog = false
                            },
                            modifier = Modifier.fillMaxWidth().padding(vertical = 4.dp)
                        ) {
                            Row(modifier = Modifier.padding(16.dp)) {
                                RadioButton(
                                    selected = state.language == lang,
                                    onClick = {
                                        viewModel.setLanguage(lang)
                                        showLanguageDialog = false
                                    }
                                )
                                Spacer(modifier = Modifier.width(8.dp))
                                Text(lang)
                            }
                        }
                    }
                }
            },
            confirmButton = {
                TextButton(onClick = { showLanguageDialog = false }) { Text("Cancel") }
            }
        )
    }
}

@Composable
private fun ProfileMenuItem(icon: androidx.compose.ui.graphics.vector.ImageVector, title: String, onClick: () -> Unit) {
    Card(
        onClick = onClick,
        modifier = Modifier.fillMaxWidth().padding(vertical = 4.dp)
    ) {
        Row(
            modifier = Modifier.padding(16.dp),
            verticalAlignment = Alignment.CenterVertically
        ) {
            Icon(icon, contentDescription = null)
            Spacer(modifier = Modifier.width(16.dp))
            Text(title, style = MaterialTheme.typography.bodyLarge)
            Spacer(modifier = Modifier.weight(1f))
            Icon(Icons.Default.ChevronRight, contentDescription = null)
        }
    }
}
