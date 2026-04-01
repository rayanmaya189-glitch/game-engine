import SwiftUI

struct Transaction: Hashable {
    let type: String
    let amount: Double
    let date: String
}

struct WalletView: View {
    @State private var balance: Double = 0.00
    @State private var transactions: [Transaction] = []
    @State private var isLoading = true

    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(spacing: 20) {
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

                    HStack(spacing: 15) {
                        WalletButton(title: "Deposit", icon: "plus.circle.fill", color: .green) {
                        }

                        WalletButton(title: "Withdraw", icon: "arrow.down.circle.fill", color: .blue) {
                        }

                        WalletButton(title: "History", icon: "clock.fill", color: .orange) {
                        }
                    }
                    .padding(.horizontal)

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
                        }
                        .foregroundColor(.orange)
                    }
                    .padding()
                    .background(Color(.systemGray6))
                    .cornerRadius(12)
                    .padding(.horizontal)

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

                    Button("Logout") {
                    }
                    .foregroundColor(.red)
                    .padding(.top, 20)
                }
            }
            .navigationTitle("Profile")
        }
    }
}
