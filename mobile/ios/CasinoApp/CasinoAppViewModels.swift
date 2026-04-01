import SwiftUI

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
