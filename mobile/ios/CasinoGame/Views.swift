import SwiftUI

struct LoginView: View {
    @Binding var isLoggedIn: Bool
    @State private var email = ""
    @State private var password = ""
    @State private var rememberMe = false
    @State private var showError = false
    @State private var isLoading = false
    
    private func loginUser() {
        guard !email.isEmpty, !password.isEmpty else {
            showError = true
            return
        }
        
        guard email.contains("@"), password.count >= 4 else {
            showError = true
            return
        }
        
        isLoading = true
        
        // Call the actual API for authentication
        Task {
            do {
                let _ = try await APIClient.shared.login(email: email, password: password)
                await MainActor.run {
                    isLoading = false
                    isLoggedIn = true
                    if rememberMe {
                        UserDefaults.standard.set(true, forKey: "rememberMe")
                    }
                }
            } catch {
                await MainActor.run {
                    isLoading = false
                    showError = true
                }
            }
        }
    }
    
    var body: some View {
        NavigationStack {
            VStack(spacing: 30) {
                Spacer()
                
                // Logo
                VStack(spacing: 10) {
                    Image(systemName: "gamecontroller.fill")
                        .font(.system(size: 60))
                        .foregroundColor(.orange)
                    
                    Text("Casino Game")
                        .font(.largeTitle)
                        .fontWeight(.bold)
                    
                    Text("Play & Win Big")
                        .font(.subheadline)
                        .foregroundColor(.secondary)
                }
                
                VStack(spacing: 20) {
                    // Email Field
                    VStack(alignment: .leading, spacing: 8) {
                        Text("Email")
                            .font(.caption)
                            .foregroundColor(.secondary)
                        
                        TextField("Enter your email", text: $email)
                            .textFieldStyle(RoundedTextFieldStyle())
                            .textContentType(.emailAddress)
                            .autocapitalization(.none)
                    }
                    
                    // Password Field
                    VStack(alignment: .leading, spacing: 8) {
                        Text("Password")
                            .font(.caption)
                            .foregroundColor(.secondary)
                        
                        SecureField("Enter your password", text: $password)
                            .textFieldStyle(RoundedTextFieldStyle())
                            .textContentType(.password)
                    }
                    
                    // Remember Me & Forgot Password
                    HStack {
                        Toggle("Remember me", isOn: $rememberMe)
                            .tint(.orange)
                        
                        Spacer()
                        
                        Button("Forgot Password?") {
                            // TODO: Implement forgot password flow
                            // For now, show alert indicating feature coming soon
                        }
                        .font(.subheadline)
                        .foregroundColor(.orange)
                    }
                    
                    // Login Button
                    Button {
                        // TODO: Replace with actual API authentication
                        // Login should call AuthService.login(email:password:completion:)
                        loginUser()
                    } label: {
                        Text("Login")
                            .font(.headline)
                            .foregroundColor(.white)
                            .frame(maxWidth: .infinity)
                            .padding()
                            .background(Color.orange)
                            .cornerRadius(12)
                    }
                    
                    // Register Link
                    HStack {
                        Text("Don't have an account?")
                            .foregroundColor(.secondary)
                        
                        Button("Register") {
                            // TODO: Navigate to registration screen
                            // For now, show alert indicating feature coming soon
                        }
                        .foregroundColor(.orange)
                        .fontWeight(.semibold)
                    }
                }
                .padding(.horizontal, 30)
                
                Spacer()
            }
            .alert("Login Failed", isPresented: $showError) {
                Button("OK", role: .cancel) { }
            } message: {
                Text("Invalid email or password. Please try again.")
            }
        }
    }
}

struct RoundedTextFieldStyle: TextFieldStyle {
    func _body(configuration: TextField<Self._Label>) -> some View {
        configuration
            .padding()
            .background(Color(.systemGray6))
            .cornerRadius(12)
    }
}

struct GamesView: View {
    @State private var searchText = ""
    @State private var selectedCategory = "All"
    
    let categories = ["All", "Slots", "Card Games", "Table Games", "Live Casino", "Jackpots"]
    
