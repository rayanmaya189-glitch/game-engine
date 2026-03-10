package com.casino.game.ui

import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.grid.GridCells
import androidx.compose.foundation.lazy.grid.LazyVerticalGrid
import androidx.compose.foundation.lazy.grid.items
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import com.casino.game.ui.home.GameCard

@Composable
fun GamesScreen(
    onGameClick: (String) -> Unit,
    viewModel: GamesViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()
    var selectedCategory by remember { mutableStateOf<String?>(null) }

    Column(modifier = Modifier.fillMaxSize()) {
        // Search Bar
        OutlinedTextField(
            value = state.searchQuery,
            onValueChange = { viewModel.searchGames(it) },
            modifier = Modifier
                .fillMaxWidth()
                .padding(16.dp),
            placeholder = { Text("Search games...") },
            leadingIcon = { Icon(Icons.Default.Search, contentDescription = null) },
            singleLine = true
        )

        // Category Filter
        ScrollableTabRow(
            selectedTabIndex = state.categories.indexOfFirst { it.id == selectedCategory }.coerceAtLeast(0),
            modifier = Modifier.fillMaxWidth()
        ) {
            Tab(
                selected = selectedCategory == null,
                onClick = { selectedCategory = null },
                text = { Text("All") }
            )
            state.categories.forEach { category ->
                Tab(
                    selected = selectedCategory == category.id,
                    onClick = { selectedCategory = category.id },
                    text = { Text(category.name) }
                )
            }
        }

        // Games Grid
        if (state.isLoading) {
            Box(
                modifier = Modifier.fillMaxSize(),
                contentAlignment = Alignment.Center
            ) {
                CircularProgressIndicator()
            }
        } else {
            LazyVerticalGrid(
                columns = GridCells.Fixed(2),
                contentPadding = PaddingValues(16.dp),
                horizontalArrangement = Arrangement.spacedBy(12.dp),
                verticalArrangement = Arrangement.spacedBy(12.dp)
            ) {
                items(state.games) { game ->
                    GameCard(
                        game = game,
                        onClick = { onGameClick(game.id) }
                    )
                }
            }
        }
    }
}

@HiltViewModel
class GamesViewModel @javax.inject.Inject constructor(
    private val gameRepository: com.casino.game.data.repository.GameRepository
) : androidx.lifecycle.ViewModel() {
    private val _state = MutableStateFlow(GamesState())
    val state: StateFlow<GamesState> = _state.asStateFlow()

    init {
        loadGames()
        loadCategories()
    }

    fun loadGames(category: String? = null) {
        androidx.lifecycle.viewModelScope.launch {
            _state.update { it.copy(isLoading = true) }
            gameRepository.getGames(category = category).fold(
                onSuccess = { response ->
                    _state.update { it.copy(games = response.games, isLoading = false) }
                },
                onFailure = {
                    _state.update { it.copy(isLoading = false) }
                }
            )
        }
    }

    private fun loadCategories() {
        androidx.lifecycle.viewModelScope.launch {
            gameRepository.getCategories().fold(
                onSuccess = { response ->
                    _state.update { it.copy(categories = response.categories) }
                },
                onFailure = { }
            )
        }
    }

    fun searchGames(query: String) {
        _state.update { it.copy(searchQuery = query) }
        // Implement search
    }
}

data class GamesState(
    val isLoading: Boolean = false,
    val games: List<com.casino.game.data.model.Game> = emptyList(),
    val categories: List<com.casino.game.data.model.Category> = emptyList(),
    val searchQuery: String = "",
    val error: String? = null
)

@Composable
fun TournamentsScreen(
    onTournamentClick: (String) -> Unit,
    viewModel: TournamentsViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()

    Column(modifier = Modifier.fillMaxSize()) {
        if (state.isLoading) {
            Box(
                modifier = Modifier.fillMaxSize(),
                contentAlignment = Alignment.Center
            ) {
                CircularProgressIndicator()
            }
        } else {
            LazyVerticalGrid(
                columns = GridCells.Fixed(1),
                contentPadding = PaddingValues(16.dp),
                verticalArrangement = Arrangement.spacedBy(12.dp)
            ) {
                items(state.tournaments) { tournament ->
                    TournamentListItem(
                        tournament = tournament,
                        onClick = { onTournamentClick(tournament.id) }
                    )
                }
            }
        }
    }
}

