package com.casino.game.ui.screens

import androidx.compose.animation.AnimatedVisibility
import androidx.compose.foundation.clickable
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
import com.casino.game.data.remote.ApiService
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class FaqItem(val id: String, val question: String, val answer: String)
data class SupportTicket(val id: String, val subject: String, val category: String, val status: String, val createdAt: String, val description: String)
data class TicketsResponse(val tickets: List<SupportTicket>)
data class TicketResponse(val ticket: SupportTicket)
data class CreateTicketRequest(val category: String, val subject: String, val description: String)
data class FaqResponse(val faqs: List<FaqItem>)

data class SupportState(
    val faqs: List<FaqItem> = emptyList(),
    val tickets: List<SupportTicket> = emptyList(),
    val isLoading: Boolean = false,
    val showCreateTicket: Boolean = false,
    val ticketCategory: String = "General",
    val ticketSubject: String = "",
    val ticketDescription: String = "",
    val error: String? = null
)

@HiltViewModel
class SupportViewModel @Inject constructor(
    private val apiService: ApiService
) : ViewModel() {
    private val _state = MutableStateFlow(SupportState())
    val state: StateFlow<SupportState> = _state.asStateFlow()

    init { loadData() }

    private fun loadData() {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            try {
                val faqResp = apiService.getFaq()
                if (faqResp.isSuccessful) _state.update { it.copy(faqs = faqResp.body()?.faqs ?: emptyList()) }
                val ticketResp = apiService.getTickets()
                if (ticketResp.isSuccessful) _state.update { it.copy(tickets = ticketResp.body()?.tickets ?: emptyList()) }
            } catch (_: Exception) {}
            _state.update { it.copy(isLoading = false) }
        }
    }

    fun showCreateDialog() { _state.update { it.copy(showCreateTicket = true) } }
    fun hideCreateDialog() { _state.update { it.copy(showCreateTicket = false, ticketSubject = "", ticketDescription = "", ticketCategory = "General") } }
    fun updateCategory(c: String) { _state.update { it.copy(ticketCategory = c) } }
    fun updateSubject(s: String) { _state.update { it.copy(ticketSubject = s) } }
    fun updateDescription(d: String) { _state.update { it.copy(ticketDescription = d) } }

    fun createTicket() {
        val s = _state.value
        if (s.ticketSubject.isBlank() || s.ticketDescription.isBlank()) return
        viewModelScope.launch {
            try {
                val resp = apiService.createTicket(CreateTicketRequest(s.ticketCategory, s.ticketSubject, s.ticketDescription))
                if (resp.isSuccessful) {
                    resp.body()?.ticket?.let { t ->
                        _state.update { it.copy(tickets = listOf(t) + it.tickets, showCreateTicket = false, ticketSubject = "", ticketDescription = "") }
                    }
                }
            } catch (_: Exception) {}
        }
    }
}

@Composable
fun SupportScreen(onBack: () -> Unit, viewModel: SupportViewModel = hiltViewModel()) {
    val state by viewModel.state.collectAsState()

    Column(modifier = Modifier.fillMaxSize()) {
        TopAppBar(title = { Text("Support") }, navigationIcon = { IconButton(onClick = onBack) { Icon(Icons.Default.ArrowBack, contentDescription = "Back") } })

        if (state.isLoading) {
            Box(modifier = Modifier.fillMaxSize(), contentAlignment = Alignment.Center) { CircularProgressIndicator() }
            return@Column
        }

        LazyColumn(modifier = Modifier.fillMaxSize(), contentPadding = PaddingValues(16.dp), verticalArrangement = Arrangement.spacedBy(12.dp)) {
            item { ContactSupportSection() }
            item { Text("FAQ", style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold) }
            items(state.faqs) { faq -> FaqCard(faq) }
            item {
                Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceBetween, verticalAlignment = Alignment.CenterVertically) {
                    Text("My Tickets", style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
                    FilledTonalButton(onClick = viewModel::showCreateDialog) {
                        Icon(Icons.Default.Add, contentDescription = null, modifier = Modifier.size(18.dp))
                        Spacer(Modifier.width(4.dp)); Text("New Ticket")
                    }
                }
            }
            if (state.tickets.isEmpty()) {
                item { Text("No tickets yet", color = MaterialTheme.colorScheme.onSurfaceVariant, modifier = Modifier.padding(vertical = 16.dp)) }
            }
            items(state.tickets) { ticket -> TicketCard(ticket) }
        }
    }

    if (state.showCreateTicket) {
        CreateTicketDialog(
            category = state.ticketCategory, subject = state.ticketSubject, description = state.ticketDescription,
            categories = listOf("General", "Account", "Payment", "Game Issue", "Bonus", "Technical"),
            onCategoryChange = viewModel::updateCategory, onSubjectChange = viewModel::updateSubject,
            onDescriptionChange = viewModel::updateDescription, onSubmit = viewModel::createTicket, onDismiss = viewModel::hideCreateDialog
        )
    }
}

