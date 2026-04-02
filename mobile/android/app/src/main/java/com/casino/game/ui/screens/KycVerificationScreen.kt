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
import com.casino.game.data.model.*
import com.casino.game.data.remote.ApiService
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class KycVerificationState(
    val currentLevel: Int = 0,
    val documents: List<KycDocument> = emptyList(),
    val isSubmitting: Boolean = false,
    val error: String? = null,
    val successMessage: String? = null
)

@HiltViewModel
class KycVerificationViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(KycVerificationState())
    val state: StateFlow<KycVerificationState> = _state.asStateFlow()

    init { loadKycStatus() }

    fun loadKycStatus() {
        viewModelScope.launch {
            try {
                val response = apiService.getKycStatus()
                if (response.isSuccessful) {
                    val body = response.body()
                    _state.update { it.copy(currentLevel = body?.level ?: 0, documents = body?.documents ?: emptyList()) }
                }
            } catch (_: Exception) { }
        }
    }

    fun uploadDocument(type: String) {
        _state.update { it.copy(isSubmitting = true) }
        viewModelScope.launch {
            try {
                val response = apiService.uploadKycDocument(mapOf("type" to type))
                if (response.isSuccessful) {
                    _state.update { it.copy(isSubmitting = false, successMessage = "$type uploaded successfully") }
                    loadKycStatus()
                } else {
                    _state.update { it.copy(isSubmitting = false, error = "Upload failed") }
                }
            } catch (e: Exception) {
                _state.update { it.copy(isSubmitting = false, error = e.message) }
            }
        }
    }
}

@Composable
fun KycVerificationScreen(
    onBack: () -> Unit,
    viewModel: KycVerificationViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()
    val scrollState = rememberScrollState()

    Column(modifier = Modifier.fillMaxSize()) {
        TopAppBar(
            title = { Text("KYC Verification") },
            navigationIcon = { IconButton(onClick = onBack) { Icon(Icons.Default.ArrowBack, contentDescription = "Back") } }
        )

        Column(
            modifier = Modifier.fillMaxSize().verticalScroll(scrollState).padding(16.dp),
            verticalArrangement = Arrangement.spacedBy(16.dp)
        ) {
            Card(modifier = Modifier.fillMaxWidth()) {
                Column(modifier = Modifier.padding(16.dp)) {
                    Text("Verification Level", style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
                    Spacer(modifier = Modifier.height(12.dp))
                    Text("Level ${state.currentLevel} of 3", style = MaterialTheme.typography.bodyMedium)
                    Spacer(modifier = Modifier.height(8.dp))
                    LinearProgressIndicator(
                        progress = { state.currentLevel / 3f },
                        modifier = Modifier.fillMaxWidth().height(8.dp)
                    )
                    Spacer(modifier = Modifier.height(8.dp))
                    Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceBetween) {
                        (0..3).forEach { lvl ->
                            Icon(
                                imageVector = if (lvl <= state.currentLevel) Icons.Default.CheckCircle else Icons.Default.RadioButtonUnchecked,
                                contentDescription = "Level $lvl",
                                tint = if (lvl <= state.currentLevel) MaterialTheme.colorScheme.primary else MaterialTheme.colorScheme.onSurfaceVariant,
                                modifier = Modifier.size(24.dp)
                            )
                        }
                    }
                }
            }

            Text("Documents", style = MaterialTheme.typography.titleMedium)

            val docTypes = listOf(
                Triple("id", "Identity Document (Passport/ID)", Icons.Default.Badge),
                Triple("address", "Proof of Address", Icons.Default.Home),
                Triple("selfie", "Selfie Verification", Icons.Default.Face)
            )

            docTypes.forEach { (type, label, icon) ->
                val doc = state.documents.find { it.type == type }
                DocumentUploadCard(
                    type = type, label = label, icon = icon, document = doc,
                    isSubmitting = state.isSubmitting,
                    onUpload = { viewModel.uploadDocument(type) }
                )
            }

            state.error?.let { error ->
                Card(colors = CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.errorContainer)) {
                    Text(error, modifier = Modifier.padding(16.dp), color = MaterialTheme.colorScheme.onErrorContainer)
                }
            }

            state.successMessage?.let { msg ->
                Card(colors = CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.primaryContainer)) {
                    Text(msg, modifier = Modifier.padding(16.dp), color = MaterialTheme.colorScheme.onPrimaryContainer)
                }
            }
        }
    }
}

@Composable
private fun DocumentUploadCard(
    type: String, label: String, icon: androidx.compose.ui.graphics.vector.ImageVector,
    document: KycDocument?,
    isSubmitting: Boolean,
    onUpload: () -> Unit
) {
    val statusColor = when (document?.status) {
        "approved" -> MaterialTheme.colorScheme.primary
        "rejected" -> MaterialTheme.colorScheme.error
        "pending" -> MaterialTheme.colorScheme.tertiary
        else -> MaterialTheme.colorScheme.onSurfaceVariant
    }

    Card(modifier = Modifier.fillMaxWidth()) {
        Column(modifier = Modifier.padding(16.dp)) {
            Row(verticalAlignment = Alignment.CenterVertically) {
                Icon(icon, contentDescription = null, tint = statusColor)
                Spacer(modifier = Modifier.width(12.dp))
                Column(modifier = Modifier.weight(1f)) {
                    Text(label, style = MaterialTheme.typography.bodyLarge, fontWeight = FontWeight.Medium)
                    document?.let { doc ->
                        Text(
                            doc.status.replaceFirstChar { it.uppercase() },
                            style = MaterialTheme.typography.bodySmall,
                            color = statusColor
                        )
                        doc.rejectionReason?.let { reason ->
                            Text("Reason: $reason", style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.error)
                        }
                    } ?: run {
                        Text("Not uploaded", style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
                    }
                }
            }
            Spacer(modifier = Modifier.height(12.dp))
            Button(
                onClick = onUpload,
                enabled = document?.status != "approved" && !isSubmitting,
                modifier = Modifier.fillMaxWidth()
            ) {
                if (isSubmitting) {
                    CircularProgressIndicator(modifier = Modifier.size(20.dp), color = MaterialTheme.colorScheme.onPrimary)
                } else {
                    Icon(Icons.Default.Upload, contentDescription = null)
                    Spacer(modifier = Modifier.width(8.dp))
                    Text(when (document?.status) {
                        "approved" -> "Verified"
                        "pending" -> "Re-upload"
                        else -> "Upload"
                    })
                }
            }
        }
    }
}
