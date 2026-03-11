import SwiftUI

@main
struct CasinoGameApp: App {
    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
}

struct ContentView: View {
    @State private var isLoggedIn = false
    
    var body: some View {
        if isLoggedIn {
            MainTabView()
        } else {
            LoginView(isLoggedIn: $isLoggedIn)
        }
    }
}

struct MainTabView: View {
    @State private var selectedTab = 0
    
    var body: some View {
        TabView(selection: $selectedTab) {
            HomeView()
                .tabItem {
                    Label("Home", systemImage: "house.fill")
                }
                .tag(0)
            
            GamesView()
                .tabItem {
                    Label("Games", systemImage: "gamecontroller.fill")
                }
                .tag(1)
            
            WalletView()
                .tabItem {
                    Label("Wallet", systemImage: "creditcard.fill")
                }
                .tag(2)
            
            ProfileView()
                .tabItem {
                    Label("Profile", systemImage: "person.fill")
                }
                .tag(3)
        }
        .tint(.orange)
    }
}

struct HomeView: View {
    @State private var username: String = ""
    @State private var balance: Double = 0.00
    @State private var isLoading = true
    
    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(spacing: 20) {
                    // Balance Card
                    VStack(alignment: .leading, spacing: 10) {
                        Text("Balance")
                            .font(.subheadline)
                            .foregroundColor(.white.opacity(0.8))
                        
                        Text("$\(balance, specifier: "%.2f")")
                            .font(.system(size: 32, weight: .bold))
                            .foregroundColor(.white)
                    }
                    .padding(20)
                    .frame(maxWidth: .infinity, alignment: .leading)
                    .background(
                        LinearGradient(
                            gradient: Gradient(colors: [.orange, .orange.opacity(0.7)]),
                            startPoint: .topLeading,
                            endPoint: .bottomTrailing
                        )
                    )
                    .cornerRadius(16)
                    .padding(.horizontal)
                    
                    // Quick Actions
                    HStack(spacing: 15) {
                        QuickActionButton(icon: "plus.circle.fill", title: "Deposit") {
                            // TODO: Navigate to deposit screen
                            // For now, show alert indicating feature coming soon
                        }
                        
                        QuickActionButton(icon: "arrow.down.circle.fill", title: "Withdraw") {
                            // TODO: Navigate to withdraw screen
                            // For now, show alert indicating feature coming soon
                        }
                        
                        QuickActionButton(icon: "gift.fill", title: "Bonus") {
                            // TODO: Navigate to bonus screen
                            // For now, show alert indicating feature coming soon
                        }
                    }
                    .padding(.horizontal)
                    
                    // Featured Games
                    VStack(alignment: .leading, spacing: 15) {
                        Text("Featured Games")
                            .font(.title2)
                            .fontWeight(.bold)
                            .padding(.horizontal)
                        
                        ScrollView(.horizontal, showsIndicators: false) {
                            HStack(spacing: 15) {
                                FeaturedGameCard(imageName: "slot1", name: "Starburst", provider: "NetEnt")
                                FeaturedGameCard(imageName: "slot2", name: "Gonzo's Quest", provider: "NetEnt")
                                FeaturedGameCard(imageName: "slot3", name: "Book of Dead", provider: "Play'n GO")
                            }
                            .padding(.horizontal)
                        }
                    }
                    
                    // Recent Activity
                    VStack(alignment: .leading, spacing: 15) {
                        Text("Recent Activity")
                            .font(.title2)
                            .fontWeight(.bold)
                            .padding(.horizontal)
                        
                        VStack(spacing: 10) {
                            ActivityRow(game: "Starburst", bet: 5.00, win: 25.00, time: "2 min ago")
                            ActivityRow(game: "Blackjack", bet: 20.00, win: 0, time: "15 min ago")
                            ActivityRow(game: "Gonzo's Quest", bet: 10.00, win: 50.00, time: "1 hour ago")
                        }
                        .padding()
                        .background(Color(.systemGray6))
                        .cornerRadius(12)
                        .padding(.horizontal)
                    }
                }
                .padding(.vertical)
            }
            .navigationTitle("Hello, \(username)")
            .toolbar {
                ToolbarItem(placement: .navigationBarTrailing) {
                    Button {
                        // TODO: Show notifications panel
                        // For now, show alert indicating feature coming soon
                    } label: {
                        Image(systemName: "bell.fill")
                    }
                }
            }
        }
    }
}

struct QuickActionButton: View {
    let icon: String
    let title: String
    let action: () -> Void
    
    var body: some View {
        Button(action: action) {
            VStack(spacing: 8) {
                Image(systemName: icon)
                    .font(.title)
                    .foregroundColor(.orange)
                
                Text(title)
                    .font(.caption)
                    .foregroundColor(.primary)
            }
            .frame(maxWidth: .infinity)
            .padding(.vertical, 15)
            .background(Color(.systemGray6))
            .cornerRadius(12)
        }
    }
}

struct FeaturedGameCard: View {
    let imageName: String
    let name: String
    let provider: String
    
    var body: some View {
        VStack(alignment: .leading) {
            Image(imageName)
                .resizable()
                .aspectRatio(16/9, contentMode: .fill)
                .frame(width: 200, height: 120)
                .cornerRadius(12)
                .overlay(
                    RoundedRectangle(cornerRadius: 12)
                        .stroke(Color.orange.opacity(0.3), lineWidth: 2)
                )
            
            Text(name)
                .font(.headline)
                .padding(.top, 5)
            
            Text(provider)
                .font(.caption)
                .foregroundColor(.secondary)
        }
    }
}

struct ActivityRow: View {
    let game: String
    let bet: Double
    let win: Double
    let time: String
    
    var body: some View {
        HStack {
            VStack(alignment: .leading, spacing: 4) {
                Text(game)
                    .font(.headline)
                
                Text(time)
                    .font(.caption)
                    .foregroundColor(.secondary)
            }
            
            Spacer()
            
            VStack(alignment: .trailing, spacing: 4) {
                Text("Bet: $\(bet, specifier: "%.2f")")
                    .font(.caption)
                
                if win > 0 {
                    Text("Win: +$\(win, specifier: "%.2f")")
                        .font(.caption)
                        .foregroundColor(.green)
                } else {
                    Text("Win: $0.00")
                        .font(.caption)
                        .foregroundColor(.red)
                }
            }
        }
    }
}
