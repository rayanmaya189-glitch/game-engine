import SwiftUI

// MARK: - Content View (Main Navigation)
struct ContentView: View {
    @EnvironmentObject var authManager: AuthManager
    
    var body: some View {
        if authManager.isLoggedIn {
            MainTabView()
        } else {
            LoginView()
        }
    }
}

// MARK: - Login View
struct LoginView: View {
    @EnvironmentObject var authManager: AuthManager
    @State private var email = ""
    @State private var password = ""
    @State private var showRegister = false
    
    var body: some View {
        NavigationStack {
            VStack(spacing: 24) {
                Image(systemName: "casino.fill")
                    .font(.system(size: 60))
                    .foregroundColor(.purple)
                
                Text("Casino Games")
                    .font(.largeTitle)
                    .fontWeight(.bold)
                
                Text("Sign in to play")
                    .foregroundColor(.secondary)
                
                if let error = authManager.error {
                    Text(error)
                        .foregroundColor(.red)
                        .font(.caption)
                }
                
                TextField("Email", text: $email)
                    .textFieldStyle(.roundedBorder)
                    .textContentType(.emailAddress)
                    .autocapitalization(.none)
                
                SecureField("Password", text: $password)
                    .textFieldStyle(.roundedBorder)
                    .textContentType(.password)
                
                Button(action: {
                    Task {
                        await authManager.login(email: email, password: password)
                    }
                }) {
                    if authManager.isLoading {
                        ProgressView()
                            .frame(maxWidth: .infinity)
                    } else {
                        Text("Sign In")
                            .frame(maxWidth: .infinity)
                    }
                }
                .buttonStyle(.borderedProminent)
                .disabled(email.isEmpty || password.isEmpty || authManager.isLoading)
                
                NavigationLink(destination: RegisterView()) {
                    Text("Don't have an account? Sign Up")
                }
            }
            .padding()
            .navigationBarHidden(true)
        }
    }
}

// MARK: - Register View
struct RegisterView: View {
    @EnvironmentObject var authManager: AuthManager
    @Environment(\.dismiss) var dismiss
    @State private var email = ""
    @State private var username = ""
    @State private var phone = ""
    @State private var password = ""
    @State private var confirmPassword = ""
    
    var body: some View {
        VStack(spacing: 16) {
            Text("Create Account")
                .font(.title)
                .fontWeight(.bold)
            
            if let error = authManager.error {
                Text(error)
                    .foregroundColor(.red)
                    .font(.caption)
            }
            
            TextField("Username", text: $username)
                .textFieldStyle(.roundedBorder)
            
            TextField("Email", text: $email)
                .textFieldStyle(.roundedBorder)
                .textContentType(.emailAddress)
                .autocapitalization(.none)
            
            TextField("Phone (Optional)", text: $phone)
                .textFieldStyle(.roundedBorder)
                .textContentType(.telephoneNumber)
            
            SecureField("Password", text: $password)
                .textFieldStyle(.roundedBorder)
            
            SecureField("Confirm Password", text: $confirmPassword)
                .textFieldStyle(.roundedBorder)
            
            Button(action: {
                Task {
                    await authManager.register(
                        email: email,
                        password: password,
                        username: username,
                        phone: phone.isEmpty ? nil : phone
                    )
                }
            }) {
                if authManager.isLoading {
                    ProgressView()
                        .frame(maxWidth: .infinity)
                } else {
                    Text("Create Account")
                        .frame(maxWidth: .infinity)
                }
            }
            .buttonStyle(.borderedProminent)
            .disabled(email.isEmpty || password.isEmpty || username.isEmpty || password != confirmPassword || authManager.isLoading)
            
            Button("Already have an account? Sign In") {
                dismiss()
            }
        }
        .padding()
        .navigationBarTitle("Register")
    }
}

// MARK: - Main Tab View
struct MainTabView: View {
    @EnvironmentObject var appState: AppState
    
    var body: some View {
        TabView(selection: $appState.selectedTab) {
            HomeView()
                .tabItem {
                    Label("Home", systemImage: "house.fill")
                }
                .tag(0)
            
            GamesView()
                .tabItem {
                    Label("Games", systemImage: "casino.fill")
                }
                .tag(1)
            
            TournamentsView()
                .tabItem {
                    Label("Tournaments", systemImage: "trophy.fill")
                }
                .tag(2)
            
            JackpotsView()
                .tabItem {
                    Label("Jackpots", systemImage: "star.fill")
                }
                .tag(3)
            
            WalletView()
                .tabItem {
                    Label("Wallet", systemImage: "creditcard.fill")
                }
                .tag(4)
            
            ProfileView()
                .tabItem {
                    Label("Profile", systemImage: "person.fill")
                }
                .tag(5)
        }
    }
}

