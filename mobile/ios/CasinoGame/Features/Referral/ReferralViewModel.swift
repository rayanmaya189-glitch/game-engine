import Foundation

struct ReferralStats: Codable {
    let totalReferrals: Int
    let totalEarnings: Double
    let pendingEarnings: Double

    enum CodingKeys: String, CodingKey {
        case totalReferrals = "total_referrals"
        case totalEarnings = "total_earnings"
        case pendingEarnings = "pending_earnings"
    }
}

struct ReferralEntry: Codable, Identifiable {
    let id: String
    let username: String
    let status: String
    let earnedAmount: Double
    let joinedAt: String

    enum CodingKeys: String, CodingKey {
        case id, username, status
        case earnedAmount = "earned_amount"
        case joinedAt = "joined_at"
    }
}

struct ReferralHistoryResponse: Codable {
    let referrals: [ReferralEntry]
    let total: Int
}

struct ReferralCodeResponse: Codable {
    let code: String
    let shareUrl: String

    enum CodingKeys: String, CodingKey {
        case code
        case shareUrl = "share_url"
    }
}

struct ReferralState {
    var isLoading: Bool = false
    var referralCode: String = ""
    var shareUrl: String = ""
    var stats: ReferralStats?
    var history: [ReferralEntry] = []
    var error: String?
}

class ReferralViewModel {

    var onStateChange: ((ReferralState) -> Void)?

    private(set) var state = ReferralState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadReferralData() {
        state.isLoading = true
        Task {
            do {
                async let codeResponse = apiClient.getReferralCode()
                async let statsResponse = apiClient.getReferralStats()
                async let historyResponse = apiClient.getReferralHistory()

                let (code, stats, history) = try await (codeResponse, statsResponse, historyResponse)
                state.isLoading = false
                state.referralCode = code.code
                state.shareUrl = code.shareUrl
                state.stats = stats
                state.history = history.referrals
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
