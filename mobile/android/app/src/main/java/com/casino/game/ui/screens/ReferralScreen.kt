package com.casino.game.ui.screens

import android.content.ClipData
import android.content.ClipboardManager
import android.content.Context
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
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

data class ReferralState(
    val referralCode: String = "",
    val totalReferrals: Int = 0,
    val totalEarnings: Double = 0.0,
    val referralHistory: List<ReferralEntry> = emptyList(),
    val tiers: List<ReferralTier> = emptyList(),
    val isLoading: Boolean = false,
    val copied: Boolean = false
)

@HiltViewModel
class ReferralViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(ReferralState())
    val state: StateFlow<ReferralState> = _state.asStateFlow()

    init { loadReferralData() }

    private fun loadReferralData() {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            try {
                val statsResp = apiService.getReferralStats()
                if (statsResp.isSuccessful) {
                    statsResp.body()?.let { s ->
                        _state.update { it.copy(referralCode = s.code, totalReferrals = s.totalReferrals, totalEarnings = s.totalEarnings) }
                    }
                }
                val historyResp = apiService.getReferralHistory()
                if (historyResp.isSuccessful) {
                    _state.update { it.copy(referralHistory = historyResp.body() ?: emptyList()) }
                }
                val tiersResp = apiService.getReferralTiers()
                if (tiersResp.isSuccessful) {
                    _state.update { it.copy(tiers = tiersResp.body() ?: emptyList(), isLoading = false) }
                }
            } catch (_: Exception) { _state.update { it.copy(isLoading = false) } }
        }
    }

    fun markCopied() { _state.update { it.copy(copied = true) } }
}

@Composable
fun ReferralScreen(
    onBack: () -> Unit,
    viewModel: ReferralViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()
    val context = LocalContext.current

    LazyColumn(modifier = Modifier.fillMaxSize(), contentPadding = PaddingValues(16.dp), verticalArrangement = Arrangement.spacedBy(12.dp)) {
        item {
            TopAppBar(title = { Text("Referral Program") }, navigationIcon = { IconButton(onClick = onBack) { Icon(Icons.Default.ArrowBack, contentDescription = "Back") } })
        }

        item {
            Card(modifier = Modifier.fillMaxWidth(), colors = CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.primaryContainer)) {
                Column(modifier = Modifier.padding(20.dp), horizontalAlignment = Alignment.CenterHorizontally) {
                    Text("Your Referral Code", style = MaterialTheme.typography.labelMedium)
                    Spacer(modifier = Modifier.height(8.dp))
                    Text(state.referralCode, style = MaterialTheme.typography.headlineMedium, fontWeight = FontWeight.Bold)
                    Spacer(modifier = Modifier.height(12.dp))
                    Row(horizontalArrangement = Arrangement.spacedBy(8.dp)) {
                        OutlinedButton(onClick = {
                            val clipboard = context.getSystemService(Context.CLIPBOARD_SERVICE) as ClipboardManager
                            clipboard.setPrimaryClip(ClipData.newPlainText("Referral Code", state.referralCode))
                            viewModel.markCopied()
                        }) {
                            Icon(if (state.copied) Icons.Default.Check else Icons.Default.ContentCopy, contentDescription = null)
                            Spacer(modifier = Modifier.width(4.dp))
                            Text(if (state.copied) "Copied!" else "Copy Code")
                        }
                        Button(onClick = { }) {
                            Icon(Icons.Default.Share, contentDescription = null)
                            Spacer(modifier = Modifier.width(4.dp))
                            Text("Share Link")
                        }
                    }
                }
            }
        }

        item {
            Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.spacedBy(12.dp)) {
                StatCard(modifier = Modifier.weight(1f), icon = Icons.Default.People, label = "Total Referrals", value = "${state.totalReferrals}")
                StatCard(modifier = Modifier.weight(1f), icon = Icons.Default.AttachMoney, label = "Total Earnings", value = "$${String.format("%,.2f", state.totalEarnings)}")
            }
        }

        if (state.tiers.isNotEmpty()) {
            item { Text("Reward Tiers", style = MaterialTheme.typography.titleMedium) }
            items(state.tiers) { tier -> RewardTierCard(tier = tier) }
        }

        if (state.referralHistory.isNotEmpty()) {
            item { Text("Referral History", style = MaterialTheme.typography.titleMedium) }
            items(state.referralHistory) { entry -> ReferralHistoryItem(entry = entry) }
        }

        if (state.isLoading) {
            item { Box(modifier = Modifier.fillMaxWidth().padding(32.dp), contentAlignment = Alignment.Center) { CircularProgressIndicator() } }
        }
    }
}

@Composable
private fun StatCard(modifier: Modifier = Modifier, icon: androidx.compose.ui.graphics.vector.ImageVector, label: String, value: String) {
    Card(modifier = modifier) {
        Column(modifier = Modifier.padding(16.dp), horizontalAlignment = Alignment.CenterHorizontally) {
            Icon(icon, contentDescription = null, tint = MaterialTheme.colorScheme.primary)
            Spacer(modifier = Modifier.height(8.dp))
            Text(value, style = MaterialTheme.typography.headlineSmall, fontWeight = FontWeight.Bold)
            Text(label, style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
        }
    }
}

@Composable
private fun RewardTierCard(tier: ReferralTier) {
    Card(modifier = Modifier.fillMaxWidth()) {
        Row(modifier = Modifier.padding(16.dp), verticalAlignment = Alignment.CenterVertically) {
            Icon(Icons.Default.Star, contentDescription = null, tint = MaterialTheme.colorScheme.tertiary)
            Spacer(modifier = Modifier.width(12.dp))
            Column(modifier = Modifier.weight(1f)) {
                Text(tier.name, style = MaterialTheme.typography.bodyLarge, fontWeight = FontWeight.Medium)
                Text("${tier.minReferrals}+ referrals", style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
            }
            Text("${tier.rewardPercent}%", style = MaterialTheme.typography.titleMedium, color = MaterialTheme.colorScheme.primary, fontWeight = FontWeight.Bold)
        }
    }
}

@Composable
private fun ReferralHistoryItem(entry: ReferralEntry) {
    Card(modifier = Modifier.fillMaxWidth()) {
        Row(modifier = Modifier.padding(12.dp), verticalAlignment = Alignment.CenterVertically) {
            Icon(Icons.Default.PersonAdd, contentDescription = null, tint = MaterialTheme.colorScheme.primary)
            Spacer(modifier = Modifier.width(12.dp))
            Column(modifier = Modifier.weight(1f)) {
                Text(entry.username, style = MaterialTheme.typography.bodyMedium)
                Text(entry.joinedAt, style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
            }
            Text("+$${String.format("%.2f", entry.earned)}", style = MaterialTheme.typography.bodyMedium, color = MaterialTheme.colorScheme.primary, fontWeight = FontWeight.Bold)
        }
    }
}
