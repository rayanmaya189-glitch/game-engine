import Foundation

struct GamesState {
    var isLoading: Bool = false
    var games: [Game] = []
    var categories: [GameCategory] = []
    var selectedCategory: String?
    var searchQuery: String = ""
    var error: String?
}

class GamesViewModel {
    
    var onStateChange: ((GamesState) -> Void)?
    
    private(set) var state = GamesState() {
        didSet {
            onStateChange?(state)
        }
    }
    
    private let apiClient = APIClient.shared
    
    func loadGames(category: String? = nil, search: String? = nil) {
        state.isLoading = true
        
        Task {
            do {
                let response = try await apiClient.getGames(
                    page: 1,
                    category: category,
                    search: search
                )
                state.isLoading = false
                state.games = response.games
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }
    
    func search(query: String) {
        state.searchQuery = query
        loadGames(search: query.isEmpty ? nil : query)
    }
    
    func selectCategory(_ category: String?) {
        state.selectedCategory = category
        loadGames(category: category)
    }
}
