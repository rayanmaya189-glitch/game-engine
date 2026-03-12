package com.game-engine.casino.ui.navigation

sealed class Screen(val route: String) {
    object Splash : Screen("splash")
    object Login : Screen("login")
    object Register : Screen("register")
    object ForgotPassword : Screen("forgot_password")
    object Main : Screen("main")
    object Home : Screen("home")
    object Games : Screen("games")
    object GameDetail : Screen("game/{gameId}") {
        fun createRoute(gameId: String) = "game/$gameId"
    }
    object Wallet : Screen("wallet")
    object Deposit : Screen("deposit")
    object Withdraw : Screen("withdraw")
    object Profile : Screen("profile")
    object EditProfile : Screen("edit_profile")
    object TransactionHistory : Screen("transaction_history")
    object Bonuses : Screen("bonuses")
}

sealed class BottomNavItem(
    val route: String,
    val title: String,
    val icon: String
) {
    object Home : BottomNavItem("home", "Home", "home")
    object Games : BottomNavItem("games", "Games", "casino")
    object Wallet : BottomNavItem("wallet", "Wallet", "wallet")
    object Profile : BottomNavItem("profile", "Profile", "person")
}
