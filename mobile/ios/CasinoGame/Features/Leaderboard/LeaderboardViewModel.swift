import Foundation

struct LeaderboardPlayer: Codable, Identifiable {
    let id: String
    let rank: Int
    let username: String
    let score: Double
    let prize: Double
    let avatarUrl: String?
    let isCurrentUser: Bool

    enum CodingKeys: String, CodingKey {
        case id, rank, username, score, prize
        case avatarUrl = "avatar_url"
        case isCurrentUser = "is_current_user"
    }
}

struct LeaderboardResponse: Codable {
    let entries: [LeaderboardPlayer]
    let period: String
}

enum LeaderboardPeriod: String, CaseIterable {
    case daily
    case weekly
    case monthly
}

struct LeaderboardState {
    var isLoading: Bool = false
    var entries: [LeaderboardPlayer] = []
    var period: LeaderboardPeriod = .daily
    var error: String?
}

class LeaderboardViewModel {

    var onStateChange: ((LeaderboardState) -> Void)?

    private(set) var state = LeaderboardState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadLeaderboard() {
        state.isLoading = true
        Task {
            do {
                let response = try await apiClient.getLeaderboard(period: state.period.rawValue)
                state.isLoading = false
                state.entries = response.entries
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func changePeriod(_ period: LeaderboardPeriod) {
        state.period = period
        loadLeaderboard()
    }
}
