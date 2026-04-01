import SwiftUI

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
