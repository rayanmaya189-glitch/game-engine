package com.gameengine.casino.ui.navigation

import androidx.compose.foundation.layout.padding
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Modifier
import androidx.hilt.navigation.compose.hiltViewModel
import androidx.navigation.NavDestination.Companion.hierarchy
import androidx.navigation.NavGraph.Companion.findStartDestination
import androidx.navigation.NavHostController
import androidx.navigation.NavType
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.currentBackStackEntryAsState
import androidx.navigation.compose.rememberNavController
import androidx.navigation.navArgument
import com.gameengine.casino.ui.auth.LoginScreen
import com.gameengine.casino.ui.auth.RegisterScreen
import com.gameengine.casino.ui.auth.ForgotPasswordScreen
import com.gameengine.casino.ui.home.HomeScreen
import com.gameengine.casino.ui.games.GamesScreen
import com.gameengine.casino.ui.games.GameDetailScreen
import com.gameengine.casino.ui.wallet.WalletScreen
import com.gameengine.casino.ui.wallet.DepositScreen
import com.gameengine.casino.ui.wallet.WithdrawScreen
import com.gameengine.casino.ui.profile.ProfileScreen
import com.gameengine.casino.ui.profile.EditProfileScreen

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun CasinoNavHost(
    navController: NavHostController = rememberNavController()
) {
    val navBackStackEntry by navController.currentBackStackEntryAsState()
    val currentDestination = navBackStackEntry?.destination
    
    val bottomNavItems = listOf(
        BottomNavItem.Home,
        BottomNavItem.Games,
        BottomNavItem.Wallet,
        BottomNavItem.Profile
    )
    
    val showBottomBar = currentDestination?.route in bottomNavItems.map { it.route }
    
    Scaffold(
        bottomBar = {
            if (showBottomBar) {
                NavigationBar {
                    bottomNavItems.forEach { item ->
                        NavigationBarItem(
                            icon = {
                                Icon(
                                    imageVector = when (item) {
                                        BottomNavItem.Home -> Icons.Default.Home
                                        BottomNavItem.Games -> Icons.Default.Casino
                                        BottomNavItem.Wallet -> Icons.Default.Wallet
                                        BottomNavItem.Profile -> Icons.Default.Person
                                    },
                                    contentDescription = item.title
                                )
                            },
                            label = { Text(item.title) },
                            selected = currentDestination?.hierarchy?.any { it.route == item.route } == true,
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
        }
    ) { innerPadding ->
        NavHost(
            navController = navController,
            startDestination = Screen.Login.route,
            modifier = Modifier.padding(innerPadding)
        ) {
            composable(Screen.Login.route) {
                LoginScreen(
                    onNavigateToRegister = {
                        navController.navigate(Screen.Register.route)
                    },
                    onNavigateToForgotPassword = {
                        navController.navigate(Screen.ForgotPassword.route)
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
            
            composable(Screen.ForgotPassword.route) {
                ForgotPasswordScreen(
                    onNavigateBack = {
                        navController.popBackStack()
                    }
                )
            }
            
            composable(Screen.Home.route) {
                HomeScreen(
                    onGameClick = { gameId ->
                        navController.navigate(Screen.GameDetail.createRoute(gameId))
                    },
                    onViewAllClick = { category ->
                        navController.navigate(Screen.Games.route)
                    }
                )
            }
            
            composable(Screen.Games.route) {
                GamesScreen(
                    onGameClick = { gameId ->
                        navController.navigate(Screen.GameDetail.createRoute(gameId))
                    }
                )
            }
            
            composable(
                route = Screen.GameDetail.route,
                arguments = listOf(navArgument("gameId") { type = NavType.StringType })
            ) { backStackEntry ->
                val gameId = backStackEntry.arguments?.getString("gameId") ?: ""
                GameDetailScreen(
                    gameId = gameId,
                    onNavigateBack = {
                        navController.popBackStack()
                    }
                )
            }
            
            composable(Screen.Wallet.route) {
                WalletScreen(
                    onNavigateToDeposit = {
                        navController.navigate(Screen.Deposit.route)
                    },
                    onNavigateToWithdraw = {
                        navController.navigate(Screen.Withdraw.route)
                    },
                    onTransactionClick = { transactionId ->
                        // Navigate to transaction detail
                    }
                )
            }
            
            composable(Screen.Deposit.route) {
                DepositScreen(
                    onNavigateBack = {
                        navController.popBackStack()
                    },
                    onDepositSuccess = {
                        navController.popBackStack()
                    }
                )
            }
            
            composable(Screen.Withdraw.route) {
                WithdrawScreen(
                    onNavigateBack = {
                        navController.popBackStack()
                    },
                    onWithdrawSuccess = {
                        navController.popBackStack()
                    }
                )
            }
            
            composable(Screen.Profile.route) {
                ProfileScreen(
                    onNavigateToEditProfile = {
                        navController.navigate(Screen.EditProfile.route)
                    },
                    onLogout = {
                        navController.navigate(Screen.Login.route) {
                            popUpTo(0) { inclusive = true }
                        }
                    }
                )
            }
            
            composable(Screen.EditProfile.route) {
                EditProfileScreen(
                    onNavigateBack = {
                        navController.popBackStack()
                    }
                )
            }
        }
    }
}
