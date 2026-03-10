import SwiftUI

// MARK: - Games View
struct GamesView: View {
    @State private var games: [Game] = []
    @State private var categories: [Category] = []
    @State private var selectedCategory: String? = nil
    @State private var searchText = ""
    @State private var isLoading = true
    
    var filteredGames: [Game] {
        if searchText.isEmpty {
            return games
        }
        return games.filter { $0.name.localizedCaseInsensitiveContains(searchText) }
    }
    
    var body: some View {
        NavigationStack {
            VStack(spacing: 0) {
                // Category Filter
                ScrollView(.horizontal, showsIndicators: false) {
                    HStack(spacing: 8) {
                        CategoryChip(title: "All", isSelected: selectedCategory == nil) {
                            selectedCategory = nil
                        }
                        ForEach(categories) { category in
                            CategoryChip(title: category.name, isSelected: selectedCategory == category.id) {
                                selectedCategory = category.id
                            }
                        }
                    }
                    .padding(.horizontal)
                }
                .padding(.vertical, 8)
                
                // Games Grid
                if isLoading {
                    Spacer()
                    ProgressView()
                    Spacer()
                } else {
                    ScrollView {
                        LazyVGrid(columns: [GridItem(.flexible()), GridItem(.flexible())], spacing: 16) {
                            ForEach(filteredGames) { game in
                                NavigationLink(destination: GameDetailView(game: game)) {
                                    GameGridCard(game: game)
                                }
                                .buttonStyle(.plain)
                            }
                        }
                        .padding()
                    }
                }
            }
            .navigationTitle("Games")
            .searchable(text: $searchText, prompt: "Search games")
        }
        .task {
            await loadData()
        }
    }
    
    private func loadData() async {
        isLoading = true
        do {
            async let gamesTask = APIClient.shared.getGames(category: selectedCategory)
            async let categoriesTask = APIClient.shared.getCategories()
            
            games = try await gamesTask.games
            categories = try await categoriesTask.categories
        } catch {
            print(error)
        }
        isLoading = false
    }
}

struct CategoryChip: View {
    let title: String
    let isSelected: Bool
    let action: () -> Void
    
    var body: some View {
        Button(action: action) {
            Text(title)
                .font(.subheadline)
                .padding(.horizontal, 16)
                .padding(.vertical, 8)
                .background(isSelected ? Color.purple : Color(.secondarySystemBackground))
                .foregroundColor(isSelected ? .white : .primary)
                .cornerRadius(20)
        }
    }
}

struct GameGridCard: View {
    let game: Game
    
    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            RoundedRectangle(cornerRadius: 8)
                .fill(Color.purple.opacity(0.2))
                .frame(height: 100)
                .overlay(
                    Image(systemName: "casino.fill")
                        .font(.largeTitle)
                        .foregroundColor(.purple)
                )
            
            Text(game.name)
                .font(.subheadline)
                .fontWeight(.medium)
                .lineLimit(1)
            
            Text(game.provider)
                .font(.caption)
                .foregroundColor(.secondary)
            
            HStack {
                Text("Min: $\(game.minBet, specifier: "%.2f")")
                    .font(.caption2)
                Spacer()
                Text("RTP: \(game.rtp, specifier: "%.1f")%")
                    .font(.caption2)
                    .foregroundColor(.green)
            }
        }
        .padding()
        .background(Color(.systemBackground))
        .cornerRadius(12)
        .shadow(radius: 2)
    }
}

