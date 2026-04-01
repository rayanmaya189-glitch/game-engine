import SwiftUI

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