// MARK: - Home View
struct HomeView: View {
    @State private var balance: Double = 0
    @State private var bonusBalance: Double = 0
    @State private var featuredGames: [Game] = []
    @State private var tournaments: [Tournament] = []
    @State private var jackpots: [Jackpot] = []
    @State private var isLoading = true
    
    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(spacing: 16) {
                    // Balance Card
                    VStack(spacing: 8) {
                        Text("Your Balance")
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
                    .padding(.horizontal)
                    
                    // Jackpots Section
                    if !jackpots.isEmpty {
                        VStack(alignment: .leading, spacing: 12) {
                            Text("Live Jackpots")
                                .font(.title2)
                                .fontWeight(.bold)
                                .padding(.horizontal)
                            
                            ScrollView(.horizontal, showsIndicators: false) {
                                HStack(spacing: 12) {
                                    ForEach(jackpots) { jackpot in
                                        JackpotCard(jackpot: jackpot)
                                    }
                                }
                                .padding(.horizontal)
                            }
                        }
                    }
                    
                    // Tournaments Section
                    if !tournaments.isEmpty {
                        VStack(alignment: .leading, spacing: 12) {
                            Text("Active Tournaments")
                                .font(.title2)
                                .fontWeight(.bold)
                                .padding(.horizontal)
                            
                            ForEach(tournaments) { tournament in
                                TournamentCard(tournament: tournament)
                                    .padding(.horizontal)
                            }
                        }
                    }
                    
                    // Featured Games Section
                    if !featuredGames.isEmpty {
                        VStack(alignment: .leading, spacing: 12) {
                            HStack {
                                Text("Featured Games")
                                    .font(.title2)
                                    .fontWeight(.bold)
                                Spacer()
                                NavigationLink("See All") {
                                    GamesView()
                                }
                            }
                            .padding(.horizontal)
                            
                            ScrollView(.horizontal, showsIndicators: false) {
                                HStack(spacing: 12) {
                                    ForEach(featuredGames) { game in
                                        GameCard(game: game)
                                    }
                                }
                                .padding(.horizontal)
                            }
                        }
                    }
                }
                .padding(.bottom, 20)
            }
            .navigationTitle("Home")
            .refreshable {
                await loadData()
            }
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
            
            let gamesResponse = try await APIClient.shared.getFeaturedGames()
            featuredGames = gamesResponse.games
            
            let tournamentsResponse = try await APIClient.shared.getTournaments(status: "active")
            tournaments = tournamentsResponse.tournaments
            
            let jackpotsResponse = try await APIClient.shared.getJackpots()
            jackpots = jackpotsResponse.jackpots
        } catch {
            print(error)
        }
        isLoading = false
    }
}

// MARK: - Game Card
struct GameCard: View {
    let game: Game
    
    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            RoundedRectangle(cornerRadius: 8)
                .fill(Color.purple.opacity(0.2))
                .frame(width: 140, height: 80)
                .overlay(
                    Image(systemName: "casino.fill")
                        .font(.title)
                        .foregroundColor(.purple)
                )
            
            Text(game.name)
                .font(.subheadline)
                .fontWeight(.medium)
                .lineLimit(1)
            
            Text(game.provider)
                .font(.caption)
                .foregroundColor(.secondary)
            
            Text("RTP: \(game.rtp, specifier: "%.1f")%")
                .font(.caption)
                .foregroundColor(.green)
        }
        .frame(width: 140)
    }
}

// MARK: - Tournament Card
struct TournamentCard: View {
    let tournament: Tournament
    
    var body: some View {
        HStack {
            VStack(alignment: .leading) {
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
                    .font(.title3)
                    .fontWeight(.bold)
                    .foregroundColor(.purple)
                Text("Prize Pool")
                    .font(.caption)
                    .foregroundColor(.secondary)
            }
        }
        .padding()
        .background(Color(.systemBackground))
        .cornerRadius(12)
        .shadow(radius: 2)
    }
}

// MARK: - Jackpot Card
struct JackpotCard: View {
    let jackpot: Jackpot
    
    var body: some View {
        VStack(alignment: .leading, spacing: 8) {
            HStack {
                Text(jackpot.name)
                    .font(.headline)
                Spacer()
                Image(systemName: "star.fill")
                    .foregroundColor(.yellow)
            }
            
            Text(jackpot.game)
                .font(.caption)
                .foregroundColor(.secondary)
            
            Text("$\(jackpot.currentAmount, specifier: "%.0f")")
                .font(.title2)
                .fontWeight(.bold)
                .foregroundColor(.purple)
            
            Text("\(jackpot.hitCount) wins")
                .font(.caption)
                .foregroundColor(.secondary)
        }
        .padding()
        .background(Color(.secondarySystemBackground))
        .cornerRadius(12)
    }
}
