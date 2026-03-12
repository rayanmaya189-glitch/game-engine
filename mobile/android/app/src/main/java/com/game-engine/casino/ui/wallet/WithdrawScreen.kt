package com.game-engine.casino.ui.wallet

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import com.game-engine.casino.ui.theme.*

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun WithdrawScreen(
    onNavigateBack: () -> Unit,
    onWithdrawSuccess: () -> Unit,
    viewModel: WalletViewModel = hiltViewModel()
) {
    val uiState by viewModel.uiState.collectAsState()
    
    var amount by remember { mutableStateOf("") }
    var selectedMethod by remember { mutableStateOf<String?>(null) }
    var paymentDetails by remember { mutableStateOf("") }
    
    Scaffold(
        topBar = {
            TopAppBar(
                title = { Text("Withdraw") },
                navigationIcon = {
                    IconButton(onClick = onNavigateBack) {
                        Icon(Icons.Default.ArrowBack, contentDescription = "Back")
                    }
                },
                colors = TopAppBarDefaults.topAppBarColors(
                    containerColor = MaterialTheme.colorScheme.background
                )
            )
        }
    ) { paddingValues ->
        Column(
            modifier = Modifier
                .fillMaxSize()
                .padding(paddingValues)
                .verticalScroll(rememberScrollState())
                .padding(16.dp)
        ) {
            // Available Balance
            Card(
                modifier = Modifier.fillMaxWidth(),
                colors = CardDefaults.cardColors(containerColor = CardBackground),
                shape = RoundedCornerShape(12.dp)
            ) {
                Row(
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(16.dp),
                    horizontalArrangement = Arrangement.SpaceBetween,
                    verticalAlignment = Alignment.CenterVertically
                ) {
                    Text(
                        text = "Available Balance",
                        style = MaterialTheme.typography.bodyMedium
                    )
                    Text(
                        text = "$${String.format("%.2f", uiState.balance?.balance ?: 0.0)}",
                        style = MaterialTheme.typography.titleLarge,
                        color = Primary,
                        fontWeight = FontWeight.Bold
                    )
                }
            }
            
            Spacer(modifier = Modifier.height(24.dp))
            
            // Amount Input
            Text(
                text = "Amount",
                style = MaterialTheme.typography.titleMedium,
                fontWeight = FontWeight.Medium
            )
            Spacer(modifier = Modifier.height(8.dp))
            
            OutlinedTextField(
                value = amount,
                onValueChange = { amount = it },
                modifier = Modifier.fillMaxWidth(),
                placeholder = { Text("Enter amount") },
                leadingIcon = { Text("$", style = MaterialTheme.typography.titleLarge) },
                keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Decimal),
                singleLine = true,
                shape = RoundedCornerShape(12.dp)
            )
            
            Spacer(modifier = Modifier.height(24.dp))
            
            // Withdrawal Methods
            Text(
                text = "Withdrawal Method",
                style = MaterialTheme.typography.titleMedium,
                fontWeight = FontWeight.Medium
            )
            Spacer(modifier = Modifier.height(8.dp))
            
            uiState.paymentMethods?.withdrawMethods?.forEach { method ->
                Card(
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(vertical = 4.dp),
                    colors = CardDefaults.cardColors(
                        containerColor = if (selectedMethod == method.id) Primary.copy(alpha = 0.2f) else CardBackground
                    ),
                    shape = RoundedCornerShape(12.dp),
                    onClick = { selectedMethod = method.id }
                ) {
                    Row(
                        modifier = Modifier
                            .fillMaxWidth()
                            .padding(16.dp),
                        verticalAlignment = Alignment.CenterVertically
                    ) {
                        Icon(
                            imageVector = Icons.Default.AccountBalance,
                            contentDescription = null,
                            tint = Primary
                        )
                        Spacer(modifier = Modifier.width(12.dp))
                        Column(modifier = Modifier.weight(1f)) {
                            Text(
                                text = method.name,
                                style = MaterialTheme.typography.bodyLarge,
                                fontWeight = FontWeight.Medium
                            )
                            Text(
                                text = "Fee: ${method.feePercentage}% • ${method.processingTime}",
                                style = MaterialTheme.typography.bodySmall,
                                color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.6f)
                            )
                        }
                        if (selectedMethod == method.id) {
                            Icon(
                                imageVector = Icons.Default.CheckCircle,
                                contentDescription = null,
                                tint = Primary
                            )
                        }
                    }
                }
            }
            
            Spacer(modifier = Modifier.height(24.dp))
            
            // Payment Details
            Text(
                text = "Payment Details",
                style = MaterialTheme.typography.titleMedium,
                fontWeight = FontWeight.Medium
            )
            Spacer(modifier = Modifier.height(8.dp))
            
            OutlinedTextField(
                value = paymentDetails,
                onValueChange = { paymentDetails = it },
                modifier = Modifier.fillMaxWidth(),
                placeholder = { Text("Enter account details") },
                leadingIcon = { Icon(Icons.Default.Info, contentDescription = null) },
                minLines = 2,
                shape = RoundedCornerShape(12.dp)
            )
            
            Spacer(modifier = Modifier.height(32.dp))
            
            // Withdraw Button
            Button(
                onClick = {
                    // Handle withdrawal
                },
                modifier = Modifier
                    .fillMaxWidth()
                    .height(56.dp),
                enabled = amount.isNotEmpty() && selectedMethod != null && paymentDetails.isNotEmpty() && !uiState.isLoading,
                shape = RoundedCornerShape(12.dp),
                colors = ButtonDefaults.buttonColors(containerColor = Primary)
            ) {
                if (uiState.isLoading) {
                    CircularProgressIndicator(
                        modifier = Modifier.size(24.dp),
                        color = MaterialTheme.colorScheme.onPrimary
                    )
                } else {
                    Text("WITHDRAW $$amount", style = MaterialTheme.typography.titleMedium)
                }
            }
        }
    }
}
