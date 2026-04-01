package com.casino.game.ui

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
    object GamePlay : Screen("game_play/{gameId}") {
        fun createRoute(gameId: String) = "game_play/$gameId"
    }
    object Tournaments : Screen("tournaments")
    object Jackpots : Screen("jackpots")
    object Wallet : Screen("wallet")
    object Deposit : Screen("deposit")
    object Withdraw : Screen("withdraw")
    object Bonuses : Screen("bonuses")
    object Profile : Screen("profile")
}

sealed class BottomNavItem(
    val route: String,
    val title: String,
    val selectedIcon: ImageVector,
    val unselectedIcon: ImageVector
) {
    object Home : BottomNavItem(
        route = Screen.Home.route,
        title = "Home",
        selectedIcon = Icons.Filled.Home,
        unselectedIcon = Icons.Outlined.Home
    )
    object Games : BottomNavItem(
        route = Screen.Games.route,
        title = "Games",
        selectedIcon = Icons.Filled.Casino,
        unselectedIcon = Icons.Outlined.Casino
    )
    object Tournaments : BottomNavItem(
        route = Screen.Tournaments.route,
        title = "Tournaments",
        selectedIcon = Icons.Filled.EmojiEvents,
        unselectedIcon = Icons.Outlined.EmojiEvents
    )
    object Jackpots : BottomNavItem(
        route = Screen.Jackpots.route,
        title = "Jackpots",
        selectedIcon = Icons.Filled.Stars,
        unselectedIcon = Icons.Outlined.Stars
    )
    object Wallet : BottomNavItem(
        route = Screen.Wallet.route,
        title = "Wallet",
        selectedIcon = Icons.Filled.AccountBalanceWallet,
        unselectedIcon = Icons.Outlined.AccountBalanceWallet
    )
    object Profile : BottomNavItem(
        route = Screen.Profile.route,
        title = "Profile",
        selectedIcon = Icons.Filled.Person,
        unselectedIcon = Icons.Outlined.Person
    )
}

@Composable
fun MainNavigation() {
    val navController = rememberNavController()
    val bottomNavItems = listOf(
        BottomNavItem.Home,
        BottomNavItem.Games,
        BottomNavItem.Tournaments,
        BottomNavItem.Jackpots,
        BottomNavItem.Wallet,
        BottomNavItem.Profile
    )

    NavHost(
        navController = navController,
        startDestination = Screen.Login.route
    ) {
        composable(Screen.Login.route) {
            LoginScreen(
                onNavigateToRegister = {
                    navController.navigate(Screen.Register.route)
                },
                onLoginSuccess = {
                    navController.navigate(Screen.Home.route) {
                        popUpTo(Screen.Login.route) { inclusive = true }
                    }
                }
            )
        }

        composable(Screen.Register.route) {
            RegisterScreen(
                onNavigateToLogin = { navController.popBackStack() },
                onRegisterSuccess = {
                    navController.navigate(Screen.Home.route) {
                        popUpTo(Screen.Login.route) { inclusive = true }
                    }
                }
            )
        }

        composable(Screen.Home.route) {
            MainScaffold(navController, bottomNavItems) {
                HomeScreen(
                    onGameClick = { gameId ->
                        navController.navigate(Screen.GamePlay.createRoute(gameId))
                    },
                    onSeeAllGames = { navController.navigate(Screen.Games.route) },
                    onJackpotClick = { navController.navigate(Screen.Jackpots.route) },
                    onTournamentClick = { navController.navigate(Screen.Tournaments.route) },
                    onDeposit = { navController.navigate(Screen.Deposit.route) },
                    onWithdraw = { navController.navigate(Screen.Withdraw.route) }
                )
            }
        }

        composable(Screen.Games.route) {
            MainScaffold(navController, bottomNavItems) {
                GamesScreen(
                    onGameClick = { gameId ->
                        navController.navigate(Screen.GamePlay.createRoute(gameId))
                    }
                )
            }
        }

        composable(
            route = Screen.GamePlay.route,
            arguments = listOf(navArgument("gameId") { type = NavType.StringType })
        ) { backStackEntry ->
            val gameId = backStackEntry.arguments?.getString("gameId") ?: ""
            GamePlayScreen(
                gameId = gameId,
                onBack = { navController.popBackStack() }
            )
        }

        composable(Screen.Tournaments.route) {
            MainScaffold(navController, bottomNavItems) {
                TournamentScreen()
            }
        }

        composable(Screen.Jackpots.route) {
            MainScaffold(navController, bottomNavItems) {
                JackpotScreen()
            }
        }

        composable(Screen.Wallet.route) {
            MainScaffold(navController, bottomNavItems) {
                WalletScreen(
                    onNavigateToDeposit = { navController.navigate(Screen.Deposit.route) },
                    onNavigateToWithdraw = { navController.navigate(Screen.Withdraw.route) }
                )
            }
        }

        composable(Screen.Deposit.route) {
            DepositScreen(
                onBack = { navController.popBackStack() },
                onDepositComplete = { navController.popBackStack() }
            )
        }

        composable(Screen.Withdraw.route) {
            WithdrawScreen(
                onBack = { navController.popBackStack() },
                onWithdrawComplete = { navController.popBackStack() }
            )
        }

        composable(Screen.Bonuses.route) {
            MainScaffold(navController, bottomNavItems) {
                BonusScreen()
            }
        }

        composable(Screen.Profile.route) {
            MainScaffold(navController, bottomNavItems) {
                ProfileScreen(
                    onLogout = {
                        navController.navigate(Screen.Login.route) {
                            popUpTo(0) { inclusive = true }
                        }
                    },
                    onNavigateToDeposit = { navController.navigate(Screen.Deposit.route) },
                    onNavigateToWithdraw = { navController.navigate(Screen.Withdraw.route) }
                )
            }
        }
    }
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun MainScaffold(
    navController: androidx.navigation.NavHostController,
    bottomNavItems: List<BottomNavItem>,
    content: @Composable () -> Unit
) {
    val navBackStackEntry by navController.currentBackStackEntryAsState()
    val currentDestination = navBackStackEntry?.destination

    Scaffold(
        bottomBar = {
            NavigationBar {
                bottomNavItems.forEach { item ->
                    val selected = currentDestination?.hierarchy?.any { it.route == item.route } == true
                    NavigationBarItem(
                        icon = {
                            Icon(
                                imageVector = if (selected) item.selectedIcon else item.unselectedIcon,
                                contentDescription = item.title
                            )
                        },
                        label = { Text(item.title) },
                        selected = selected,
                        onClick = {
                            navController.navigate(item.route) {
                                popUpTo(navController.graph.findStartDestination().id) {
                                    saveState = true
                                }
                                launchSingleTop = true
                                restoreState = true
                            }
                        }
                    )
                }
            }
        }
    ) { innerPadding ->
        androidx.compose.foundation.layout.Box(modifier = Modifier.padding(innerPadding)) {
            content()
        }
    }
}
