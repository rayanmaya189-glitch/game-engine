import Foundation

struct GameDetailState {
    var isLoading: Bool = false
    var game: Game?
    var recentWins: [RecentWin] = []
    var relatedGames: [Game] = []
    var error: String?
}

class GameDetailViewModel {

    var onStateChange: ((GameDetailState) -> Void)?

    private(set) var state = GameDetailState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared
    let gameId: String

    init(gameId: String) {
        self.gameId = gameId
    }

    func loadGameDetail() {
        state.isLoading = true

        Task {
            do {
                let response = try await apiClient.getGameDetail(id: gameId)
                state.isLoading = false
                state.game = response.game
                state.recentWins = response.recentWins
                state.relatedGames = response.relatedGames
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func startGame(isFunPlay: Bool = false) async throws -> GameSession {
        return try await apiClient.startGame(id: gameId, isFunPlay: isFunPlay)
    }
}
