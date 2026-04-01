package com.game_engine.casino.ui.tournament

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import java.text.NumberFormat
import java.time.format.DateTimeFormatter
import java.util.Locale

@Composable
fun TournamentCard(
    tournament: Tournament,
    onClick: () -> Unit
) {
    val currencyFormat = remember { NumberFormat.getCurrencyInstance(Locale.US) }
    val dateFormatter = remember { DateTimeFormatter.ofPattern("MMM dd, HH:mm") }

    Card(
        modifier = Modifier
            .fillMaxWidth()
            .clickable(onClick = onClick),
        shape = RoundedCornerShape(12.dp)
    ) {
        Column(
            modifier = Modifier.padding(16.dp)
        ) {
            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween,
                verticalAlignment = Alignment.CenterVertically
            ) {
                Text(
                    text = tournament.name,
                    style = MaterialTheme.typography.titleMedium,
                    fontWeight = FontWeight.Bold
                )
                StatusChip(tournament.status)
            }

            Spacer(modifier = Modifier.height(8.dp))

            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.spacedBy(16.dp)
            ) {
                InfoItem(
                    icon = Icons.Default.SportsPoker,
                    label = "Game",
                    value = tournament.gameType
                )
                InfoItem(
                    icon = Icons.Default.AttachMoney,
                    label = "Buy-in",
                    value = currencyFormat.format(tournament.buyIn)
                )
            }

            Spacer(modifier = Modifier.height(8.dp))

            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.spacedBy(16.dp)
            ) {
                InfoItem(
                    icon = Icons.Default.EmojiEvents,
                    label = "Prize Pool",
                    value = currencyFormat.format(tournament.prizePool)
                )
                InfoItem(
                    icon = Icons.Default.People,
                    label = "Players",
                    value = "${tournament.registeredPlayers}/${tournament.maxPlayers}"
                )
            }

            Spacer(modifier = Modifier.height(12.dp))

            Row(
                modifier = Modifier.fillMaxWidth(),
                horizontalArrangement = Arrangement.SpaceBetween,
                verticalAlignment = Alignment.CenterVertically
            ) {
                Text(
                    text = "${tournament.startTime.format(dateFormatter)} - ${tournament.endTime.format(dateFormatter)}",
                    style = MaterialTheme.typography.bodySmall,
                    color = MaterialTheme.colorScheme.onSurfaceVariant
                )

                if (tournament.isRegistered) {
                    Button(
                        onClick = { /* Play */ },
                        colors = ButtonDefaults.buttonColors(
                            containerColor = MaterialTheme.colorScheme.primary
                        )
                    ) {
                        Text("Play Now")
                    }
                } else if (tournament.status == TournamentStatus.REGISTRATION_OPEN) {
                    Button(onClick = { /* Register */ }) {
                        Text("Register")
                    }
                }
            }
        }
    }
}

@Composable
fun StatusChip(status: TournamentStatus) {
    val (color, text) = when (status) {
        TournamentStatus.UPCOMING -> MaterialTheme.colorScheme.secondary to "Upcoming"
        TournamentStatus.REGISTRATION_OPEN -> MaterialTheme.colorScheme.primary to "Register"
        TournamentStatus.REGISTRATION_CLOSED -> MaterialTheme.colorScheme.tertiary to "Closed"
        TournamentStatus.IN_PROGRESS -> MaterialTheme.colorScheme.primary to "Live"
        TournamentStatus.COMPLETED -> MaterialTheme.colorScheme.surfaceVariant to "Completed"
        TournamentStatus.CANCELLED -> MaterialTheme.colorScheme.error to "Cancelled"
    }

    Surface(
        shape = RoundedCornerShape(16.dp),
        color = color.copy(alpha = 0.1f)
    ) {
        Text(
            text = text,
            modifier = Modifier.padding(horizontal = 12.dp, vertical = 4.dp),
            style = MaterialTheme.typography.labelSmall,
            color = color
        )
    }
}

@Composable
fun InfoItem(
    icon: ImageVector,
    label: String,
    value: String
) {
    Row(
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(4.dp)
    ) {
        Icon(
            imageVector = icon,
            contentDescription = null,
            modifier = Modifier.size(16.dp),
            tint = MaterialTheme.colorScheme.onSurfaceVariant
        )
        Column {
            Text(
                text = label,
                style = MaterialTheme.typography.labelSmall,
                color = MaterialTheme.colorScheme.onSurfaceVariant
            )
            Text(
                text = value,
                style = MaterialTheme.typography.bodyMedium,
                fontWeight = FontWeight.Medium
            )
        }
    }
}