struct GameDetailView: View {
    let game: Game
    @State private var gameDetails: GameDetails?
    
    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 16) {
                RoundedRectangle(cornerRadius: 16)
                    .fill(Color.purple.opacity(0.2))
                    .frame(height: 200)
                    .overlay(
                        Image(systemName: "casino.fill")
                            .font(.system(size: 60))
                            .foregroundColor(.purple)
                    )
                
                Text(game.name)
                    .font(.title)
                    .fontWeight(.bold)
                
                Text("Provider: \(game.provider)")
                    .foregroundColor(.secondary)
                
                HStack(spacing: 24) {
                    VStack {
                        Text("Min Bet")
                            .font(.caption)
                            .foregroundColor(.secondary)
                        Text("$\(game.minBet, specifier: "%.2f")")
                            .fontWeight(.bold)
                    }
                    VStack {
                        Text("Max Bet")
                            .font(.caption)
                            .foregroundColor(.secondary)
                        Text("$\(game.maxBet, specifier: "%.2f")")
                            .fontWeight(.bold)
                    }
                    VStack {
                        Text("RTP")
                            .font(.caption)
                            .foregroundColor(.secondary)
                        Text("\(game.rtp, specifier: "%.1f")%")
                            .fontWeight(.bold)
                            .foregroundColor(.green)
                    }
                }
                .padding()
                .background(Color(.secondarySystemBackground))
                .cornerRadius(12)
                
                Button(action: {}) {
                    Text("Play Now")
                        .font(.headline)
                        .frame(maxWidth: .infinity)
                        .padding()
                        .background(Color.purple)
                        .foregroundColor(.white)
                        .cornerRadius(12)
                }
            }
            .padding()
        }
        .navigationTitle("Game Details")
        .navigationBarTitleDisplayMode(.inline)
    }
}

// MARK: - Tournaments View
struct TournamentsView: View {
    @State private var tournaments: [Tournament] = []
    @State private var isLoading = true
    
    var body: some View {
        NavigationStack {
            if isLoading {
                ProgressView()
            } else if tournaments.isEmpty {
                Text("No tournaments available")
                    .foregroundColor(.secondary)
            } else {
                List(tournaments) { tournament in
                    NavigationLink(destination: TournamentDetailView(tournament: tournament)) {
                        TournamentListRow(tournament: tournament)
                    }
                }
                .listStyle(.plain)
            }
        }
        .task {
            await loadData()
        }
    }
    
    private func loadData() async {
        isLoading = true
        do {
            let response = try await APIClient.shared.getTournaments()
            tournaments = response.tournaments
        } catch {
            print(error)
        }
        isLoading = false
    }
}

struct TournamentListRow: View {
    let tournament: Tournament
    
    var body: some View {
        HStack {
            VStack(alignment: .leading, spacing: 4) {
                Text(tournament.name)
                    .font(.headline)
                Text(tournament.game)
                    .font(.subheadline)
                    .foregroundColor(.secondary)
                Text("\(tournament.playerCount) players")
                    .font(.caption)
                    .foregroundColor(.secondary)
            }
            Spacer()
            VStack(alignment: .trailing) {
                Text("$\(tournament.prizePool, specifier: "%.0f")")
                    .font(.headline)
                    .foregroundColor(.purple)
                Text(tournament.status.capitalized)
                    .font(.caption)
                    .padding(.horizontal, 8)
                    .padding(.vertical, 4)
                    .background(tournament.status == "active" ? Color.green.opacity(0.2) : Color.gray.opacity(0.2))
                    .cornerRadius(4)
            }
        }
        .padding(.vertical, 8)
    }
}

struct TournamentDetailView: View {
    let tournament: Tournament
    
    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 16) {
                VStack(alignment: .leading, spacing: 8) {
                    Text(tournament.name)
                        .font(.title)
                        .fontWeight(.bold)
                    Text(tournament.game)
                        .foregroundColor(.secondary)
                }
                
                HStack(spacing: 24) {
                    VStack {
                        Text("Prize Pool")
                            .font(.caption)
                            .foregroundColor(.secondary)
                        Text("$\(tournament.prizePool, specifier: "%.0f")")
                            .font(.title2)
                            .fontWeight(.bold)
                            .foregroundColor(.purple)
                    }
                    VStack {
                        Text("Players")
                            .font(.caption)
                            .foregroundColor(.secondary)
                        Text("\(tournament.playerCount)")
                            .font(.title2)
                            .fontWeight(.bold)
                    }
                    VStack {
                        Text("Min Bet")
                            .font(.caption)
                            .foregroundColor(.secondary)
                        Text("$\(tournament.minBet, specifier: "%.2f")")
                            .font(.title2)
                            .fontWeight(.bold)
                    }
                }
                .padding()
                .background(Color(.secondarySystemBackground))
                .cornerRadius(12)
                
                if let description = tournament.description {
                    Text(description)
                        .foregroundColor(.secondary)
                }
            }
            .padding()
        }
        .navigationTitle("Tournament")
        .navigationBarTitleDisplayMode(.inline)
    }
}

