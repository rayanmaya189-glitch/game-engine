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
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.casino.game.data.remote.ApiService
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class GamingLimits(
    val dailyDepositLimit: Double = 1000.0,
    val weeklyDepositLimit: Double = 5000.0,
    val monthlyDepositLimit: Double = 20000.0,
    val sessionTimeLimitMinutes: Int = 120,
    val dailyLossLimit: Double = 500.0,
    val selfExclusionDays: Int = 0,
    val coolOffEnabled: Boolean = false,
    val realityCheckIntervalMinutes: Int = 30
)

data class GamingLimitsResponse(val limits: GamingLimits)
data class UpdateGamingLimitsRequest(val limits: GamingLimits)

data class ResponsibleGamingState(
    val limits: GamingLimits = GamingLimits(),
    val isLoading: Boolean = false,
    val isSaving: Boolean = false,
    val showExclusionDialog: Boolean = false,
    val saveSuccess: Boolean = false,
    val error: String? = null
)

@HiltViewModel
class ResponsibleGamingViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(ResponsibleGamingState())
    val state: StateFlow<ResponsibleGamingState> = _state.asStateFlow()

    init { loadLimits() }

    private fun loadLimits() {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            try {
                val resp = apiService.getGamingLimits()
                if (resp.isSuccessful) resp.body()?.let { r -> _state.update { it.copy(limits = r.limits, isLoading = false) } }
            } catch (_: Exception) { _state.update { it.copy(isLoading = false) } }
        }
    }

    fun updateDailyDeposit(v: Float) { _state.update { it.copy(limits = it.limits.copy(dailyDepositLimit = v.toDouble())) } }
    fun updateWeeklyDeposit(v: Float) { _state.update { it.copy(limits = it.limits.copy(weeklyDepositLimit = v.toDouble())) } }
    fun updateMonthlyDeposit(v: Float) { _state.update { it.copy(limits = it.limits.copy(monthlyDepositLimit = v.toDouble())) } }
    fun updateSessionTime(v: Float) { _state.update { it.copy(limits = it.limits.copy(sessionTimeLimitMinutes = v.toInt())) } }
    fun updateLossLimit(v: Float) { _state.update { it.copy(limits = it.limits.copy(dailyLossLimit = v.toDouble())) } }
    fun toggleCoolOff() { _state.update { it.copy(limits = it.limits.copy(coolOffEnabled = !it.limits.coolOffEnabled)) } }
    fun updateRealityCheck(v: Float) { _state.update { it.copy(limits = it.limits.copy(realityCheckIntervalMinutes = v.toInt())) } }
    fun showExclusionDialog() { _state.update { it.copy(showExclusionDialog = true) } }
    fun hideExclusionDialog() { _state.update { it.copy(showExclusionDialog = false) } }

    fun setSelfExclusion(days: Int) {
        _state.update { it.copy(limits = it.limits.copy(selfExclusionDays = days), showExclusionDialog = false) }
    }

    fun saveLimits() {
        _state.update { it.copy(isSaving = true, saveSuccess = false) }
        viewModelScope.launch {
            try {
                val resp = apiService.updateGamingLimits(UpdateGamingLimitsRequest(_state.value.limits))
                _state.update { it.copy(isSaving = false, saveSuccess = resp.isSuccessful) }
            } catch (_: Exception) { _state.update { it.copy(isSaving = false) } }
        }
    }
}

