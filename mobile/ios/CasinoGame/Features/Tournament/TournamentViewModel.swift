import Foundation
import Combine

/**
 * Tournament View Model
 * 
 * Manages tournament state and operations.
 */
class TournamentViewModel: ObservableObject {
    
    @Published var uiState: TournamentUiState = .loading
    @Published var selectedTournament: Tournament?
    @Published var leaderboard: [LeaderboardEntry] = []
    
    private var cancellables = Set<AnyCancellable>()
    
    func loadTournaments() {
        uiState = .loading
        
        // API call would go here
        // For demo, simulate delay and empty list
        DispatchQueue.main.asyncAfter(deadline: .now() + 0.5) { [weak self] in
            self?.uiState = .success([])
        }
    }
    
    func filterTournaments(_ filter: TournamentFilter) {
        // Filter and reload
    }
    
    func selectTournament(_ tournamentId: String) {
        // Load tournament details
    }
    
    func registerForTournament(_ tournamentId: String) {
        // API call to register
    }
    
    func unregisterFromTournament(_ tournamentId: String) {
        // API call to unregister
    }
}

// MARK: - Models

enum TournamentUiState {
    case loading
    case success([Tournament])
    case error(String)
}

struct Tournament: Identifiable {
    let id: String
    let name: String
    let gameType: String
    let status: TournamentStatus
    let startTime: Date
    let endTime: Date
    let buyIn: Double
    let prizePool: Double
    let maxPlayers: Int
    let registeredPlayers: Int
    var isRegistered: Bool = false
}

enum TournamentStatus: String {
    case upcoming
    case registrationOpen = "registration_open"
    case registrationClosed = "registration_closed"
    case inProgress = "in_progress"
    case completed
    case cancelled
}

struct LeaderboardEntry: Identifiable {
    let id = UUID()
    let rank: Int
    let playerName: String
    let avatarUrl: String?
    let chipCount: Double
    var isCurrentPlayer: Bool = false
}

enum TournamentFilter: Int, CaseIterable {
    case all
    case upcoming
    case active
    case myTournaments
}
