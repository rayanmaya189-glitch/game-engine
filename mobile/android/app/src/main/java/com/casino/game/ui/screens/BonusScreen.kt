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
import com.casino.game.data.repository.BonusRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class BonusState(
    val bonuses: List<Bonus> = emptyList(),
    val availableBonuses: List<Bonus> = emptyList(),
    val isLoading: Boolean = false,
    val claimingId: String? = null,
    val message: String? = null
)

@HiltViewModel
class BonusViewModel @Inject constructor(
    private val bonusRepository: BonusRepository
) : ViewModel() {
    private val _state = MutableStateFlow(BonusState())
    val state: StateFlow<BonusState> = _state.asStateFlow()

    init { loadBonuses() }

    fun loadBonuses() {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            bonusRepository.getBonuses().fold(
                onSuccess = { r -> _state.update { it.copy(bonuses = r.bonuses) } },
                onFailure = { }
            )
            bonusRepository.getAvailableBonuses().fold(
                onSuccess = { r -> _state.update { it.copy(availableBonuses = r.bonuses, isLoading = false) } },
                onFailure = { _state.update { it.copy(isLoading = false) } }
            )
        }
    }

    fun claimBonus(bonusId: String) {
        _state.update { it.copy(claimingId = bonusId) }
        viewModelScope.launch {
            bonusRepository.claimBonus(bonusId).fold(
                onSuccess = { r ->
                    _state.update {
                        it.copy(claimingId = null, message = r.message)
                    }
                    loadBonuses()
                },
                onFailure = { e ->
                    _state.update {
                        it.copy(claimingId = null, message = e.message ?: "Failed to claim bonus")
                    }
                }
            )
        }
    }
}

@Composable
fun BonusScreen(viewModel: BonusViewModel = hiltViewModel()) {
    val state by viewModel.state.collectAsState()

    LazyColumn(
        modifier = Modifier.fillMaxSize(),
        contentPadding = PaddingValues(16.dp),
        verticalArrangement = Arrangement.spacedBy(12.dp)
    ) {
        item {
            Text("Bonuses", style = MaterialTheme.typography.headlineSmall)
            Spacer(modifier = Modifier.height(4.dp))
        }

        if (state.availableBonuses.isNotEmpty()) {
            item {
                Text("Available Bonuses", style = MaterialTheme.typography.titleMedium)
            }
            items(state.availableBonuses) { bonus ->
                AvailableBonusCard(
                    bonus = bonus,
                    isClaiming = state.claimingId == bonus.id,
                    onClaim = { viewModel.claimBonus(bonus.id) }
                )
            }
        }

        if (state.bonuses.isNotEmpty()) {
            item {
                Text("Active Bonuses", style = MaterialTheme.typography.titleMedium)
            }
            items(state.bonuses.filter { it.status == "active" }) { bonus ->
                ActiveBonusCard(bonus = bonus)
            }
        }

        if (state.isLoading) {
            item {
                Box(
                    modifier = Modifier.fillMaxWidth().padding(32.dp),
                    contentAlignment = Alignment.Center
                ) { CircularProgressIndicator() }
            }
        }

        if (state.availableBonuses.isEmpty() && state.bonuses.isEmpty() && !state.isLoading) {
            item {
                Box(
                    modifier = Modifier.fillMaxWidth().padding(32.dp),
                    contentAlignment = Alignment.Center
                ) {
                    Column(horizontalAlignment = Alignment.CenterHorizontally) {
                        Icon(Icons.Default.CardGiftcard, contentDescription = null, modifier = Modifier.size(64.dp), tint = MaterialTheme.colorScheme.onSurfaceVariant)
                        Spacer(modifier = Modifier.height(8.dp))
                        Text("No bonuses available", style = MaterialTheme.typography.bodyLarge)
                    }
                }
            }
        }
    }

    state.message?.let { msg ->
        LaunchedEffect(msg) { kotlinx.coroutines.delay(3000) }
    }
}

@Composable
private fun AvailableBonusCard(bonus: Bonus, isClaiming: Boolean, onClaim: () -> Unit) {
    Card(modifier = Modifier.fillMaxWidth()) {
        Column(modifier = Modifier.padding(16.dp)) {
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween,
                verticalAlignment = Alignment.CenterVertically
            ) {
                Column(modifier = Modifier.weight(1f)) {
                    Text(bonus.name, style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
                    Text(bonus.type.replaceFirstChar { it.uppercase() }, style = MaterialTheme.typography.bodySmall)
                }
                Text(
                    "${bonus.amount.toInt()}%",
                    style = MaterialTheme.typography.headlineSmall,
                    color = MaterialTheme.colorScheme.primary,
                    fontWeight = FontWeight.Bold
                )
            }
            bonus.description?.let {
                Spacer(modifier = Modifier.height(8.dp))
                Text(it, style = MaterialTheme.typography.bodySmall)
            }
            Spacer(modifier = Modifier.height(12.dp))
            Button(
                onClick = onClaim,
                enabled = !isClaiming,
                modifier = Modifier.fillMaxWidth()
            ) {
                if (isClaiming) {
                    CircularProgressIndicator(modifier = Modifier.size(20.dp), color = MaterialTheme.colorScheme.onPrimary)
                } else {
                    Icon(Icons.Default.Redeem, contentDescription = null)
                    Spacer(modifier = Modifier.width(8.dp))
                    Text("Claim Bonus")
                }
            }
        }
    }
}

@Composable
private fun ActiveBonusCard(bonus: Bonus) {
    Card(modifier = Modifier.fillMaxWidth()) {
        Column(modifier = Modifier.padding(16.dp)) {
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween
            ) {
                Text(bonus.name, style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
                AssistChip(onClick = { }, label = { Text("Active") })
            }
            bonus.wagerRequirement?.let { req ->
                Spacer(modifier = Modifier.height(8.dp))
                Text("Wager Requirement: ${req}x", style = MaterialTheme.typography.bodySmall)
                LinearProgressIndicator(
                    progress = { 0.4f },
                    modifier = Modifier.fillMaxWidth().padding(top = 4.dp)
                )
            }
            bonus.expiresAt?.let {
                Spacer(modifier = Modifier.height(4.dp))
                Text("Expires: $it", style = MaterialTheme.typography.bodySmall)
            }
        }
    }
}
