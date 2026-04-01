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


