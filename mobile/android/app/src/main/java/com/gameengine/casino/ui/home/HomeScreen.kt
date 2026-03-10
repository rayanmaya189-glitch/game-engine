package com.gameengine.casino.ui.home

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.LazyRow
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import coil.compose.AsyncImage
import com.gameengine.casino.data.model.Game
import com.gameengine.casino.ui.theme.*

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
                    // Balance display
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
            // Jackpot Banner
            item {
                JackpotBanner(
                    jackpotAmount = uiState.jackpotGames.firstOrNull()?.jackpotAmount ?: 0.0
                )
            }
            
            // Categories
            item {
                CategoriesRow(
                    categories = uiState.categories.map { it.name },
                    onCategoryClick = { /* Navigate to category */ }
                )
            }
            
            // Featured Games
            item {
                GameSection(
                    title = "Featured Games",
                    games = uiState.featuredGames,
                    onGameClick = onGameClick,
                    onViewAllClick = { onViewAllClick("featured") }
                )
            }
            
            // Popular Games
            item {
                GameSection(
                    title = "Popular Games",
                    games = uiState.popularGames,
                    onGameClick = onGameClick,
                    onViewAllClick = { onViewAllClick("popular") }
                )
            }
            
            // Jackpot Games
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

@Composable
fun JackpotBanner(jackpotAmount: Double) {
    Box(
        modifier = Modifier
            .fillMaxWidth()
            .padding(16.dp)
            .background(
                brush = Brush.horizontalGradient(
                    colors = listOf(Primary, Gold)
                ),
                RoundedCornerShape(16.dp)
            )
            .padding(20.dp)
    ) {
        Row(
            modifier = Modifier.fillMaxWidth(),
            horizontalArrangement = Arrangement.SpaceBetween,
            verticalAlignment = Alignment.CenterVertically
        ) {
            Column {
                Text(
                    text = "JACKPOT",
                    style = MaterialTheme.typography.titleMedium,
                    color = MaterialTheme.colorScheme.onPrimary
                )
                Text(
                    text = "$${String.format("%,.2f", jackpotAmount)}",
                    style = MaterialTheme.typography.headlineLarge,
                    color = MaterialTheme.colorScheme.onPrimary,
                    fontWeight = FontWeight.Bold
                )
            }
            Icon(
                imageVector = Icons.Default.EmojiEvents,
                contentDescription = null,
                modifier = Modifier.size(48.dp),
                tint = MaterialTheme.colorScheme.onPrimary
            )
        }
    }
}

@Composable
fun CategoriesRow(
    categories: List<String>,
    onCategoryClick: (String) -> Unit
) {
    LazyRow(
        modifier = Modifier.padding(vertical = 8.dp),
        contentPadding = PaddingValues(horizontal = 16.dp),
        horizontalArrangement = Arrangement.spacedBy(12.dp)
    ) {
        items(categories) { category ->
            CategoryChip(
                category = category,
                onClick = { onCategoryClick(category) }
            )
        }
    }
}

@Composable
fun CategoryChip(
    category: String,
    onClick: () -> Unit
) {
    val icon = when (category.lowercase()) {
        "slots" -> Icons.Default.Casino
        "cards" -> Icons.Default.style
        "dice" -> Icons.Default.Casino
        "table" -> Icons.Default.TableRestaurant
        "live" -> Icons.Default.Videocam
        else -> Icons.Default.Games
    }
    
    Surface(
        modifier = Modifier.clickable(onClick = onClick),
        shape = RoundedCornerShape(20.dp),
        color = MaterialTheme.colorScheme.surface
    ) {
        Row(
            modifier = Modifier.padding(horizontal = 16.dp, vertical = 10.dp),
            verticalAlignment = Alignment.CenterVertically
        ) {
            Icon(
                imageVector = icon,
                contentDescription = null,
                modifier = Modifier.size(20.dp),
                tint = Primary
            )
            Spacer(modifier = Modifier.width(8.dp))
            Text(
                text = category,
                style = MaterialTheme.typography.bodyMedium,
                color = MaterialTheme.colorScheme.onSurface
            )
        }
    }
}