@HiltViewModel
class TournamentsViewModel @javax.inject.Inject constructor(
    private val tournamentRepository: com.casino.game.data.repository.TournamentRepository
) : androidx.lifecycle.ViewModel() {
    private val _state = MutableStateFlow(TournamentsState())
    val state: StateFlow<TournamentsState> = _state.asStateFlow()

    init {
        loadTournaments()
    }

    fun loadTournaments() {
        androidx.lifecycle.viewModelScope.launch {
            _state.update { it.copy(isLoading = true) }
            tournamentRepository.getTournaments().fold(
                onSuccess = { response ->
                    _state.update { it.copy(tournaments = response.tournaments, isLoading = false) }
                },
                onFailure = {
                    _state.update { it.copy(isLoading = false) }
                }
            )
        }
    }
}

data class TournamentsState(
    val isLoading: Boolean = false,
    val tournaments: List<com.casino.game.data.model.Tournament> = emptyList(),
    val error: String? = null
)

@Composable
fun TournamentListItem(
    tournament: com.casino.game.data.model.Tournament,
    onClick: () -> Unit
) {
    Card(
        onClick = onClick,
        modifier = Modifier.fillMaxWidth()
    ) {
        Row(
            modifier = Modifier.padding(16.dp),
            horizontalArrangement = Arrangement.SpaceBetween
        ) {
            Column {
                Text(tournament.name, style = MaterialTheme.typography.titleMedium)
                Text(tournament.game, style = MaterialTheme.typography.bodySmall)
            }
            Column(horizontalAlignment = Alignment.End) {
                Text("$${tournament.prizePool}", style = MaterialTheme.typography.titleMedium)
                Text("${tournament.playerCount} players", style = MaterialTheme.typography.bodySmall)
            }
        }
    }
}

@Composable
fun JackpotsScreen(
    onJackpotClick: (String) -> Unit,
    viewModel: JackpotsViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()

    Column(modifier = Modifier.fillMaxSize()) {
        if (state.isLoading) {
            Box(
                modifier = Modifier.fillMaxSize(),
                contentAlignment = Alignment.Center
            ) {
                CircularProgressIndicator()
            }
        } else {
            LazyVerticalGrid(
                columns = GridCells.Fixed(1),
                contentPadding = PaddingValues(16.dp),
                verticalArrangement = Arrangement.spacedBy(12.dp)
            ) {
                items(state.jackpots) { jackpot ->
                    JackpotListItem(
                        jackpot = jackpot,
                        onClick = { onJackpotClick(jackpot.id) }
                    )
                }
            }
        }
    }
}

@HiltViewModel
class JackpotsViewModel @javax.inject.Inject constructor(
    private val jackpotRepository: com.casino.game.data.repository.JackpotRepository
) : androidx.lifecycle.ViewModel() {
    private val _state = MutableStateFlow(JackpotsState())
    val state: StateFlow<JackpotsState> = _state.asStateFlow()

    init {
        loadJackpots()
    }

    fun loadJackpots() {
        androidx.lifecycle.viewModelScope.launch {
            _state.update { it.copy(isLoading = true) }
            jackpotRepository.getJackpots().fold(
                onSuccess = { response ->
                    _state.update { it.copy(jackpots = response.jackpots, isLoading = false) }
                },
                onFailure = {
                    _state.update { it.copy(isLoading = false) }
                }
            )
        }
    }
}

data class JackpotsState(
    val isLoading: Boolean = false,
    val jackpots: List<com.casino.game.data.model.Jackpot> = emptyList(),
    val error: String? = null
)

@Composable
fun JackpotListItem(
    jackpot: com.casino.game.data.model.Jackpot,
    onClick: () -> Unit
) {
    Card(
        onClick = onClick,
        modifier = Modifier.fillMaxWidth()
    ) {
        Row(
            modifier = Modifier.padding(16.dp),
            horizontalArrangement = Arrangement.SpaceBetween
        ) {
            Column {
                Text(jackpot.name, style = MaterialTheme.typography.titleMedium)
                Text(jackpot.game, style = MaterialTheme.typography.bodySmall)
            }
            Column(horizontalAlignment = Alignment.End) {
                Text("$${jackpot.currentAmount}", style = MaterialTheme.typography.titleMedium)
                Text("${jackpot.hitCount} wins", style = MaterialTheme.typography.bodySmall)
            }
        }
    }
}