@Composable
private fun ContactSupportSection() {
    Card(modifier = Modifier.fillMaxWidth()) {
        Column(modifier = Modifier.padding(16.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
            Text("Contact Us", style = MaterialTheme.typography.titleMedium, fontWeight = FontWeight.Bold)
            Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.spacedBy(8.dp)) {
                FilledTonalButton(onClick = {}, modifier = Modifier.weight(1f)) {
                    Icon(Icons.Default.Chat, contentDescription = null, modifier = Modifier.size(18.dp))
                    Spacer(Modifier.width(4.dp)); Text("Live Chat")
                }
                OutlinedButton(onClick = {}, modifier = Modifier.weight(1f)) {
                    Icon(Icons.Default.Email, contentDescription = null, modifier = Modifier.size(18.dp))
                    Spacer(Modifier.width(4.dp)); Text("Email")
                }
            }
            Row(verticalAlignment = Alignment.CenterVertically) {
                Icon(Icons.Default.Phone, contentDescription = null, modifier = Modifier.size(18.dp), tint = MaterialTheme.colorScheme.primary)
                Spacer(Modifier.width(8.dp)); Text("+1-800-CASINO-1", style = MaterialTheme.typography.bodyMedium)
            }
        }
    }
}

@Composable
private fun FaqCard(faq: FaqItem) {
    var expanded by remember { mutableStateOf(false) }
    Card(modifier = Modifier.fillMaxWidth().clickable { expanded = !expanded }) {
        Column(modifier = Modifier.padding(16.dp)) {
            Row(verticalAlignment = Alignment.CenterVertically) {
                Text(faq.question, style = MaterialTheme.typography.bodyLarge, fontWeight = FontWeight.Medium, modifier = Modifier.weight(1f))
                Icon(if (expanded) Icons.Default.ExpandLess else Icons.Default.ExpandMore, contentDescription = null)
            }
            AnimatedVisibility(visible = expanded) {
                Text(faq.answer, style = MaterialTheme.typography.bodyMedium, color = MaterialTheme.colorScheme.onSurfaceVariant, modifier = Modifier.padding(top = 8.dp))
            }
        }
    }
}

@Composable
private fun TicketCard(ticket: SupportTicket) {
    val statusColor = when (ticket.status.lowercase()) {
        "open" -> MaterialTheme.colorScheme.primary
        "resolved" -> MaterialTheme.colorScheme.tertiary
        else -> MaterialTheme.colorScheme.onSurfaceVariant
    }
    Card(modifier = Modifier.fillMaxWidth()) {
        Column(modifier = Modifier.padding(16.dp)) {
            Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.SpaceBetween, verticalAlignment = Alignment.CenterVertically) {
                Text(ticket.subject, style = MaterialTheme.typography.bodyLarge, fontWeight = FontWeight.Medium)
                AssistChip(onClick = {}, label = { Text(ticket.status, style = MaterialTheme.typography.labelSmall) }, colors = AssistChipDefaults.assistChipColors(labelColor = statusColor))
            }
            Spacer(Modifier.height(4.dp))
            Text("Category: ${ticket.category}", style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
            Text(ticket.createdAt, style = MaterialTheme.typography.bodySmall, color = MaterialTheme.colorScheme.onSurfaceVariant)
        }
    }
}

@Composable
private fun CreateTicketDialog(
    category: String, subject: String, description: String, categories: List<String>,
    onCategoryChange: (String) -> Unit, onSubjectChange: (String) -> Unit,
    onDescriptionChange: (String) -> Unit, onSubmit: () -> Unit, onDismiss: () -> Unit
) {
    var showCategoryDropdown by remember { mutableStateOf(false) }
    AlertDialog(onDismissRequest = onDismiss, title = { Text("Create Ticket") }, text = {
        Column(verticalArrangement = Arrangement.spacedBy(12.dp)) {
            Box {
                OutlinedTextField(value = category, onValueChange = {}, label = { Text("Category") }, modifier = Modifier.fillMaxWidth(), readOnly = true, trailingIcon = { Icon(Icons.Default.ArrowDropDown, null) })
                DropdownMenu(expanded = showCategoryDropdown, onDismissRequest = { showCategoryDropdown = false }) {
                    categories.forEach { c -> DropdownMenuItem(text = { Text(c) }, onClick = { onCategoryChange(c); showCategoryDropdown = false }) }
                }
            }
            OutlinedTextField(value = subject, onValueChange = onSubjectChange, label = { Text("Subject") }, modifier = Modifier.fillMaxWidth())
            OutlinedTextField(value = description, onValueChange = onDescriptionChange, label = { Text("Description") }, modifier = Modifier.fillMaxWidth().height(120.dp), maxLines = 5)
        }
    }, confirmButton = { Button(onClick = onSubmit, enabled = subject.isNotBlank() && description.isNotBlank()) { Text("Submit") } }, dismissButton = { TextButton(onClick = onDismiss) { Text("Cancel") } })
}
