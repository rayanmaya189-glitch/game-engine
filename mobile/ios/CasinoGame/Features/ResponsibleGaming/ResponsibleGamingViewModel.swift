import Foundation

struct GamingLimits: Codable {
    var depositLimit: Double
    var lossLimit: Double
    var sessionLimit: Int
    var coolOffEnabled: Bool
    var selfExclusionDays: Int

    enum CodingKeys: String, CodingKey {
        case depositLimit = "deposit_limit"
        case lossLimit = "loss_limit"
        case sessionLimit = "session_limit"
        case coolOffEnabled = "cool_off_enabled"
        case selfExclusionDays = "self_exclusion_days"
    }
}

struct GamingLimitsResponse: Codable {
    let limits: GamingLimits
}

struct ResponsibleGamingState {
    var isLoading: Bool = false
    var limits: GamingLimits = GamingLimits(depositLimit: 1000, lossLimit: 500, sessionLimit: 60, coolOffEnabled: false, selfExclusionDays: 0)
    var isSaved: Bool = false
    var error: String?
}

class ResponsibleGamingViewModel {

    var onStateChange: ((ResponsibleGamingState) -> Void)?

    private(set) var state = ResponsibleGamingState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadLimits() {
        state.isLoading = true
        Task {
            do {
                let response = try await apiClient.getGamingLimits()
                state.isLoading = false
                state.limits = response.limits
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func updateDepositLimit(_ value: Double) {
        state.limits.depositLimit = value
    }

    func updateLossLimit(_ value: Double) {
        state.limits.lossLimit = value
    }

    func updateSessionLimit(_ value: Int) {
        state.limits.sessionLimit = value
    }

    func toggleCoolOff(_ enabled: Bool) {
        state.limits.coolOffEnabled = enabled
    }

    func setSelfExclusion(days: Int) {
        state.limits.selfExclusionDays = days
    }

    func saveLimits() {
        state.isLoading = true
        Task {
            do {
                let limits = try await apiClient.updateGamingLimits(limits: state.limits)
                state.isLoading = false
                state.limits = limits
                state.isSaved = true
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