    var body: some View {
        NavigationStack {
            VStack(spacing: 0) {
                // Search Bar
                HStack {
                    Image(systemName: "magnifyingglass")
                        .foregroundColor(.secondary)
                    
                    TextField("Search games", text: $searchText)
                }
                .padding()
                .background(Color(.systemGray6))
                .cornerRadius(12)
                .padding()
                
                // Categories
                ScrollView(.horizontal, showsIndicators: false) {
                    HStack(spacing: 12) {
                        ForEach(categories, id: \.self) { category in
                            CategoryButton(
                                title: category,
                                isSelected: selectedCategory == category
                            ) {
                                selectedCategory = category
                            }
                        }
                    }
                    .padding(.horizontal)
                }
                
                // Games Grid
                ScrollView {
                    LazyVGrid(columns: [
                        GridItem(.flexible(), spacing: 15),
                        GridItem(.flexible(), spacing: 15)
                    ], spacing: 15) {
                        GameCard(name: "Starburst", provider: "NetEnt", imageColor: .purple)
                        GameCard(name: "Gonzo's Quest", provider: "NetEnt", imageColor: .blue)
                        GameCard(name: "Book of Dead", provider: "Play'n GO", imageColor: .orange)
                        GameCard(name: "Mega Moolah", provider: "Microgaming", imageColor: .yellow)
                        GameCard(name: "Blackjack", provider: "Evolution", imageColor: .green)
                        GameCard(name: "Roulette", provider: "Evolution", imageColor: .red)
                    }
                    .padding()
                }
            }
            .navigationTitle("Games")
        }
    }
}

struct CategoryButton: View {
    let title: String
    let isSelected: Bool
    let action: () -> Void
    
    var body: some View {
        Button(action: action) {
            Text(title)
                .font(.subheadline)
                .fontWeight(isSelected ? .bold : .regular)
                .foregroundColor(isSelected ? .white : .primary)
                .padding(.horizontal, 16)
                .padding(.vertical, 8)
                .background(isSelected ? Color.orange : Color(.systemGray6))
                .cornerRadius(20)
        }
    }
}

struct GameCard: View {
    let name: String
    let provider: String
    let imageColor: Color
    
    var body: some View {
        VStack(alignment: .leading) {
            ZStack {
                RoundedRectangle(cornerRadius: 12)
                    .fill(imageColor.opacity(0.3))
                    .aspectRatio(1, contentMode: .fit)
                
                Image(systemName: "gamecontroller.fill")
                    .font(.largeTitle)
                    .foregroundColor(imageColor)
            }
            
            Text(name)
                .font(.headline)
                .lineLimit(1)
            
            Text(provider)
                .font(.caption)
                .foregroundColor(.secondary)
        }
        .onTapGesture {
            // TODO: Navigate to game detail screen
            // For now, show alert or log action
        }
    }
}

struct WalletView: View {
    @State private var balance: Double = 0.00
    @State private var transactions: [Transaction] = []
    @State private var isLoading = true
    
    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(spacing: 20) {
                    // Balance Card
                    VStack(spacing: 10) {
                        Text("Total Balance")
                            .font(.subheadline)
                            .foregroundColor(.white.opacity(0.8))
                        
                        if isLoading {
                            ProgressView()
                                .progressViewStyle(CircularProgressViewStyle(tint: .white))
                        } else {
                            Text("$\(balance, specifier: "%.2f")")
                                .font(.system(size: 40, weight: .bold))
                                .foregroundColor(.white)
                        }
                    }
                    .padding(30)
                    .frame(maxWidth: .infinity)
                    .background(
                        LinearGradient(
                            gradient: Gradient(colors: [.orange, .purple]),
                            startPoint: .topLeading,
                            endPoint: .bottomTrailing
                        )
                    )
                    .cornerRadius(20)
                    .padding()
                    
                    // Action Buttons
                    HStack(spacing: 15) {
                        WalletButton(title: "Deposit", icon: "plus.circle.fill", color: .green) {
                            // TODO: Navigate to deposit screen
                            // For now, show alert indicating feature coming soon
                        }
                        
                        WalletButton(title: "Withdraw", icon: "arrow.down.circle.fill", color: .blue) {
                            // TODO: Navigate to withdraw screen
                            // For now, show alert indicating feature coming soon
                        }
                        
                        WalletButton(title: "History", icon: "clock.fill", color: .orange) {
                            // TODO: Navigate to transaction history
                            // For now, show alert indicating feature coming soon
                        }
                    }
                    .padding(.horizontal)
                    
                    // Transactions
                    VStack(alignment: .leading, spacing: 15) {
                        Text("Recent Transactions")
                            .font(.title2)
                            .fontWeight(.bold)
                            .padding(.horizontal)
                        
                        ForEach(transactions, id: \.self) { transaction in
                            TransactionRow(transaction: transaction)
                        }
                    }
                }
            }
            .navigationTitle("Wallet")
        }
    }
}

struct Transaction: Hashable {
    let type: String
    let amount: Double
    let date: String
}

struct WalletButton: View {
    let title: String
    let icon: String
    let color: Color
    let action: () -> Void
    
