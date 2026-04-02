import Foundation

enum PaymentFilter: String, CaseIterable {
    case all = "all"
    case deposits = "deposit"
    case withdrawals = "withdrawal"

    var title: String {
        switch self {
        case .all: return "All"
        case .deposits: return "Deposits"
        case .withdrawals: return "Withdrawals"
        }
    }
}

struct PaymentHistoryState {
    var isLoading: Bool = false
    var transactions: [Transaction] = []
    var filter: PaymentFilter = .all
    var error: String?
}

class PaymentHistoryViewModel {

    var onStateChange: ((PaymentHistoryState) -> Void)?

    private(set) var state = PaymentHistoryState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadTransactions() {
        state.isLoading = true
        let filterType = state.filter == .all ? nil : state.filter.rawValue

        Task {
            do {
                let response = try await apiClient.getPaymentHistory(type: filterType)
                state.isLoading = false
                state.transactions = response.transactions
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func setFilter(_ filter: PaymentFilter) {
        state.filter = filter
        state.transactions = []
        loadTransactions()
    }
}
