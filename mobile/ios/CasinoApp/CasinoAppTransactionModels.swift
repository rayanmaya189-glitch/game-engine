import Foundation

struct BalanceResponse: Codable {
    let balance: Double
    let bonusBalance: Double
    let pendingBalance: Double
    let currency: String

    enum CodingKeys: String, CodingKey {
        case balance, currency
        case bonusBalance = "bonus_balance"
        case pendingBalance = "pending_balance"
    }
}

struct Transaction: Codable, Identifiable {
    let id: String
    let type: String
    let amount: Double
    let status: String
    let method: String?
    let createdAt: String
    let transactionId: String?

    enum CodingKeys: String, CodingKey {
        case id, type, amount, status, method
        case createdAt = "created_at"
        case transactionId = "transaction_id"
    }
}

struct TransactionsResponse: Codable {
    let transactions: [Transaction]
    let total: Int
    let page: Int
    let pages: Int
}

struct Bonus: Codable, Identifiable {
    let id: String
    let name: String
    let type: String
    let description: String?
    let amount: Double
    let maxBonus: Double
    let minDeposit: Double?
    let wagerRequirement: Int?
    let expiresAt: String?
    let status: String

    enum CodingKeys: String, CodingKey {
        case id, name, type, description, amount, status
        case maxBonus = "max_bonus"
        case minDeposit = "min_deposit"
        case wagerRequirement = "wager_requirement"
        case expiresAt = "expires_at"
    }
}

struct BonusesResponse: Codable {
    let bonuses: [Bonus]
}

struct ClaimBonusResponse: Codable {
    let success: Bool
    let message: String
    let bonusAmount: Double?

    enum CodingKeys: String, CodingKey {
        case success, message
        case bonusAmount = "bonus_amount"
    }
}
