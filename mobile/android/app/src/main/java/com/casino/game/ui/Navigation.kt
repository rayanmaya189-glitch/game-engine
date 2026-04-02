package com.casino.game.ui

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.padding
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material.icons.outlined.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.navigation.NavDestination.Companion.hierarchy
import androidx.navigation.NavGraph.Companion.findStartDestination
import androidx.navigation.NavType
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.currentBackStackEntryAsState
import androidx.navigation.compose.rememberNavController
import androidx.navigation.navArgument
import com.casino.game.ui.auth.LoginScreen
import com.casino.game.ui.auth.RegisterScreen
import com.casino.game.ui.home.HomeScreen
import com.casino.game.ui.screens.*

sealed class Screen(val route: String) {
    object Login : Screen("login")
    object Register : Screen("register")
    object Home : Screen("home")
    object Games : Screen("games")
    object GamePlay : Screen("game_play/{gameId}") { fun createRoute(gameId: String) = "game_play/$gameId" }
    object Tournaments : Screen("tournaments")
    object Jackpots : Screen("jackpots")
    object Wallet : Screen("wallet")
    object Deposit : Screen("deposit")
    object Withdraw : Screen("withdraw")
    object Bonuses : Screen("bonuses")
    object Profile : Screen("profile")
    object Chat : Screen("chat")
    object Notifications : Screen("notifications")
    object Settings : Screen("settings")
    object Referral : Screen("referral")
    object LiveDealer : Screen("live_dealer")
    object Blackjack : Screen("blackjack")
    object Poker : Screen("poker")
    object PaymentHistory : Screen("payment_history")
    object KycVerification : Screen("kyc_verification")
    object GameDetail : Screen("game_detail/{gameId}") { fun createRoute(gameId: String) = "game_detail/$gameId" }
    object BetHistory : Screen("bet_history")
    object Support : Screen("support")
    object Leaderboard : Screen("leaderboard")
    object ResponsibleGaming : Screen("responsible_gaming")
}

sealed class BottomNavItem(val route: String, val title: String, val selectedIcon: ImageVector, val unselectedIcon: ImageVector) {
    object Home : BottomNavItem(Screen.Home.route, "Home", Icons.Filled.Home, Icons.Outlined.Home)
    object Games : BottomNavItem(Screen.Games.route, "Games", Icons.Filled.Casino, Icons.Outlined.Casino)
    object Tournaments : BottomNavItem(Screen.Tournaments.route, "Tournaments", Icons.Filled.EmojiEvents, Icons.Outlined.EmojiEvents)
    object Jackpots : BottomNavItem(Screen.Jackpots.route, "Jackpots", Icons.Filled.Stars, Icons.Outlined.Stars)
    object Wallet : BottomNavItem(Screen.Wallet.route, "Wallet", Icons.Filled.AccountBalanceWallet, Icons.Outlined.AccountBalanceWallet)
    object Profile : BottomNavItem(Screen.Profile.route, "Profile", Icons.Filled.Person, Icons.Outlined.Person)
}

