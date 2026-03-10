package com.gameengine.casino.ui.profile

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import com.gameengine.casino.ui.theme.*

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun ProfileScreen(
    onNavigateToEditProfile: () -> Unit,
    onLogout: () -> Unit,
    viewModel: ProfileViewModel = hiltViewModel()
) {
    val uiState by viewModel.uiState.collectAsState()
    var showLogoutDialog by remember { mutableStateOf(false) }
    
    LaunchedEffect(uiState.isLoggedOut) {
        if (uiState.isLoggedOut) {
            onLogout()
        }
    }
    
    if (showLogoutDialog) {
        AlertDialog(
            onDismissRequest = { showLogoutDialog = false },
            title = { Text("Logout") },
            text = { Text("Are you sure you want to logout?") },
            confirmButton = {
                TextButton(
                    onClick = {
                        showLogoutDialog = false
                        viewModel.logout()
                    }
                ) {
                    Text("Logout", color = Error)
                }
            },
            dismissButton = {
                TextButton(onClick = { showLogoutDialog = false }) {
                    Text("Cancel")
                }
            }
        )
    }
    
    Column(
        modifier = Modifier
            .fillMaxSize()
            .background(MaterialTheme.colorScheme.background)
    ) {
        // Profile Header
        Box(
            modifier = Modifier
                .fillMaxWidth()
                .background(Primary.copy(alpha = 0.1f))
                .padding(24.dp)
        ) {
            Column(
                horizontalAlignment = Alignment.CenterHorizontally,
                modifier = Modifier.fillMaxWidth()
            ) {
                // Avatar
                Box(
                    modifier = Modifier
                        .size(80.dp)
                        .clip(CircleShape)
                        .background(Primary),
                    contentAlignment = Alignment.Center
                ) {
                    Text(
                        text = uiState.user?.username?.take(2)?.uppercase() ?: "?",
                        style = MaterialTheme.typography.headlineMedium,
                        color = OnPrimary
                    )
                }
                
                Spacer(modifier = Modifier.height(12.dp))
                
                Text(
                    text = uiState.user?.username ?: "User",
                    style = MaterialTheme.typography.titleLarge,
                    fontWeight = FontWeight.Bold
                )
                Text(
                    text = uiState.user?.email ?: "",
                    style = MaterialTheme.typography.bodyMedium,
                    color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.7f)
                )
                
                Spacer(modifier = Modifier.height(8.dp))
                
                // KYC Badge
                Surface(
                    shape = RoundedCornerShape(16.dp),
                    color = if (uiState.user?.isVerified == true) Success.copy(alpha = 0.2f) else Warning.copy(alpha = 0.2f)
                ) {
                    Row(
                        modifier = Modifier.padding(horizontal = 12.dp, vertical = 6.dp),
                        verticalAlignment = Alignment.CenterVertically
                    ) {
                        Icon(
                            imageVector = if (uiState.user?.isVerified == true) Icons.Default.Verified else Icons.Default.Warning,
                            contentDescription = null,
                            modifier = Modifier.size(16.dp),
                            tint = if (uiState.user?.isVerified == true) Success else Warning
                        )
                        Spacer(modifier = Modifier.width(4.dp))
                        Text(
                            text = if (uiState.user?.isVerified == true) "Verified" else "Pending Verification",
                            style = MaterialTheme.typography.bodySmall,
                            color = if (uiState.user?.isVerified == true) Success else Warning
                        )
                    }
                }
            }
        }
        
        Column(
            modifier = Modifier
                .fillMaxSize()
                .verticalScroll(rememberScrollState())
                .padding(16.dp)
        ) {
            // Edit Profile
            ProfileMenuItem(
                icon = Icons.Default.Edit,
                title = "Edit Profile",
                onClick = onNavigateToEditProfile
            )
            
            // KYC Verification
            ProfileMenuItem(
                icon = Icons.Default.Badge,
                title = "KYC Verification",
                subtitle = "Level ${uiState.user?.kycLevel ?: 0}",
                onClick = { /* Navigate to KYC */ }
            )
            
            // Bonuses
            ProfileMenuItem(
                icon = Icons.Default.CardGiftcard,
                title = "My Bonuses",
                onClick = { /* Navigate to Bonuses */ }
            )
            
            // Transaction History
            ProfileMenuItem(
                icon = Icons.Default.Receipt,
                title = "Transaction History",
                onClick = { /* Navigate to History */ }
            )
            
            // Responsible Gaming
            ProfileMenuItem(
                icon = Icons.Default.Security,
                title = "Responsible Gaming",
                onClick = { /* Navigate to RG */ }
            )
            
            // Settings
            ProfileMenuItem(
                icon = Icons.Default.Settings,
                title = "Settings",
                onClick = { /* Navigate to Settings */ }
            )
            
            // Help & Support
            ProfileMenuItem(
                icon = Icons.Default.Help,
                title = "Help & Support",
                onClick = { /* Navigate to Help */ }
            )
            
            Spacer(modifier = Modifier.height(16.dp))
            
            // Logout Button
            OutlinedButton(
                onClick = { showLogoutDialog = true },
                modifier = Modifier
                    .fillMaxWidth()
                    .height(56.dp),
                shape = RoundedCornerShape(12.dp),
                colors = ButtonDefaults.outlinedButtonColors(contentColor = Error)
            ) {
                Icon(
                    imageVector = Icons.Default.Logout,
                    contentDescription = null,
                    tint = Error
                )
                Spacer(modifier = Modifier.width(8.dp))
                Text("Logout", color = Error)
            }
        }
    }
}

@Composable
fun ProfileMenuItem(
    icon: androidx.compose.ui.graphics.vector.ImageVector,
    title: String,
    subtitle: String? = null,
    onClick: () -> Unit
) {
    Card(
        modifier = Modifier
            .fillMaxWidth()
            .padding(vertical = 4.dp)
            .clickable(onClick = onClick),
        colors = CardDefaults.cardColors(containerColor = CardBackground),
        shape = RoundedCornerShape(12.dp)
    ) {
        Row(
            modifier = Modifier
                .fillMaxWidth()
                .padding(16.dp),
            verticalAlignment = Alignment.CenterVertically
        ) {
            Icon(
                imageVector = icon,
                contentDescription = null,
                tint = Primary,
                modifier = Modifier.size(24.dp)
            )
            Spacer(modifier = Modifier.width(12.dp))
            Column(modifier = Modifier.weight(1f)) {
                Text(
                    text = title,
                    style = MaterialTheme.typography.bodyLarge,
                    fontWeight = FontWeight.Medium
                )
                subtitle?.let {
                    Text(
                        text = it,
                        style = MaterialTheme.typography.bodySmall,
                        color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.6f)
                    )
                }
            }
            Icon(
                imageVector = Icons.Default.ChevronRight,
                contentDescription = null,
                tint = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.5f)
            )
        }
    }
}