@Composable
fun GameSection(
    title: String,
    games: List<Game>,
    onGameClick: (String) -> Unit,
    onViewAllClick: () -> Unit
) {
    Column(modifier = Modifier.padding(vertical = 8.dp)) {
        Row(
            modifier = Modifier
                .fillMaxWidth()
                .padding(horizontal = 16.dp),
            horizontalArrangement = Arrangement.SpaceBetween,
            verticalAlignment = Alignment.CenterVertically
        ) {
            Text(
                text = title,
                style = MaterialTheme.typography.titleMedium,
                color = MaterialTheme.colorScheme.onBackground,
                fontWeight = FontWeight.Bold
            )
            TextButton(onClick = onViewAllClick) {
                Text("View All", color = Primary)
                Icon(
                    imageVector = Icons.Default.ArrowForward,
                    contentDescription = null,
                    tint = Primary,
                    modifier = Modifier.size(16.dp)
                )
            }
        }
        
        LazyRow(
            contentPadding = PaddingValues(horizontal = 16.dp),
            horizontalArrangement = Arrangement.spacedBy(12.dp)
        ) {
            items(games) { game ->
                GameCard(
                    game = game,
                    onClick = { onGameClick(game.id) }
                )
            }
        }
    }
}

@Composable
fun GameCard(
    game: Game,
    onClick: () -> Unit
) {
    Card(
        modifier = Modifier
            .width(140.dp)
            .height(180.dp)
            .clickable(onClick = onClick),
        shape = RoundedCornerShape(12.dp),
        colors = CardDefaults.cardColors(containerColor = CardBackground)
    ) {
        Box(modifier = Modifier.fillMaxSize()) {
            // Game thumbnail
            if (game.thumbnailUrl != null) {
                AsyncImage(
                    model = game.thumbnailUrl,
                    contentDescription = game.name,
                    modifier = Modifier.fillMaxSize(),
                    contentScale = ContentScale.Crop
                )
            } else {
                Box(
                    modifier = Modifier
                        .fillMaxSize()
                        .background(MaterialTheme.colorScheme.surface),
                    contentAlignment = Alignment.Center
                ) {
                    Icon(
                        imageVector = Icons.Default.Casino,
                        contentDescription = null,
                        modifier = Modifier.size(48.dp),
                        tint = MaterialTheme.colorScheme.onSurface.copy(alpha = 0.5f)
                    )
                }
            }
            
            // Gradient overlay
            Box(
                modifier = Modifier
                    .fillMaxSize()
                    .background(
                        brush = Brush.verticalGradient(
                            colors = listOf(
                                android.graphics.Color.TRANSPARENT.let { androidx.compose.ui.graphics.Color(it) },
                                androidx.compose.ui.graphics.Color.Black.copy(alpha = 0.8f)
                            )
                        )
                    )
            )
            
            // Game info
            Column(
                modifier = Modifier
                    .align(Alignment.BottomStart)
                    .padding(12.dp)
            ) {
                Text(
                    text = game.name,
                    style = MaterialTheme.typography.bodyMedium,
                    color = MaterialTheme.colorScheme.onBackground,
                    maxLines = 1,
                    overflow = TextOverflow.Ellipsis
                )
                if (game.jackpotAmount != null && game.jackpotAmount > 0) {
                    Row(verticalAlignment = Alignment.CenterVertically) {
                        Icon(
                            imageVector = Icons.Default.EmojiEvents,
                            contentDescription = null,
                            modifier = Modifier.size(14.dp),
                            tint = Gold
                        )
                        Spacer(modifier = Modifier.width(4.dp))
                        Text(
                            text = "$${String.format("%,.0f", game.jackpotAmount)}",
                            style = MaterialTheme.typography.bodySmall,
                            color = Gold
                        )
                    }
                }
            }
            
            // Featured badge
            if (game.isFeatured || game.isNew || game.isHot) {
                Surface(
                    modifier = Modifier
                        .align(Alignment.TopEnd)
                        .padding(8.dp),
                    shape = RoundedCornerShape(4.dp),
                    color = when {
                        game.isNew -> Success
                        game.isHot -> Primary
                        else -> Gold
                    }
                ) {
                    Text(
                        text = when {
                            game.isNew -> "NEW"
                            game.isHot -> "HOT"
                            else -> "FEATURED"
                        },
                        style = MaterialTheme.typography.labelSmall,
                        color = if (game.isNew) OnSuccess else OnPrimary,
                        modifier = Modifier.padding(horizontal = 6.dp, vertical = 2.dp)
                    )
                }
            }
        }
    }
}