@Composable
fun MainNavigation() {
    val navController = rememberNavController()
    val bottomNavItems = listOf(BottomNavItem.Home, BottomNavItem.Games, BottomNavItem.Tournaments, BottomNavItem.Jackpots, BottomNavItem.Wallet, BottomNavItem.Profile)

    NavHost(navController = navController, startDestination = Screen.Login.route) {
        composable(Screen.Login.route) {
            LoginScreen(onNavigateToRegister = { navController.navigate(Screen.Register.route) }, onLoginSuccess = {
                navController.navigate(Screen.Home.route) { popUpTo(Screen.Login.route) { inclusive = true } }
            })
        }
        composable(Screen.Register.route) {
            RegisterScreen(onNavigateToLogin = { navController.popBackStack() }, onRegisterSuccess = {
                navController.navigate(Screen.Home.route) { popUpTo(Screen.Login.route) { inclusive = true } }
            })
        }
        composable(Screen.Home.route) {
            MainScaffold(navController, bottomNavItems) {
                HomeScreen(onGameClick = { navController.navigate(Screen.GamePlay.createRoute(it)) }, onSeeAllGames = { navController.navigate(Screen.Games.route) }, onJackpotClick = { navController.navigate(Screen.Jackpots.route) }, onTournamentClick = { navController.navigate(Screen.Tournaments.route) }, onDeposit = { navController.navigate(Screen.Deposit.route) }, onWithdraw = { navController.navigate(Screen.Withdraw.route) })
            }
        }
        composable(Screen.Games.route) { MainScaffold(navController, bottomNavItems) { GamesScreen(onGameClick = { navController.navigate(Screen.GamePlay.createRoute(it)) }) } }
        composable(route = Screen.GamePlay.route, arguments = listOf(navArgument("gameId") { type = NavType.StringType })) { backStackEntry ->
            GamePlayScreen(gameId = backStackEntry.arguments?.getString("gameId") ?: "", onBack = { navController.popBackStack() })
        }
        composable(Screen.Tournaments.route) { MainScaffold(navController, bottomNavItems) { TournamentScreen() } }
        composable(Screen.Jackpots.route) { MainScaffold(navController, bottomNavItems) { JackpotScreen() } }
        composable(Screen.Wallet.route) { MainScaffold(navController, bottomNavItems) { WalletScreen(onNavigateToDeposit = { navController.navigate(Screen.Deposit.route) }, onNavigateToWithdraw = { navController.navigate(Screen.Withdraw.route) }) } }
        composable(Screen.Deposit.route) { DepositScreen(onBack = { navController.popBackStack() }, onDepositComplete = { navController.popBackStack() }) }
        composable(Screen.Withdraw.route) { WithdrawScreen(onBack = { navController.popBackStack() }, onWithdrawComplete = { navController.popBackStack() }) }
        composable(Screen.Bonuses.route) { MainScaffold(navController, bottomNavItems) { BonusScreen() } }
        composable(Screen.Profile.route) {
            MainScaffold(navController, bottomNavItems) {
                ProfileScreen(onLogout = { navController.navigate(Screen.Login.route) { popUpTo(0) { inclusive = true } } }, onNavigateToDeposit = { navController.navigate(Screen.Deposit.route) }, onNavigateToWithdraw = { navController.navigate(Screen.Withdraw.route) })
            }
        }
        composable(Screen.Chat.route) { MainScaffold(navController, bottomNavItems) { ChatScreen(onBack = { navController.popBackStack() }) } }
        composable(Screen.Notifications.route) { MainScaffold(navController, bottomNavItems) { NotificationScreen(onBack = { navController.popBackStack() }) } }
        composable(Screen.Settings.route) {
            MainScaffold(navController, bottomNavItems) {
                SettingsScreen(onBack = { navController.popBackStack() }, onLogout = { navController.navigate(Screen.Login.route) { popUpTo(0) { inclusive = true } } })
            }
        }
        composable(Screen.Referral.route) { MainScaffold(navController, bottomNavItems) { ReferralScreen(onBack = { navController.popBackStack() }) } }
        composable(Screen.LiveDealer.route) { MainScaffold(navController, bottomNavItems) { LiveDealerScreen(onBack = { navController.popBackStack() }) } }
        composable(Screen.Blackjack.route) { MainScaffold(navController, bottomNavItems) { BlackjackScreen(onBack = { navController.popBackStack() }) } }
        composable(Screen.Poker.route) { MainScaffold(navController, bottomNavItems) { PokerScreen(onBack = { navController.popBackStack() }) } }
        composable(Screen.PaymentHistory.route) { MainScaffold(navController, bottomNavItems) { PaymentHistoryScreen(onBack = { navController.popBackStack() }) } }
        composable(Screen.KycVerification.route) { MainScaffold(navController, bottomNavItems) { KycVerificationScreen(onBack = { navController.popBackStack() }) } }
        composable(route = Screen.GameDetail.route, arguments = listOf(navArgument("gameId") { type = NavType.StringType })) { backStackEntry ->
            GameDetailScreen(
                gameId = backStackEntry.arguments?.getString("gameId") ?: "",
                onBack = { navController.popBackStack() },
                onPlayGame = { navController.navigate(Screen.GamePlay.createRoute(it)) },
                onGameClick = { navController.navigate(Screen.GameDetail.createRoute(it)) }
            )
        }
        composable(Screen.BetHistory.route) { MainScaffold(navController, bottomNavItems) { BetHistoryScreen(onBack = { navController.popBackStack() }) } }
        composable(Screen.Support.route) { MainScaffold(navController, bottomNavItems) { SupportScreen(onBack = { navController.popBackStack() }) } }
        composable(Screen.Leaderboard.route) { MainScaffold(navController, bottomNavItems) { LeaderboardScreen(onBack = { navController.popBackStack() }) } }
        composable(Screen.ResponsibleGaming.route) { MainScaffold(navController, bottomNavItems) { ResponsibleGamingScreen(onBack = { navController.popBackStack() }) } }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun MainScaffold(navController: androidx.navigation.NavHostController, bottomNavItems: List<BottomNavItem>, content: @Composable () -> Unit) {
    val navBackStackEntry by navController.currentBackStackEntryAsState()
    val currentDestination = navBackStackEntry?.destination
    Scaffold(bottomBar = {
        NavigationBar { bottomNavItems.forEach { item ->
            val selected = currentDestination?.hierarchy?.any { it.route == item.route } == true
            NavigationBarItem(icon = { Icon(imageVector = if (selected) item.selectedIcon else item.unselectedIcon, contentDescription = item.title) }, label = { Text(item.title) }, selected = selected, onClick = {
                navController.navigate(item.route) { popUpTo(navController.graph.findStartDestination().id) { saveState = true }; launchSingleTop = true; restoreState = true }
            })
        }}
    }) { innerPadding -> Box(modifier = Modifier.padding(innerPadding)) { content() } }
}