// MARK: - Jackpots View
struct JackpotsView: View {
    @State private var jackpots: [Jackpot] = []
    @State private var isLoading = true
    
    var body: some View {
        NavigationStack {
            if isLoading {
                ProgressView()
            } else if jackpots.isEmpty {
                Text("No jackpots available")
                    .foregroundColor(.secondary)
            } else {
                List(jackpots) { jackpot in
                    NavigationLink(destination: JackpotDetailView(jackpot: jackpot)) {
                        JackpotListRow(jackpot: jackpot)
                    }
                }
                .listStyle(.plain)
            }
        }
        .task {
            await loadData()
        }
    }
    
    private func loadData() async {
        isLoading = true
        do {
            let response = try await APIClient.shared.getJackpots()
            jackpots = response.jackpots
        } catch {
            print(error)
        }
        isLoading = false
    }
}

struct JackpotListRow: View {
    let jackpot: Jackpot
    
    var body: some View {
        HStack {
            VStack(alignment: .leading, spacing: 4) {
                Text(jackpot.name)
                    .font(.headline)
                Text(jackpot.game)
                    .font(.subheadline)
                    .foregroundColor(.secondary)
            }
            Spacer()
            VStack(alignment: .trailing) {
                Text("$\(jackpot.currentAmount, specifier: "%.0f")")
                    .font(.headline)
                    .foregroundColor(.purple)
                Text("\(jackpot.hitCount) wins")
                    .font(.caption)
                    .foregroundColor(.secondary)
            }
        }
        .padding(.vertical, 8)
    }
}

struct JackpotDetailView: View {
    let jackpot: Jackpot
    
    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 16) {
                VStack(alignment: .leading, spacing: 8) {
                    Text(jackpot.name)
                        .font(.title)
                        .fontWeight(.bold)
                    Text(jackpot.game)
                        .foregroundColor(.secondary)
                }
                
                VStack(spacing: 8) {
                    Text("Current Amount")
                        .font(.caption)
                        .foregroundColor(.secondary)
                    Text("$\(jackpot.currentAmount, specifier: "%.0f")")
                        .font(.system(size: 40, weight: .bold))
                        .foregroundColor(.purple)
                }
                .frame(maxWidth: .infinity)
                .padding()
                .background(Color(.secondarySystemBackground))
                .cornerRadius(12)
                
                HStack(spacing: 24) {
                    VStack {
                        Text("Min Bet")
                            .font(.caption)
                            .foregroundColor(.secondary)
                        Text("$\(jackpot.minBet, specifier: "%.2f")")
                            .fontWeight(.bold)
                    }
                    VStack {
                        Text("Max Win")
                            .font(.caption)
                            .foregroundColor(.secondary)
                        Text("$\(jackpot.maxWin, specifier: "%.0f")")
                            .fontWeight(.bold)
                    }
                    VStack {
                        Text("Total Wins")
                            .font(.caption)
                            .foregroundColor(.secondary)
                        Text("\(jackpot.hitCount)")
                            .fontWeight(.bold)
                    }
                }
                .padding()
                .background(Color(.secondarySystemBackground))
                .cornerRadius(12)
            }
            .padding()
        }
        .navigationTitle("Jackpot")
        .navigationBarTitleDisplayMode(.inline)
    }
}