@Composable
fun WalletScreen(
    viewModel: WalletViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()

    Column(modifier = Modifier.fillMaxSize().padding(16.dp)) {
        // Balance Card
        Card(
            modifier = Modifier.fillMaxWidth(),
            colors = CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.primaryContainer)
        ) {
            Column(modifier = Modifier.padding(24.dp), horizontalAlignment = Alignment.CenterHorizontally) {
                Text("Balance", style = MaterialTheme.typography.bodyLarge)
                Text("$${state.balance}", style = MaterialTheme.typography.headlineLarge)
                if (state.bonusBalance > 0) {
                    Text("Bonus: $${state.bonusBalance}", style = MaterialTheme.typography.bodyMedium)
                }
            }
        }

        Spacer(modifier = Modifier.height(24.dp))

        // Actions
        Row(
            modifier = Modifier.fillMaxWidth(),
            horizontalArrangement = Arrangement.spacedBy(12.dp)
        ) {
            Button(
                onClick = { viewModel.showDepositDialog() },
                modifier = Modifier.weight(1f)
            ) {
                Icon(Icons.Default.Add, contentDescription = null)
                Spacer(modifier = Modifier.width(8.dp))
                Text("Deposit")
            }
            OutlinedButton(
                onClick = { viewModel.showWithdrawDialog() },
                modifier = Modifier.weight(1f)
            ) {
                Icon(Icons.Default.Remove, contentDescription = null)
                Spacer(modifier = Modifier.width(8.dp))
                Text("Withdraw")
            }
        }

        Spacer(modifier = Modifier.height(24.dp))

        // Transactions
        Text("Recent Transactions", style = MaterialTheme.typography.titleMedium)
        Spacer(modifier = Modifier.height(12.dp))

        state.transactions.forEach { transaction ->
            TransactionItem(transaction = transaction)
            Spacer(modifier = Modifier.height(8.dp))
        }
    }
}

@HiltViewModel
class WalletViewModel @javax.inject.Inject constructor(
    private val walletRepository: com.casino.game.data.repository.WalletRepository
) : androidx.lifecycle.ViewModel() {
    private val _state = MutableStateFlow(WalletState())
    val state: StateFlow<WalletState> = _state.asStateFlow()

    init {
        loadWallet()
    }

    fun loadWallet() {
        androidx.lifecycle.viewModelScope.launch {
            walletRepository.getBalance().fold(
                onSuccess = { balance ->
                    _state.update { it.copy(balance = balance.balance, bonusBalance = balance.bonusBalance) }
                },
                onFailure = { }
            )
            walletRepository.getTransactions().fold(
                onSuccess = { response ->
                    _state.update { it.copy(transactions = response.transactions) }
                },
                onFailure = { }
            )
        }
    }

    fun showDepositDialog() { }
    fun showWithdrawDialog() { }
}

data class WalletState(
    val balance: Double = 0.0,
    val bonusBalance: Double = 0.0,
    val transactions: List<com.casino.game.data.model.Transaction> = emptyList(),
    val isLoading: Boolean = false
)

@Composable
fun TransactionItem(transaction: com.casino.game.data.model.Transaction) {
    Card(modifier = Modifier.fillMaxWidth()) {
        Row(
            modifier = Modifier.padding(12.dp),
            horizontalArrangement = Arrangement.SpaceBetween,
            verticalAlignment = Alignment.CenterVertically
        ) {
            Row(verticalAlignment = Alignment.CenterVertically) {
                Icon(
                    imageVector = if (transaction.type == "deposit") Icons.Default.ArrowDownward else Icons.Default.ArrowUpward,
                    contentDescription = null,
                    tint = if (transaction.type == "deposit") MaterialTheme.colorScheme.primary else MaterialTheme.colorScheme.error
                )
                Spacer(modifier = Modifier.width(12.dp))
                Column {
                    Text(transaction.type.replaceFirstChar { it.uppercase() }, style = MaterialTheme.typography.bodyMedium)
                    Text(transaction.createdAt, style = MaterialTheme.typography.bodySmall)
                }
            }
            Text(
                "${if (transaction.type == "deposit") "+" else "-"}$${transaction.amount}",
                style = MaterialTheme.typography.titleMedium,
                color = if (transaction.type == "deposit") MaterialTheme.colorScheme.primary else MaterialTheme.colorScheme.error
            )
        }
    }
}