@Composable
fun ResponsibleGamingScreen(onBack: () -> Unit, viewModel: ResponsibleGamingViewModel = hiltViewModel()) {
    val state by viewModel.state.collectAsState()
    val limits = state.limits

    Column(modifier = Modifier.fillMaxSize()) {
        TopAppBar(title = { Text("Responsible Gaming") }, navigationIcon = { IconButton(onClick = onBack) { Icon(Icons.Default.ArrowBack, contentDescription = "Back") } })

        if (state.isLoading) {
            Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) { CircularProgressIndicator() }
            return@Column
        }

        Column(modifier = Modifier.fillMaxSize().verticalScroll(rememberScrollState()).padding(16.dp), verticalArrangement = Arrangement.spacedBy(16.dp)) {
            LimitCard(title = "Deposit Limits", icon = Icons.Default.AccountBalanceWallet) {
                LimitSlider("Daily Limit", limits.dailyDepositLimit.toFloat(), 100f, 50000f, "$", viewModel::updateDailyDeposit)
                LimitSlider("Weekly Limit", limits.weeklyDepositLimit.toFloat(), 500f, 100000f, "$", viewModel::updateWeeklyDeposit)
                LimitSlider("Monthly Limit", limits.monthlyDepositLimit.toFloat(), 1000f, 300000f, "$", viewModel::updateMonthlyDeposit)
            }

            LimitCard(title = "Session & Loss Limits", icon = Icons.Default.Timer) {
                LimitSlider("Session Time (min)", limits.sessionTimeLimitMinutes.toFloat(), 15f, 480f, "", viewModel::updateSessionTime, step = 15)
                LimitSlider("Daily Loss Limit", limits.dailyLossLimit.toFloat(), 50f, 25000f, "$", viewModel::updateLossLimit)
            }

            LimitCard(title = "Reality Check", icon = Icons.Default.NotificationsActive) {
                LimitSlider("Check Interval (min)", limits.realityCheckIntervalMinutes.toFloat(), 5f, 120f, "", viewModel::updateRealityCheck, step = 5)
            }

            Card(modifier = Modifier.fillMaxWidth()) {
                Column(modifier = Modifier.padding(16.dp), verticalArrangement = Arrangement.spacedBy(12.dp)) {
                    Row(verticalAlignment = Alignment.CenterVertically) {
                        Icon(Icons.Default.AcUnit, contentDescription = null, tint = MaterialTheme.colorScheme.primary)
                        Spacer(Modifier.width(12.dp))
                        Column(modifier = Modifier.weight(1f)) {
                            Text("Cool-Off Period", style = MaterialTheme.typography.bodyLarge, fontWeight = FontWeight.Medium)
                            Text("Temporarily restrict your account", style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
                        }
                        Switch(checked = limits.coolOffEnabled, onCheckedChange = { viewModel.toggleCoolOff() })
                    }
                }
            }

            Card(modifier = Modifier.fillMaxWidth()) {
                Column(modifier = Modifier.padding(16.dp), verticalArrangement = Arrangement.spacedBy(12.dp)) {
                    Row(verticalAlignment = Alignment.CenterVertically) {
                        Icon(Icons.Default.Block, contentDescription = null, tint = MaterialTheme.colorScheme.error)
                        Spacer(Modifier.width(12.dp))
                        Column(modifier = Modifier.weight(1f)) {
                            Text("Self-Exclusion", style = MaterialTheme.typography.bodyLarge, fontWeight = FontWeight.Medium)
                            val exclusionText = when (limits.selfExclusionDays) {
                                0 -> "Not active"
                                -1 -> "Permanent"
                                else -> "${limits.selfExclusionDays} days"
                            }
                            Text("Current: $exclusionText", style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
                        }
                        OutlinedButton(onClick = viewModel::showExclusionDialog) { Text("Configure") }
                    }
                }
            }

            if (state.saveSuccess) {
                Card(colors = CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.tertiaryContainer)) {
                    Row(modifier = Modifier.padding(16.dp), verticalAlignment = Alignment.CenterVertically) {
                        Icon(Icons.Default.CheckCircle, contentDescription = null, tint = MaterialTheme.colorScheme.tertiary)
                        Spacer(Modifier.width(8.dp)); Text("Limits saved successfully")
                    }
                }
            }

            Button(onClick = viewModel::saveLimits, modifier = Modifier.fillMaxWidth().height(50.dp), enabled = !state.isSaving) {
                if (state.isSaving) CircularProgressIndicator(modifier = Modifier.size(20.dp), strokeWidth = 2.dp) else Text("Save Limits")
            }

            Spacer(modifier = Modifier.height(16.dp))
        }
    }

    if (state.showExclusionDialog) {
        AlertDialog(
            onDismissRequest = viewModel::hideExclusionDialog,
            title = { Text("Self-Exclusion") },
            text = {
                Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
                    Text("Choose a self-exclusion period. This cannot be undone before the period expires.")
                    Spacer(Modifier.height(4.dp))
                    listOf(1 to "1 Day", 7 to "7 Days", 30 to "30 Days", -1 to "Permanent").forEach { (days, label) ->
                        OutlinedButton(onClick = { viewModel.setSelfExclusion(days) }, modifier = Modifier.fillMaxWidth()) { Text(label) }
                    }
                }
            },
            confirmButton = {}, dismissButton = { TextButton(onClick = viewModel::hideExclusionDialog) { Text("Cancel") } }
        )
    }
}

@Composable
private fun LimitCard(title: String, icon: androidx.compose.ui.graphics.vector.ImageVector, content: @Composable ColumnScope.() -> Unit) {
    Card(modifier = Modifier.fillMaxWidth()) {
        Column(modifier = Modifier.padding(16.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
            Row(verticalAlignment = Alignment.CenterVertically) {
                Icon(icon, contentDescription = null, tint = MaterialTheme.colorScheme.primary)
                Spacer(Modifier.width(8.dp)); Text(title, style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
            }
            HorizontalDivider(); content()
        }
    }
}

@Composable
private fun LimitSlider(label: String, value: Float, min: Float, max: Float, prefix: String, onValueChange: (Float) -> Unit, step: Int = 0) {
    Column {
        Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceBetween) {
            Text(label, style = MaterialTheme.typography.bodyMedium)
            Text(if (prefix == "$") "$${String.format("%,.0f", value)}" else "${value.toInt()}", style = MaterialTheme.typography.bodyMedium, fontWeight = FontWeight.Bold)
        }
        Slider(value = value, onValueChange = onValueChange, valueRange = min..max, steps = if (step > 0) ((max - min) / step - 1).toInt() else 0)
    }
}