// MARK: - Wallet View
struct WalletView: View {
    @State private var balance: Double = 0
    @State private var bonusBalance: Double = 0
    @State private var transactions: [Transaction] = []
    @State private var isLoading = true
    
    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(spacing: 16) {
                    // Balance Card
                    VStack(spacing: 8) {
                        Text("Balance")
                            .font(.subheadline)
                            .foregroundColor(.white.opacity(0.8))
                        Text("$\(balance, specifier: "%.2f")")
                            .font(.system(size: 36, weight: .bold))
                            .foregroundColor(.white)
                        if bonusBalance > 0 {
                            Text("Bonus: $\(bonusBalance, specifier: "%.2f")")
                                .font(.caption)
                                .foregroundColor(.white.opacity(0.7))
                        }
                    }
                    .frame(maxWidth: .infinity)
                    .padding(.vertical, 24)
                    .background(Color.purple)
                    .cornerRadius(16)
                    
                    // Actions
                    HStack(spacing: 12) {
                        Button(action: {}) {
                            Label("Deposit", systemImage: "plus.circle.fill")
                                .frame(maxWidth: .infinity)
                        }
                        .buttonStyle(.borderedProminent)
                        
                        Button(action: {}) {
                            Label("Withdraw", systemImage: "minus.circle.fill")
                                .frame(maxWidth: .infinity)
                        }
                        .buttonStyle(.bordered)
                    }
                    
                    // Transactions
                    VStack(alignment: .leading, spacing: 12) {
                        Text("Recent Transactions")
                            .font(.headline)
                        
                        if transactions.isEmpty {
                            Text("No transactions yet")
                                .foregroundColor(.secondary)
                                .frame(maxWidth: .infinity, alignment: .center)
                                .padding()
                        } else {
                            ForEach(transactions) { transaction in
                                TransactionRow(transaction: transaction)
                            }
                        }
                    }
                }
                .padding()
            }
            .navigationTitle("Wallet")
        }
        .task {
            await loadData()
        }
    }
    
    private func loadData() async {
        isLoading = true
        do {
            let balanceResponse = try await APIClient.shared.getBalance()
            balance = balanceResponse.balance
            bonusBalance = balanceResponse.bonusBalance
            
            let transactionsResponse = try await APIClient.shared.getTransactions()
            transactions = transactionsResponse.transactions
        } catch {
            print(error)
        }
        isLoading = false
    }
}

struct TransactionRow: View {
    let transaction: Transaction
    
    var body: some View {
        HStack {
            Image(systemName: transaction.type == "deposit" ? "arrow.down.circle.fill" : "arrow.up.circle.fill")
                .foregroundColor(transaction.type == "deposit" ? .green : .red)
                .font(.title2)
            
            VStack(alignment: .leading, spacing: 2) {
                Text(transaction.type.capitalized)
                    .font(.subheadline)
                    .fontWeight(.medium)
                Text(transaction.createdAt)
                    .font(.caption)
                    .foregroundColor(.secondary)
            }
            
            Spacer()
            
            Text("\(transaction.type == "deposit" ? "+" : "-")$\(transaction.amount, specifier: "%.2f")")
                .font(.subheadline)
                .fontWeight(.bold)
                .foregroundColor(transaction.type == "deposit" ? .green : .red)
        }
        .padding()
        .background(Color(.secondarySystemBackground))
        .cornerRadius(8)
    }
}

// MARK: - Profile View
struct ProfileView: View {
    @EnvironmentObject var authManager: AuthManager
    @State private var profile: UserProfile?
    @State private var isLoading = true
    
    var body: some View {
        NavigationStack {
            List {
                Section {
                    HStack {
                        Image(systemName: "person.circle.fill")
                            .font(.system(size: 60))
                            .foregroundColor(.purple)
                        
                        VStack(alignment: .leading, spacing: 4) {
                            Text(profile?.username ?? "Player")
                                .font(.headline)
                            Text(profile?.email ?? "")
                                .font(.subheadline)
                                .foregroundColor(.secondary)
                        }
                    }
                    .padding(.vertical, 8)
                }
                
                Section("Account") {
                    NavigationLink("Edit Profile") { }
                    NavigationLink("Security") { }
                    NavigationLink("Notifications") { }
                }
                
                Section("Support") {
                    NavigationLink("Help & Support") { }
                    NavigationLink("About") { }
                }
                
                Section {
                    Button("Logout") {
                        Task {
                            await authManager.logout()
                        }
                    }
                    .foregroundColor(.red)
                }
            }
            .navigationTitle("Profile")
        }
        .task {
            await loadData()
        }
    }
    
    private func loadData() async {
        isLoading = true
        do {
            profile = try await APIClient.shared.getProfile()
        } catch {
            print(error)
        }
        isLoading = false
    }
}
