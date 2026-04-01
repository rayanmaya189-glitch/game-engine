package com.casino.game.ui.home

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun HomeScreen(
    onGameClick: (String) -> Unit,
    onSeeAllGames: () -> Unit,
    onJackpotClick: (String) -> Unit,
    onTournamentClick: (String) -> Unit,
    onDeposit: () -> Unit,
    onWithdraw: () -> Unit,
    viewModel: HomeViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()

    Box(modifier = Modifier.fillMaxSize()) {
        Column(
            modifier = Modifier
                .fillMaxSize()
                .verticalScroll(rememberScrollState())
        ) {
            state.balance?.let { balance ->
                BalanceCard(
                    balance = balance.balance,
                    bonusBalance = balance.bonusBalance,
                    onDeposit = onDeposit,
                    onWithdraw = onWithdraw
                )
            }

            if (state.currentJackpots.isNotEmpty()) {
                JackpotsSection(
                    jackpots = state.currentJackpots,
                    onJackpotClick = onJackpotClick
                )
            }

            if (state.activeTournaments.isNotEmpty()) {
                TournamentsSection(
                    tournaments = state.activeTournaments,
                    onTournamentClick = onTournamentClick
                )
            }

            if (state.featuredGames.isNotEmpty()) {
                GamesSection(
                    title = "Featured Games",
                    games = state.featuredGames,
                    onGameClick = onGameClick,
                    onSeeAll = onSeeAllGames
                )
            }

            if (state.popularGames.isNotEmpty()) {
                GamesSection(
                    title = "Popular Games",
                    games = state.popularGames,
                    onGameClick = onGameClick,
                    onSeeAll = onSeeAllGames
                )
            }

            Spacer(modifier = Modifier.height(80.dp))
        }

        if (state.isLoading) {
            CircularProgressIndicator(
                modifier = Modifier.align(Alignment.Center)
            )
        }
    }
}
