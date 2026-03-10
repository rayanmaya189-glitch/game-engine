package com.gameengine.casino.ui.wallet

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import com.gameengine.casino.data.model.Transaction
import com.gameengine.casino.ui.theme.*

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun WalletScreen(
    onNavigateToDeposit: () -> Unit,
    onNavigateToWithdraw: () -> Unit,
    onTransactionClick: (String) -> Unit,
    viewModel: WalletViewModel = hiltViewModel()
) {
    val uiState by viewModel.uiState.collectAsState()
    
    Column(
        modifier = Modifier
            .fillMaxSize()
            .background(MaterialTheme.colorScheme.background)
    ) {
        // Balance Card
        Box(
            modifier = Modifier
                .fillMaxWidth()
                .background(
                    brush = Brush.horizontalGradient(
                        colors = listOf(Primary, Gold)
                    )
                )
                .padding(24.dp)
        ) {
            Column {
                Text(
                    text = "Total Balance",
                    style = MaterialTheme.typography.bodyMedium,
                    color = MaterialTheme.colorScheme.onPrimary.copy(alpha = 0.9f)
                )
                Spacer(modifier = Modifier.height(8.dp))
                Text(
                    text = "$${String.format("%.2f", uiState.balance?.balance ?: 0.0)}",
                    style = MaterialTheme.typography.displaySmall,
                    color = MaterialTheme.colorScheme.onPrimary,
                    fontWeight = FontWeight.Bold
                )
                
                Spacer(modifier = Modifier.height(16.dp))
                
                Row(
                    modifier = Modifier.fillMaxWidth(),
                    horizontalArrangement = Arrangement.spacedBy(12.dp)
                ) {
                    // Deposit Button
                    Button(
                        onClick = onNavigateToDeposit,
                        modifier = Modifier
                            .weight(1f)
                            .height(48.dp),
                        shape = RoundedCornerShape(12.dp),
                        colors = ButtonDefaults.buttonColors(
                            containerColor = MaterialTheme.colorScheme.onPrimary.copy(alpha = 0.2f)
                        )
                    ) {
                        Icon(
                            imageVector = Icons.Default.Add,
                            contentDescription = null,
                            tint = MaterialTheme.colorScheme.onPrimary
                        )
                        Spacer(modifier = Modifier.width(8.dp))
                        Text(
                            text = "Deposit",
                            color = MaterialTheme.colorScheme.onPrimary
                        )
                    }
                    
                    // Withdraw Button
                    OutlinedButton(
                        onClick = onNavigateToWithdraw,
                        modifier = Modifier
                            .weight(1f)
                            .height(48.dp),
                        shape = RoundedCornerShape(12.dp)
                    ) {
                        Icon(
                            imageVector = Icons.Default.Remove,
                            contentDescription = null,
                            tint = MaterialTheme.colorScheme.onPrimary
                        )
                        Spacer(modifier = Modifier.width(8.dp))
                        Text(
                            text = "Withdraw",
                            color = MaterialTheme.colorScheme.onPrimary
                        )
                    }
                }
            }
        }
        
        // Bonus Balance
        uiState.balance?.let { balance ->
            if (balance.bonusBalance > 0) {
                Card(
                    modifier = Modifier
                        .fillMaxWidth()
                        .padding(16.dp),
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
                        Column {
                            Text(
                                text = "Bonus Balance",
                                style = MaterialTheme.typography.bodyMedium,
                                color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.7f)
                            )
                            Text(
                                text = "$${String.format("%.2f", balance.bonusBalance)}",
                                style = MaterialTheme.typography.titleMedium,
                                color = Gold,
                                fontWeight = FontWeight.Bold
                            )
                        }
                        Icon(
                            imageVector = Icons.Default.CardGiftcard,
                            contentDescription = null,
                            tint = Gold,
                            modifier = Modifier.size(32.dp)
                        )
                    }
                }
            }
        }
        
        // Transaction History
        Text(
            text = "Recent Transactions",
            style = MaterialTheme.typography.titleMedium,
            modifier = Modifier.padding(horizontal = 16.dp, vertical = 8.dp)
        )
        
        if (uiState.transactions.isEmpty()) {
            Box(
                modifier = Modifier
                    .fillMaxWidth()
                    .padding(32.dp),
                contentAlignment = Alignment.Center
            ) {
                Column(horizontalAlignment = Alignment.CenterHorizontally) {
                    Icon(
                        imageVector = Icons.Default.Receipt,
                        contentDescription = null,
                        modifier = Modifier.size(48.dp),
                        tint = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.5f)
                    )
                    Spacer(modifier = Modifier.height(8.dp))
                    Text(
                        text = "No transactions yet",
                        color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.7f)
                    )
                }
            }
        } else {
            LazyColumn(
                contentPadding = PaddingValues(horizontal = 16.dp),
                verticalArrangement = Arrangement.spacedBy(8.dp)
            ) {
                items(uiState.transactions) { transaction ->
                    TransactionItem(
                        transaction = transaction,
                        onClick = { onTransactionClick(transaction.id) }
                    )
                }
            }
        }
    }
}

@Composable
fun TransactionItem(
    transaction: Transaction,
    onClick: () -> Unit
) {
    Card(
        modifier = Modifier
            .fillMaxWidth()
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
            // Icon
            val (icon, color) = when (transaction.type.lowercase()) {
                "deposit" -> Icons.Default.ArrowDownward to Success
                "withdraw" -> Icons.Default.ArrowUpward to Primary
                "bet" -> Icons.Default.Casino to Error
                "win" -> Icons.Default.EmojiEvents to Gold
                else -> Icons.Default.SwapHoriz to MaterialTheme.colorScheme.onBackground
            }
            
            Box(
                modifier = Modifier
                    .size(40.dp)
                    .background(color.copy(alpha = 0.2f), RoundedCornerShape(8.dp)),
                contentAlignment = Alignment.Center
            ) {
                Icon(
                    imageVector = icon,
                    contentDescription = null,
                    tint = color,
                    modifier = Modifier.size(20.dp)
                )
            }
            
            Spacer(modifier = Modifier.width(12.dp))
            
            Column(modifier = Modifier.weight(1f)) {
                Text(
                    text = transaction.type.replaceFirstChar { it.uppercase() },
                    style = MaterialTheme.typography.bodyMedium,
                    fontWeight = FontWeight.Medium
                )
                Text(
                    text = transaction.createdAt.substring(0, 10),
                    style = MaterialTheme.typography.bodySmall,
                    color = MaterialTheme.colorScheme.onBackground.copy(alpha = 0.6f)
                )
            }
            
            Column(horizontalAlignment = Alignment.End) {
                Text(
                    text = "${if (transaction.type.lowercase() in listOf("deposit", "win")) "+" else "-"}$${String.format("%.2f", transaction.amount)}",
                    style = MaterialTheme.typography.bodyMedium,
                    fontWeight = FontWeight.Bold,
                    color = when (transaction.type.lowercase()) {
                        "deposit", "win" -> Success
                        "withdraw" -> Primary
                        else -> MaterialTheme.colorScheme.onBackground
                    }
                )
                Text(
                    text = transaction.status.replaceFirstChar { it.uppercase() },
                    style = MaterialTheme.typography.bodySmall,
                    color = when (transaction.status.lowercase()) {
                        "completed" -> Success
                        "pending" -> Warning
                        "failed" -> Error
                        else -> MaterialTheme.colorScheme.onBackground.copy(alpha = 0.6f)
                    }
                )
            }
        }
    }
}
