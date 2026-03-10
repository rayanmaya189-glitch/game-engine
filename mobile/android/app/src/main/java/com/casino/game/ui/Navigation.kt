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
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.currentBackStackEntryAsState
import androidx.navigation.compose.rememberNavController
import com.casino.game.ui.auth.LoginScreen
import com.casino.game.ui.auth.RegisterScreen
import com.casino.game.ui.home.HomeScreen

sealed class Screen(val route: String) {
    object Login : Screen("login")
    object Register : Screen("register")
    object Home : Screen("home")
    object Games : Screen("games")
    object Tournaments : Screen("tournaments")
    object Jackpots : Screen("jackpots")
    object Wallet : Screen("wallet")
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
                onNavigateToLogin = {
                    navController.popBackStack()
                },
                onRegisterSuccess = {
                    navController.navigate(Screen.Home.route) {
                        popUpTo(Screen.Login.route) { inclusive = true }
                    }
                }
            )
        }

        composable(Screen.Home.route) {
            MainScaffold(
                navController = navController,
                bottomNavItems = bottomNavItems
            ) {
                HomeScreen(
                    onGameClick = { gameId ->
                        // Navigate to game detail
                    },
                    onSeeAllGames = {
                        navController.navigate(Screen.Games.route)
                    },
                    onJackpotClick = { jackpotId ->
                        navController.navigate(Screen.Jackpots.route)
                    },
                    onTournamentClick = { tournamentId ->
                        navController.navigate(Screen.Tournaments.route)
                    },
                    onDeposit = {
                        navController.navigate(Screen.Wallet.route)
                    },
                    onWithdraw = {
                        navController.navigate(Screen.Wallet.route)
                    }
                )
            }
        }

        composable(Screen.Games.route) {
            MainScaffold(
                navController = navController,
                bottomNavItems = bottomNavItems
            ) {
                GamesScreen(
                    onGameClick = { gameId ->
                        // Navigate to game detail
                    }
                )
            }
        }

        composable(Screen.Tournaments.route) {
            MainScaffold(
                navController = navController,
                bottomNavItems = bottomNavItems
            ) {
                TournamentsScreen(
                    onTournamentClick = { tournamentId ->
                        // Navigate to tournament detail
                    }
                )
            }
        }

        composable(Screen.Jackpots.route) {
            MainScaffold(
                navController = navController,
                bottomNavItems = bottomNavItems
            ) {
                JackpotsScreen(
                    onJackpotClick = { jackpotId ->
                        // Navigate to jackpot detail
                    }
                )
            }
        }

        composable(Screen.Wallet.route) {
            MainScaffold(
                navController = navController,
                bottomNavItems = bottomNavItems
            ) {
                WalletScreen()
            }
        }

        composable(Screen.Profile.route) {
            MainScaffold(
                navController = navController,
                bottomNavItems = bottomNavItems
            ) {
                ProfileScreen(
                    onLogout = {
                        navController.navigate(Screen.Login.route) {
                            popUpTo(0) { inclusive = true }
                        }
                    }
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
