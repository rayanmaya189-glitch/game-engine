package com.game_engine.casino.ui.home

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.AccountBalanceWallet
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import com.game_engine.casino.ui.theme.Gold

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun HomeScreen(
    onGameClick: (String) -> Unit,
    onViewAllClick: (String) -> Unit,
    viewModel: HomeViewModel = hiltViewModel()
) {
    val uiState by viewModel.uiState.collectAsState()

    Scaffold(
        topBar = {
            TopAppBar(
                title = {
                    Row(verticalAlignment = Alignment.CenterVertically) {
                        Text(
                            text = "CASINO",
                            style = MaterialTheme.typography.titleLarge,
                            color = Gold,
                            fontWeight = FontWeight.Bold
                        )
                        Text(
                            text = "GAME",
                            style = MaterialTheme.typography.titleLarge,
                            color = MaterialTheme.colorScheme.onBackground
                        )
                    }
                },
                actions = {
                    uiState.balance?.let { balance ->
                        Row(
                            modifier = Modifier
                                .background(
                                    MaterialTheme.colorScheme.surface,
                                    RoundedCornerShape(20.dp)
                                )
                                .padding(horizontal = 12.dp, vertical = 6.dp),
                            verticalAlignment = Alignment.CenterVertically
                        ) {
                            Icon(
                                imageVector = Icons.Default.AccountBalanceWallet,
                                contentDescription = null,
                                tint = Gold,
                                modifier = Modifier.size(18.dp)
                            )
                            Spacer(modifier = Modifier.width(4.dp))
                            Text(
                                text = "$${String.format("%.2f", balance.balance)}",
                                style = MaterialTheme.typography.titleSmall,
                                color = MaterialTheme.colorScheme.onSurface
                            )
                        }
                    }
                    Spacer(modifier = Modifier.width(8.dp))
                },
                colors = TopAppBarDefaults.topAppBarColors(
                    containerColor = MaterialTheme.colorScheme.background
                )
            )
        }
    ) { paddingValues ->
        LazyColumn(
            modifier = Modifier
                .fillMaxSize()
                .padding(paddingValues)
        ) {
            item {
                JackpotBanner(
                    jackpotAmount = uiState.jackpotGames.firstOrNull()?.jackpotAmount ?: 0.0
                )
            }

            item {
                CategoriesRow(
                    categories = uiState.categories.map { it.name },
                    onCategoryClick = { /* Navigate to category */ }
                )
            }

            item {
                GameSection(
                    title = "Featured Games",
                    games = uiState.featuredGames,
                    onGameClick = onGameClick,
                    onViewAllClick = { onViewAllClick("featured") }
                )
            }

            item {
                GameSection(
                    title = "Popular Games",
                    games = uiState.popularGames,
                    onGameClick = onGameClick,
                    onViewAllClick = { onViewAllClick("popular") }
                )
            }

            item {
                GameSection(
                    title = "Jackpot Games",
                    games = uiState.jackpotGames,
                    onGameClick = onGameClick,
                    onViewAllClick = { onViewAllClick("jackpot") }
                )
            }
        }
    }
}
