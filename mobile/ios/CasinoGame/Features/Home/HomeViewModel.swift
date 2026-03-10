import Foundation

struct HomeState {
    var isLoading: Bool = false
    var user: User?
    var balance: WalletBalance?
    var featuredGames: [Game] = []
    var popularGames: [Game] = []
    var jackpotGames: [Game] = []
    var categories: [GameCategory] = []
    var error: String?
}

class HomeViewModel {
    
    var onStateChange: ((HomeState) -> Void)?
    
    private(set) var state = HomeState() {
        didSet {
            onStateChange?(state)
        }
    }
    
    private let apiClient = APIClient.shared
    
    func loadData() {
        loadBalance()
        loadFeaturedGames()
    }
    
    private func loadBalance() {
        Task {
            do {
                let balance = try await apiClient.getBalance()
                state.balance = balance
            } catch {
                // Handle error silently
            }
        }
    }
    
    private func loadFeaturedGames() {
        state.isLoading = true
        
        Task {
            do {
                let response = try await apiClient.getFeaturedGames()
                state.isLoading = false
                state.featuredGames = response.featured
                state.popularGames = response.popular
                state.jackpotGames = response.jackpot
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }
    
    private func loadCategories() {
        Task {
            do {
                let response = try await apiClient.getCategories()
                state.categories = response.categories
            } catch {
                // Handle error silently
            }
        }
    }
}
