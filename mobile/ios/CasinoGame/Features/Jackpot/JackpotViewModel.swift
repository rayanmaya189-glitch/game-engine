import Foundation

struct JackpotWinner: Codable {
    let username: String
    let amount: Double
    let gameName: String
    let createdAt: String

    enum CodingKeys: String, CodingKey {
        case username, amount
        case gameName = "game_name"
        case createdAt = "created_at"
    }
}

struct JackpotState {
    var isLoading: Bool = false
    var jackpotAmount: Double = 0
    var winners: [JackpotWinner] = []
    var error: String?
}

class JackpotViewModel {

    var onStateChange: ((JackpotState) -> Void)?

    private(set) var state = JackpotState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadData() {
        state.isLoading = true
        Task {
            do {
                let response = try await apiClient.getFeaturedGames()
                let totalJackpot = response.jackpot.compactMap { $0.jackpotAmount }.reduce(0, +)
                state.isLoading = false
                state.jackpotAmount = totalJackpot
                loadWinners()
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    private func loadWinners() {
        Task {
            do {
                let winners = try await apiClient.getJackpotWinners()
                state.winners = winners
            } catch { }
        }
    }
}
