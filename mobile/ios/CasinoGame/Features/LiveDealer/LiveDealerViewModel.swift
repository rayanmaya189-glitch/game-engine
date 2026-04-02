import Foundation

struct LiveDealerTable: Codable, Identifiable {
    let id: String
    let dealerName: String
    let gameType: String
    let currentPlayers: Int
    let maxPlayers: Int
    let minBet: Double
    let maxBet: Double
    let isActive: Bool
    let streamUrl: String?

    enum CodingKeys: String, CodingKey {
        case id
        case dealerName = "dealer_name"
        case gameType = "game_type"
        case currentPlayers = "current_players"
        case maxPlayers = "max_players"
        case minBet = "min_bet"
        case maxBet = "max_bet"
        case isActive = "is_active"
        case streamUrl = "stream_url"
    }
}

struct LiveDealerTablesResponse: Codable {
    let tables: [LiveDealerTable]
    let total: Int
}

struct LiveDealerState {
    var isLoading: Bool = false
    var tables: [LiveDealerTable] = []
    var selectedTableId: String?
    var error: String?
}

class LiveDealerViewModel {

    var onStateChange: ((LiveDealerState) -> Void)?

    private(set) var state = LiveDealerState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadTables() {
        state.isLoading = true
        Task {
            do {
                let response = try await apiClient.getLiveDealerTables()
                state.isLoading = false
                state.tables = response.tables
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func joinTable(_ tableId: String) {
        state.selectedTableId = tableId
        Task {
            do {
                let _ = try await apiClient.joinLiveDealerTable(id: tableId)
            } catch let error as APIError {
                state.error = error.errorDescription
            } catch {
                state.error = error.localizedDescription
            }
        }
    }
}