    var body: some View {
        Button(action: action) {
            VStack(spacing: 8) {
                Image(systemName: icon)
                    .font(.title2)
                    .foregroundColor(color)
                
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

struct TransactionRow: View {
    let transaction: Transaction
    
    var body: some View {
        HStack {
            Image(systemName: transactionIcon)
                .foregroundColor(transactionColor)
                .frame(width: 40, height: 40)
                .background(transactionColor.opacity(0.1))
                .cornerRadius(10)
            
            VStack(alignment: .leading, spacing: 4) {
                Text(transactionTitle)
                    .font(.headline)
                
                Text(transaction.date)
                    .font(.caption)
                    .foregroundColor(.secondary)
            }
            
            Spacer()
            
            Text("\(transaction.amount >= 0 ? "+" : "")$\(transaction.amount, specifier: "%.2f")")
                .font(.headline)
                .foregroundColor(transactionColor)
        }
        .padding()
        .background(Color(.systemGray6))
        .cornerRadius(12)
        .padding(.horizontal)
    }
    
    var transactionIcon: String {
        switch transaction.type {
        case "deposit": return "arrow.down.circle.fill"
        case "withdraw": return "arrow.up.circle.fill"
        case "bet": return "gamecontroller.fill"
        case "win": return "star.fill"
        default: return "dollarsign.circle.fill"
        }
    }
    
    var transactionColor: Color {
        switch transaction.type {
        case "deposit", "win": return .green
        case "withdraw", "bet": return .red
        default: return .orange
        }
    }
    
    var transactionTitle: String {
        switch transaction.type {
        case "deposit": return "Deposit"
        case "withdraw": return "Withdrawal"
        case "bet": return "Bet Placed"
        case "win": return "Winning"
        default: return "Transaction"
        }
    }
}

struct ProfileView: View {
    @State private var username: String = ""
    @State private var email: String = ""
    @State private var tier: String = ""
    @State private var points: Int = 0
    @State private var isLoading = true
    
    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(spacing: 20) {
                    // Profile Header
                    VStack(spacing: 15) {
                        Image(systemName: "person.circle.fill")
                            .font(.system(size: 80))
                            .foregroundColor(.orange)
                        
                        if isLoading {
                            ProgressView()
                        } else {
                            Text(username)
                                .font(.title)
                                .fontWeight(.bold)
                            
                            // Tier Badge
                            if !tier.isEmpty {
                                HStack {
                                    Image(systemName: "star.fill")
                                        .foregroundColor(.yellow)
                                    Text(tier)
                                        .fontWeight(.semibold)
                                }
                                .padding(.horizontal, 16)
                                .padding(.vertical, 6)
                                .background(Color.orange.opacity(0.2))
                                .cornerRadius(20)
                            }
                        }
                    }
                    .padding(.top)
                    
                    // Points Card
                    HStack {
                        VStack(alignment: .leading, spacing: 5) {
                            Text("Loyalty Points")
                                .font(.subheadline)
                                .foregroundColor(.secondary)
                            
                            if isLoading {
                                ProgressView()
                            } else {
                                Text("\(points)")
                                    .font(.title)
                                    .fontWeight(.bold)
                            }
                        }
                        
                        Spacer()
                        
                        Button("Redeem") {
                            // TODO: Navigate to reward redemption screen
                            // For now, show alert indicating feature coming soon
                        }
                        .foregroundColor(.orange)
                    }
                    .padding()
                    .background(Color(.systemGray6))
                    .cornerRadius(12)
                    .padding(.horizontal)
                    
                    // Menu Items
                    VStack(spacing: 0) {
                        ProfileMenuItem(icon: "person.fill", title: "Edit Profile") { }
                        ProfileMenuItem(icon: "bell.fill", title: "Notifications") { }
                        ProfileMenuItem(icon: "lock.fill", title: "Security") { }
                        ProfileMenuItem(icon: "creditcard.fill", title: "Payment Methods") { }
                        ProfileMenuItem(icon: "doc.text.fill", title: "Transaction History") { }
                        ProfileMenuItem(icon: "questionmark.circle.fill", title: "Help & Support") { }
                        ProfileMenuItem(icon: "info.circle.fill", title: "About") { }
                    }
                    .padding(.horizontal)
                    
                    // Logout Button
                    Button("Logout") {
                        // Logout
                    }
                    .foregroundColor(.red)
                    .padding(.top, 20)
                }
            }
            .navigationTitle("Profile")
        }
    }
}

struct ProfileMenuItem: View {
    let icon: String
    let title: String
    let action: () -> Void
    
    var body: some View {
        Button(action: action) {
            HStack {
                Image(systemName: icon)
                    .foregroundColor(.orange)
                    .frame(width: 30)
                
                Text(title)
                    .foregroundColor(.primary)
                
                Spacer()
                
                Image(systemName: "chevron.right")
                    .foregroundColor(.secondary)
                    .font(.caption)
            }
            .padding()
        }
    }
}