@Composable
fun ProfileScreen(
    onLogout: () -> Unit,
    viewModel: ProfileViewModel = hiltViewModel()
) {
    val state by viewModel.state.collectAsState()

    Column(modifier = Modifier.fillMaxSize().padding(16.dp)) {
        // Profile Header
        Card(modifier = Modifier.fillMaxWidth()) {
            Column(
                modifier = Modifier.padding(24.dp),
                horizontalAlignment = Alignment.CenterHorizontally
            ) {
                Icon(
                    imageVector = Icons.Default.AccountCircle,
                    contentDescription = null,
                    modifier = Modifier.size(80.dp),
                    tint = MaterialTheme.colorScheme.primary
                )
                Spacer(modifier = Modifier.height(12.dp))
                Text(state.username, style = MaterialTheme.typography.headlineSmall)
                Text(state.email, style = MaterialTheme.typography.bodyMedium)
                Spacer(modifier = Modifier.height(8.dp))
                AssistChip(
                    onClick = { },
                    label = { Text("KYC: ${state.kycLevel}") }
                )
            }
        }

        Spacer(modifier = Modifier.height(24.dp))

        // Menu Items
        ProfileMenuItem(icon = Icons.Default.Person, title = "Edit Profile", onClick = { })
        ProfileMenuItem(icon = Icons.Default.Security, title = "Security", onClick = { })
        ProfileMenuItem(icon = Icons.Default.Notifications, title = "Notifications", onClick = { })
        ProfileMenuItem(icon = Icons.Default.Help, title = "Help & Support", onClick = { })

        Spacer(modifier = Modifier.weight(1f))

        Button(
            onClick = { viewModel.logout(onLogout) },
            modifier = Modifier.fillMaxWidth(),
            colors = ButtonDefaults.buttonColors(containerColor = MaterialTheme.colorScheme.error)
        ) {
            Icon(Icons.Default.Logout, contentDescription = null)
            Spacer(modifier = Modifier.width(8.dp))
            Text("Logout")
        }
    }
}

@Composable
fun ProfileMenuItem(
    icon: androidx.compose.ui.graphics.vector.ImageVector,
    title: String,
    onClick: () -> Unit
) {
    Card(
        onClick = onClick,
        modifier = Modifier.fillMaxWidth().padding(vertical = 4.dp)
    ) {
        Row(
            modifier = Modifier.padding(16.dp),
            verticalAlignment = Alignment.CenterVertically
        ) {
            Icon(icon, contentDescription = null)
            Spacer(modifier = Modifier.width(16.dp))
            Text(title, style = MaterialTheme.typography.bodyLarge)
            Spacer(modifier = Modifier.weight(1f))
            Icon(Icons.Default.ChevronRight, contentDescription = null)
        }
    }
}

@HiltViewModel
class ProfileViewModel @javax.inject.Inject constructor(
    private val authRepository: com.casino.game.data.repository.AuthRepository
) : androidx.lifecycle.ViewModel() {
    private val _state = MutableStateFlow(ProfileState())
    val state: StateFlow<ProfileState> = _state.asStateFlow()

    init {
        loadProfile()
    }

    private fun loadProfile() {
        _state.update { it.copy(username = "Player", email = "player@casino.com", kycLevel = "verified") }
    }

    fun logout(onSuccess: () -> Unit) {
        androidx.lifecycle.viewModelScope.launch {
            authRepository.logout()
            onSuccess()
        }
    }
}

data class ProfileState(
    val username: String = "",
    val email: String = "",
    val kycLevel: String = "",
    val isLoading: Boolean = false
)
