import Foundation

struct Bonus: Codable, Identifiable {
    let id: String
    let name: String
    let description: String
    let type: String
    let amount: Double
    let wageringRequirement: Double
    let wageringProgress: Double
    let expiresAt: String?
    let isActive: Bool

    enum CodingKeys: String, CodingKey {
        case id, name, description, type, amount
        case wageringRequirement = "wagering_requirement"
        case wageringProgress = "wagering_progress"
        case expiresAt = "expires_at"
        case isActive = "is_active"
    }
}

struct BonusState {
    var isLoading: Bool = false
    var availableBonuses: [Bonus] = []
    var activeBonuses: [Bonus] = []
    var claimedBonusId: String?
    var error: String?
}

class BonusViewModel {

    var onStateChange: ((BonusState) -> Void)?

    private(set) var state = BonusState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadBonuses() {
        state.isLoading = true
        Task {
            do {
                let bonuses = try await apiClient.getBonuses()
                state.isLoading = false
                state.availableBonuses = bonuses.filter { !$0.isActive }
                state.activeBonuses = bonuses.filter { $0.isActive }
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func claimBonus(_ bonusId: String) {
        state.isLoading = true
        state.error = nil
        Task {
            do {
                try await apiClient.claimBonus(id: bonusId)
                state.isLoading = false
                state.claimedBonusId = bonusId
                loadBonuses()
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }
}
