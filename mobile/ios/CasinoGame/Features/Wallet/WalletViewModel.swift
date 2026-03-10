import Foundation

struct WalletState {
    var isLoading: Bool = false
    var wallet: Wallet?
    var balance: WalletBalance?
    var transactions: [Transaction] = []
    var paymentMethods: PaymentMethodsResponse?
    var error: String?
}

class WalletViewModel {
    
    var onStateChange: ((WalletState) -> Void)?
    
    private(set) var state = WalletState() {
        didSet {
            onStateChange?(state)
        }
    }
    
    private let apiClient = APIClient.shared
    
    func loadData() {
        loadWallet()
        loadTransactions()
    }
    
    private func loadWallet() {
        Task {
            do {
                let balance = try await apiClient.getBalance()
                state.balance = balance
            } catch {
                // Handle error silently
            }
        }
    }
    
    private func loadTransactions() {
        state.isLoading = true
        
        Task {
            do {
                let response = try await apiClient.getTransactions()
                state.isLoading = false
                state.transactions = response.transactions
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
            }
        }
    }
}
