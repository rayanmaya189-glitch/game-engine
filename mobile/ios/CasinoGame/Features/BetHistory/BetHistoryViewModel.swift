import Foundation

enum BetResultFilter: String, CaseIterable {
    case all = "all"
    case won = "won"
    case lost = "lost"
    case pending = "pending"

    var title: String {
        switch self {
        case .all: return "All"
        case .won: return "Won"
        case .lost: return "Lost"
        case .pending: return "Pending"
        }
    }
}

struct BetHistoryState {
    var isLoading: Bool = false
    var bets: [Bet] = []
    var filter: BetResultFilter = .all
    var error: String?
}

class BetHistoryViewModel {

    var onStateChange: ((BetHistoryState) -> Void)?

    private(set) var state = BetHistoryState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadBets() {
        state.isLoading = true
        let filterResult = state.filter == .all ? nil : state.filter.rawValue

        Task {
            do {
                let response = try await apiClient.getBetHistory(result: filterResult)
                state.isLoading = false
                state.bets = response.bets
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func setFilter(_ filter: BetResultFilter) {
        state.filter = filter
        state.bets = []
        loadBets()
    }
}
